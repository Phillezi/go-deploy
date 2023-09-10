package e2e

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go-deploy/pkg/app"
	"go-deploy/pkg/conf"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

const TestUserID = "955f0f87-37fd-4792-90eb-9bf6989e698e"

var deployApp *app.App

func Setup() {
	requiredEnvs := []string{
		"DEPLOY_CONFIG_FILE",
	}

	for _, env := range requiredEnvs {
		_, result := os.LookupEnv(env)
		if !result {
			log.Fatalln("required environment variable not set: " + env)
		}
	}

	_, result := os.LookupEnv("DEPLOY_CONFIG_FILE")
	if result {
		conf.SetupEnvironment()
	}

	conf.Env.TestMode = true
	conf.Env.DB.Name = conf.Env.DB.Name + "-test"

	deployApp = app.Create(&app.Options{
		API:           true,
		Confirmer:     true,
		StatusUpdater: true,
		JobExecutor:   true,
		Repairer:      true,
		Pinger:        true,
		Snapshotter:   true,
		TestMode:      true,
	})
	if deployApp == nil {
		log.Fatalln("failed to create app")
	}

	// TODO: wait for server to start instead of using this "hack"
	time.Sleep(3 * time.Second)
}

func Shutdown() {
	if deployApp != nil {
		deployApp.Stop()
	}
}

func GenName(base string) string {
	return base + "-" + strings.ReplaceAll(uuid.NewString()[:10], "-", "")
}

func DoPlainGetRequest(t *testing.T, path string) *http.Response {
	t.Helper()

	req, err := http.NewRequest("GET", path, nil)
	assert.NoError(t, err)

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err)

	return resp
}

func CreateServerURL(subPath string) string {
	return CreateServerUrlWithProtocol("http", subPath)
}

func CreateServerUrlWithProtocol(protocol, subPath string) string {
	return protocol + "://localhost:8080/v1" + subPath
}

func DoGetRequest(t *testing.T, subPath string) *http.Response {
	t.Helper()

	req, err := http.NewRequest("GET", CreateServerURL(subPath), nil)
	assert.NoError(t, err)

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err)

	return resp
}

func DoPostRequest(t *testing.T, subPath string, body interface{}) *http.Response {
	t.Helper()

	jsonBody, err := json.Marshal(body)
	assert.NoError(t, err)
	assert.NotEmpty(t, jsonBody)

	req, err := http.NewRequest("POST", CreateServerURL(subPath), bytes.NewBuffer(jsonBody))
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err)

	return resp
}

func DoDeleteRequest(t *testing.T, subPath string) *http.Response {
	t.Helper()

	req, err := http.NewRequest("DELETE", CreateServerURL(subPath), nil)
	assert.NoError(t, err)

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err)

	return resp
}

func ReadResponseBody(t *testing.T, resp *http.Response, body interface{}) error {
	t.Cleanup(func() {
		err := resp.Body.Close()
		assert.NoError(t, err)
	})

	return json.NewDecoder(resp.Body).Decode(body)
}
