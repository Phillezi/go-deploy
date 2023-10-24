package k8s_service

import (
	"fmt"
	"go-deploy/models/sys/deployment/storage_manager"
	"go-deploy/pkg/subsystems/k8s"
	"go-deploy/service/deployment_service/base"
	"go-deploy/service/resources"
	"go-deploy/utils/subsystemutils"
)

type StorageManagerContext struct {
	base.StorageManagerContext

	Client    *k8s.Client
	Generator *resources.K8sGenerator
}

func NewStorageManagerContext(storageManagerID string) (*StorageManagerContext, error) {
	makeError := func(err error) error {
		return fmt.Errorf("error creating context in deployment helper client. details: %w", err)
	}

	baseContext, err := base.NewStorageManagerBaseContext(storageManagerID)
	if err != nil {
		return nil, makeError(err)
	}

	k8sClient, err := withClient(baseContext.Zone, getStorageManagerNamespaceName(baseContext.StorageManager.OwnerID))
	if err != nil {
		return nil, makeError(err)
	}

	return &StorageManagerContext{
		StorageManagerContext: *baseContext,
		Client:                k8sClient,
		Generator:             baseContext.Generator.K8s(k8sClient),
	}, nil
}

func getStorageManagerNamespaceName(userID string) string {
	return subsystemutils.GetPrefixedName("system-" + userID)
}

func (c *StorageManagerContext) WithCreateParams(params *storage_manager.CreateParams) *StorageManagerContext {
	c.CreateParams = params
	c.Generator.WithStorageManagerCreateParams(params)
	return c
}