package v1_vm

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-deploy/models/dto/body"
	"go-deploy/models/dto/query"
	"go-deploy/models/dto/uri"
	gpuModel "go-deploy/models/sys/gpu"
	jobModel "go-deploy/models/sys/job"
	vmModel "go-deploy/models/sys/vm"
	zoneModel "go-deploy/models/sys/zone"
	"go-deploy/pkg/sys"
	v1 "go-deploy/routers/api/v1"
	"go-deploy/service"
	sErrors "go-deploy/service/errors"
	"go-deploy/service/job_service"
	"go-deploy/service/user_service"
	"go-deploy/service/vm_service"
	"go-deploy/service/vm_service/client"
	"go-deploy/service/zone_service"
	"go-deploy/utils"
)

// List
// @Summary Get list of VMs
// @Description Get list of VMs
// @Tags VM
// @Accept  json
// @Produce  json
// @Param all query bool false "Get all"
// @Param userId query string false "Filter by user id"
// @Param page query int false "Page number"
// @Param pageSize query int false "Number of items per page"
// @Success 200 {array} body.VmRead
// @Failure 500 {object} sys.ErrorResponse
// @Router /vm [get]
func List(c *gin.Context) {
	context := sys.NewContext(c)

	var requestQuery query.VmList
	if err := context.GinContext.Bind(&requestQuery); err != nil {
		context.BindingError(v1.CreateBindingError(err))
		return
	}

	auth, err := v1.WithAuth(&context)
	if err != nil {
		context.ServerError(err, v1.AuthInfoNotAvailableErr)
		return
	}

	vsc := vm_service.New().WithAuth(auth)

	var userID string
	if requestQuery.UserID != nil {
		userID = *requestQuery.UserID
	} else if !requestQuery.All {
		userID = auth.UserID
	}

	vms, err := vsc.List(&client.ListOptions{
		Pagination: &service.Pagination{
			Page:     requestQuery.Page,
			PageSize: requestQuery.PageSize,
		},
		UserID: userID,
		Shared: true,
	})
	if err != nil {
		context.ServerError(err, v1.InternalError)
		return
	}

	if vms == nil {
		context.Ok([]interface{}{})
		return
	}

	dtoVMs := make([]body.VmRead, len(vms))
	for i, vm := range vms {
		connectionString, _ := vsc.GetConnectionString(vm.ID)

		var gpuRead *body.GpuRead
		if gpu := vm.GetGpu(); gpu != nil {
			gpuDTO := gpu.ToDTO(true)
			gpuRead = &gpuDTO
		}

		mapper, err := vsc.GetExternalPortMapper(vm.ID)
		if err != nil {
			utils.PrettyPrintError(fmt.Errorf("failed to get external port mapper for vm when listing. details: %w", err))
			continue
		}

		dtoVMs[i] = vm.ToDTO(vm.StatusMessage, connectionString, getTeamIDs(vm.ID, auth), gpuRead, mapper)
	}

	context.Ok(dtoVMs)
}

// Get
// @Summary Get VM by id
// @Description Get VM by id
// @Tags VM
// @Accept  json
// @Produce  json
// @Param vmId path string true "VM ID"
// @Success 200 {object} body.VmRead
// @Failure 400 {object} sys.ErrorResponse
// @Failure 404 {object} sys.ErrorResponse
// @Failure 500 {object} sys.ErrorResponse
// @Router /vm/{vmId} [get]
func Get(c *gin.Context) {
	context := sys.NewContext(c)

	var requestURI uri.VmGet
	if err := context.GinContext.ShouldBindUri(&requestURI); err != nil {
		context.BindingError(v1.CreateBindingError(err))
		return
	}

	auth, err := v1.WithAuth(&context)
	if err != nil {
		context.ServerError(err, v1.AuthInfoNotAvailableErr)
		return
	}

	vsc := vm_service.New().WithAuth(auth)

	vm, err := vsc.Get(requestURI.VmID, &client.GetOptions{Shared: true})
	if err != nil {
		context.ServerError(err, v1.InternalError)
		return
	}

	if vm == nil {
		context.NotFound("VM not found")
		return
	}

	connectionString, _ := vsc.GetConnectionString(requestURI.VmID)
	var gpuRead *body.GpuRead
	if gpu := vm.GetGpu(); gpu != nil {
		gpuDTO := gpu.ToDTO(true)
		gpuRead = &gpuDTO
	}

	mapper, err := vsc.GetExternalPortMapper(vm.ID)
	if err != nil {
		utils.PrettyPrintError(fmt.Errorf("failed to get external port mapper for vm %s. details: %w", vm.ID, err))
	}

	context.Ok(vm.ToDTO(vm.StatusMessage, connectionString, getTeamIDs(vm.ID, auth), gpuRead, mapper))
}

