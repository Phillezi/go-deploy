package role

import (
	"github.com/fatih/structs"
	"go-deploy/models/dto/body"
)

func (r *Role) ToDTO(includeQuota bool) body.Role {
	permissionsStructMap := structs.Map(r.Permissions)
	permissions := make([]string, 0)
	for name, value := range permissionsStructMap {
		hasPermission, ok := value.(bool)
		if ok && hasPermission {
			permissions = append(permissions, name)
		}
	}

	var quota *body.Quota
	if includeQuota {
		dto := r.Quotas.ToDTO()
		quota = &dto
	}

	return body.Role{
		Name:        r.Name,
		Description: r.Description,
		Permissions: permissions,
		Quota:       quota,
	}
}

func (q *Quotas) ToDTO() body.Quota {
	return body.Quota{
		Deployments:      q.Deployments,
		CpuCores:         q.CpuCores,
		RAM:              q.RAM,
		DiskSize:         q.DiskSize,
		Snapshots:        q.Snapshots,
		GpuLeaseDuration: q.GpuLeaseDuration,
	}
}