package resources

import (
	deploymentModels "go-deploy/models/sys/deployment"
	"go-deploy/models/sys/deployment/storage_manager"
	"go-deploy/models/sys/enviroment"
	vmModels "go-deploy/models/sys/vm"
	"go-deploy/pkg/subsystems/k8s"
)

type Deployment struct {
	deployment   *deploymentModels.Deployment
	createParams *deploymentModels.CreateParams
	updateParams *deploymentModels.UpdateParams
	zone         *enviroment.DeploymentZone
}

type StorageManager struct {
	storageManager *storage_manager.StorageManager
	createParams   *storage_manager.CreateParams
	zone           *enviroment.DeploymentZone
}

type VM struct {
	vm             *vmModels.VM
	createParams   *vmModels.CreateParams
	updateParams   *vmModels.UpdateParams
	vmZone         *enviroment.VmZone
	deploymentZone *enviroment.DeploymentZone
}

type PublicGeneratorType struct {
	d Deployment
	v VM
	s StorageManager
}

func PublicGenerator() *PublicGeneratorType {
	return &PublicGeneratorType{}
}

func (pc *PublicGeneratorType) WithDeploymentZone(zone *enviroment.DeploymentZone) *PublicGeneratorType {
	pc.d.zone = zone
	pc.s.zone = zone
	pc.v.deploymentZone = zone
	return pc
}

func (pc *PublicGeneratorType) WithVmZone(zone *enviroment.VmZone) *PublicGeneratorType {
	pc.v.vmZone = zone
	return pc
}

func (pc *PublicGeneratorType) WithDeployment(deployment *deploymentModels.Deployment) *PublicGeneratorType {
	pc.d.deployment = deployment
	return pc
}

func (pc *PublicGeneratorType) WithStorageManager(storageManager *storage_manager.StorageManager) *PublicGeneratorType {
	pc.s.storageManager = storageManager
	return pc
}

func (pc *PublicGeneratorType) WithVM(vm *vmModels.VM) *PublicGeneratorType {
	pc.v.vm = vm
	return pc
}

func (pc *PublicGeneratorType) WithDeploymentCreateParams(params *deploymentModels.CreateParams) *PublicGeneratorType {
	pc.d.createParams = params
	return pc
}

func (pc *PublicGeneratorType) WithDeploymentUpdateParams(params *deploymentModels.UpdateParams) *PublicGeneratorType {
	pc.d.updateParams = params
	return pc
}

func (pc *PublicGeneratorType) WithStorageManagerCreateParams(params *storage_manager.CreateParams) *PublicGeneratorType {
	pc.s.createParams = params
	return pc
}

func (pc *PublicGeneratorType) WithVmCreateParams(params *vmModels.CreateParams) *PublicGeneratorType {
	pc.v.createParams = params
	return pc
}

func (pc *PublicGeneratorType) WithVmUpdateParams(params *vmModels.UpdateParams) *PublicGeneratorType {
	pc.v.updateParams = params
	return pc
}

func (pc *PublicGeneratorType) K8s(client *k8s.Client) *K8sGenerator {
	return &K8sGenerator{
		PublicGeneratorType: pc,
		namespace:           client.Namespace,
		client:              client,
	}
}

func (pc *PublicGeneratorType) Harbor(project string) *HarborGenerator {
	return &HarborGenerator{
		PublicGeneratorType: pc,
		project:             project,
	}
}

func (pc *PublicGeneratorType) CS() *CsGenerator {
	return &CsGenerator{
		PublicGeneratorType: pc,
	}
}