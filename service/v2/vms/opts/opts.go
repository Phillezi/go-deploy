package opts

import (
	configModels "go-deploy/models/config"
	"go-deploy/models/dto/v2/body"
	roleModels "go-deploy/models/sys/role"
	vmModels "go-deploy/models/sys/vm"
	"go-deploy/service/v2/utils"
)

// Opts is used to specify which resources to get.
type Opts struct {
	VmID      string
	Client    bool
	Generator bool

	ExtraOpts
}

// ExtraOpts is used to specify the extra options when getting a VM.
// This is useful when overwriting the implicit options,
// such as where user ID is by default taken from VM.OwnerID.
type ExtraOpts struct {
	UserID string
	Zone   *configModels.DeploymentZone
}

// GetOpts is used to specify the options when getting a VM.
type GetOpts struct {
	TransferCode *string
	Shared       bool
}

// ListOpts is used to specify the options when listing VMs.
type ListOpts struct {
	Pagination *utils.Pagination
	UserID     *string
	Shared     bool
}

// GetGpuOpts is used to specify the options when getting a VM's gpu.
type GetGpuOpts struct {
	Zone          *string
	AvailableGPUs bool
}

// ListGpuOpts is used to specify the options when listing VMs' gpus.
type ListGpuOpts struct {
	Pagination    *utils.Pagination
	Zone          *string
	AvailableGPUs bool
}

// GetSnapshotOpts is used to specify the options when getting a VM's snapshot.
type GetSnapshotOpts struct {
}

// ListSnapshotOpts is used to specify the options when listing VMs' snapshots.
type ListSnapshotOpts struct {
	Pagination *utils.Pagination
}

// CreateSnapshotOpts is used to specify the options when creating a VM's snapshot.
type CreateSnapshotOpts struct {
	System *vmModels.CreateSnapshotParams
	User   *body.VmSnapshotCreate
}

// QuotaOpts is used to specify the options when getting a VM's quota.
type QuotaOpts struct {
	Quota          *roleModels.Quotas
	Create         *body.VmCreate
	Update         *body.VmUpdate
	CreateSnapshot *body.VmSnapshotCreate
}