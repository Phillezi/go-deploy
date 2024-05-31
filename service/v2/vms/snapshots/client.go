package snapshots

import (
	"go-deploy/service/clients"
	"go-deploy/service/core"
	"go-deploy/service/v2/vms/client"
)

type Client struct {
	V2 clients.V2

	client.BaseClient[Client]
}

func New(v2 clients.V2, cache ...*core.Cache) *Client {
	var ca *core.Cache
	if len(cache) > 0 {
		ca = cache[0]
	} else {
		ca = core.NewCache()
	}

	c := &Client{V2: v2, BaseClient: client.NewBaseClient[Client](ca)}
	c.BaseClient.SetParent(c)
	return c
}
