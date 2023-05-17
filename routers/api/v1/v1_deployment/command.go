package v1_deployment

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-deploy/models/dto/body"
	"go-deploy/models/dto/uri"
	"go-deploy/pkg/app"
	"go-deploy/pkg/status_codes"
	v1 "go-deploy/routers/api/v1"
	"go-deploy/service/deployment_service"
	"net/http"
)

// DoCommand
// @Summary Do command
// @Description Do command
// @Tags VM
// @Accept  json
// @Produce  json
// @Param vmId path string true "VM ID"
// @Param body body body.DoCommand true "Command body"
// @Success 200 {empty} empty
// @Failure 400 {object} app.ErrorResponse
// @Failure 404 {object} app.ErrorResponse
// @Failure 423 {object} app.ErrorResponse
// @Failure 500 {object} app.ErrorResponse
// @Router /api/v1/vms/{vmId}/command [post]
func DoCommand(c *gin.Context) {
	context := app.NewContext(c)

	var requestURI uri.DeploymentCommand
	if err := context.GinContext.BindUri(&requestURI); err != nil {
		context.JSONResponse(http.StatusBadRequest, v1.CreateBindingError(err))
		return
	}

	var requestBody body.DeploymentCommand
	if err := context.GinContext.BindJSON(&requestBody); err != nil {
		context.JSONResponse(http.StatusBadRequest, v1.CreateBindingError(err))
		return
	}

	auth, err := v1.WithAuth(&context)
	if err != nil {
		context.ErrorResponse(http.StatusInternalServerError, status_codes.Error, fmt.Sprintf("Failed to get auth info: %s", err))
		return
	}

	deployment, err := deployment_service.GetByID(auth.UserID, requestURI.DeploymentID, auth.IsAdmin())
	if err != nil {
		context.ErrorResponse(http.StatusInternalServerError, status_codes.ResourceValidationFailed, "Failed to validate")
		return
	}

	if deployment == nil {
		context.ErrorResponse(http.StatusNotFound, status_codes.ResourceNotFound, fmt.Sprintf("Resource with id %s not found", requestURI.DeploymentID))
		return
	}

	if !deployment.Ready() {
		context.ErrorResponse(http.StatusLocked, status_codes.ResourceNotReady, fmt.Sprintf("Resource %s is not ready", requestURI.DeploymentID))
		return
	}

	deployment_service.DoCommand(deployment, requestBody.Command)

	context.OkDeleted()
}
