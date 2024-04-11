package teams

import (
	"github.com/stretchr/testify/assert"
	"go-deploy/dto/v1/body"
	"go-deploy/models/model"
	"go-deploy/test/e2e"
	"go-deploy/test/e2e/v1"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	e2e.Setup()
	code := m.Run()
	e2e.Shutdown()
	os.Exit(code)
}

func TestCreateEmptyTeam(t *testing.T) {
	t.Parallel()

	requestBody := body.TeamCreate{
		Name:        e2e.GenName(),
		Description: e2e.GenName(),
		Resources:   nil,
		Members:     nil,
	}

	_ = v1.WithTeam(t, requestBody)
}

func TestCreateWithMembers(t *testing.T) {
	t.Parallel()

	requestBody := body.TeamCreate{
		Name:        e2e.GenName(),
		Description: e2e.GenName(),
		Resources:   nil,
		Members: []body.TeamMemberCreate{
			{ID: e2e.PowerUserID, TeamRole: model.TeamMemberRoleAdmin},
		},
	}

	// Create team
	_ = v1.WithTeam(t, requestBody)

	// Fetch TestUser2's teams
	teams := v1.ListTeams(t, "?userId="+e2e.PowerUserID)
	found := false
	for _, team := range teams {
		if team.Name == requestBody.Name {
			found = true
			break
		}
	}
	assert.True(t, found, "user has no teams")
}

func TestCreateWithResources(t *testing.T) {
	t.Parallel()

	resource, _ := v1.WithDeployment(t, body.DeploymentCreate{
		Name: e2e.GenName(),
	})

	requestBody := body.TeamCreate{
		Name:        e2e.GenName(),
		Description: e2e.GenName(),
		Resources:   []string{resource.ID},
		Members:     nil,
	}

	// Create Team
	team := v1.WithTeam(t, requestBody)

	// Fetch deployment's teams
	resource = v1.GetDeployment(t, resource.ID)
	assert.EqualValues(t, []string{team.ID}, resource.Teams, "invalid teams on model")

}

func TestCreateFull(t *testing.T) {
	t.Parallel()

	resource, _ := v1.WithDeployment(t, body.DeploymentCreate{
		Name: e2e.GenName(),
	})

	requestBody := body.TeamCreate{
		Name:        e2e.GenName(),
		Description: e2e.GenName(),
		Resources:   []string{resource.ID},
		Members:     []body.TeamMemberCreate{{ID: e2e.PowerUserID, TeamRole: model.TeamMemberRoleAdmin}},
	}

	// Create team
	team := v1.WithTeam(t, requestBody)

	// Fetch TestUser2's teams
	teams := v1.ListTeams(t, "?userId="+e2e.PowerUserID)
	assert.NotEmpty(t, teams, "user has no teams")

	// Fetch deployment's teams
	resource = v1.GetDeployment(t, resource.ID)
	assert.EqualValues(t, []string{team.ID}, resource.Teams, "invalid teams on model")
}

func TestCreateWithInvitation(t *testing.T) {
	t.Parallel()

	team := v1.WithTeam(t, body.TeamCreate{
		Name:        e2e.GenName(),
		Description: e2e.GenName(),
		Resources:   nil,
		Members:     []body.TeamMemberCreate{{ID: e2e.PowerUserID, TeamRole: model.TeamMemberRoleAdmin}},
	}, e2e.DefaultUserID)

	assert.Equal(t, 2, len(team.Members), "invalid number of members")

	found := false
	for _, member := range team.Members {
		if member.ID == e2e.PowerUserID {
			assert.Equal(t, model.TeamMemberRoleAdmin, member.TeamRole, "invalid member role")
			assert.Equal(t, model.TeamMemberStatusInvited, member.MemberStatus, "invalid member status")

			found = true
			break
		}
	}

	if !found {
		assert.Fail(t, "user was not invited")
	}

	notifications := v1.ListNotifications(t, "?userId="+e2e.PowerUserID)
	assert.NotEmpty(t, notifications, "user has no notifications")

	found = false
	for _, notification := range notifications {
		if notification.Type == model.NotificationTeamInvite {
			for key, val := range notification.Content {
				if key == "id" && val == team.ID {
					return
				}
			}
		}
	}

	assert.Fail(t, "user has no team invite notification")
}

