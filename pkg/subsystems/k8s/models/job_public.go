package models

import (
	"go-deploy/pkg/subsystems/k8s/keys"
	v1 "k8s.io/api/batch/v1"
)

type JobPublic struct {
	ID        string   `json:"id" bson:"id"`
	Name      string   `json:"name" bson:"name"`
	Namespace string   `json:"namespace" bson:"namespace"`
	Image     string   `json:"image" bson:"image"`
	Command   []string `json:"command" bson:"command"`
	Args      []string `json:"args" bson:"args"`
	Volumes   []Volume `json:"volumes" bson:"volumes"`
}

func (job *JobPublic) Created() bool {
	return job.ID != ""
}

func CreateJobPublicFromRead(job *v1.Job) *JobPublic {

	var volumes []Volume

	for _, k8sVolume := range job.Spec.Template.Spec.Volumes {
		volumes = append(volumes, Volume{
			Name:    k8sVolume.Name,
			PvcName: k8sVolume.PersistentVolumeClaim.ClaimName,
		})
	}

	if len(job.Spec.Template.Spec.Containers) > 0 {
		firstContainer := job.Spec.Template.Spec.Containers[0]
		volumeMounts := firstContainer.VolumeMounts

		for _, volumeMount := range volumeMounts {
			// if we cannot find the volume mount in the volumes list, then it is not a volume we care about
			for _, volume := range volumes {
				if volume.Name == volumeMount.Name {
					volume.MountPath = volumeMount.MountPath
				}
			}
		}
	}

	// delete any k8sVolumes that does not have a mount path, they need to be recreated
	for i := len(volumes) - 1; i >= 0; i-- {
		if volumes[i].MountPath == "" {
			volumes = append(volumes[:i], volumes[i+1:]...)
		}
	}

	return &JobPublic{
		ID:      job.Labels[keys.ManifestLabelID],
		Name:    job.Labels[keys.ManifestLabelName],
		Image:   job.Spec.Template.Spec.Containers[0].Image,
		Command: job.Spec.Template.Spec.Containers[0].Command,
		Args:    job.Spec.Template.Spec.Containers[0].Args,
		Volumes: volumes,
	}
}
