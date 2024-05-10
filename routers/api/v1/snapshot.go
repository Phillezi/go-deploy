package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-deploy/dto/v1/body"
	"go-deploy/dto/v1/query"
	"go-deploy/dto/v1/uri"
	"go-deploy/models/model"
	"go-deploy/models/version"
	"go-deploy/pkg/sys"
	"go-deploy/service"
	sErrors "go-deploy/service/errors"
	v12 "go-deploy/service/v1/utils"
	"go-deploy/service/v1/vms/opts"
)

// GetSnapshot
// @Summary Get snapshot
// @Description Get snapshot
// @Tags VM
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param vmId path string true "VM ID"
// @Param snapshotId path string true "Snapshot ID"
// @Success 200 {object} body.VmSnapshotRead
// @Failure 400 {object} sys.ErrorResponse
// @Failure 404 {object} sys.ErrorResponse
// @Failure 500 {object} sys.ErrorResponse
// @Router /v1/vms/{vmId}/snapshot/{snapshotId} [post]
func GetSnapshot(c *gin.Context) {
	context := sys.NewContext(c)

	var requestURI uri.VmSnapshotGet
	if err := context.GinContext.ShouldBindUri(&requestURI); err != nil {
		context.BindingError(CreateBindingError(err))
		return
	}

	auth, err := WithAuth(&context)
	if err != nil {
		context.ServerError(err, AuthInfoNotAvailableErr)
		return
	}

	deployV1 := service.V1(auth)

	vm, err := deployV1.VMs().Get(requestURI.VmID, opts.GetOpts{Shared: true})
	if err != nil {
		context.ServerError(err, InternalError)
		return
	}

	if vm == nil {
		context.NotFound("VM not found")
		return
	}

	snapshot, err := deployV1.VMs().GetSnapshot(requestURI.VmID, requestURI.SnapshotID, opts.GetSnapshotOpts{})
	if err != nil {
		context.ServerError(err, InternalError)
		return
	}

	if snapshot == nil {
		context.NotFound("Snapshot not found")
		return
	}

	context.Ok(snapshot.ToDTOv1())
}

// ListSnapshots
// @Summary List snapshots
// @Description List snapshots
// @Tags VM
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param vmId path string true "VM ID"
// @Param page query int false "Page number"
// @Param pageSize query int false "Number of items per page"
// @Success 200 {array} body.VmSnapshotRead
// @Failure 400 {object} sys.ErrorResponse
// @Failure 404 {object} sys.ErrorResponse
// @Failure 423 {object} sys.ErrorResponse
// @Failure 500 {object} sys.ErrorResponse
// @Router /v1/vms/{vmId}/snapshots [get]
func ListSnapshots(c *gin.Context) {
	context := sys.NewContext(c)

	var requestQuery query.VmSnapshotList
	if err := context.GinContext.ShouldBindQuery(&requestQuery); err != nil {
		context.BindingError(CreateBindingError(err))
		return
	}

	var requestURI uri.VmSnapshotList
	if err := context.GinContext.ShouldBindUri(&requestURI); err != nil {
		context.BindingError(CreateBindingError(err))
		return
	}

	snapshots, err := service.V1().VMs().ListSnapshots(requestURI.VmID, opts.ListSnapshotOpts{
		Pagination: v12.GetOrDefaultPagination(requestQuery.Pagination),
	})
	if err != nil {
		context.ServerError(err, InternalError)
		return
	}

	if snapshots == nil {
		context.Ok([]interface{}{})
		return
	}

	dtoSnapshots := make([]body.VmSnapshotRead, len(snapshots))
	for i, snapshot := range snapshots {
		dtoSnapshots[i] = snapshot.ToDTOv1()
	}

	context.Ok(dtoSnapshots)
}