// Create
// @Summary Create VM
// @Description Create VM
// @Tags VM
// @Accept  json
// @Produce  json
// @Param body body body.VmCreate true "VM body"
// @Success 200 {object} body.VmCreated
// @Failure 400 {object} sys.ErrorResponse
// @Failure 401 {object} sys.ErrorResponse
// @Failure 404 {object} sys.ErrorResponse
// @Failure 423 {object} sys.ErrorResponse
// @Failure 500 {object} sys.ErrorResponse
// @Router /vm [post]
func Create(c *gin.Context) {
	context := sys.NewContext(c)

	var requestBody body.VmCreate
	if err := context.GinContext.ShouldBindJSON(&requestBody); err != nil {
		context.BindingError(v1.CreateBindingError(err))
		return
	}

	auth, err := v1.WithAuth(&context)
	if err != nil {
		context.ServerError(err, v1.AuthInfoNotAvailableErr)
		return
	}

	vsc := vm_service.New().WithAuth(auth)

	unique, err := vm_service.NameAvailable(requestBody.Name)
	if err != nil {
		context.ServerError(err, v1.InternalError)
		return
	}

	if !unique {
		context.UserError("VM already exists")
		return
	}

	if requestBody.Zone != nil {
		zone := zone_service.GetZone(*requestBody.Zone, zoneModel.ZoneTypeVM)
		if zone == nil {
			context.NotFound("Zone not found")
			return
		}
	}

	err = vsc.CheckQuota("", auth.UserID, &auth.GetEffectiveRole().Quotas, &client.QuotaOptions{
		Create: &requestBody,
	})
	if err != nil {
		var quotaExceedErr sErrors.QuotaExceededError
		if errors.As(err, &quotaExceedErr) {
			context.Forbidden(quotaExceedErr.Error())
			return
		}

		context.ServerError(err, v1.InternalError)
		return
	}

	vmID := uuid.New().String()
	jobID := uuid.New().String()
	err = job_service.Create(jobID, auth.UserID, jobModel.TypeCreateVM, map[string]interface{}{
		"id":      vmID,
		"ownerId": auth.UserID,
		"params":  requestBody,
	})
	if err != nil {
		context.ServerError(err, v1.InternalError)
		return
	}

	context.Ok(body.VmCreated{
		ID:    vmID,
		JobID: jobID,
	})
}

// Delete
// @Summary Delete VM
// @Description Delete VM
// @Tags VM
// @Accept  json
// @Produce  json
// @Param vmId path string true "VM ID"
// @Success 200 {object} body.VmDeleted
// @Failure 400 {object} sys.ErrorResponse
// @Failure 401 {object} sys.ErrorResponse
// @Failure 404 {object} sys.ErrorResponse
// @Failure 423 {object} sys.ErrorResponse
// @Failure 500 {object} sys.ErrorResponse
// @Router /vm/{vmId} [delete]
func Delete(c *gin.Context) {
	context := sys.NewContext(c)

	var requestURI uri.VmDelete
	if err := context.GinContext.ShouldBindUri(&requestURI); err != nil {
		context.BindingError(v1.CreateBindingError(err))
		return
	}

	auth, err := v1.WithAuth(&context)
	if err != nil {
		context.ServerError(err, v1.AuthInfoNotAvailableErr)
		return
	}

	vsc := vm_service.New().WithAuth(auth)

	vm, err := vsc.Get(requestURI.VmID, &client.GetOptions{Shared: true})
	if err != nil {
		context.ServerError(err, v1.InternalError)
		return
	}

	if vm == nil {
		context.NotFound("VM not found")
		return
	}

	if vm.OwnerID != auth.UserID && !auth.IsAdmin {
		context.Forbidden("VMs can only be deleted by their owner")
		return
	}

	err = vsc.StartActivity(vm.ID, vmModel.ActivityBeingDeleted)
	if err != nil {
		var failedToStartActivityErr sErrors.FailedToStartActivityError
		if errors.As(err, &failedToStartActivityErr) {
			context.Locked(failedToStartActivityErr.Error())
			return
		}

		if errors.Is(err, sErrors.VmNotFoundErr) {
			context.NotFound("Deployment not found")
			return
		}

		context.ServerError(err, v1.InternalError)
		return
	}

	jobID := uuid.New().String()
	err = job_service.Create(jobID, auth.UserID, jobModel.TypeDeleteVM, map[string]interface{}{
		"id": vm.ID,
	})
	if err != nil {
		context.ServerError(err, v1.InternalError)
		return
	}

	context.Ok(body.VmDeleted{
		ID:    vm.ID,
		JobID: jobID,
	})
}

