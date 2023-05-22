package confirmers

import (
	"go-deploy/models/sys/deployment"
	"go-deploy/models/sys/vm"
)

func getDeploymentCreatedConfirmers() []func(*deployment.Deployment) (bool, error) {
	return []func(*deployment.Deployment) (bool, error){
		K8sCreated,
		NPMCreated,
		HarborCreated,
		GitHubCreated,
	}
}

func getDeploymentDeletedConfirmers() []func(*deployment.Deployment) (bool, error) {
	return []func(*deployment.Deployment) (bool, error){
		K8sDeleted,
		NPMDeleted,
		HarborDeleted,
		GitHubDeleted,
	}
}

func getVmCreatedConfirmers() []func(*vm.VM) (bool, error) {
	return []func(*vm.VM) (bool, error){
		CSCreated,
	}
}

func getVmDeletedConfirmers() []func(*vm.VM) (bool, error) {
	return []func(*vm.VM) (bool, error){
		CSDeleted,
	}
}

func DeploymentCreated(deployment *deployment.Deployment) bool {
	confirmers := getDeploymentCreatedConfirmers()
	for _, confirmer := range confirmers {
		created, _ := confirmer(deployment)
		if !created {
			return false
		}
	}
	return true
}

func DeploymentDeleted(deployment *deployment.Deployment) bool {
	confirmers := getDeploymentDeletedConfirmers()
	for _, confirmer := range confirmers {
		deleted, _ := confirmer(deployment)
		if !deleted {
			return false
		}
	}
	return true
}

func VmCreated(vm *vm.VM) bool {
	confirmers := getVmCreatedConfirmers()
	for _, confirmer := range confirmers {
		created, _ := confirmer(vm)
		if !created {
			return false
		}
	}
	return true
}

func VmDeleted(vm *vm.VM) bool {
	confirmers := getVmDeletedConfirmers()
	for _, confirmer := range confirmers {
		deleted, _ := confirmer(vm)
		if !deleted {
			return false
		}
	}
	return true
}
