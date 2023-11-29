package cs_service

import (
	"errors"
	"fmt"
	configModels "go-deploy/models/config"
	gpu2 "go-deploy/models/sys/gpu"
	vmModel "go-deploy/models/sys/vm"
	"go-deploy/pkg/config"
	"go-deploy/pkg/subsystems/cs/commands"
	csModels "go-deploy/pkg/subsystems/cs/models"
	"go-deploy/service"
	"go-deploy/service/resources"
	"go-deploy/service/vm_service/base"
	"go-deploy/service/vm_service/service_errors"
	"golang.org/x/exp/slices"
	"log"
)

func Create(vmID string, params *vmModel.CreateParams) error {
	log.Println("setting up cs for", params.Name)

	makeError := func(err error) error {
		return fmt.Errorf("failed to setup cs for vm %s. details: %w", params.Name, err)
	}

	context, err := NewContext(vmID)
	if err != nil {
		if errors.Is(err, base.VmDeletedErr) {
			return nil
		}

		return makeError(err)
	}

	context.Client.WithUserSshPublicKey(params.SshPublicKey)
	context.Client.WithAdminSshPublicKey(config.Config.VM.AdminSshPublicKey)

	// Service offering
	for _, soPublic := range context.Generator.SOs() {
		err = resources.SsCreator(context.Client.CreateServiceOffering).
			WithDbFunc(dbFunc(vmID, "serviceOffering")).
			WithPublic(&soPublic).
			Exec()

		if err != nil {
			return makeError(err)
		}

		context.VM.Subsystems.CS.ServiceOffering = soPublic
	}

	// VM
	for _, vmPublic := range context.Generator.VMs() {
		err = resources.SsCreator(context.Client.CreateVM).
			WithDbFunc(dbFunc(vmID, "vm")).
			WithPublic(&vmPublic).
			Exec()

		if err != nil {
			_ = resources.SsDeleter(context.Client.DeleteServiceOffering).
				WithDbFunc(dbFunc(vmID, "serviceOffering")).
				Exec()

			return makeError(err)
		}

		context.VM.Subsystems.CS.VM = vmPublic
	}

	// Port-forwarding rules
	for _, pfrPublic := range context.Generator.PFRs() {
		if pfrPublic.PublicPort == 0 {
			pfrPublic.PublicPort, err = context.Client.GetFreePort(
				context.Zone.PortRange.Start,
				context.Zone.PortRange.End,
			)

			if err != nil {
				return makeError(err)
			}
		}

		err = resources.SsCreator(context.Client.CreatePortForwardingRule).
			WithDbFunc(dbFunc(vmID, "portForwardingRuleMap."+pfrPublic.Name)).
			WithPublic(&pfrPublic).
			Exec()

		if err != nil {
			return makeError(err)
		}
	}

	return nil
}

func Delete(id string) error {
	log.Println("deleting cs for", id)

	makeError := func(err error) error {
		return fmt.Errorf("failed to delete cs for vm %s. details: %w", id, err)
	}

	context, err := NewContext(id)
	if err != nil {
		if errors.Is(err, base.VmDeletedErr) {
			return nil
		}

		return makeError(err)
	}

	for mapName, pfr := range context.VM.Subsystems.CS.GetPortForwardingRuleMap() {
		err = resources.SsDeleter(context.Client.DeletePortForwardingRule).
			WithResourceID(pfr.ID).
			WithDbFunc(dbFunc(id, "portForwardingRuleMap."+mapName)).
			Exec()
	}

	err = resources.SsDeleter(context.Client.DeleteVM).
		WithResourceID(context.VM.Subsystems.CS.VM.ID).
		WithDbFunc(dbFunc(id, "vm")).
		Exec()

	if err != nil {
		return makeError(err)
	}

	err = resources.SsDeleter(context.Client.DeleteServiceOffering).
		WithResourceID(context.VM.Subsystems.CS.ServiceOffering.ID).
		WithDbFunc(dbFunc(id, "serviceOffering")).
		Exec()

	if err != nil {
		return makeError(err)
	}

	return nil
}