// Update
// @Summary Update VM
// @Description Update VM
// @Tags VM
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Param vmId path string true "VM ID"
// @Param body body body.VmUpdate true "VM update"
// @Success 200 {object} body.VmUpdated
// @Failure 400 {object} sys.ErrorResponse
// @Failure 401 {object} sys.ErrorResponse
// @Failure 404 {object} sys.ErrorResponse
// @Failure 423 {object} sys.ErrorResponse
// @Failure 500 {object} sys.ErrorResponse
// @Router /vm/{vmId} [post]
func Update(c *gin.Context) {
	context := sys.NewContext(c)

	var requestURI uri.VmUpdate
	if err := context.GinContext.ShouldBindUri(&requestURI); err != nil {
		context.BindingError(v1.CreateBindingError(err))
		return
	}

	var requestBody body.VmUpdate
	if err := context.GinContext.ShouldBindJSON(&requestBody); err != nil {
		context.BindingError(v1.CreateBindingError(err))
		return
	}

	auth, err := v1.WithAuth(&context)
	if err != nil {
		context.ServerError(err, v1.AuthInfoNotAvailableErr)
		return
	}

	vsc := vm_service.New().WithAuth(auth)

	var vm *vmModel.VM
	if requestBody.TransferCode != nil {
		vm, err = vsc.Get("", &client.GetOptions{TransferCode: *requestBody.TransferCode})
		if err != nil {
			context.ServerError(err, v1.InternalError)
			return
		}

		if requestBody.OwnerID == nil {
			requestBody.OwnerID = &auth.UserID
		}

	} else {
		vm, err = vsc.Get(requestURI.VmID, &client.GetOptions{Shared: true})
		if err != nil {
			context.ServerError(err, v1.InternalError)
			return
		}
	}

	if vm == nil {
		context.NotFound("VM not found")
		return
	}

	if requestBody.OwnerID != nil {
		if *requestBody.OwnerID == "" {
			err = vsc.ClearUpdateOwner(vm.ID)
			if err != nil {
				if errors.Is(err, sErrors.VmNotFoundErr) {
					context.NotFound("VM not found")
					return
				}

				context.ServerError(err, v1.InternalError)
				return
			}

			context.Ok(body.VmUpdated{
				ID: vm.ID,
			})
			return
		}

		if *requestBody.OwnerID == vm.OwnerID {
			context.UserError("Owner already set")
			return
		}

		exists, err := user_service.New().Exists(*requestBody.OwnerID)
		if err != nil {
			context.ServerError(err, v1.InternalError)
			return
		}

		if !exists {
			context.UserError("User not found")
			return
		}

		jobID, err := vsc.UpdateOwnerSetup(vm.ID, &body.VmUpdateOwner{
			NewOwnerID:   *requestBody.OwnerID,
			OldOwnerID:   vm.OwnerID,
			TransferCode: requestBody.TransferCode,
		})
		if err != nil {
			if errors.Is(err, sErrors.VmNotFoundErr) {
				context.NotFound("VM not found")
				return
			}

			if errors.Is(err, sErrors.InvalidTransferCodeErr) {
				context.Forbidden("Bad transfer code")
				return
			}

			context.ServerError(err, v1.InternalError)
			return
		}

		context.Ok(body.VmUpdated{
			ID:    vm.ID,
			JobID: jobID,
		})
		return
	}

	if requestBody.GpuID != nil {
		updateGPU(&context, &requestBody, auth, vm)
		return
	}

	if requestBody.Name != nil {
		available, err := vm_service.NameAvailable(*requestBody.Name)
		if err != nil {
			context.ServerError(err, v1.InternalError)
			return
		}

		if !available {
			context.UserError("Name already taken")
			return
		}
	}

	if requestBody.Ports != nil {
		for _, port := range *requestBody.Ports {
			if port.HttpProxy != nil {
				available, err := vm_service.HttpProxyNameAvailable(requestURI.VmID, port.HttpProxy.Name)
				if err != nil {
					context.ServerError(err, v1.InternalError)
					return
				}

				if !available {
					context.UserError("Http proxy name already taken")
					return
				}
			}
		}
	}

	err = vsc.CheckQuota(auth.UserID, vm.ID, &auth.GetEffectiveRole().Quotas, &client.QuotaOptions{Update: &requestBody})
	if err != nil {
		var quotaExceededErr sErrors.QuotaExceededError
		if errors.As(err, &quotaExceededErr) {
			context.Forbidden(quotaExceededErr.Error())
			return
		}

		context.ServerError(err, v1.InternalError)
		return
	}

	err = vsc.StartActivity(vm.ID, vmModel.ActivityUpdating)
	if err != nil {
		var failedToStartActivityErr sErrors.FailedToStartActivityError
		if errors.As(err, &failedToStartActivityErr) {
			context.Locked(failedToStartActivityErr.Error())
			return
		}

		if errors.Is(err, sErrors.VmNotFoundErr) {
			context.NotFound("Deployment not found")
			return
		}

		context.ServerError(err, v1.InternalError)
		return
	}

	jobID := uuid.New().String()
	err = job_service.Create(jobID, auth.UserID, jobModel.TypeUpdateVM, map[string]interface{}{
		"id":     vm.ID,
		"params": requestBody,
	})

	if err != nil {
		context.ServerError(err, v1.InternalError)
		return
	}

	context.Ok(body.VmUpdated{
		ID:    vm.ID,
		JobID: &jobID,
	})
}

