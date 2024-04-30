package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"go-deploy/dto/v1/body"
	"go-deploy/dto/v1/query"
	"go-deploy/dto/v1/uri"
	"go-deploy/models/model"
	"go-deploy/pkg/sys"
	"go-deploy/service"
	sErrors "go-deploy/service/errors"
	deploymentOpts "go-deploy/service/v1/deployments/opts"
	"go-deploy/service/v1/teams/opts"
	v12 "go-deploy/service/v1/utils"
	vmOptsV1 "go-deploy/service/v1/vms/opts"
	vmOptsV2 "go-deploy/service/v2/vms/opts"
	"go-deploy/utils"
	"net/http"
	"time"
)

// GetTeam
// @Summary Get team
// @Description Get team
// @Tags Team
// @Accept json
// @Produce json
// @Param teamId path string true "Team ID"
// @Success 200 {object} body.TeamRead
// @Failure 400 {object} body.BindingError
// @Failure 500 {object} sys.ErrorResponse
// @Router /v1/teams/{teamId} [get]
func GetTeam(c *gin.Context) {
	context := sys.NewContext(c)

	var requestURI uri.TeamGet
	if err := context.GinContext.ShouldBindUri(&requestURI); err != nil {
		context.BindingError(CreateBindingError(err))
		return
	}

	auth, err := WithAuth(&context)
	if err != nil {
		context.ServerError(err, AuthInfoNotAvailableErr)
		return
	}

	team, err := service.V1(auth).Teams().Get(requestURI.TeamID)
	if err != nil {
		context.ServerError(err, InternalError)
		return
	}

	context.Ok(team.ToDTO(getMember, getResourceName))
}

// ListTeams
// @Summary List teams
// @Description List teams
// @Tags Team
// @Accept json
// @Produce json
// @Param all query bool false "All teams"
// @Param userId query string false "User ID"
// @Param page query int false "Page"
// @Param pageSize query int false "Page Size"
// @Success 200 {array} body.TeamRead
// @Failure 400 {object} body.BindingError
// @Failure 500 {object} sys.ErrorResponse
// @Router /v1/teams [get]
func ListTeams(c *gin.Context) {
	context := sys.NewContext(c)

	var requestQuery query.TeamList
	if err := context.GinContext.ShouldBind(&requestQuery); err != nil {
		context.BindingError(CreateBindingError(err))
		return
	}

	auth, err := WithAuth(&context)
	if err != nil {
		context.ServerError(err, AuthInfoNotAvailableErr)
		return
	}

	var userID string
	if requestQuery.UserID != nil {
		userID = *requestQuery.UserID
	} else if !requestQuery.All {
		userID = auth.UserID
	}

	teamList, err := service.V1(auth).Teams().List(opts.ListOpts{
		Pagination: v12.GetOrDefaultPagination(requestQuery.Pagination),
		UserID:     userID,
	})
	if err != nil {
		context.ServerError(err, InternalError)
		return
	}

	teamListDTO := make([]body.TeamRead, len(teamList))
	for i, team := range teamList {
		teamListDTO[i] = team.ToDTO(getMember, getResourceName)
	}

	context.Ok(teamListDTO)
}

// CreateTeam
// @Summary Create team
// @Description Create team
// @Tags Team
// @Accept json
// @Produce json
// @Param body body body.TeamCreate true "Team"
// @Success 200 {object} body.TeamRead
// @Failure 400 {object} body.BindingError
// @Failure 500 {object} sys.ErrorResponse
// @Router /v1/teams [post]
func CreateTeam(c *gin.Context) {
	context := sys.NewContext(c)

	var requestQuery body.TeamCreate
	if err := context.GinContext.ShouldBindJSON(&requestQuery); err != nil {
		context.BindingError(CreateBindingError(err))
		return
	}

	auth, err := WithAuth(&context)
	if err != nil {
		context.ServerError(err, AuthInfoNotAvailableErr)
		return
	}

	team, err := service.V1(auth).Teams().Create(uuid.NewString(), auth.UserID, &requestQuery)
	if err != nil {
		if errors.Is(err, sErrors.TeamNameTakenErr) {
			context.UserError("Team name is taken")
			return
		}

		context.ServerError(err, InternalError)
		return
	}

	context.JSONResponse(http.StatusCreated, team.ToDTO(getMember, getResourceName))
}

