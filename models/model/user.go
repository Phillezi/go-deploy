package model

import (
	"time"
)

type User struct {
	ID        string `bson:"id"`
	Username  string `bson:"username"`
	FirstName string `bson:"firstName"`
	LastName  string `bson:"lastName"`
	Email     string `bson:"email"`

	IsAdmin       bool          `bson:"isAdmin"`
	EffectiveRole EffectiveRole `bson:"effectiveRole"`

	PublicKeys []PublicKey `bson:"publicKeys"`
	UserData   []UserData  `bson:"userData"`

	LastAuthenticatedAt time.Time `bson:"lastAuthenticatedAt"`
}

type PublicKey struct {
	Name string `bson:"name"`
	Key  string `bson:"key"`
}

type UserData struct {
	Key   string `bson:"key"`
	Value string `bson:"value"`
}

type UserUsage struct {
	Deployments int `bson:"deployments"`
	CpuCores    int `bson:"cpuCores"`
	RAM         int `bson:"ram"`
	DiskSize    int `bson:"diskSize"`
	Snapshots   int `bson:"snapshots"`
}

type EffectiveRole struct {
	Name        string `bson:"name"`
	Description string `bson:"description"`
}
