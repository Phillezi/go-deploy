package k8s

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go-deploy/pkg/subsystems/k8s/keys"
	"go-deploy/pkg/subsystems/k8s/models"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

func (client *Client) ReadJob(id string) (*models.JobPublic, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to read k8s job %s. details: %w", id, err)
	}

	if id == "" {
		return nil, nil
	}

	if client.Namespace == "" {
		return nil, nil
	}

	list, err := client.K8sClient.BatchV1().Jobs(client.Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, makeError(err)
	}

	for _, item := range list.Items {
		idLabel := GetLabel(item.ObjectMeta.Labels, keys.ManifestLabelID)
		if idLabel == id {
			return models.CreateJobPublicFromRead(&item), nil
		}
	}

	return nil, nil
}

func (client *Client) CreateJob(public *models.JobPublic) (string, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to create k8s job %s. details: %w", public.Name, err)
	}

	if public.Namespace == "" {
		return "", nil
	}

	list, err := client.K8sClient.BatchV1().Jobs(public.Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return "", makeError(err)
	}

	for _, item := range list.Items {
		if FindLabel(item.ObjectMeta.Labels, keys.ManifestLabelName, public.Name) {
			idLabel := GetLabel(item.ObjectMeta.Labels, keys.ManifestLabelID)
			if idLabel != "" {
				return idLabel, nil
			}
		}
	}

	public.ID = uuid.New().String()
	public.CreatedAt = time.Now()

	manifest := CreateJobManifest(public)
	_, err = client.K8sClient.BatchV1().Jobs(public.Namespace).Create(context.TODO(), manifest, metav1.CreateOptions{})
	if err != nil {
		return "", makeError(err)
	}

	return public.ID, nil
}

func (client *Client) DeleteJob(id string) error {
	makeError := func(err error) error {
		return fmt.Errorf("failed to delete k8s job %s. details: %w", id, err)
	}

	if id == "" {
		return nil
	}

	if client.Namespace == "" {
		return nil
	}

	list, err := client.K8sClient.BatchV1().Jobs(client.Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return makeError(err)
	}

	for _, item := range list.Items {
		idLabel := GetLabel(item.ObjectMeta.Labels, keys.ManifestLabelID)
		if idLabel == id {
			err = client.K8sClient.BatchV1().Jobs(client.Namespace).Delete(context.TODO(), item.Name, metav1.DeleteOptions{})
			if err != nil {
				return makeError(err)
			}
			return nil
		}
	}

	return nil
}