package notification_service

import (
	notificationModels "go-deploy/models/sys/notification"
	"go-deploy/service"
	"sort"
)

type Client struct {
	// Cache is used to cache the resources fetched inside the service.
	Cache *service.Cache

	// Auth is the authentication information for the client.
	Auth *service.AuthInfo
}

func New() *Client {
	return &Client{
		Cache: service.NewCache(),
	}
}

// WithAuth sets the auth on the context.
func (c *Client) WithAuth(auth *service.AuthInfo) *Client {
	c.Auth = auth
	return c
}

func (c *Client) Notification(id string, nmc *notificationModels.Client) (*notificationModels.Notification, error) {
	notification := c.Cache.GetNotification(id)
	if notification == nil {
		var err error
		notification, err = nmc.GetByID(id)
		if err != nil {
			return nil, err
		}

		c.Cache.StoreNotification(notification)
	}

	return notification, nil
}

func (c *Client) Notifications(nmc *notificationModels.Client) ([]notificationModels.Notification, error) {
	// Right now we don't have a way to skip fetching when requesting a list of resources
	notifications, err := nmc.List()
	if err != nil {
		return nil, err
	}

	for _, user := range notifications {
		c.Cache.StoreNotification(&user)
	}

	sort.Slice(notifications, func(i, j int) bool {
		return notifications[i].CreatedAt.After(notifications[j].CreatedAt)
	})

	return notifications, nil
}

func (c *Client) RefreshNotification(id string, umc *notificationModels.Client) (*notificationModels.Notification, error) {
	notification, err := umc.GetByID(id)
	if err != nil {
		return nil, err
	}

	c.Cache.StoreNotification(notification)
	return notification, nil
}