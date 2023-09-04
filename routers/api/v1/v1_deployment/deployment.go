package v1_deployment

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-deploy/models/dto/body"
	"go-deploy/models/dto/query"
	"go-deploy/models/dto/uri"
	deploymentModels "go-deploy/models/sys/deployment"
	jobModel "go-deploy/models/sys/job"
	vmModel "go-deploy/models/sys/vm"
	zoneModel "go-deploy/models/sys/zone"
	"go-deploy/pkg/status_codes"
	"go-deploy/pkg/sys"
	v1 "go-deploy/routers/api/v1"
	"go-deploy/service"
	"go-deploy/service/deployment_service"
	"go-deploy/service/job_service"
	"go-deploy/service/zone_service"
	"net/http"
)

func getURL(deployment *deploymentModels.Deployment) *string {
	ingress, ok := deployment.Subsystems.K8s.IngressMap["main"]
	if !ok || !ingress.Created() {
		return nil
	}

	if len(ingress.Hosts) > 0 && len(ingress.Hosts[0]) > 0 {
		return &ingress.Hosts[0]
	}
	return nil
}

func getStorageManagerURL(auth *service.AuthInfo) *string {
	storageManager, err := deployment_service.GetStorageManagerByOwnerID(auth.UserID, auth)
	if err != nil {
		return nil
	}

	if storageManager == nil {
		return nil
	}

	ingress, ok := storageManager.Subsystems.K8s.IngressMap["oauth-proxy"]
	if !ok || !ingress.Created() {
		return nil
	}

	if len(ingress.Hosts) > 0 && len(ingress.Hosts[0]) > 0 {
		return &ingress.Hosts[0]
	}

	return nil
}

func getAll(context *sys.ClientContext, auth *service.AuthInfo) {
	deployments, _ := deployment_service.GetAllAuth(auth)

	dtoDeployments := make([]body.DeploymentRead, len(deployments))
	for i, deployment := range deployments {
		var storageManagerURL *string
		if mainApp := deployment.GetMainApp(); mainApp != nil && len(mainApp.Volumes) > 0 {
			storageManagerURL = getStorageManagerURL(auth)
		}

		dtoDeployments[i] = deployment.ToDTO(getURL(&deployment), storageManagerURL)
	}

	context.JSONResponse(http.StatusOK, dtoDeployments)
}

// GetList
// @Summary Get list of deployments
// @Description Get list of deployments
// @BasePath /api/v1
// @Tags Deployment
// @Accept  json
// @Produce  json
// @Param Authorization header string true "With the bearer started"
// @Param wantAll query bool false "Get all deployments"
// @Success 200 {array} body.DeploymentRead
// @Failure 400 {object} sys.ErrorResponse
// @Failure 500 {object} sys.ErrorResponse
// @Router /api/v1/deployments [get]
func GetList(c *gin.Context) {
	context := sys.NewContext(c)

	var requestQuery query.DeploymentList
	if err := context.GinContext.Bind(&requestQuery); err != nil {
		context.JSONResponse(http.StatusBadRequest, v1.CreateBindingError(err))
		return
	}

	auth, err := v1.WithAuth(&context)
	if err != nil {
		context.ErrorResponse(http.StatusInternalServerError, status_codes.Error, fmt.Sprintf("Failed to get auth info: %s", err.Error()))
		return
	}

	if requestQuery.WantAll {
		getAll(&context, auth)
		return
	}

	deployments, _ := deployment_service.GetByOwnerIdAuth(auth.UserID, auth)
	if deployments == nil {
		context.JSONResponse(200, []interface{}{})
		return
	}

	dtoDeployments := make([]body.DeploymentRead, len(deployments))
	for i, deployment := range deployments {
		var storageManagerURL *string
		if mainApp := deployment.GetMainApp(); mainApp != nil && len(mainApp.Volumes) > 0 {
			storageManagerURL = getStorageManagerURL(auth)
		}

		dtoDeployments[i] = deployment.ToDTO(getURL(&deployment), storageManagerURL)
	}

	context.JSONResponse(200, dtoDeployments)
}

