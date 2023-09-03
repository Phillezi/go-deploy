package deployment

import (
	"context"
	"fmt"
	"go-deploy/models"
	"go-deploy/models/sys/deployment/subsystems"
	"go-deploy/pkg/status_codes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func (client *Client) Create(deploymentID, ownerID string, params *CreateParams) (bool, error) {
	appName := "main"
	mainApp := App{
		Name:         appName,
		Private:      params.Private,
		Envs:         params.Envs,
		Volumes:      params.Volumes,
		InitCommands: params.InitCommands,
		ExtraDomains: make([]string, 0),
		PingResult:   0,
	}

	deployment := Deployment{
		ID:           deploymentID,
		Name:         params.Name,
		OwnerID:      ownerID,
		Zone:         params.Zone,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Time{},
		RepairedAt:   time.Time{},
		RestartedAt:  time.Time{},
		Private:      false,
		Envs:         make([]Env, 0),
		Volumes:      make([]Volume, 0),
		InitCommands: make([]string, 0),
		Apps:         map[string]App{appName: mainApp},
		Activities:   []string{ActivityBeingCreated},
		Subsystems: Subsystems{
			GitLab: subsystems.GitLab{
				LastBuild: subsystems.GitLabBuild{
					ID:        0,
					ProjectID: 0,
					Trace:     []string{"created by go-deploy"},
					Status:    "initialized",
					Stage:     "initialization",
					CreatedAt: time.Now(),
				},
			},
		},
		StatusCode:    status_codes.ResourceBeingCreated,
		StatusMessage: status_codes.GetMsg(status_codes.ResourceBeingCreated),
	}

	filter := bson.D{{"name", params.Name}, {"deletedAt", bson.D{{"$in", []interface{}{time.Time{}, nil}}}}}
	result, err := client.Collection.UpdateOne(context.TODO(), filter, bson.D{
		{"$setOnInsert", deployment},
	}, options.Update().SetUpsert(true))
	if err != nil {
		return false, fmt.Errorf("failed to create deployment. details: %w", err)
	}

	if result.UpsertedCount == 0 {
		if result.MatchedCount == 1 {
			fetchedDeployment, err := client.GetByName(params.Name)
			if err != nil {
				return false, err
			}

			if fetchedDeployment == nil {
				log.Println(fmt.Errorf("failed to fetch deployment %s after creation. assuming it was deleted", params.Name))
				return false, nil
			}

			if fetchedDeployment.ID == deploymentID {
				return true, nil
			}
		}

		return false, nil
	}

	return true, nil
}

func (client *Client) GetAllByGitHubWebhookID(id int64) ([]Deployment, error) {
	return models.GetManyResources[Deployment](client.Collection, bson.D{{"subsystems.github.webhookId", id}}, false)
}

func (client *Client) GetByOwnerID(ownerID string) ([]Deployment, error) {
	return models.GetManyResources[Deployment](client.Collection, bson.D{{"ownerId", ownerID}}, false)
}

func (client *Client) DeleteByID(deploymentID string) error {
	_, err := client.Collection.UpdateOne(context.TODO(),
		bson.D{{"id", deploymentID}},
		bson.D{
			{"$set", bson.D{{"deletedAt", time.Now()}}},
			{"$pull", bson.D{{"activities", ActivityBeingDeleted}}},
		},
	)
	if err != nil {
		return fmt.Errorf("failed to delete deployment %s. details: %w", deploymentID, err)
	}

	return nil
}

func (client *Client) CountByOwnerID(ownerID string) (int, error) {
	return models.CountResources(client.Collection, bson.D{{"ownerId", ownerID}}, false)
}

func (client *Client) UpdateWithParamsByID(id string, update *UpdateParams) error {
	deployment, err := client.GetByID(id)
	if err != nil {
		return err
	}

	if deployment == nil {
		log.Println("deployment not found when updating deployment", id, ". assuming it was deleted")
		return nil
	}

	mainApp := deployment.GetMainApp()
	if mainApp == nil {
		log.Println("main app not found when updating deployment", id, ". assuming it was deleted")
		return nil
	}

	if update.Envs != nil {
		mainApp.Envs = *update.Envs
	}

	if update.Private != nil {
		mainApp.Private = *update.Private
	}

	if update.ExtraDomains != nil {
		mainApp.ExtraDomains = *update.ExtraDomains
	}

	deployment.Apps["main"] = *mainApp

	_, err = client.Collection.UpdateOne(context.TODO(),
		bson.D{{"id", id}},
		bson.D{{"$set", bson.D{{"apps", deployment.Apps}}}},
	)
	if err != nil {
		return fmt.Errorf("failed to update deployment %s. details: %w", id, err)
	}

	return nil

}

func (client *Client) UpdateSubsystemByName(name, subsystem string, key string, update interface{}) error {
	subsystemKey := fmt.Sprintf("subsystems.%s.%s", subsystem, key)
	return client.UpdateWithBsonByName(name, bson.D{{subsystemKey, update}})
}

func (client *Client) UpdateSubsystemByID(id, subsystem string, key string, update interface{}) error {
	subsystemKey := fmt.Sprintf("subsystems.%s.%s", subsystem, key)
	return client.UpdateWithBsonByID(id, bson.D{{subsystemKey, update}})
}

func (client *Client) MarkRepaired(id string) error {
	filter := bson.D{{"id", id}}
	update := bson.D{
		{"$set", bson.D{{"repairedAt", time.Now()}}},
		{"$pull", bson.D{{"activities", "repairing"}}},
	}

	_, err := client.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) MarkUpdated(id string) error {
	filter := bson.D{{"id", id}}
	update := bson.D{
		{"$set", bson.D{{"updatedAt", time.Now()}}},
		{"$pull", bson.D{{"activities", "updating"}}},
	}

	_, err := client.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) UpdateGitLabBuild(deploymentID string, build subsystems.GitLabBuild) error {
	filter := bson.D{
		{"id", deploymentID},
		{"subsystems.gitlab.lastBuild.createdAt", bson.M{"$lte": build.CreatedAt}},
	}

	update := bson.D{
		{"$set", bson.D{
			{"subsystems.gitlab.lastBuild", build},
		}},
	}

	_, err := client.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) GetLastGitLabBuild(deploymentID string) (*subsystems.GitLabBuild, error) {
	// fetch only subsystem.gitlab.lastBuild
	projection := bson.D{
		{"subsystems.gitlab.lastBuild", 1},
	}

	var deployment Deployment
	err := client.Collection.FindOne(context.TODO(),
		bson.D{{"id", deploymentID}},
		options.FindOne().SetProjection(projection),
	).Decode(&deployment)
	if err != nil {
		return &subsystems.GitLabBuild{}, err
	}

	return &deployment.Subsystems.GitLab.LastBuild, nil
}

func (client *Client) SavePing(id string, pingResult int) error {
	deployment, err := client.GetByID(id)
	if err != nil {
		return err
	}

	if deployment == nil {
		log.Println("deployment not found when saving ping result", id, ". assuming it was deleted")
		return nil
	}

	app := deployment.GetMainApp()
	if app == nil {
		return fmt.Errorf("failed to find main app for deployment %s", id)
	}

	app.PingResult = pingResult

	deployment.Apps["main"] = *app

	_, err = client.Collection.UpdateOne(context.TODO(),
		bson.D{{"id", id}},
		bson.D{{"$set", bson.D{{"apps.main.pingResult", pingResult}}}},
	)
	if err != nil {
		return fmt.Errorf("failed to update deployment ping result %s. details: %w", id, err)
	}

	return nil
}
