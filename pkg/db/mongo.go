package db

import (
	"context"
	"fmt"
	"go-deploy/pkg/config"
	"go-deploy/pkg/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"time"
)

// CollectionDefinition is a struct that defines a collection.
// It is used as an abstraction layer between MongoDB and the application.
// It contains the collection name, as well as any indexes that should be
// created on the collection.
type CollectionDefinition struct {
	Name    string
	Indexes []string
	// unique only for non-deleted documents
	UniqueIndexes [][]string
	// unique even for deleted documents
	TotallyUniqueIndexes [][]string
	TextIndexFields      []string
}

// setupMongo initializes the MongoDB connection.
// It should be called once at the start of the application.
func (dbCtx *Context) setupMongo() error {
	makeError := func(err error) error {
		return fmt.Errorf("failed to set up mongodb. details: %w", err)
	}

	log.Println("Setting up MongoDB")

	var err error
	dbCtx.MongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(config.Config.MongoDB.URL))
	if err != nil {
		return makeError(err)
	}

	err = dbCtx.MongoClient.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalln(makeError(err))
	}

	// Find collections
	DB.CollectionMap = make(map[string]*mongo.Collection)

	DB.CollectionDefinitionMap = getCollectionDefinitions()

	for _, def := range DB.CollectionDefinitionMap {
		DB.CollectionMap[def.Name] = dbCtx.MongoClient.Database(config.Config.MongoDB.Name).Collection(def.Name)
	}

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

	log.Printf(" - Ensured %d indexes", ensureCount)

	ensureCount = 0
	for _, def := range DB.CollectionDefinitionMap {
		for _, indexName := range def.UniqueIndexes {
			var keys bson.D
			for _, key := range indexName {
				keys = append(keys, bson.E{Key: key, Value: 1})
			}

			_, err = DB.GetCollection(def.Name).Indexes().CreateOne(context.Background(), mongo.IndexModel{
				Keys:    keys,
				Options: options.Index().SetUnique(true).SetPartialFilterExpression(bson.D{{"deletedAt", bson.D{{"$in", []interface{}{nil, time.Time{}}}}}}),
			})
			if err != nil && !isIndexExistsError(err) {
				return makeError(err)
			}

			ensureCount++
		}
	}

	log.Printf(" - Ensured %d unique indexes", ensureCount)

	ensureCount = 0
	for _, def := range DB.CollectionDefinitionMap {
		for _, indexName := range def.TotallyUniqueIndexes {
			keys := bson.D{}
			for _, key := range indexName {
				keys = append(keys, bson.E{Key: key, Value: 1})
			}

			_, err = DB.GetCollection(def.Name).Indexes().CreateOne(context.Background(), mongo.IndexModel{
				Keys:    keys,
				Options: options.Index().SetUnique(true),
			})
			if err != nil && !isIndexExistsError(err) {
				return makeError(err)
			}

			ensureCount++
		}
	}

	log.Printf(" - Ensured %d totally unique indexes", ensureCount)

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

	log.Printf(" - Ensured %d text indexes", ensureCount)

	return nil
}

// shutdownMongo closes the MongoDB connection.
// It should be called once at the end of the application.
func (dbCtx *Context) shutdownMongo() error {
	makeError := func(err error) error {
		return fmt.Errorf("failed to shutdown database. details: %w", err)
	}

	err := dbCtx.MongoClient.Disconnect(context.Background())
	if err != nil {
		return makeError(err)
	}

	dbCtx.CollectionMap = nil

	return nil
}

// getCollectionDefinitions returns a map of all collections and their definitions,
// including any indexes that should be created on the collection.
func getCollectionDefinitions() map[string]CollectionDefinition {
	return map[string]CollectionDefinition{
		"deployments": {
			Name:                 "deployments",
			Indexes:              []string{"ownerId", "type", "statusCode", "createdAt", "deletedAt", "repairedAt", "restartedAt", "zone"},
			UniqueIndexes:        [][]string{{"name"}},
			TotallyUniqueIndexes: [][]string{{"id"}},
		},
		"events": {
			Name:                 "events",
			Indexes:              []string{"type", "createdAt", "source.userId"},
			TotallyUniqueIndexes: [][]string{{"id"}},
		},
		"gpuGroups": {
			Name:                 "gpuGroups",
			Indexes:              []string{},
			UniqueIndexes:        [][]string{{"name", "zone"}},
			TotallyUniqueIndexes: [][]string{{"id"}},
		},
		"gpuLeases": {
			Name:    "gpuLeases",
			Indexes: []string{"groupName", "vmId", "createdAt"},
			// Right now we only allow a single lease per user, this might change in the future
			UniqueIndexes:        [][]string{{"userId"}},
			TotallyUniqueIndexes: [][]string{{"id"}},
		},
		"jobs": {
			Name:                 "jobs",
			Indexes:              []string{"userId", "type", "args.id", "status", "createdAt", "runAfter"},
			TotallyUniqueIndexes: [][]string{{"id"}},
		},
		"notifications": {
			Name:                 "notifications",
			Indexes:              []string{"userId", "type", "createdAt", "readAt", "deletedAt"},
			TotallyUniqueIndexes: [][]string{{"id"}},
		},
		"resourceMigrations": {
			Name:                 "resourceMigrations",
			Indexes:              []string{"resourceId", "type", "resourceType", "userId", "createdAt", "deletedAt"},
			TotallyUniqueIndexes: [][]string{{"id"}},
		},
		"storageManagers": {
			Name:                 "storageManagers",
			Indexes:              []string{"ownerId", "createdAt", "deletedAt", "repairedAt", "zone"},
			TotallyUniqueIndexes: [][]string{{"id"}},
		},
		"teams": {
			Name:                 "teams",
			Indexes:              []string{"name", "ownerId", "createdAt", "deletedAt"},
			TotallyUniqueIndexes: [][]string{{"id"}},
			TextIndexFields:      []string{"name"},
		},
		"users": {
			Name:                 "users",
			Indexes:              []string{"username", "email", "firstName", "lastName", "effectiveRole.name", "lastAuthenticatedAt", "apiKeys.key"},
			TotallyUniqueIndexes: [][]string{{"id"}},
			TextIndexFields:      []string{"username", "email", "firstName", "lastName"},
		},
		"vms": {
			Name:                 "vms",
			Indexes:              []string{"ownerId", "gpuId", "statusCode", "createdAt", "deletedAt", "repairedAt", "restartedAt", "zone"},
			UniqueIndexes:        [][]string{{"name"}},
			TotallyUniqueIndexes: [][]string{{"id"}},
		},
		"vmPorts": {
			Name:          "vmPorts",
			Indexes:       []string{"publicPort", "zone", "lease.privatePort", "lease.userId", "lease.vmId"},
			UniqueIndexes: [][]string{{"publicPort", "zone"}},
		},
		"workerStatus": {
			Name:    "workerStatus",
			Indexes: []string{"name", "status"},
		},
	}
}

// isIndexExistsError returns true if the given error is an "index exists error".
// This is used to ignore errors when creating indexes that already exist.
func isIndexExistsError(err error) bool {
	if mongo.IsDuplicateKeyError(err) {
		return true
	}

	if strings.Contains(err.Error(), "An existing index has the same name as the requested index") {
		return true
	}

	return false
}