// Get
// @Summary Get deployment by id
// @Description Get deployment by id
// @BasePath /api/v1
// @Tags Deployment
// @Accept  json
// @Produce  json
// @Param Authorization header string true "With the bearer started"
// @Param deploymentId path string true "Deployment ID"
// @Success 200 {object} body.DeploymentRead
// @Failure 400 {object} sys.ErrorResponse
// @Failure 500 {object} sys.ErrorResponse
// @Router /api/v1/deployments/{deployment_id} [get]
func Get(c *gin.Context) {
	context := sys.NewContext(c)

	var requestURI uri.DeploymentGet
	if err := context.GinContext.BindUri(&requestURI); err != nil {
		context.JSONResponse(http.StatusBadRequest, v1.CreateBindingError(err))
		return
	}

	auth, err := v1.WithAuth(&context)
	if err != nil {
		context.ErrorResponse(http.StatusInternalServerError, status_codes.Error, fmt.Sprintf("Failed to get auth info: %s", err))
		return
	}

	deployment, err := deployment_service.GetByIdAuth(requestURI.DeploymentID, auth)
	if err != nil {
		context.ErrorResponse(http.StatusInternalServerError, status_codes.Error, fmt.Sprintf("%s", err))
		return
	}

	if deployment == nil {
		context.NotFound()
		return
	}

	var storageManagerURL *string
	if len(deployment.GetMainApp().Volumes) > 0 {
		storageManagerURL = getStorageManagerURL(auth)
	}

	context.JSONResponse(200, deployment.ToDTO(getURL(deployment), storageManagerURL))
}

