package internal_service

import (
	"errors"
	"fmt"
	vmModel "go-deploy/models/sys/vm"
	gpu2 "go-deploy/models/sys/vm/gpu"
	"go-deploy/pkg/conf"
	"go-deploy/pkg/subsystems/cs"
	"go-deploy/pkg/subsystems/cs/commands"
	csModels "go-deploy/pkg/subsystems/cs/models"
	"log"
	"strings"
)

type CsCreated struct {
	VM *csModels.VmPublic
}

func withClient() (*cs.Client, error) {
	return cs.New(&cs.ClientConf{
		URL:         conf.Env.CS.URL,
		ApiKey:      conf.Env.CS.ApiKey,
		Secret:      conf.Env.CS.Secret,
		IpAddressID: conf.Env.CS.IpAddressID,
		NetworkID:   conf.Env.CS.NetworkID,
		ProjectID:   conf.Env.CS.ProjectID,
		ZoneID:      conf.Env.CS.ZoneID,
	})
}

func CreateCS(params *vmModel.CreateParams) (*CsCreated, error) {
	log.Println("setting up cs for", params.Name)

	userSshPublicKey := params.SshPublicKey
	adminSshPublicKey := conf.Env.VM.AdminSshPublicKey

	makeError := func(err error) error {
		return fmt.Errorf("failed to setup cs for vm %s. details: %s", params.Name, err)
	}

	client, err := withClient()
	if err != nil {
		return nil, makeError(err)
	}

	vm, err := vmModel.GetByName(params.Name)
	if err != nil {
		return nil, makeError(err)
	}

	if vm == nil {
		// if vm does not exist, don't treat as error, don't create -> job will not fail
		return nil, nil
	}

	// service offering
	var csServiceOffering *csModels.ServiceOfferingPublic
	if vm.Subsystems.CS.ServiceOffering.ID == "" {
		// create service offering
		id, err := client.CreateServiceOffering(&csModels.ServiceOfferingPublic{
			Name:        params.Name,
			Description: fmt.Sprintf("Auto-generated by deploy for vm %s", params.Name),
			CpuCores:    params.CpuCores,
			RAM:         params.RAM,
			DiskSize:    params.DiskSize,
		})
		if err != nil {
			return nil, makeError(err)
		}

		csServiceOffering, err = client.ReadServiceOffering(id)
		if err != nil {
			return nil, makeError(err)
		}

		if csServiceOffering == nil {
			return nil, makeError(errors.New("failed to read service offering after creation"))
		}

		err = vmModel.UpdateSubsystemByName(params.Name, "cs", "serviceOffering", *csServiceOffering)
		if err != nil {
			return nil, makeError(err)
		}
	} else {
		csServiceOffering = &vm.Subsystems.CS.ServiceOffering
	}

	// vm
	var csVM *csModels.VmPublic
	if vm.Subsystems.CS.VM.ID == "" {
		id, err := client.CreateVM(&csModels.VmPublic{
			Name:              params.Name,
			ServiceOfferingID: csServiceOffering.ID,
			TemplateID:        "fb6b6b11-6196-42d9-a12d-038bdeecb6f6", // deploy-template-cloud-init-ubuntu2204, temporary
			Tags: []csModels.Tag{
				{Key: "name", Value: params.Name},
				{Key: "managedBy", Value: conf.Env.Manager},
				{Key: "deployName", Value: params.Name},
			},
		},
			userSshPublicKey, adminSshPublicKey,
		)
		if err != nil {
			return nil, makeError(err)
		}

		csVM, err = client.ReadVM(id)
		if err != nil {
			return nil, makeError(errors.New("failed to read vm after creation"))
		}

		err = vmModel.UpdateSubsystemByName(params.Name, "cs", "vm", *csVM)
		if err != nil {
			return nil, makeError(err)
		}
	} else {
		csVM = &vm.Subsystems.CS.VM
	}

	// port-forwarding rule map
	addDeploySshToPortMap(&params.Ports)

	ruleMap := vm.Subsystems.CS.PortForwardingRuleMap
	if ruleMap == nil {
		ruleMap = map[string]csModels.PortForwardingRulePublic{}
	}

	for _, port := range params.Ports {
		rule, hasRule := ruleMap[port.Name]
		if !hasRule || rule.ID == "" {
			freePort, err := client.GetFreePort(conf.Env.CS.PortRange.Start, conf.Env.CS.PortRange.End)
			if err != nil {
				return nil, makeError(err)
			}

			if freePort == -1 {
				return nil, makeError(fmt.Errorf("no free port found"))
			}

			id, err := client.CreatePortForwardingRule(&csModels.PortForwardingRulePublic{
				VmID:        csVM.ID,
				Name:        params.Name,
				Protocol:    port.Protocol,
				PublicPort:  freePort,
				PrivatePort: port.Port,
				Tags: []csModels.Tag{
					{Key: "name", Value: port.Name},
					{Key: "managedBy", Value: conf.Env.Manager},
					{Key: "deployName", Value: params.Name},
				},
			})
			if err != nil {
				return nil, makeError(err)
			}

			rule, err := client.ReadPortForwardingRule(id)
			if err != nil {
				return nil, makeError(err)
			}

			if ruleMap == nil {
				return nil, makeError(fmt.Errorf("failed to read port forwarding rule after creation"))
			}

			ruleMap[port.Name] = *rule

			err = vmModel.UpdateSubsystemByName(params.Name, "cs", "portForwardingRuleMap", ruleMap)
			if err != nil {
				return nil, makeError(err)
			}
		}
	}

	return &CsCreated{
		VM: csVM,
	}, nil
}

