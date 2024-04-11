package v1

import (
	"github.com/stretchr/testify/assert"
	"go-deploy/dto/v1/body"
	"go-deploy/pkg/app/status_codes"
	"go-deploy/test"
	"go-deploy/test/e2e"
	"golang.org/x/net/idna"
	"net/http"
	"strconv"
	"testing"
)

const (
	DeploymentPath  = "/v1/deployments/"
	DeploymentsPath = "/v1/deployments"
)

func CheckUpURL(t *testing.T, url string) bool {
	t.Helper()

	resp, err := http.Get(url)
	if err == nil {
		if resp.StatusCode == http.StatusOK {
			return true
		}
	}

	return false
}

func GetDeployment(t *testing.T, id string, userID ...string) body.DeploymentRead {
	resp := e2e.DoGetRequest(t, DeploymentPath+id, userID...)
	return e2e.Parse[body.DeploymentRead](t, resp)
}

func ListDeployments(t *testing.T, query string, userID ...string) []body.DeploymentRead {
	resp := e2e.DoGetRequest(t, DeploymentsPath+query, userID...)
	return e2e.Parse[[]body.DeploymentRead](t, resp)
}

func UpdateDeployment(t *testing.T, id string, requestBody body.DeploymentUpdate, userID ...string) body.DeploymentRead {
	resp := e2e.DoPostRequest(t, DeploymentPath+id, requestBody, userID...)
	deploymentUpdated := e2e.Parse[body.DeploymentUpdated](t, resp)

	if deploymentUpdated.JobID != nil {
		WaitForJobFinished(t, *deploymentUpdated.JobID, nil)
	}
	WaitForDeploymentRunning(t, id, func(read *body.DeploymentRead) bool {
		if read.Private {
			return true
		}

		// Make sure it is accessible
		if read.URL != nil {
			return CheckUpURL(t, *read.URL)
		}
		return false
	})

	updated := GetDeployment(t, id, userID...)

	if requestBody.CustomDomain != nil {
		punyEncoded, err := idna.New().ToASCII("https://" + *requestBody.CustomDomain)
		assert.NoError(t, err, "custom domain was not puny encoded")
		assert.Equal(t, punyEncoded, *updated.CustomDomainURL)
	}

	if requestBody.Image != nil {
		assert.Equal(t, *requestBody.Image, *updated.Image)
	}

	if requestBody.InitCommands != nil {
		test.EqualOrEmpty(t, *requestBody.InitCommands, updated.InitCommands)
	}

	if requestBody.Envs != nil {
		// PORT env is generated, so it will always be in the read, but not necessarily in the request
		customPort := false
		for _, env := range *requestBody.Envs {
			if env.Name == "PORT" {
				customPort = true
				break
			}
		}

		// If it was not requested, we add it to the request body so we can compare the rest
		if !customPort {
			*requestBody.Envs = append(*requestBody.Envs, body.Env{
				Name:  "PORT",
				Value: "8080",
			})
		}

		test.EqualOrEmpty(t, *requestBody.Envs, updated.Envs)
	}

	if requestBody.Volumes != nil {
		test.EqualOrEmpty(t, *requestBody.Volumes, updated.Volumes)
	}

	if requestBody.Private != nil {
		assert.Equal(t, *requestBody.Private, updated.Private)
	}

	if requestBody.HealthCheckPath != nil {
		assert.Equal(t, *requestBody.HealthCheckPath, updated.HealthCheckPath)
	}

	if requestBody.Replicas != nil {
		assert.Equal(t, *requestBody.Replicas, updated.Replicas)
	}

	return updated
}

