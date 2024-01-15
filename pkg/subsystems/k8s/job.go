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

// ReadJob reads a Job from Kubernetes.
func (client *Client) ReadJob(id string) (*models.JobPublic, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to read k8s job %s. details: %w", id, err)
	}

	if id == "" {
		log.Println("no id supplied when reading k8s job. assuming it was deleted")
		return nil, nil
	}

	list, err := client.K8sClient.BatchV1().Jobs(client.Namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: fmt.Sprintf("%s=%s", keys.ManifestLabelID, id),
	})
	if err != nil {
		return nil, makeError(err)
	}

	if len(list.Items) > 0 {
		return models.CreateJobPublicFromRead(&list.Items[0]), nil
	}

	return nil, nil
}

// CreateJob creates a Job in Kubernetes.
func (client *Client) CreateJob(public *models.JobPublic) (*models.JobPublic, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to create k8s job %s. details: %w", public.Name, err)
	}

	job, err := client.K8sClient.BatchV1().Jobs(public.Namespace).Get(context.TODO(), public.Name, metav1.GetOptions{})
	if err != nil && !IsNotFoundErr(err) {
		return nil, makeError(err)
	}

	if err == nil {
		return models.CreateJobPublicFromRead(job), nil
	}

	public.ID = uuid.New().String()
	public.CreatedAt = time.Now()

	manifest := CreateJobManifest(public)
	res, err := client.K8sClient.BatchV1().Jobs(public.Namespace).Create(context.TODO(), manifest, metav1.CreateOptions{})
	if err != nil {
		return nil, makeError(err)
	}

	return models.CreateJobPublicFromRead(res), nil
}

// DeleteJob deletes a Job in Kubernetes.
func (client *Client) DeleteJob(id string) error {
	makeError := func(err error) error {
		return fmt.Errorf("failed to delete k8s job %s. details: %w", id, err)
	}

	if id == "" {
		log.Println("no id supplied when deleting k8s job. assuming it was deleted")
		return nil
	}

	list, err := client.K8sClient.BatchV1().Jobs(client.Namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: fmt.Sprintf("%s=%s", keys.ManifestLabelID, id),
	})
	if err != nil {
		return makeError(err)
	}

	for _, job := range list.Items {
		err = client.K8sClient.BatchV1().Jobs(client.Namespace).Delete(context.TODO(), job.Name, metav1.DeleteOptions{
			PropagationPolicy: &[]metav1.DeletionPropagation{metav1.DeletePropagationBackground}[0],
		})
		if err != nil && !IsNotFoundErr(err) {
			return makeError(err)
		}
	}

	return nil
}