func DeleteCS(name string) error {
	log.Println("deleting cs for", name)

	makeError := func(err error) error {
		return fmt.Errorf("failed to delete cs for vm %s. details: %s", name, err)
	}

	client, err := withClient()
	if err != nil {
		return makeError(err)
	}

	vm, err := vmModel.GetByName(name)
	if err != nil {
		return makeError(err)
	}

	if vm == nil {
		return nil
	}

	ruleMap := vm.Subsystems.CS.PortForwardingRuleMap

	for _, rule := range ruleMap {
		err = client.DeletePortForwardingRule(rule.ID)
		if err != nil {
			return makeError(err)
		}
	}

	err = vmModel.UpdateSubsystemByName(name, "cs", "portForwardingRuleMap", map[string]csModels.PortForwardingRulePublic{})
	if err != nil {
		return makeError(err)
	}

	if vm.Subsystems.CS.VM.ID != "" {
		err = client.DeleteVM(vm.Subsystems.CS.VM.ID)
		if err != nil {
			return makeError(err)
		}

		err = vmModel.UpdateSubsystemByName(name, "cs", "vm", csModels.VmPublic{})
		if err != nil {
			return makeError(err)
		}
	}

	if vm.Subsystems.CS.ServiceOffering.ID != "" {
		err = client.DeleteServiceOffering(vm.Subsystems.CS.ServiceOffering.ID)
		if err != nil {
			return makeError(err)
		}

		err = vmModel.UpdateSubsystemByName(name, "cs", "serviceOffering", csModels.ServiceOfferingPublic{})
		if err != nil {
			return makeError(err)
		}
	}

	return nil
}