func WithDeployment(t *testing.T, requestBody body.DeploymentCreate) (body.DeploymentRead, string) {
	resp := e2e.DoPostRequest(t, DeploymentsPath, requestBody)
	assert.Equal(t, http.StatusOK, resp.StatusCode, "deployment was not created")

	var deploymentCreated body.DeploymentCreated
	err := e2e.ReadResponseBody(t, resp, &deploymentCreated)
	assert.NoError(t, err, "deployment was not created")

	t.Cleanup(func() {
		CleanUpDeployment(t, deploymentCreated.ID)
	})

	WaitForJobFinished(t, deploymentCreated.JobID, nil)
	WaitForDeploymentRunning(t, deploymentCreated.ID, func(deploymentRead *body.DeploymentRead) bool {
		//make sure it is accessible
		if deploymentRead.URL != nil {
			return CheckUpURL(t, *deploymentRead.URL)
		}
		return false
	})

	deploymentRead := GetDeployment(t, deploymentCreated.ID)
	assert.NotEmpty(t, deploymentRead.ID)
	assert.Equal(t, requestBody.Name, deploymentRead.Name)
	assert.Equal(t, requestBody.Private, deploymentRead.Private)

	if requestBody.Zone == nil {
		// Some zone is set by default
		assert.NotEmpty(t, deploymentRead.Zone)
	} else {
		assert.Equal(t, requestBody.Zone, deploymentRead.Zone)
	}

	// PORT env is generated, so we check for that first, then delete and match exact with the rest
	assert.NotEmpty(t, deploymentRead.Envs)

	portRequested := 0
	customPort := false
	for _, env := range requestBody.Envs {
		if env.Name == "PORT" {
			portRequested, _ = strconv.Atoi(env.Value)
			customPort = true
			break
		}
	}

	if portRequested == 0 {
		portRequested = 8080
	}

	portReceived := 0
	for _, env := range deploymentRead.Envs {
		if env.Name == "PORT" {
			portReceived, _ = strconv.Atoi(env.Value)
			break
		}
	}

	assert.NotZero(t, portReceived)
	assert.Equal(t, portRequested, portReceived)

	if !customPort {
		// remove PORT from read, since we won't be able to match the rest of the envs other wise
		for i, env := range deploymentRead.Envs {
			if env.Name == "PORT" {
				deploymentRead.Envs = append(deploymentRead.Envs[:i], deploymentRead.Envs[i+1:]...)
				break
			}
		}
	}

	test.EqualOrEmpty(t, requestBody.InitCommands, deploymentRead.InitCommands, "invalid init commands")
	test.EqualOrEmpty(t, requestBody.Args, deploymentRead.Args, "invalid args")
	test.EqualOrEmpty(t, requestBody.Envs, deploymentRead.Envs, "invalid envs")
	test.EqualOrEmpty(t, requestBody.Volumes, deploymentRead.Volumes, "invalid volumes")

	if requestBody.CustomDomain != nil {
		punyEncoded, err := idna.New().ToASCII("https://" + *requestBody.CustomDomain)
		assert.NoError(t, err, "custom domain was not puny encoded")
		assert.Equal(t, punyEncoded, *deploymentRead.CustomDomainURL)
	}

	return deploymentRead, deploymentCreated.JobID
}

func WithAssumedFailedDeployment(t *testing.T, requestBody body.DeploymentCreate) {
	resp := e2e.DoPostRequest(t, DeploymentsPath, requestBody)
	if resp.StatusCode == http.StatusBadRequest {
		return
	}

	var deploymentCreated body.DeploymentCreated
	err := e2e.ReadResponseBody(t, resp, &deploymentCreated)
	assert.NoError(t, err, "deployment created body was not read")

	t.Cleanup(func() { CleanUpDeployment(t, deploymentCreated.ID) })

	assert.FailNow(t, "deployment was created but should have failed")
}

func CleanUpDeployment(t *testing.T, id string) {
	resp := e2e.DoDeleteRequest(t, DeploymentPath+id)
	if resp.StatusCode == http.StatusNotFound {
		return
	}

	if resp.StatusCode == http.StatusOK {
		var vmDeleted body.VmDeleted
		err := e2e.ReadResponseBody(t, resp, &vmDeleted)
		assert.NoError(t, err, "deleted body was not read")
		assert.Equal(t, id, vmDeleted.ID)

		WaitForJobFinished(t, vmDeleted.JobID, nil)
		WaitForDeploymentDeleted(t, vmDeleted.ID, nil)

		return
	}

	assert.FailNow(t, "deployment was not deleted")
}

func WaitForDeploymentRunning(t *testing.T, id string, callback func(*body.DeploymentRead) bool) {
	e2e.FetchUntil(t, DeploymentPath+id, func(resp *http.Response) bool {
		deploymentRead := e2e.Parse[body.DeploymentRead](t, resp)
		if deploymentRead.Status == status_codes.GetMsg(status_codes.ResourceRunning) {
			if callback == nil || callback(&deploymentRead) {
				return true
			}
		}

		return false
	})
}

func WaitForDeploymentDeleted(t *testing.T, id string, callback func() bool) {
	e2e.FetchUntil(t, DeploymentPath+id, func(resp *http.Response) bool {
		return resp.StatusCode == http.StatusNotFound
	})
}
