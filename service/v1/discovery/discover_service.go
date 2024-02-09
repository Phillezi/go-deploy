package discovery

import (
	"go-deploy/models/sys/discover"
	"go-deploy/pkg/config"
)

// Discover returns the discover information.
// This is stored locally in the config.
func (c *Client) Discover() (*discover.Discover, error) {
	return &discover.Discover{
		Version: config.Config.Version,
		Roles:   config.Config.Roles,
	}, nil
}