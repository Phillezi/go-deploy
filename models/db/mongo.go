package db

import (
	"context"
	"fmt"
	"go-deploy/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strings"
	"time"
)

type CollectionDefinition struct {
	Name    string
	Indexes []string
	// unique only for non-deleted documents
	UniqueIndexes []string
	// unique even for deleted documents
	TotallyUniqueIndexes []string
	TextIndexFields      []string
}

func (dbCtx *Context) setupMongo() error {
	makeError := func(err error) error {
		return fmt.Errorf("failed to setup mongodb. details: %w", err)
	}

	log.Println("setting up mongodb")

	var err error
	dbCtx.mongoClient, err = mongo.NewClient(options.Client().ApplyURI(config.Config.MongoDB.URL))
	if err != nil {
		return makeError(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = dbCtx.mongoClient.Connect(ctx)
	if err != nil {
		return makeError(err)
	}

	err = dbCtx.mongoClient.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalln(makeError(err))
	}

	log.Println("connected to mongodb")

	// Find collections
	DB.CollectionMap = make(map[string]*mongo.Collection)

	DB.CollectionDefinitionMap = getCollectionDefinitions()

	for _, def := range DB.CollectionDefinitionMap {
		DB.CollectionMap[def.Name] = dbCtx.mongoClient.Database(config.Config.MongoDB.Name).Collection(def.Name)
	}

	log.Println("found", len(DB.CollectionDefinitionMap), "collections")

	ensureCount := 0
	for _, def := range DB.CollectionDefinitionMap {
		for _, indexName := range def.Indexes {
			_, err = DB.GetCollection(def.Name).Indexes().CreateOne(context.Background(), mongo.IndexModel{
				Keys:    map[string]int{indexName: 1},
				Options: options.Index().SetUnique(false),
			})
			if err != nil && !isIndexExistsError(err) {
				return makeError(err)
			}

			ensureCount++
		}
	}

	log.Println("ensured", ensureCount, "indexes")

	ensureCount = 0
	for _, def := range DB.CollectionDefinitionMap {
		for _, indexName := range def.UniqueIndexes {
			_, err = DB.GetCollection(def.Name).Indexes().CreateOne(context.Background(), mongo.IndexModel{
				Keys:    map[string]int{indexName: 1},
				Options: options.Index().SetUnique(true).SetPartialFilterExpression(bson.D{{"deletedAt", bson.D{{"$in", []interface{}{nil, time.Time{}}}}}}),
			})
			if err != nil && !isIndexExistsError(err) {
				return makeError(err)
			}

			ensureCount++
		}
	}

	log.Println("ensured", ensureCount, "unique indexes")

	ensureCount = 0
	for _, def := range DB.CollectionDefinitionMap {
		for _, indexName := range def.UniqueIndexes {
			_, err = DB.GetCollection(def.Name).Indexes().CreateOne(context.Background(), mongo.IndexModel{
				Keys:    map[string]int{indexName: 1},
				Options: options.Index().SetUnique(true),
			})
			if err != nil && !isIndexExistsError(err) {
				return makeError(err)
			}

			ensureCount++
		}
	}

	log.Println("ensured", ensureCount, "totally unique indexes")

	ensureCount = 0
	for _, def := range DB.CollectionDefinitionMap {
		if def.TextIndexFields == nil {
			continue
		}

		keys := bson.D{}
		for _, indexName := range def.TextIndexFields {
			keys = append(keys, bson.E{Key: indexName, Value: "text"})
		}

		_, err = DB.GetCollection(def.Name).Indexes().CreateOne(context.Background(), mongo.IndexModel{
			Keys: keys,
		})
		if err != nil && !isIndexExistsError(err) {
			return makeError(err)
		}

		ensureCount++
	}

	log.Println("ensured", ensureCount, "text indexes")

	return nil
}

func (dbCtx *Context) shutdownMongo() error {
	makeError := func(err error) error {
		return fmt.Errorf("failed to shutdown database. details: %w", err)
	}

	err := dbCtx.mongoClient.Disconnect(context.Background())
	if err != nil {
		return makeError(err)
	}

	dbCtx.CollectionMap = nil

	return nil
}

func getCollectionDefinitions() map[string]CollectionDefinition {
	return map[string]CollectionDefinition{
		"deployments": {
			Name:                 "deployments",
			Indexes:              []string{"ownerId", "type", "statusCode", "createdAt", "deletedAt", "repairedAt", "restartedAt", "zone"},
			UniqueIndexes:        []string{"name"},
			TotallyUniqueIndexes: []string{"id"},
		},
		"storageManagers": {
			Name:                 "storageManagers",
			Indexes:              []string{"ownerId", "createdAt", "deletedAt", "repairedAt", "zone"},
			TotallyUniqueIndexes: []string{"id"},
		},
		"vms": {
			Name:                 "vms",
			Indexes:              []string{"ownerId", "gpuId", "statusCode", "createdAt", "deletedAt", "repairedAt", "restartedAt", "zone"},
			UniqueIndexes:        []string{"name"},
			TotallyUniqueIndexes: []string{"id"},
		},
		"gpus": {
			Name:                 "gpus",
			Indexes:              []string{"name", "host", "lease.vmId", "lease.user", "lease.end"},
			TotallyUniqueIndexes: []string{"id"},
		},
		"users": {
			Name:                 "users",
			Indexes:              []string{"username", "email", "firstName", "lastName", "effectiveRole.name"},
			TotallyUniqueIndexes: []string{"id"},
			TextIndexFields:      []string{"username", "email", "firstName", "lastName"},
		},
		"teams": {
			Name:                 "teams",
			Indexes:              []string{"name", "ownerId", "createdAt", "deletedAt"},
			TotallyUniqueIndexes: []string{"id"},
			TextIndexFields:      []string{"name"},
		},
		"jobs": {
			Name:                 "jobs",
			Indexes:              []string{"userId", "type", "args.id", "status", "createdAt", "runAfter"},
			TotallyUniqueIndexes: []string{"id"},
		},
		"notifications": {
			Name:                 "notifications",
			Indexes:              []string{"userId", "type", "createdAt", "readAt", "deletedAt"},
			TotallyUniqueIndexes: []string{"id"},
		},
		"events": {
			Name:                 "events",
			Indexes:              []string{"type", "createdAt", "source.userId"},
			TotallyUniqueIndexes: []string{"id"},
		},
	}
}

func isIndexExistsError(err error) bool {
	if mongo.IsDuplicateKeyError(err) {
		return true
	}

	if strings.Contains(err.Error(), "An existing index has the same name as the requested index") {
		return true
	}

	return false
}