func Update(vmID string, updateParams *vmModel.UpdateParams) error {
	makeError := func(err error) error {
		return fmt.Errorf("failed to update cs for vm %s. details: %w", vmID, err)
	}

	context, err := NewContext(vmID)
	if err != nil {
		if errors.Is(err, base.VmDeletedErr) {
			return nil
		}

		return makeError(err)
	}

	// port-forwarding rule
	if updateParams.Ports != nil {
		pfrs := context.Generator.PFRs()

		for _, currentPfr := range context.VM.Subsystems.CS.GetPortForwardingRuleMap() {
			if slices.IndexFunc(pfrs, func(p csModels.PortForwardingRulePublic) bool { return p.Name == currentPfr.Name }) == -1 {
				err = resources.SsDeleter(context.Client.DeletePortForwardingRule).
					WithResourceID(currentPfr.ID).
					WithDbFunc(dbFunc(vmID, "portForwardingRuleMap."+currentPfr.Name)).
					Exec()
			}
		}

		for _, pfrPublic := range pfrs {
			if _, ok := context.VM.Subsystems.CS.PortForwardingRuleMap[pfrPublic.Name]; !ok {
				if pfrPublic.PublicPort == 0 {
					pfrPublic.PublicPort, err = context.Client.GetFreePort(
						context.Zone.PortRange.Start,
						context.Zone.PortRange.End,
					)

					if err != nil {
						return makeError(err)
					}
				}

				err = resources.SsCreator(context.Client.CreatePortForwardingRule).
					WithDbFunc(dbFunc(vmID, "portForwardingRuleMap."+pfrPublic.Name)).
					WithPublic(&pfrPublic).
					Exec()

				if err != nil {
					return makeError(err)
				}
			}
		}
	}

	// service offering
	var soID *string
	if so := &context.VM.Subsystems.CS.ServiceOffering; service.Created(so) {
		var requiresUpdate bool
		if updateParams.CpuCores != nil {
			requiresUpdate = true
		}

		if updateParams.RAM != nil {
			requiresUpdate = true
		}

		if requiresUpdate {
			err = resources.SsDeleter(context.Client.DeleteServiceOffering).
				WithResourceID(so.ID).
				WithDbFunc(dbFunc(vmID, "serviceOffering")).
				Exec()

			if err != nil {
				return makeError(err)
			}

			err = context.Refresh()
			if err != nil {
				if errors.Is(err, base.VmDeletedErr) {
					return nil
				}

				return makeError(err)
			}

			for _, soPublic := range context.Generator.SOs() {
				err = resources.SsCreator(context.Client.CreateServiceOffering).
					WithDbFunc(dbFunc(vmID, "serviceOffering")).
					WithPublic(&soPublic).
					Exec()

				if err != nil {
					return makeError(err)
				}

				soID = &soPublic.ID
			}
		} else {
			soID = &so.ID
		}
	} else {
		for _, soPublic := range context.Generator.SOs() {
			err = resources.SsCreator(context.Client.CreateServiceOffering).
				WithDbFunc(dbFunc(vmID, "serviceOffering")).
				WithPublic(&soPublic).
				Exec()

			if err != nil {
				return makeError(err)
			}

			soID = &soPublic.ID
		}
	}

	serviceOfferingUpdated := false

	// make sure the vm is using the latest service offering
	if soID != nil && context.VM.Subsystems.CS.VM.ServiceOfferingID != *soID {
		serviceOfferingUpdated = true

		deferFunc, err := stopVmIfRunning(context)
		if err != nil {
			return makeError(err)
		}

		defer deferFunc()
	}

	if updateParams.Name != nil || serviceOfferingUpdated {
		err = context.Refresh()
		if err != nil {
			if errors.Is(err, base.VmDeletedErr) {
				return nil
			}

			return makeError(err)
		}

		vms := context.Generator.VMs()
		for _, vmPublic := range vms {
			err = resources.SsUpdater(context.Client.UpdateVM).
				WithPublic(&vmPublic).
				WithDbFunc(dbFunc(vmID, "vm")).
				Exec()

			if err != nil {
				return makeError(err)
			}
		}
	}

	return nil
}