func UpdateCS(vmID string, updateParams *vmModel.UpdateParams) error {
	makeError := func(err error) error {
		return fmt.Errorf("failed to update cs for vm %s. details: %s", vmID, err)
	}

	client, err := withClient()
	if err != nil {
		return makeError(err)
	}

	vm, err := vmModel.GetByID(vmID)
	if err != nil {
		return makeError(err)
	}

	if vm == nil {
		return nil
	}

	if vm.Subsystems.CS.VM.ID == "" {
		return nil
	}

	// port-forwarding rule
	if updateParams.Ports != nil {
		removeDeploySshFromPortMap(updateParams.Ports)

		/// delete old rules and create new ones
		ruleMap := vm.Subsystems.CS.PortForwardingRuleMap

		currentPortForwardingRules, err := client.ReadPortForwardingRules(vm.Subsystems.CS.VM.ID)
		if err != nil {
			return makeError(err)
		}

		currentPorts := convertToPorts(currentPortForwardingRules)
		for i, port := range currentPorts {
			if port.Name == "__ssh" || port.Port == 22 {
				continue
			}

			err = client.DeletePortForwardingRule(currentPortForwardingRules[i].ID)
			if err != nil {
				return makeError(err)
			}

			delete(ruleMap, port.Name)

			err = vmModel.UpdateSubsystemByName(vm.Name, "cs", "portForwardingRuleMap", ruleMap)
			if err != nil {
				return makeError(err)
			}
		}

		for _, port := range *updateParams.Ports {
			freePort, err := client.GetFreePort(conf.Env.CS.PortRange.Start, conf.Env.CS.PortRange.End)
			if err != nil {
				return makeError(err)
			}

			if freePort == -1 {
				return makeError(fmt.Errorf("no free port found"))
			}

			id, err := client.CreatePortForwardingRule(&csModels.PortForwardingRulePublic{
				Name:        port.Name,
				VmID:        vm.Subsystems.CS.VM.ID,
				Protocol:    port.Protocol,
				PrivatePort: port.Port,
				PublicPort:  freePort,
				Tags: []csModels.Tag{
					{Key: "name", Value: port.Name},
					{Key: "managedBy", Value: conf.Env.Manager},
					{Key: "deployName", Value: vm.Name},
				},
			})
			if err != nil {
				return makeError(err)
			}

			rule, err := client.ReadPortForwardingRule(id)
			if err != nil {
				return makeError(err)
			}

			ruleMap[port.Name] = *rule

			err = vmModel.UpdateSubsystemByName(vm.Name, "cs", "portForwardingRuleMap", ruleMap)
			if err != nil {
				return makeError(err)
			}
		}
	}

	// service offering
	var serviceOffering *csModels.ServiceOfferingPublic
	if vm.Subsystems.CS.ServiceOffering.ID != "" {
		requiresUpdate := updateParams.CpuCores != nil && *updateParams.CpuCores != vm.Subsystems.CS.ServiceOffering.CpuCores ||
			updateParams.RAM != nil && *updateParams.RAM != vm.Subsystems.CS.ServiceOffering.RAM

		if requiresUpdate {
			err = client.DeleteServiceOffering(vm.Subsystems.CS.ServiceOffering.ID)
			if err != nil {
				return makeError(err)
			}

			id, err := client.CreateServiceOffering(&csModels.ServiceOfferingPublic{
				Name:        vm.Name,
				Description: vm.Subsystems.CS.ServiceOffering.Description,
				CpuCores:    *updateParams.CpuCores,
				RAM:         *updateParams.RAM,
				DiskSize:    vm.Specs.DiskSize,
			})
			if err != nil {
				return makeError(err)
			}

			serviceOffering, err = client.ReadServiceOffering(id)
			if err != nil {
				return makeError(err)
			}

			if serviceOffering == nil {
				return makeError(fmt.Errorf("failed to read service offering after creation"))
			}

			err = vmModel.UpdateSubsystemByName(vm.Name, "cs", "serviceOffering", *serviceOffering)
			if err != nil {
				return makeError(err)
			}
		} else {
			serviceOffering = &vm.Subsystems.CS.ServiceOffering
		}

		// make sure the vm is using the latest service offering
		if vm.Subsystems.CS.VM.ServiceOfferingID != serviceOffering.ID {
			vm.Subsystems.CS.VM.ServiceOfferingID = serviceOffering.ID

			// turn it off if it is on, but remember the status
			status, err := client.GetVmStatus(vm.Subsystems.CS.VM.ID)
			if err != nil {
				return makeError(err)
			}

			if status == "Running" {
				err = client.DoVmCommand(vm.Subsystems.CS.VM.ID, nil, "stop")
				if err != nil {
					return makeError(err)
				}
			}

			// update the service offering
			err = client.UpdateVM(&vm.Subsystems.CS.VM)
			if err != nil {
				return makeError(err)
			}

			err = vmModel.UpdateSubsystemByName(vm.Name, "cs", "vm", vm.Subsystems.CS.VM)
			if err != nil {
				return makeError(err)
			}

			// turn it on if it was on
			if status == "Running" {
				var requiredHost *string
				if vm.GpuID != "" {
					requiredHost, err = getRequiredHost(vm.GpuID)
					if err != nil {
						return makeError(err)
					}
				}

				err = client.DoVmCommand(vm.Subsystems.CS.VM.ID, requiredHost, "start")
				if err != nil {
					return makeError(err)
				}
			}

		}
	}

	return nil
}

