package k8s

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go-deploy/pkg/subsystems/k8s/keys"
	"go-deploy/pkg/subsystems/k8s/models"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"time"
)

func (client *Client) ReadPV(id string) (*models.PvPublic, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to read k8s pv %s. details: %w", id, err)
	}

	if id == "" {
		log.Println("no id supplied when reading k8s pv. assuming it was deleted")
		return nil, nil
	}

	list, err := client.K8sClient.CoreV1().PersistentVolumes().List(context.TODO(), metav1.ListOptions{
		LabelSelector: fmt.Sprintf("%s=%s", keys.ManifestLabelID, id),
	})
	if err != nil {
		return nil, makeError(err)
	}

	if len(list.Items) > 0 {
		return models.CreatePvPublicFromRead(&list.Items[0]), nil
	}

	return nil, nil
}

func (client *Client) CreatePV(public *models.PvPublic) (*models.PvPublic, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to create k8s pv %s. details: %w", public.Name, err)
	}

	pv, err := client.K8sClient.CoreV1().PersistentVolumes().Get(context.TODO(), public.Name, metav1.GetOptions{})
	if err != nil && !IsNotFoundErr(err) {
		return nil, makeError(err)
	}

	if pv != nil {
		return models.CreatePvPublicFromRead(pv), nil
	}

	public.ID = uuid.New().String()
	public.CreatedAt = time.Now()

	manifest := CreatePvManifest(public)
	res, err := client.K8sClient.CoreV1().PersistentVolumes().Create(context.TODO(), manifest, metav1.CreateOptions{})
	if err != nil {
		return nil, makeError(err)
	}

	return models.CreatePvPublicFromRead(res), nil
}

func (client *Client) DeletePV(id string) error {
	makeError := func(err error) error {
		return fmt.Errorf("failed to delete k8s pv %s. details: %w", id, err)
	}

	if id == "" {
		log.Println("no id supplied when deleting k8s pv. assuming it was deleted")
		return nil
	}

	list, err := client.K8sClient.CoreV1().PersistentVolumes().List(context.TODO(), metav1.ListOptions{
		LabelSelector: fmt.Sprintf("%s=%s", keys.ManifestLabelID, id),
	})
	if err != nil {
		return makeError(err)
	}

	for _, item := range list.Items {
		err = client.K8sClient.CoreV1().PersistentVolumes().Delete(context.TODO(), item.Name, metav1.DeleteOptions{})
		if err != nil {
			return makeError(err)
		}
	}

	return nil
}