func EnsureOwner(id, oldOwnerID string) error {
	// nothing needs to be done, but the method is kept as there is a project for networks,
	// and this could be implemented as user-specific networks

	return nil
}

func Repair(id string) error {
	makeError := func(err error) error {
		return fmt.Errorf("failed to repair cs %s. details: %w", id, err)
	}

	context, err := NewContext(id)
	if err != nil {
		if errors.Is(err, base.VmDeletedErr) {
			return nil
		}

		return makeError(err)
	}

	// Service offering
	so := context.Generator.SOs()[0]
	err = resources.SsRepairer(
		context.Client.ReadServiceOffering,
		context.Client.CreateServiceOffering,
		context.Client.UpdateServiceOffering,
		context.Client.DeleteServiceOffering,
	).WithResourceID(so.ID).WithGenPublic(&so).WithDbFunc(dbFunc(id, "serviceOffering")).Exec()

	if err != nil {
		return makeError(err)
	}

	err = context.Refresh()
	if err != nil {
		if errors.Is(err, base.VmDeletedErr) {
			return nil
		}

		return makeError(err)
	}

	// VM
	vm := context.Generator.VMs()[0]
	status, err := context.Client.GetVmStatus(context.VM.Subsystems.CS.VM.ID)
	if err != nil {
		return makeError(err)
	}

	// only repair if the vm is stopped to prevent downtime for the user
	if status == "Stopped" {
		var gpu *gpu2.GPU
		if gpuID := context.VM.GetGpuID(); gpuID != nil {
			gpu, err = gpu2.New().GetByID(*gpuID)
			if err != nil {
				return makeError(err)
			}
		}

		if gpu != nil {
			vm.ExtraConfig = CreateExtraConfig(gpu)
		}

		err = resources.SsRepairer(
			context.Client.ReadVM,
			context.Client.CreateVM,
			context.Client.UpdateVM,
			func(id string) error { return nil },
		).WithResourceID(vm.ID).WithGenPublic(&vm).WithDbFunc(dbFunc(id, "vm")).Exec()

		if err != nil {
			return makeError(err)
		}
	}

	// Port-forwarding rules
	pfrs := context.Generator.PFRs()
	for mapName, pfr := range context.VM.Subsystems.CS.GetPortForwardingRuleMap() {
		idx := slices.IndexFunc(pfrs, func(p csModels.PortForwardingRulePublic) bool { return p.Name == mapName })
		if idx == -1 {
			err = resources.SsDeleter(context.Client.DeletePortForwardingRule).
				WithResourceID(pfr.ID).
				WithDbFunc(dbFunc(id, "portForwardingRuleMap."+pfr.Name)).
				Exec()

			if err != nil {
				return makeError(err)
			}

			continue
		}
	}
	for _, pfr := range pfrs {
		err = resources.SsRepairer(
			context.Client.ReadPortForwardingRule,
			context.Client.CreatePortForwardingRule,
			context.Client.UpdatePortForwardingRule,
			context.Client.DeletePortForwardingRule,
		).WithResourceID(pfr.ID).WithDbFunc(dbFunc(id, "portForwardingRuleMap."+pfr.Name)).WithGenPublic(&pfr).Exec()
	}

	return nil
}

