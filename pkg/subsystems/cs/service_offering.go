package cs

import (
	"fmt"
	"go-deploy/pkg/imp/cloudstack"
	"go-deploy/pkg/subsystems/cs/models"
	"log"
	"strings"
)

func (client *Client) ReadServiceOffering(id string) (*models.ServiceOfferingPublic, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to read service offering %s. details: %w", id, err)
	}

	if id == "" {
		log.Println("cs service offering id not supplied when updating. assuming it was deleted")
		return nil, nil
	}

	serviceOffering, _, err := client.CsClient.ServiceOffering.GetServiceOfferingByID(id)
	if err != nil {
		errString := err.Error()
		if !strings.Contains(errString, "No match found for") {
			return nil, makeError(err)
		}
	}

	if serviceOffering == nil {
		return nil, nil
	}

	return models.CreateServiceOfferingPublicFromGet(serviceOffering), nil
}

func (client *Client) CreateServiceOffering(public *models.ServiceOfferingPublic) (string, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to create service offering. details: %w", err)
	}

	params := client.CsClient.ServiceOffering.NewListServiceOfferingsParams()
	params.SetName(public.Name)
	params.SetListall(true)

	serviceOfferings, err := client.CsClient.ServiceOffering.ListServiceOfferings(params)
	if err != nil {
		return "", makeError(err)
	}

	if len(serviceOfferings.ServiceOfferings) > 0 {
		return serviceOfferings.ServiceOfferings[0].Id, nil
	}

	createParams := cloudstack.CreateServiceOfferingParams{}
	createParams.SetName(public.Name)
	createParams.SetDisplaytext(public.Name)
	createParams.SetCpunumber(public.CpuCores)
	createParams.SetCpuspeed(1)
	createParams.SetMemory(public.RAM * 1024)
	createParams.SetOfferha(false)
	createParams.SetLimitcpuuse(false)
	createParams.SetRootdisksize(int64(public.DiskSize))

	serviceOffering, err := client.CsClient.ServiceOffering.CreateServiceOffering(&createParams)
	if err != nil {
		return "", makeError(err)
	}

	return serviceOffering.Id, nil
}

func (client *Client) UpdateServiceOffering(public *models.ServiceOfferingPublic) error {
	makeError := func(err error) error {
		return fmt.Errorf("failed to update service offering %s. details: %w", public.ID, err)
	}

	if public.ID == "" {
		log.Println("cs service offering id not supplied when updating. assuming it was deleted")
		return nil
	}

	updateParams := client.CsClient.ServiceOffering.NewUpdateServiceOfferingParams(public.ID)
	updateParams.SetName(public.Name)
	updateParams.SetDisplaytext(public.Name)

	_, err := client.CsClient.ServiceOffering.UpdateServiceOffering(updateParams)
	if err != nil {
		return makeError(err)
	}

	return nil
}

func (client *Client) DeleteServiceOffering(id string) error {
	makeError := func(err error) error {
		return fmt.Errorf("failed to delete service offering %s. details: %w", id, err)
	}

	if id == "" {
		log.Println("cs service offering id not supplied when updating. assuming it was deleted")
		return nil
	}

	params := client.CsClient.ServiceOffering.NewDeleteServiceOfferingParams(id)

	_, err := client.CsClient.ServiceOffering.DeleteServiceOffering(params)
	if err != nil {
		return makeError(err)
	}

	return nil
}
