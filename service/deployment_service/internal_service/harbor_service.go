package internal_service

import (
	"errors"
	"fmt"
	deploymentModel "go-deploy/models/sys/deployment"
	"go-deploy/pkg/conf"
	"go-deploy/pkg/subsystems/harbor"
	harborModels "go-deploy/pkg/subsystems/harbor/models"
	"go-deploy/utils/subsystemutils"
	"log"
)

func createProjectPublic(projectName string) *harborModels.ProjectPublic {
	return &harborModels.ProjectPublic{
		Name:   projectName,
		Public: true,
	}
}

func createRobotPublic(projectID int, projectName, name string) *harborModels.RobotPublic {
	return &harborModels.RobotPublic{
		Name:        name,
		ProjectID:   projectID,
		ProjectName: projectName,
		Description: "Auto created with Go Deploy",
		Disable:     false,
	}
}

func createRepositoryPublic(projectID int, projectName string, name string) *harborModels.RepositoryPublic {
	return &harborModels.RepositoryPublic{
		Name:        name,
		ProjectID:   projectID,
		ProjectName: projectName,
		Seeded:      false,
		Placeholder: &harborModels.PlaceHolder{
			ProjectName:    conf.Env.DockerRegistry.Placeholder.Project,
			RepositoryName: conf.Env.DockerRegistry.Placeholder.Repository,
		},
	}
}

func createWebhookPublic(projectID int, projectName, name string) *harborModels.WebhookPublic {
	webhookTarget := fmt.Sprintf("%s/hooks/deployments/harbor", conf.Env.ExternalUrl)
	return &harborModels.WebhookPublic{
		Name:        name,
		ProjectID:   projectID,
		ProjectName: projectName,
		Target:      webhookTarget,
		Token:       conf.Env.Harbor.WebhookSecret,
	}
}

func CreateHarbor(name, userID string) error {
	log.Println("setting up harbor for", name)

	makeError := func(err error) error {
		return fmt.Errorf("failed to setup harbor for deployment %s. details: %s", name, err)
	}

	client, err := harbor.New(&harbor.ClientConf{
		ApiUrl:   conf.Env.Harbor.Url,
		Username: conf.Env.Harbor.User,
		Password: conf.Env.Harbor.Password,
	})
	if err != nil {
		return makeError(err)
	}

	deployment, err := deploymentModel.GetByName(name)
	if err != nil {
		return makeError(err)
	}

	if deployment == nil {
		return nil
	}

	// Project
	var project *harborModels.ProjectPublic
	if deployment.Subsystems.Harbor.Project.ID == 0 {
		projectName := subsystemutils.GetPrefixedName(userID)
		id, err := client.CreateProject(createProjectPublic(projectName))
		if err != nil {
			return makeError(err)
		}

		project, err = client.ReadProject(id)
		if err != nil {
			return makeError(err)
		}

		if project == nil {
			return makeError(errors.New("failed to read project after creation"))
		}

		err = deploymentModel.UpdateSubsystemByName(name, "harbor", "project", project)
		if err != nil {
			return makeError(err)
		}
	} else {
		project, err = client.ReadProject(deployment.Subsystems.Harbor.Project.ID)
		if err != nil {
			return makeError(err)
		}
	}

	// Robot
	if deployment.Subsystems.Harbor.Robot.ID == 0 {
		created, err := client.CreateRobot(createRobotPublic(project.ID, project.Name, name))
		if err != nil {
			return makeError(err)
		}

		robot, err := client.ReadRobot(created.ID)
		if err != nil {
			return makeError(err)
		}

		if robot == nil {
			return makeError(errors.New("failed to read robot after creation"))
		}

		robot.Secret = created.Secret

		err = deploymentModel.UpdateSubsystemByName(name, "harbor", "robot", robot)
		if err != nil {
			return makeError(err)
		}
	}

	// Repository
	if deployment.Subsystems.Harbor.Repository.ID == 0 {
		_, err := client.CreateRepository(createRepositoryPublic(project.ID, project.Name, name))
		if err != nil {
			return makeError(err)
		}

		repository, err := client.ReadRepository(project.Name, name)
		if err != nil {
			return makeError(err)
		}

		if repository == nil {
			return makeError(errors.New("failed to read repository after creation"))
		}

		err = deploymentModel.UpdateSubsystemByName(name, "harbor", "repository", repository)
		if err != nil {
			return makeError(err)
		}
	}

	// Webhook
	if deployment.Subsystems.Harbor.Webhook.ID == 0 {
		id, err := client.CreateWebhook(createWebhookPublic(project.ID, project.Name, name))
		if err != nil {
			return makeError(err)
		}

		webhook, err := client.ReadWebhook(project.ID, id)
		if err != nil {
			return makeError(err)
		}

		if webhook == nil {
			return makeError(errors.New("failed to read webhook after creation"))
		}

		err = deploymentModel.UpdateSubsystemByName(name, "harbor", "webhook", webhook)
		if err != nil {
			return makeError(err)
		}
	}

	return nil
}

func DeleteHarbor(name string) error {
	log.Println("deleting harbor for", name)

	makeError := func(err error) error {
		return fmt.Errorf("failed to delete harbor for deployment %s. details: %s", name, err)
	}

	client, err := harbor.New(&harbor.ClientConf{
		ApiUrl:   conf.Env.Harbor.Url,
		Username: conf.Env.Harbor.User,
		Password: conf.Env.Harbor.Password,
	})
	if err != nil {
		return makeError(err)
	}

	deployment, err := deploymentModel.GetByName(name)
	if err != nil {
		return makeError(err)
	}

	if deployment == nil {
		return nil
	}

	if deployment.Subsystems.Harbor.Webhook.ID != 0 {
		err = client.DeleteWebhook(deployment.Subsystems.Harbor.Webhook.ProjectID, deployment.Subsystems.Harbor.Webhook.ID)
		if err != nil {
			return makeError(err)
		}

		err = deploymentModel.UpdateSubsystemByName(name, "harbor", "webhook", harborModels.WebhookPublic{})
		if err != nil {
			return makeError(err)
		}
	}

	if deployment.Subsystems.Harbor.Repository.ID != 0 {
		err = client.DeleteRepository(deployment.Subsystems.Harbor.Repository.ProjectName, deployment.Subsystems.Harbor.Repository.Name)
		if err != nil {
			return makeError(err)
		}

		err = deploymentModel.UpdateSubsystemByName(name, "harbor", "repository", harborModels.RepositoryPublic{})
		if err != nil {
			return makeError(err)
		}
	}

	if deployment.Subsystems.Harbor.Robot.ID != 0 {
		err = client.DeleteRobot(deployment.Subsystems.Harbor.Robot.ID)
		if err != nil {
			return makeError(err)
		}

		err = deploymentModel.UpdateSubsystemByName(name, "harbor", "robot", harborModels.RobotPublic{})
		if err != nil {
			return makeError(err)
		}
	}

	if deployment.Subsystems.Harbor.Project.ID != 0 {
		empty, err := client.IsProjectEmpty(deployment.Subsystems.Harbor.Project.ID)
		if err != nil {
			return makeError(err)
		}

		if empty {
			err = client.DeleteProject(deployment.Subsystems.Harbor.Project.ID)
			if err != nil {
				return makeError(err)
			}
		}

		err = deploymentModel.UpdateSubsystemByName(name, "harbor", "project", harborModels.ProjectPublic{})
		if err != nil {
			return makeError(err)
		}
	}

	return nil
}