func DoCommand(csVmID string, gpuID *string, command, zoneName string) error {
	makeError := func(err error) error {
		return fmt.Errorf("failed to execute command %s for cs vm %s. details: %w", command, csVmID, err)
	}

	context, err := NewContextWithoutVM(zoneName)
	if err != nil {
		return makeError(err)
	}

	var requiredHost *string
	if gpuID != nil {
		requiredHost, err = GetRequiredHost(*gpuID)
		if err != nil {
			return makeError(err)
		}
	}

	err = context.Client.DoVmCommand(csVmID, requiredHost, commands.Command(command))
	if err != nil {
		return makeError(err)
	}

	return nil
}

func CanStart(csVmID, hostName, zoneName string) error {
	makeError := func(err error) error {
		return fmt.Errorf("failed to check if cs vm %s can be started on host %s. details: %w", csVmID, hostName, err)
	}

	context, err := NewContextWithoutVM(zoneName)
	if err != nil {
		return makeError(err)
	}

	hasCapacity, err := context.Client.HasCapacity(csVmID, hostName)
	if err != nil {
		return makeError(err)
	}

	if !hasCapacity {
		return service_errors.VmTooLargeErr
	}

	err = HostInCorrectState(hostName, context.Zone)
	if err != nil {
		if errors.Is(err, service_errors.HostNotAvailableErr) {
			return service_errors.VmTooLargeErr
		}

		return makeError(err)
	}

	return nil
}

func GetHostByName(hostName string, zone string) (*csModels.HostPublic, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to get host %s. details: %w", hostName, err)
	}

	context, err := NewContextWithoutVM(zone)
	if err != nil {
		return nil, makeError(err)
	}

	host, err := context.Client.ReadHostByName(hostName)
	if err != nil {
		return nil, makeError(err)
	}

	return host, nil
}

func HostInCorrectState(hostName string, zone *configModels.VmZone) error {
	makeError := func(err error) error {
		return fmt.Errorf("failed to check if host %s is in correct state. details: %w", zone.Name, err)
	}

	context, err := NewContextWithoutVM(zone.Name)
	if err != nil {
		return makeError(err)
	}

	host, err := context.Client.ReadHostByName(hostName)
	if err != nil {
		return makeError(err)
	}

	if host.State != "Up" || host.ResourceState != "Enabled" {
		return service_errors.HostNotAvailableErr
	}

	return nil
}

func GetConfiguration(zone string) (*csModels.ConfigurationPublic, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to get configuration. details: %w", err)
	}

	context, err := NewContextWithoutVM(zone)
	if err != nil {
		return nil, makeError(err)
	}

	configuration, err := context.Client.ReadConfiguration()
	if err != nil {
		return nil, makeError(err)
	}

	return configuration, nil
}

func stopVmIfRunning(context *Context) (func(), error) {
	// turn it off if it is on, but remember the status
	status, err := context.Client.GetVmStatus(context.VM.Subsystems.CS.VM.ID)
	if err != nil {
		return nil, err
	}

	if status == "Running" {
		err = context.Client.DoVmCommand(context.VM.Subsystems.CS.VM.ID, nil, commands.Stop)
		if err != nil {
			return nil, err
		}
	}

	return func() {
		// turn it on if it was on
		if status == "Running" {
			var requiredHost *string
			if gpuID := context.VM.GetGpuID(); gpuID != nil {
				requiredHost, err = GetRequiredHost(*gpuID)
				if err != nil {
					log.Println("failed to get required host for vm", context.VM.Name, "in zone", context.Zone.Name, ". details:", err)
					return
				}
			}

			err = context.Client.DoVmCommand(context.VM.Subsystems.CS.VM.ID, requiredHost, commands.Start)
			if err != nil {
				log.Println("failed to start vm", context.VM.Name, "in zone", context.Zone.Name, ". details:", err)
				return
			}
		}
	}, nil
}

func dbFunc(vmID, key string) func(interface{}) error {
	return func(data interface{}) error {
		if data == nil {
			return vmModel.New().DeleteSubsystemByID(vmID, "cs."+key)
		}
		return vmModel.New().UpdateSubsystemByID(vmID, "cs."+key, data)
	}
}