func updateGPU(context *sys.ClientContext, requestBody *body.VmUpdate, auth *service.AuthInfo, vm *vmModel.VM) {
	decodedGpuID, decodeErr := decodeGpuID(*requestBody.GpuID)
	if decodeErr != nil {
		context.UserError("Invalid GPU ID")
		return
	}

	requestBody.GpuID = &decodedGpuID

	if *requestBody.GpuID == "" {
		detachGPU(context, auth, vm)
		return
	} else {
		attachGPU(context, requestBody, auth, vm)
		return
	}
}

func detachGPU(context *sys.ClientContext, auth *service.AuthInfo, vm *vmModel.VM) {
	if !vm.HasGPU() {
		context.UserError("VM does not have a GPU attached")
		return
	}

	vsc := vm_service.New().WithAuth(auth)

	err := vsc.StartActivity(vm.ID, vmModel.ActivityDetachingGPU)
	if err != nil {
		var failedToStartActivityErr sErrors.FailedToStartActivityError
		if errors.As(err, &failedToStartActivityErr) {
			context.Locked(failedToStartActivityErr.Error())
			return
		}

		if errors.Is(err, sErrors.VmNotFoundErr) {
			context.NotFound("Deployment not found")
			return
		}

		context.ServerError(err, v1.InternalError)
		return
	}

	jobID := uuid.New().String()
	err = job_service.Create(jobID, auth.UserID, jobModel.TypeDetachGPU, map[string]interface{}{
		"id": vm.ID,
	})
	if err != nil {
		context.ServerError(err, v1.InternalError)
		return
	}

	context.Ok(body.GpuDetached{
		ID:    vm.ID,
		JobID: jobID,
	})
}