func AttachGPU(gpuID, vmID string) error {
	makeError := func(err error) error {
		return fmt.Errorf("failed to attach gpu %s to cs vm %s. details: %s", gpuID, vmID, err)
	}

	client, err := withClient()
	if err != nil {
		return makeError(err)
	}

	vm, err := vmModel.GetByID(vmID)
	if err != nil {
		return makeError(err)
	}

	if vm == nil {
		return makeError(fmt.Errorf("vm %s not found", vmID))
	}

	if vm.Subsystems.CS.VM.ID == "" {
		return makeError(fmt.Errorf("vm is not created yet"))
	}

	gpu, err := gpu2.GetGpuByID(gpuID)
	if err != nil {
		return makeError(err)
	}

	status, err := client.GetVmStatus(vm.Subsystems.CS.VM.ID)
	if err != nil {
		return makeError(err)
	}

	if status == "Running" {
		err = client.DoVmCommand(vm.Subsystems.CS.VM.ID, nil, "stop")
		if err != nil {
			return makeError(err)
		}
	}

	vm.Subsystems.CS.VM.ExtraConfig = createExtraConfig(gpu)

	err = client.UpdateVM(&vm.Subsystems.CS.VM)
	if err != nil {
		return makeError(err)
	}

	err = vmModel.UpdateSubsystemByName(vm.Name, "cs", "vm.extraConfig", vm.Subsystems.CS.VM.ExtraConfig)
	if err != nil {
		return makeError(err)
	}

	// always start the vm after attaching gpu, to make sure the vm can be started on the host
	requiredHost, err := getRequiredHost(gpuID)
	if err != nil {
		return makeError(err)
	}

	err = client.DoVmCommand(vm.Subsystems.CS.VM.ID, requiredHost, "start")
	if err != nil {
		return makeError(err)
	}

	return nil
}

func DetachGPU(vmID string) error {
	makeError := func(err error) error {
		return fmt.Errorf("failed to detach gpu from cs vm %s. details: %s", vmID, err)
	}

	client, err := withClient()
	if err != nil {
		return makeError(err)
	}

	vm, err := vmModel.GetByID(vmID)
	if err != nil {
		return makeError(err)
	}

	if vm == nil {
		return makeError(fmt.Errorf("vm %s not found", vmID))
	}

	if vm.Subsystems.CS.VM.ID == "" {
		return makeError(fmt.Errorf("vm is not created yet"))
	}

	// turn it off if it is on, but remember the status
	status, err := client.GetVmStatus(vm.Subsystems.CS.VM.ID)
	if err != nil {
		return makeError(err)
	}

	if status == "Running" {
		err = client.DoVmCommand(vm.Subsystems.CS.VM.ID, nil, "stop")
		if err != nil {
			return makeError(err)
		}
	}

	vm.Subsystems.CS.VM.ExtraConfig = ""

	err = client.UpdateVM(&vm.Subsystems.CS.VM)
	if err != nil {
		return makeError(err)
	}

	err = vmModel.UpdateSubsystemByName(vm.Name, "cs", "vm.extraConfig", vm.Subsystems.CS.VM.ExtraConfig)
	if err != nil {
		return makeError(err)
	}

	// turn it on if it was on
	if status == "Running" {
		err = client.DoVmCommand(vm.Subsystems.CS.VM.ID, nil, "start")
		if err != nil {
			return makeError(err)
		}
	}

	return nil
}

