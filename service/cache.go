package service

import (
	deploymentModels "go-deploy/models/sys/deployment"
	gpuModels "go-deploy/models/sys/gpu"
	jobModels "go-deploy/models/sys/job"
	smModels "go-deploy/models/sys/sm"
	teamModels "go-deploy/models/sys/team"
	userModels "go-deploy/models/sys/user"
	vmModels "go-deploy/models/sys/vm"
)

type Cache struct {
	deploymentStore map[string]*deploymentModels.Deployment
	vmStore         map[string]*vmModels.VM
	gpuStore        map[string]*gpuModels.GPU
	smStore         map[string]*smModels.SM
	userStore       map[string]*userModels.User
	teamStore       map[string]*teamModels.Team
	jobStore        map[string]*jobModels.Job
}

func NewCache() *Cache {
	return &Cache{
		deploymentStore: make(map[string]*deploymentModels.Deployment),
		vmStore:         make(map[string]*vmModels.VM),
		gpuStore:        make(map[string]*gpuModels.GPU),
		smStore:         make(map[string]*smModels.SM),
		userStore:       make(map[string]*userModels.User),
		teamStore:       make(map[string]*teamModels.Team),
		jobStore:        make(map[string]*jobModels.Job),
	}
}

func (c *Cache) StoreDeployment(deployment *deploymentModels.Deployment) {
	c.deploymentStore[deployment.ID] = deployment
}

func (c *Cache) GetDeployment(id string) *deploymentModels.Deployment {
	r, ok := c.deploymentStore[id]
	if !ok {
		return nil
	}

	return r
}

func (c *Cache) StoreVM(vm *vmModels.VM) {
	c.vmStore[vm.ID] = vm
}

func (c *Cache) GetVM(id string) *vmModels.VM {
	r, ok := c.vmStore[id]
	if !ok {
		return nil
	}

	return r
}

func (c *Cache) StoreGPU(gpu *gpuModels.GPU) {
	c.gpuStore[gpu.ID] = gpu
}

func (c *Cache) GetGPU(id string) *gpuModels.GPU {
	r, ok := c.gpuStore[id]
	if !ok {
		return nil
	}

	return r
}

func (c *Cache) StoreSM(sm *smModels.SM) {
	c.smStore[sm.ID] = sm
	c.smStore[sm.OwnerID] = sm
}

func (c *Cache) GetSM(id string) *smModels.SM {
	r, ok := c.smStore[id]
	if !ok {
		return nil
	}

	return r
}

func (c *Cache) StoreUser(user *userModels.User) {
	c.userStore[user.ID] = user
}

func (c *Cache) GetUser(id string) *userModels.User {
	r, ok := c.userStore[id]
	if !ok {
		return nil
	}

	return r
}

func (c *Cache) StoreTeam(team *teamModels.Team) {
	c.teamStore[team.ID] = team
}

func (c *Cache) GetTeam(id string) *teamModels.Team {
	r, ok := c.teamStore[id]
	if !ok {
		return nil
	}

	return r
}

func (c *Cache) StoreJob(job *jobModels.Job) {
	c.jobStore[job.ID] = job
}

func (c *Cache) GetJob(id string) *jobModels.Job {
	r, ok := c.jobStore[id]
	if !ok {
		return nil
	}

	return r
}