func attachGPU(context *sys.ClientContext, requestBody *body.VmUpdate, auth *service.AuthInfo, vm *vmModel.VM) {
	if !auth.GetEffectiveRole().Permissions.UseGPUs {
		context.Forbidden("Tier does not include GPU access")
		return
	}

	vsc := vm_service.New().WithAuth(auth)
	currentGPU := vm.GetGpu()

	var gpus []gpuModel.GPU
	if *requestBody.GpuID == "any" {
		if currentGPU != nil {
			if !currentGPU.Lease.IsExpired() {
				context.UserError("GPU lease not expired")
				return
			}

			gpus = []gpuModel.GPU{*currentGPU}
		} else {
			availableGpus, err := vsc.ListGPUs(&client.ListGpuOptions{
				AvailableGPUs: true,
			})
			if err != nil {
				context.ServerError(err, v1.InternalError)
				return
			}

			if availableGpus == nil {
				context.ServerUnavailableError(fmt.Errorf("no available gpus when attaching gpu to vm %s", vm.ID), v1.NoAvailableGpuErr)
				return
			}

			gpus = availableGpus
		}
	} else {
		if !auth.GetEffectiveRole().Permissions.ChooseGPU {
			context.Forbidden("Tier does not include GPU selection")
			return
		}

		privilegedGPU, err := vsc.IsGpuPrivileged(*requestBody.GpuID)
		if err != nil {
			context.ServerError(err, v1.InternalError)
			return
		}

		if privilegedGPU && !auth.GetEffectiveRole().Permissions.UsePrivilegedGPUs {
			context.NotFound("GPU not found")
			return
		}

		requestedGPU, err := vsc.GetGPU(*requestBody.GpuID, &client.GetGpuOptions{})
		if err != nil {
			context.ServerError(err, v1.InternalError)
			return
		}

		if requestedGPU == nil {
			context.NotFound("GPU not found")
			return
		}

		if currentGPU != nil && currentGPU.ID != requestedGPU.ID {
			context.UserError("VM already has a GPU attached")
			return
		}

		if !requestedGPU.Lease.IsExpired() {
			context.UserError("GPU lease not expired")
			return
		}

		err = vsc.CheckGpuHardwareAvailable(requestedGPU.ID)
		if err != nil {
			switch {
			case errors.Is(err, sErrors.HostNotAvailableErr):
				context.ServerUnavailableError(fmt.Errorf("host not available when attaching gpu to vm %s. details: %w", vm.ID, err), v1.HostNotAvailableErr)
			default:
				context.ServerError(err, v1.InternalError)
			}
			return
		}

		gpus = []gpuModel.GPU{*requestedGPU}
	}

	if len(gpus) == 0 {
		context.ServerUnavailableError(fmt.Errorf("no available gpus when attaching gpu to vm %s", vm.ID), v1.NoAvailableGpuErr)
		return
	}

	// do this check to give a nice error to the user if the gpu cannot be attached
	// otherwise it will be silently ignored
	if len(gpus) == 1 {
		if err := vsc.CheckSuitableHost(vm.ID, gpus[0].Host, gpus[0].Zone); err != nil {
			switch {
			case errors.Is(err, sErrors.HostNotAvailableErr):
				context.ServerUnavailableError(fmt.Errorf("host not available when attaching gpu to vm %s. details: %w", vm.ID, err), v1.HostNotAvailableErr)
			case errors.Is(err, sErrors.VmTooLargeErr):
				tooLargeErr := v1.VmTooLargeForHostErr
				caps, err := vm_service.GetCloudStackHostCapabilities(gpus[0].Host, vm.Zone)
				if err == nil && caps != nil {
					tooLargeErr = v1.MakeVmToLargeForHostErr(caps.CpuCoresTotal-caps.CpuCoresUsed, caps.RamTotal-caps.RamUsed)
				}
				context.ServerUnavailableError(fmt.Errorf("vm %s too large when attaching gpu", vm.ID), tooLargeErr)
			case errors.Is(err, sErrors.VmNotCreatedErr):
				context.ServerUnavailableError(fmt.Errorf("vm %s not created when attaching gpu to vm %s. details: %w", vm.ID, vm.ID, err), v1.VmNotReadyErr)
			default:
				context.ServerError(err, v1.InternalError)
			}
			return
		}
	}

	gpuIds := make([]string, len(gpus))
	for i, gpu := range gpus {
		gpuIds[i] = gpu.ID
	}

	err := vsc.StartActivity(vm.ID, vmModel.ActivityAttachingGPU)
	if err != nil {
		var failedToStartActivityErr sErrors.FailedToStartActivityError
		if errors.As(err, &failedToStartActivityErr) {
			context.Locked(failedToStartActivityErr.Error())
			return
		}

		if errors.Is(err, sErrors.VmNotFoundErr) {
			context.NotFound("Deployment not found")
			return
		}

		context.ServerError(err, v1.InternalError)
		return
	}

	jobID := uuid.New().String()
	err = job_service.Create(jobID, auth.UserID, jobModel.TypeAttachGPU, map[string]interface{}{
		"id":            vm.ID,
		"gpuIds":        gpuIds,
		"userId":        auth.UserID,
		"leaseDuration": auth.GetEffectiveRole().Quotas.GpuLeaseDuration,
	})

	if err != nil {
		context.ServerError(err, v1.InternalError)
		return
	}

	context.Ok(body.GpuAttached{
		ID:    vm.ID,
		JobID: jobID,
	})
}

func decodeGpuID(gpuID string) (string, error) {
	if gpuID == "any" {
		return gpuID, nil
	}

	res, err := base64.StdEncoding.DecodeString(gpuID)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func getTeamIDs(resourceID string, auth *service.AuthInfo) []string {
	teams, err := user_service.New().ListTeams(&user_service.ListTeamsOpts{ResourceID: resourceID, UserID: auth.UserID})

	if err != nil {
		return []string{}
	}

	teamIDs := make([]string, len(teams))
	for idx, team := range teams {
		teamIDs[idx] = team.ID
	}

	return teamIDs
}