func TestJoin(t *testing.T) {
	t.Parallel()

	team := v1.WithTeam(t, body.TeamCreate{
		Name:        e2e.GenName(),
		Description: e2e.GenName(),
		Resources:   nil,
		Members:     []body.TeamMemberCreate{{ID: e2e.PowerUserID, TeamRole: model.TeamMemberRoleAdmin}},
	}, e2e.DefaultUserID)

	notifications := v1.ListNotifications(t, "?userId="+e2e.PowerUserID)

	for _, notification := range notifications {
		if notification.Type == model.NotificationTeamInvite {
			for key, val := range notification.Content {
				if key == "id" && val == team.ID {
					code := notification.Content["code"].(string)
					v1.JoinTeam(t, team.ID, code, e2e.PowerUserID)
					return
				}
			}
		}
	}

	assert.Fail(t, "user has no team invite notification")
}

func TestJoinWithBadCode(t *testing.T) {
	t.Parallel()

	team := v1.WithTeam(t, body.TeamCreate{
		Name:        e2e.GenName(),
		Description: e2e.GenName(),
		Resources:   nil,
		Members:     []body.TeamMemberCreate{{ID: e2e.DefaultUserID, TeamRole: model.TeamMemberRoleAdmin}},
	})

	resp := e2e.DoPostRequest(t, v1.TeamPath+team.ID, body.TeamJoin{InvitationCode: "bad-code"}, e2e.DefaultUserID)
	assert.Equal(t, 400, resp.StatusCode, "bad code was not detected")
}

func TestUpdate(t *testing.T) {
	t.Parallel()

	team := v1.WithTeam(t, body.TeamCreate{
		Name:        e2e.GenName(),
		Description: e2e.GenName(),
		Resources:   nil,
		Members:     nil,
	})

	requestBody := body.TeamUpdate{
		Name:        e2e.StrPtr(e2e.GenName("new-team")),
		Description: e2e.StrPtr(e2e.GenName("new-description")),
		Resources:   nil,
		Members:     nil,
	}

	v1.UpdateTeam(t, team.ID, requestBody)
}

func TestUpdateResources(t *testing.T) {
	t.Parallel()

	team := v1.WithTeam(t, body.TeamCreate{
		Name:        e2e.GenName(),
		Description: e2e.GenName(),
		Resources:   nil,
		Members:     nil,
	})

	resource, _ := v1.WithDeployment(t, body.DeploymentCreate{
		Name: e2e.GenName("deployment"),
	})

	requestBody := body.TeamUpdate{
		Name:        nil,
		Description: nil,
		Resources:   &[]string{resource.ID},
		Members:     nil,
	}

	v1.UpdateTeam(t, team.ID, requestBody)

	// Fetch deployment's teams
	resource = v1.GetDeployment(t, resource.ID)
	assert.EqualValues(t, []string{team.ID}, resource.Teams, "invalid teams on model")
}

func TestUpdateMembers(t *testing.T) {
	t.Parallel()

	team := v1.WithTeam(t, body.TeamCreate{
		Name:        e2e.GenName(),
		Description: e2e.GenName(),
		Resources:   nil,
		Members:     nil,
	})

	requestBody := body.TeamUpdate{
		Name:        nil,
		Description: nil,
		Resources:   nil,
		Members:     &[]body.TeamMemberUpdate{{ID: e2e.PowerUserID, TeamRole: model.TeamMemberRoleAdmin}},
	}

	v1.UpdateTeam(t, team.ID, requestBody)

	// Fetch TestUser2's teams
	teams := v1.ListTeams(t, "?userId="+e2e.PowerUserID)
	assert.NotEmpty(t, teams, "user has no teams")
}

func TestDelete(t *testing.T) {
	t.Parallel()

	team := v1.WithTeam(t, body.TeamCreate{
		Name:        e2e.GenName(),
		Description: e2e.GenName(),
		Resources:   nil,
		Members:     nil,
	})

	v1.DeleteTeam(t, team.ID)
}

func TestDeleteAsNonOwner(t *testing.T) {
	t.Parallel()

	team := v1.WithTeam(t, body.TeamCreate{
		Name:        e2e.GenName(),
		Description: e2e.GenName(),
		Resources:   nil,
		Members:     nil,
	})

	resp := e2e.DoDeleteRequest(t, v1.TeamPath+team.ID, e2e.DefaultUserID)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode, "team was deleted by non-owner member")
}
