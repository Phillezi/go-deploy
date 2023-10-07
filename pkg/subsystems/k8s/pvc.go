package k8s

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go-deploy/pkg/subsystems/k8s/keys"
	"go-deploy/pkg/subsystems/k8s/models"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"time"
)

func (client *Client) ReadPVC(id string) (*models.PvcPublic, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to read k8s pvc %s. details: %w", id, err)
	}

	if id == "" {
		log.Println("no id supplied when reading k8s pvc. assuming it was deleted")
		return nil, nil
	}

	list, err := client.K8sClient.CoreV1().PersistentVolumeClaims(client.Namespace).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		return nil, makeError(err)
	}

	for _, item := range list.Items {
		idLabel := GetLabel(item.ObjectMeta.Labels, keys.ManifestLabelID)
		if idLabel == id {
			return models.CreatePvcPublicFromRead(&item), nil
		}
	}

	return nil, nil
}

func (client *Client) CreatePVC(public *models.PvcPublic) (*models.PvcPublic, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to create k8s pvc %s. details: %w", public.Name, err)
	}

	if public.Name == "" {
		log.Println("no name supplied when creating k8s pvc. assuming it was deleted")
		return nil, nil
	}

	list, err := client.K8sClient.CoreV1().PersistentVolumeClaims(public.Namespace).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		return nil, makeError(err)
	}

	for _, item := range list.Items {
		if FindLabel(item.ObjectMeta.Labels, keys.ManifestLabelName, public.Name) {
			idLabel := GetLabel(item.ObjectMeta.Labels, keys.ManifestLabelID)
			if idLabel != "" {
				return models.CreatePvcPublicFromRead(&item), nil
			}
		}
	}

	public.ID = uuid.New().String()
	public.CreatedAt = time.Now()

	manifest := CreatePvcManifest(public)
	res, err := client.K8sClient.CoreV1().PersistentVolumeClaims(public.Namespace).Create(context.TODO(), manifest, v1.CreateOptions{})
	if err != nil {
		return nil, makeError(err)
	}

	return models.CreatePvcPublicFromRead(res), nil
}

func (client *Client) DeletePVC(id string) error {
	makeError := func(err error) error {
		return fmt.Errorf("failed to delete k8s pvc %s. details: %w", id, err)
	}

	if id == "" {
		log.Println("no id supplied when deleting k8s pvc. assuming it was deleted")
		return nil
	}

	list, err := client.K8sClient.CoreV1().PersistentVolumeClaims(client.Namespace).List(context.TODO(), v1.ListOptions{
		LabelSelector: fmt.Sprintf("%s=%s", keys.ManifestLabelID, id),
	})
	if err != nil {
		return makeError(err)
	}

	for _, item := range list.Items {
		err = client.K8sClient.CoreV1().PersistentVolumeClaims(client.Namespace).Delete(context.TODO(), item.Name, v1.DeleteOptions{})
		if err != nil {
			return makeError(err)
		}
	}

	return nil
}
