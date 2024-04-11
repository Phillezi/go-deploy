package zones

import (
	"go-deploy/test/e2e"
	v1 "go-deploy/test/e2e/v1"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	e2e.Setup()
	code := m.Run()
	e2e.Shutdown()
	os.Exit(code)
}

func TestList(t *testing.T) {
	queries := []string{
		"?page=1&pageSize=10",
	}

	for _, query := range queries {
		v1.ListZones(t, query)
	}
}
