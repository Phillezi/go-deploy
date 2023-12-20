package resources

import (
	"fmt"
	"go-deploy/models/sys/vm"
	"go-deploy/pkg/config"
	"go-deploy/pkg/subsystems"
	"go-deploy/pkg/subsystems/cs/models"
	"go-deploy/utils/subsystemutils"
	"golang.org/x/exp/slices"
	"sort"
	"time"
)

type CsGenerator struct {
	*PublicGeneratorType
}

func (cr *CsGenerator) SOs() []models.ServiceOfferingPublic {
	var res []models.ServiceOfferingPublic

	if cr.v.vm != nil {
		so := models.ServiceOfferingPublic{
			Name:        subsystemutils.GetPrefixedName(cr.v.vm.Name),
			Description: fmt.Sprintf("Auto-generated by deploy for vm %s", cr.v.vm.Name),
			CpuCores:    cr.v.vm.Specs.CpuCores,
			RAM:         cr.v.vm.Specs.RAM,
			DiskSize:    cr.v.vm.Specs.DiskSize,
		}

		if s := &cr.v.vm.Subsystems.CS.ServiceOffering; subsystems.Created(s) {
			so.ID = s.ID
			so.CreatedAt = s.CreatedAt
		}

		res = append(res, so)
		return res
	}

	return nil
}

func (cr *CsGenerator) VMs() []models.VmPublic {
	var res []models.VmPublic

	if cr.v.vm != nil {
		csVM := models.VmPublic{
			Name:              cr.v.vm.Name,
			ServiceOfferingID: cr.v.vm.Subsystems.CS.ServiceOffering.ID,
			TemplateID:        cr.v.vmZone.TemplateID,
			ExtraConfig:       cr.v.vm.Subsystems.CS.VM.ExtraConfig,
			Tags:              createTags(cr.v.vm.Name, cr.v.vm.Name),
		}

		if v := &cr.v.vm.Subsystems.CS.VM; subsystems.Created(v) {
			csVM.ID = v.ID
			csVM.CreatedAt = v.CreatedAt

			for idx, tag := range csVM.Tags {
				if tag.Key == "createdAt" {
					for _, createdTag := range v.Tags {
						if createdTag.Key == "createdAt" {
							csVM.Tags[idx].Value = createdTag.Value
						}
					}
				}
			}
		}

		res = append(res, csVM)
		return res
	}

	return nil
}

func (cr *CsGenerator) PFRs() []models.PortForwardingRulePublic {
	var res []models.PortForwardingRulePublic

	if cr.v.vm != nil {
		ports := cr.v.vm.Ports

		for _, port := range ports {
			res = append(res, models.PortForwardingRulePublic{
				Name:        port.Name,
				VmID:        cr.v.vm.Subsystems.CS.VM.ID,
				NetworkID:   cr.v.vmZone.NetworkID,
				IpAddressID: cr.v.vmZone.IpAddressID,
				PublicPort:  0, // this is set externally
				PrivatePort: port.Port,
				Protocol:    port.Protocol,
				Tags:        createTags(port.Name, cr.v.vm.Name),
			})
		}

		for mapName, pfr := range cr.v.vm.Subsystems.CS.GetPortForwardingRuleMap() {
			idx := slices.IndexFunc(ports, func(p vm.Port) bool {
				if p.Name == "__ssh" {
					return p.Name == mapName
				} else {
					return pfrName(p.Port, p.Protocol) == mapName
				}
			})

			if idx != -1 {
				res[idx].ID = pfr.ID
				res[idx].CreatedAt = pfr.CreatedAt
				res[idx].PublicPort = pfr.PublicPort

				for jdx, tag := range res[idx].Tags {
					if tag.Key == "createdAt" {
						for _, createdTag := range pfr.Tags {
							if createdTag.Key == "createdAt" {
								res[idx].Tags[jdx].Value = createdTag.Value
							}
						}
					}
				}
			}
		}

		return res
	}

	return nil
}

func createTags(name string, deployName string) []models.Tag {
	tags := []models.Tag{
		{Key: "name", Value: name},
		{Key: "managedBy", Value: config.Config.Manager},
		{Key: "deployName", Value: deployName},
		{Key: "createdAt", Value: time.Now().Format(time.RFC3339)},
	}

	sort.Slice(tags, func(i, j int) bool {
		return tags[i].Key < tags[j].Key
	})

	return tags
}

func pfrName(privatePort int, protocol string) string {
	return fmt.Sprintf("priv-%d-prot-%s", privatePort, protocol)
}
