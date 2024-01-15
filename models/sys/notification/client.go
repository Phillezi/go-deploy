package notification

import (
	"go-deploy/models/db"
	"go-deploy/models/sys/base"
	"go-deploy/models/sys/base/resource"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Client is used to manage notifications in the database.
type Client struct {
	Collection     *mongo.Collection
	RestrictUserID *string

	resource.ResourceClient[Notification]
}

// New returns a new notification client.
func New() *Client {
	return &Client{
		Collection: db.DB.GetCollection("notifications"),

		ResourceClient: resource.ResourceClient[Notification]{
			Collection:     db.DB.GetCollection("notifications"),
			IncludeDeleted: false,
		},
	}
}

// WithPagination adds pagination to the client.
func (client *Client) WithPagination(page, pageSize int) *Client {
	client.ResourceClient.Pagination = &base.Pagination{
		Page:     page,
		PageSize: pageSize,
	}

	return client
}

// WithUserID adds a filter to the client to only include notifications with the given user ID.
func (client *Client) WithUserID(ownerID string) *Client {
	client.ResourceClient.AddExtraFilter(bson.D{{"userId", ownerID}})
	client.RestrictUserID = &ownerID

	return client
}

// WithType adds a filter to the client to only include notifications with the given type.
func (client *Client) WithType(notificationType string) *Client {
	client.ResourceClient.AddExtraFilter(bson.D{{"type", notificationType}})

	return client
}

// FilterContent adds a filter to the client to only include notifications with the given content.
func (client *Client) FilterContent(contentName string, filter interface{}) *Client {
	client.AddExtraFilter(bson.D{{"content." + contentName, filter}})

	return client
}
