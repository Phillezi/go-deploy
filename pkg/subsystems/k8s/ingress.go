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

func (client *Client) ReadIngress(id string) (*models.IngressPublic, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to read k8s ingress %s. details: %w", id, err)
	}

	if id == "" {
		log.Println("no id supplied when reading k8s ingress. assuming it was deleted")
		return nil, nil
	}

	list, err := client.K8sClient.NetworkingV1().Ingresses(client.Namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: fmt.Sprintf("%s=%s", keys.ManifestLabelID, id),
	})
	if err != nil {
		return nil, makeError(err)
	}

	if len(list.Items) > 0 {
		return models.CreateIngressPublicFromRead(&list.Items[0]), nil
	}

	return nil, nil
}

func (client *Client) CreateIngress(public *models.IngressPublic) (*models.IngressPublic, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to create k8s ingress %s. details: %w", public.Name, err)
	}

	list, err := client.K8sClient.NetworkingV1().Ingresses(public.Namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: fmt.Sprintf("%s=%s", keys.ManifestLabelID, public.ID),
	})
	if err != nil {
		return nil, err
	}

	if len(list.Items) > 0 {
		return models.CreateIngressPublicFromRead(&list.Items[0]), nil
	}

	public.ID = uuid.New().String()
	public.CreatedAt = time.Now()

	manifest := CreateIngressManifest(public)
	res, err := client.K8sClient.NetworkingV1().Ingresses(public.Namespace).Create(context.TODO(), manifest, metav1.CreateOptions{})
	if err != nil {
		return nil, makeError(err)
	}

	return models.CreateIngressPublicFromRead(res), nil
}

func (client *Client) UpdateIngress(public *models.IngressPublic) (*models.IngressPublic, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to update k8s ingress %s. details: %w", public.Name, err)
	}

	if public.ID == "" {
		log.Println("no id supplied when updating k8s ingress. assuming it was deleted")
		return nil, nil
	}

	list, err := client.K8sClient.NetworkingV1().Ingresses(public.Namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: fmt.Sprintf("%s=%s", keys.ManifestLabelID, public.ID),
	})
	if err != nil {
		return nil, makeError(err)
	}

	if len(list.Items) > 0 {
		manifest := CreateIngressManifest(public)
		res, err := client.K8sClient.NetworkingV1().Ingresses(public.Namespace).Update(context.TODO(), manifest, metav1.UpdateOptions{})
		if err != nil {
			return nil, makeError(err)
		}
		return models.CreateIngressPublicFromRead(res), nil
	}

	log.Println("k8s ingress", public.Name, "not found when updating. assuming it was deleted")
	return nil, nil
}

func (client *Client) DeleteIngress(id string) error {
	makeError := func(err error) error {
		return fmt.Errorf("failed to delete k8s ingress %s. details: %w", id, err)
	}

	if id == "" {
		log.Println("no id supplied when deleting k8s ingress. assuming it was deleted")
		return nil
	}

	list, err := client.K8sClient.NetworkingV1().Ingresses(client.Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return makeError(err)
	}

	for _, item := range list.Items {
		idLabel := GetLabel(item.ObjectMeta.Labels, keys.ManifestLabelID)
		if idLabel == id {
			err = client.K8sClient.NetworkingV1().Ingresses(client.Namespace).Delete(context.TODO(), item.Name, metav1.DeleteOptions{})
			if err != nil {
				return makeError(err)
			}

			return nil
		}
	}

	return nil
}