// Create
// @Summary Create deployment
// @Description Create deployment
// @BasePath /api/v1
// @Tags Deployment
// @Accept  json
// @Produce  json
// @Param Authorization header string true "With the bearer started"
// @Param body body body.DeploymentCreate true "Deployment body"
// @Success 200 {object} body.DeploymentRead
// @Failure 400 {object} sys.ErrorResponse
// @Failure 500 {object} sys.ErrorResponse
// @Router /api/v1/deployments [post]
func Create(c *gin.Context) {
	context := sys.NewContext(c)

	var requestBody body.DeploymentCreate
	if err := context.GinContext.BindJSON(&requestBody); err != nil {
		context.JSONResponse(http.StatusBadRequest, v1.CreateBindingError(err))
		return
	}

	auth, err := v1.WithAuth(&context)
	if err != nil {
		context.ErrorResponse(http.StatusInternalServerError, status_codes.Error, fmt.Sprintf("Failed to get auth info: %s", err))
		return
	}

	effectiveRole := auth.GetEffectiveRole()
	if effectiveRole == nil {
		context.ErrorResponse(http.StatusInternalServerError, status_codes.Error, fmt.Sprintf("Failed to get effective role"))
		return
	}

	if effectiveRole.Quotas.Deployments <= 0 {
		context.ErrorResponse(http.StatusForbidden, status_codes.Error, "User is not allowed to create deployments")
		return
	}

	if requestBody.Zone != nil {
		zone := zone_service.GetZone(*requestBody.Zone, zoneModel.ZoneTypeDeployment)
		if zone == nil {
			context.ErrorResponse(http.StatusNotFound, status_codes.ResourceNotFound, "Zone not found")
			return
		}
	}

	deployment, err := deployment_service.GetByName(requestBody.Name)
	if err != nil {
		context.ErrorResponse(http.StatusInternalServerError, status_codes.ResourceValidationFailed, "Failed to validate")
		return
	}

	if requestBody.GitHub != nil {
		validGhToken, reason, err := deployment_service.ValidGitHubToken(requestBody.GitHub.Token)
		if err != nil {
			context.ErrorResponse(http.StatusInternalServerError, status_codes.ResourceValidationFailed, "Failed to validate GitHub token")
			return
		}

		if !validGhToken {
			context.ErrorResponse(http.StatusBadRequest, status_codes.ResourceValidationFailed, reason)
			return
		}

		validGitHubRepository, reason, err := deployment_service.ValidGitHubRepository(requestBody.GitHub.Token, requestBody.GitHub.RepositoryID)
		if err != nil {
			context.ErrorResponse(http.StatusInternalServerError, status_codes.ResourceValidationFailed, "Failed to validate GitHub repository")
			return
		}

		if !validGitHubRepository {
			context.ErrorResponse(http.StatusBadRequest, status_codes.ResourceValidationFailed, reason)
			return
		}
	}

	if deployment != nil {
		if deployment.OwnerID != auth.UserID {
			context.ErrorResponse(http.StatusBadRequest, status_codes.ResourceNotCreated, "Resource already exists")
			return
		}

		started, reason, err := deployment_service.StartActivity(deployment.ID, vmModel.ActivityBeingCreated)
		if err != nil {
			context.ErrorResponse(http.StatusInternalServerError, status_codes.Error, fmt.Sprintf("Failed to start activity: %s", err))
			return
		}

		if !started {
			context.ErrorResponse(http.StatusLocked, status_codes.ResourceNotReady, reason)
			return
		}

		jobID := uuid.New().String()
		err = job_service.Create(jobID, auth.UserID, jobModel.TypeCreateDeployment, map[string]interface{}{
			"id":      deployment.ID,
			"ownerId": auth.UserID,
			"params":  requestBody,
		})
		if err != nil {
			context.ErrorResponse(http.StatusInternalServerError, status_codes.Error, fmt.Sprintf("Failed to create job: %s", err))
			return
		}

		context.JSONResponse(http.StatusCreated, body.DeploymentCreated{
			ID:    deployment.ID,
			JobID: jobID,
		})
		return
	}

	deploymentCount, err := deployment_service.GetCountAuth(auth.UserID, auth)
	if err != nil {
		context.ErrorResponse(http.StatusInternalServerError, status_codes.Error, fmt.Sprintf("%s", err))
		return
	}

	if deploymentCount >= effectiveRole.Quotas.Deployments {
		context.ErrorResponse(http.StatusForbidden, status_codes.Error, fmt.Sprintf("User is not allowed to create more than %d deployments", effectiveRole.Quotas.Deployments))
		return
	}

	deploymentID := uuid.New().String()
	jobID := uuid.New().String()
	err = job_service.Create(jobID, auth.UserID, jobModel.TypeCreateDeployment, map[string]interface{}{
		"id":      deploymentID,
		"ownerId": auth.UserID,
		"params":  requestBody,
	})

	if err != nil {
		context.ErrorResponse(http.StatusInternalServerError, status_codes.Error, fmt.Sprintf("Failed to create job: %s", err))
		return
	}

	context.JSONResponse(http.StatusCreated, body.DeploymentCreated{
		ID:    deploymentID,
		JobID: jobID,
	})
}