// CreateSnapshot
// @Summary Create snapshot
// @Description Create snapshot
// @Tags VM
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param vmId path string true "VM ID"
// @Success 200 {object} body.VmSnapshotRead
// @Failure 400 {object} sys.ErrorResponse
// @Failure 404 {object} sys.ErrorResponse
// @Failure 500 {object} sys.ErrorResponse
// @Router /v1/vms/{vmId}/snapshots [post]
func CreateSnapshot(c *gin.Context) {
	context := sys.NewContext(c)

	var requestURI uri.VmSnapshotCreate
	if err := context.GinContext.ShouldBindUri(&requestURI); err != nil {
		context.BindingError(CreateBindingError(err))
		return
	}

	var requestBody body.VmSnapshotCreate
	if err := context.GinContext.ShouldBindJSON(&requestBody); err != nil {
		context.BindingError(CreateBindingError(err))
		return
	}

	auth, err := WithAuth(&context)
	if err != nil {
		context.ServerError(err, AuthInfoNotAvailableErr)
		return
	}

	deployV1 := service.V1(auth)

	err = deployV1.VMs().CheckQuota(requestURI.VmID, auth.User.ID, &auth.GetEffectiveRole().Quotas, opts.QuotaOpts{
		CreateSnapshot: &requestBody,
	})
	if err != nil {
		var quotaExceededErr sErrors.QuotaExceededError
		if errors.As(err, &quotaExceededErr) {
			context.Forbidden(quotaExceededErr.Error())
			return
		}

		context.ServerError(err, InternalError)
		return
	}

	current, err := deployV1.VMs().GetSnapshotByName(requestURI.VmID, requestBody.Name, opts.GetSnapshotOpts{})
	if err != nil {
		context.ServerError(err, InternalError)
		return
	}

	if current != nil {
		context.UserError("Snapshot already exists")
		return
	}

	vm, err := deployV1.VMs().Get(requestURI.VmID, opts.GetOpts{Shared: true})
	if err != nil {
		context.ServerError(err, InternalError)
		return
	}

	if vm == nil {
		context.NotFound("VM not found")
		return
	}

	jobID := uuid.New().String()
	err = deployV1.Jobs().Create(jobID, auth.User.ID, model.JobCreateVmUserSnapshot, version.V1, map[string]interface{}{
		"id": vm.ID,
		"params": body.VmSnapshotCreate{
			Name: requestBody.Name,
		},
		"authInfo": auth,
	})

	if err != nil {
		context.ServerError(err, InternalError)
		return
	}

	context.Ok(body.VmSnapshotCreated{
		ID:    vm.ID,
		JobID: jobID,
	})
}

// DeleteSnapshot
// @Summary Delete snapshot
// @Description Delete snapshot
// @Tags VM
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param vmId path string true "VM ID"
// @Param snapshotId path string true "Snapshot ID"
// @Success 200 {object} body.VmSnapshotRead
// @Failure 400 {object} sys.ErrorResponse
// @Failure 404 {object} sys.ErrorResponse
// @Failure 500 {object} sys.ErrorResponse
// @Router /v1/vms/{vmId}/snapshot/{snapshotId} [delete]
func DeleteSnapshot(c *gin.Context) {
	context := sys.NewContext(c)

	var requestURI uri.VmSnapshotDelete
	if err := context.GinContext.ShouldBindUri(&requestURI); err != nil {
		context.BindingError(CreateBindingError(err))
		return
	}

	auth, err := WithAuth(&context)
	if err != nil {
		context.ServerError(err, AuthInfoNotAvailableErr)
		return
	}

	deployV1 := service.V1(auth)

	vm, err := deployV1.VMs().Get(requestURI.VmID, opts.GetOpts{Shared: true})
	if err != nil {
		context.ServerError(err, InternalError)
		return
	}

	if vm == nil {
		context.NotFound("VM not found")
		return
	}

	snapshot, err := deployV1.VMs().GetSnapshot(requestURI.VmID, requestURI.SnapshotID, opts.GetSnapshotOpts{})
	if err != nil {
		context.ServerError(err, InternalError)
		return
	}

	if snapshot == nil {
		context.NotFound("Snapshot not found")
		return
	}

	jobID := uuid.New().String()
	err = deployV1.Jobs().Create(jobID, auth.User.ID, model.JobDeleteVmSnapshot, version.V1, map[string]interface{}{
		"id":         vm.ID,
		"snapshotId": snapshot.ID,
		"authInfo":   auth,
	})

	if err != nil {
		context.ServerError(err, InternalError)
		return
	}

	context.Ok(body.VmSnapshotDeleted{
		ID:    vm.ID,
		JobID: jobID,
	})
}