// UpdateTeam godoc
// @Summary Update team
// @Description Update team
// @Tags Team
// @Accept json
// @Produce json
// @Param teamId path string true "Team ID"
// @Param body body body.TeamUpdate true "Team"
// @Success 200 {object} body.TeamRead
// @Failure 400 {object} body.BindingError
// @Failure 500 {object} sys.ErrorResponse
// @Router /v1/teams/{teamId} [post]
func UpdateTeam(c *gin.Context) {
	context := sys.NewContext(c)

	var requestURI uri.TeamUpdate
	if err := context.GinContext.ShouldBindUri(&requestURI); err != nil {
		context.BindingError(CreateBindingError(err))
		return
	}

	var requestQueryJoin body.TeamJoin
	if err := context.GinContext.ShouldBindBodyWith(&requestQueryJoin, binding.JSON); err == nil {
		joinTeam(context, requestURI.TeamID, &requestQueryJoin)
		return
	}

	var requestQuery body.TeamUpdate
	if err := context.GinContext.ShouldBindBodyWith(&requestQuery, binding.JSON); err != nil {
		context.BindingError(CreateBindingError(err))
		return
	}

	auth, err := WithAuth(&context)
	if err != nil {
		context.ServerError(err, AuthInfoNotAvailableErr)
		return
	}

	updated, err := service.V1(auth).Teams().Update(requestURI.TeamID, &requestQuery)
	if err != nil {
		if errors.Is(err, sErrors.TeamNameTakenErr) {
			context.UserError("Team name is taken")
			return
		}

		context.ServerError(err, InternalError)
		return
	}

	if updated == nil {
		context.NotFound("Team not found")
		return
	}

	context.JSONResponse(http.StatusOK, updated.ToDTO(getMember, getResourceName))
}

// DeleteTeam
// @Summary Delete team
// @Description Delete team
// @Tags Team
// @Accept json
// @Produce json
// @Param teamId path string true "Team ID"
// @Success 204 "No Content"
// @Failure 400 {object} body.BindingError
// @Failure 500 {object} sys.ErrorResponse
// @Router /v1/teams/{teamId} [delete]
func DeleteTeam(c *gin.Context) {
	context := sys.NewContext(c)

	var requestURI uri.TeamUpdate
	if err := context.GinContext.ShouldBindUri(&requestURI); err != nil {
		context.BindingError(CreateBindingError(err))
		return
	}

	auth, err := WithAuth(&context)
	if err != nil {
		context.ServerError(err, AuthInfoNotAvailableErr)
		return
	}

	err = service.V1(auth).Teams().Delete(requestURI.TeamID)
	if err != nil {
		if errors.Is(err, sErrors.TeamNotFoundErr) {
			context.NotFound("Team not found")
			return
		}

		context.ServerError(err, InternalError)
		return
	}

	context.OkNoContent()
}

// joinTeam is an alternate entrypoint for UpdateTeam that allows a user to join a team
// It is called if a body.TeamJoin is passed in the request body, instead of a body.TeamUpdate
func joinTeam(context sys.ClientContext, id string, requestBody *body.TeamJoin) {
	auth, err := WithAuth(&context)
	if err != nil {
		context.ServerError(err, AuthInfoNotAvailableErr)
	}

	team, err := service.V1(auth).Teams().Join(id, requestBody)
	if err != nil {
		if errors.Is(err, sErrors.NotInvitedErr) {
			context.UserError("User not invited to team")
			return
		}

		if errors.Is(err, sErrors.BadInviteCodeErr) {
			context.Forbidden("Bad invite code")
			return
		}

		context.ServerError(err, InternalError)
		return
	}

	if team == nil {
		context.NotFound("Team not found")
		return
	}

	context.JSONResponse(http.StatusCreated, team.ToDTO(getMember, getResourceName))
}

// getMember is a helper function for converting a team member to a team member DTO
func getMember(member *model.TeamMember) *body.TeamMember {
	user, err := service.V1().Users().Get(member.ID)
	if err != nil {
		utils.PrettyPrintError(fmt.Errorf("failed to get user when getting team member for team: %s", err))
		return nil
	}

	if user == nil {
		return nil
	}

	var joinedAt *time.Time
	if !member.JoinedAt.IsZero() {
		joinedAt = &member.JoinedAt
	}

	var addedAt *time.Time
	if !member.AddedAt.IsZero() {
		addedAt = &member.AddedAt
	}

	return &body.TeamMember{
		ID:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		TeamRole:     member.TeamRole,
		JoinedAt:     joinedAt,
		AddedAt:      addedAt,
		MemberStatus: member.MemberStatus,
	}
}

// getResourceName is a helper function for converting a team model to a model name
// It checks the model type and gets the model name from the appropriate service
func getResourceName(resource *model.TeamResource) *string {
	if resource == nil {
		return nil
	}

	deployV1 := service.V1()
	deployV2 := service.V2()

	switch resource.Type {
	case model.TeamResourceDeployment:
		d, err := deployV1.Deployments().Get(resource.ID, deploymentOpts.GetOpts{Shared: true})
		if err != nil {
			utils.PrettyPrintError(fmt.Errorf("failed to get deployment when getting team model name: %s", err))
			return nil
		}

		if d == nil {
			return nil
		}

		return &d.Name
	case model.TeamResourceVM:
		vm, err := deployV1.VMs().Get(resource.ID, vmOptsV1.GetOpts{Shared: true})
		if err != nil {
			utils.PrettyPrintError(fmt.Errorf("failed to get vm when getting team model name: %s", err))
			return nil
		}

		if vm != nil {
			return &vm.Name
		}

		vm, err = deployV2.VMs().Get(resource.ID, vmOptsV2.GetOpts{Shared: true})
		if err != nil {
			utils.PrettyPrintError(fmt.Errorf("failed to get vm when getting team model name: %s", err))
			return nil
		}

		if vm != nil {
			return &vm.Name
		}

		return nil
	}

	return nil

}