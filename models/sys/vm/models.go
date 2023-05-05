package vm

import (
	csModels "go-deploy/pkg/subsystems/cs/models"
	psModels "go-deploy/pkg/subsystems/pfsense/models"
)

type Port struct {
	Name     string `bson:"name"`
	Port     int    `bson:"port"`
	Protocol string `bson:"protocol"`
}

type VM struct {
	ID        string `bson:"id"`
	Name      string `bson:"name"`
	OwnerID   string `bson:"ownerId"`
	ManagedBy string `bson:"managedBy"`

	GpuID        string `bson:"gpuId"`
	SshPublicKey string `bson:"sshPublicKey"`

	Ports []struct {
		Name string `bson:"name"`
		Port int    `bson:"port"`
	} `bson:"ports"`

	BeingCreated  bool       `bson:"beingCreated"`
	BeingDeleted bool       `bson:"beingDeleted"`
	Subsystems   Subsystems `bson:"subsystems"`
	StatusCode   int        `bson:"statusCode"`
	StatusMessage string     `bson:"statusMessage"`
}

type Subsystems struct {
	CS      CS      `bson:"cs"`
	PfSense PfSense `bson:"pfSense"`
}

type CS struct {
	VM                    csModels.VmPublic                            `bson:"vm"`
	PortForwardingRuleMap map[string]csModels.PortForwardingRulePublic `bson:"portForwardingRuleMap"`
	PublicIpAddress       csModels.PublicIpAddressPublic               `bson:"publicIpAddress"`
}

type PfSense struct {
	PortForwardingRuleMap map[string]psModels.PortForwardingRulePublic `bson:"portForwardingRuleMap"`
}


type VmUpdate struct {
	Ports *[]Port `json:"ports"`
}