func IsGpuAttachedCS(host string, bus string) (bool, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to check if gpu %s:%s is attached to any cs vm. details: %s", host, bus, err)
	}

	client, err := withClient()
	if err != nil {
		return false, makeError(err)
	}

	params := client.CsClient.VirtualMachine.NewListVirtualMachinesParams()
	params.SetListall(true)

	vms, err := client.CsClient.VirtualMachine.ListVirtualMachines(params)
	if err != nil {
		return false, makeError(err)
	}

	for _, vm := range vms.VirtualMachines {
		if vm.Details != nil && vm.Hostname == host {
			extraConfig, ok := vm.Details["extraconfig-1"]
			if ok {
				if strings.Contains(extraConfig, fmt.Sprintf("bus='0x%s'", bus)) {
					return true, nil
				}
			}
		}
	}

	return false, nil
}

func DoCommandCS(vmID string, gpuID *string, command string) error {
	makeError := func(err error) error {
		return fmt.Errorf("failed to execute command %s for cs vm %s. details: %s", command, vmID, err)
	}

	client, err := withClient()
	if err != nil {
		return makeError(err)
	}

	var requiredHost *string
	if gpuID != nil {
		requiredHost, err = getRequiredHost(*gpuID)
		if err != nil {
			return makeError(err)
		}
	}

	err = client.DoVmCommand(vmID, requiredHost, commands.Command(command))
	if err != nil {
		return makeError(err)
	}

	return nil
}

func CanStartCS(vmID, host string) (bool, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to check if cs vm %s can be started on host %s. details: %s", vmID, host, err)
	}

	client, err := withClient()
	if err != nil {
		return false, makeError(err)
	}

	canStart, err := client.CanStartOn(vmID, host)
	if err != nil {
		return false, err
	}

	return canStart, nil
}

func createExtraConfig(gpu *gpu2.GPU) string {
	data := fmt.Sprintf(`
<devices> <hostdev mode='subsystem' type='pci' managed='yes'> <driver name='vfio' />
	<source> <address domain='0x0000' bus='0x%s' slot='0x00' function='0x0' /> </source> 
	<alias name='nvidia0' /> <address type='pci' domain='0x0000' bus='0x00' slot='0x00' function='0x0' /> 
</hostdev> </devices>`, gpu.Data.Bus)

	data = strings.Replace(data, "\n", "", -1)
	data = strings.Replace(data, "\t", "", -1)

	return data
}

func getRequiredHost(gpuID string) (*string, error) {
	gpu, err := gpu2.GetGpuByID(gpuID)
	if err != nil {
		return nil, err
	}

	if gpu.Host == "" {
		return nil, fmt.Errorf("no host found for gpu %s", gpu.ID)
	}

	return &gpu.Host, nil
}

func addDeploySshToPortMap(portMap *[]vmModel.Port) {
	for i, port := range *portMap {
		if (port.Port == 22 || port.Name == "__ssh") && port.Protocol == "tcp" {
			*portMap = append((*portMap)[:i], (*portMap)[i+1:]...)
			break
		}
	}

	*portMap = append(*portMap, vmModel.Port{
		Port:     22,
		Name:     "__ssh",
		Protocol: "tcp",
	})
}

func removeDeploySshFromPortMap(portMap *[]vmModel.Port) {
	for i, port := range *portMap {
		if (port.Port == 22 || port.Name == "__ssh") && port.Protocol == "tcp" {
			*portMap = append((*portMap)[:i], (*portMap)[i+1:]...)
			break
		}
	}
}

func convertToPorts(rules []csModels.PortForwardingRulePublic) []vmModel.Port {
	var ports []vmModel.Port

	for _, rule := range rules {
		var name string
		for _, tag := range rule.Tags {
			if tag.Key == "name" {
				name = tag.Value
				break
			}
		}

		ports = append(ports, vmModel.Port{
			Port:     rule.PublicPort,
			Name:     name,
			Protocol: rule.Protocol,
		})
	}

	return ports
}