// Delete
// @Summary Delete deployment
// @Description Delete deployment
// @BasePath /api/v1
// @Tags Deployment
// @Accept  json
// @Produce  json
// @Param Authorization header string true "With the bearer started"
// @Param deploymentId path string true "Deployment ID"
// @Success 200 {object} body.DeploymentCreated
// @Failure 400 {object} sys.ErrorResponse
// @Failure 401 {object} sys.ErrorResponse
// @Failure 404 {object} sys.ErrorResponse
// @Failure 500 {object} sys.ErrorResponse
// @Router /api/v1/deployments/{deploymentId} [delete]
func Delete(c *gin.Context) {
	context := sys.NewContext(c)

	var requestURI uri.DeploymentDelete
	if err := context.GinContext.BindUri(&requestURI); err != nil {
		context.JSONResponse(http.StatusBadRequest, v1.CreateBindingError(err))
		return
	}

	auth, err := v1.WithAuth(&context)
	if err != nil {
		context.ErrorResponse(http.StatusInternalServerError, status_codes.Error, fmt.Sprintf("Failed to get auth info: %s", err))
		return
	}

	currentDeployment, err := deployment_service.GetByIdAuth(requestURI.DeploymentID, auth)
	if err != nil {
		context.ErrorResponse(http.StatusInternalServerError, status_codes.ResourceValidationFailed, "Failed to validate")
		return
	}

	if currentDeployment == nil {
		context.ErrorResponse(http.StatusNotFound, status_codes.ResourceNotFound, "Resource not found")
		return
	}

	started, reason, err := deployment_service.StartActivity(currentDeployment.ID, deploymentModels.ActivityBeingDeleted)
	if err != nil {
		context.ErrorResponse(http.StatusInternalServerError, status_codes.Error, fmt.Sprintf("Failed to start activity: %s", err))
		return
	}

	if !started {
		context.ErrorResponse(http.StatusNotModified, status_codes.ResourceNotUpdated, fmt.Sprintf("Could not delete resource: %s", reason))
		return
	}

	jobID := uuid.New().String()
	err = job_service.Create(jobID, auth.UserID, jobModel.TypeDeleteDeployment, map[string]interface{}{
		"name": currentDeployment.Name,
	})
	if err != nil {
		context.ErrorResponse(http.StatusInternalServerError, status_codes.Error, fmt.Sprintf("Failed to create job: %s", err))
		return
	}

	context.JSONResponse(http.StatusOK, body.DeploymentDeleted{
		ID:    currentDeployment.ID,
		JobID: jobID,
	})
}

// Update
// @Summary Update deployment
// @Description Update deployment
// @BasePath /api/v1
// @Tags Deployment
// @Accept  json
// @Produce  json
// @Param Authorization header string true "With the bearer started"
// @Param deploymentId path string true "Deployment ID"
// @Param body body body.DeploymentUpdate true "Deployment update"
// @Success 200 {object} body.DeploymentUpdated
// @Failure 400 {object} sys.ErrorResponse
// @Failure 500 {object} sys.ErrorResponse
// @Router /api/v1/deployments/{deploymentId} [put]
func Update(c *gin.Context) {
	context := sys.NewContext(c)

	var requestURI uri.DeploymentUpdate
	if err := context.GinContext.BindUri(&requestURI); err != nil {
		context.JSONResponse(http.StatusBadRequest, v1.CreateBindingError(err))
		return
	}

	var requestBody body.DeploymentUpdate
	if err := context.GinContext.BindJSON(&requestBody); err != nil {
		context.JSONResponse(http.StatusBadRequest, v1.CreateBindingError(err))
		return
	}

	auth, err := v1.WithAuth(&context)
	if err != nil {
		context.ErrorResponse(http.StatusInternalServerError, status_codes.Error, fmt.Sprintf("Failed to get auth info: %s", err))
		return
	}

	deployment, err := deployment_service.GetByIdAuth(requestURI.DeploymentID, auth)
	if err != nil {
		context.ErrorResponse(http.StatusInternalServerError, status_codes.ResourceValidationFailed, fmt.Sprintf("Failed to get vm: %s", err))
		return
	}

	if deployment == nil {
		context.ErrorResponse(http.StatusNotFound, status_codes.ResourceNotFound, fmt.Sprintf("Deployment with id %s not found", requestURI.DeploymentID))
		return
	}

	if deployment.BeingCreated() {
		context.ErrorResponse(http.StatusLocked, status_codes.ResourceBeingCreated, "Resource is currently being created")
		return
	}

	if deployment.BeingDeleted() {
		context.ErrorResponse(http.StatusLocked, status_codes.ResourceBeingDeleted, "Resource is currently being deleted")
		return
	}

	jobID := uuid.New().String()
	err = job_service.Create(jobID, auth.UserID, jobModel.TypeUpdateDeployment, map[string]interface{}{
		"id":     deployment.ID,
		"update": requestBody,
	})

	if err != nil {
		context.ErrorResponse(http.StatusInternalServerError, status_codes.Error, fmt.Sprintf("Failed to create job: %s", err))
		return
	}

	context.JSONResponse(http.StatusOK, body.DeploymentUpdated{
		ID:    deployment.ID,
		JobID: jobID,
	})

}
