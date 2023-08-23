package k8s

import (
	"go-deploy/pkg/subsystems/k8s/keys"
	"go-deploy/pkg/subsystems/k8s/models"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func int32Ptr(i int32) *int32 { return &i }

func CreateNamespaceManifest(public *models.NamespacePublic) *apiv1.Namespace {
	return &apiv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: public.FullName,
			Labels: map[string]string{
				keys.ManifestLabelID:   public.ID,
				keys.ManifestLabelName: public.Name,
			},
		},
	}
}

func CreateDeploymentManifest(public *models.DeploymentPublic) *appsv1.Deployment {
	var envs []apiv1.EnvVar
	for _, env := range public.EnvVars {
		envs = append(envs, env.ToK8sEnvVar())
	}

	limits := createResourceList(public.Resources.Limits.CPU, public.Resources.Limits.Memory)
	requests := createResourceList(public.Resources.Requests.CPU, public.Resources.Requests.Memory)

	var lifecycle *apiv1.Lifecycle
	if len(public.InitCommands) > 0 {
		lifecycle = &apiv1.Lifecycle{
			PostStart: &apiv1.LifecycleHandler{
				Exec: &apiv1.ExecAction{
					Command: public.InitCommands,
				},
			},
		}
	}

	volumes := make([]apiv1.Volume, len(public.Volumes))
	for i, volume := range public.Volumes {
		volumes[i] = apiv1.Volume{
			Name: volume.Name,
			VolumeSource: apiv1.VolumeSource{
				PersistentVolumeClaim: &apiv1.PersistentVolumeClaimVolumeSource{
					ClaimName: volume.PvcName,
				},
			},
		}
	}

	normalContainerMounts := make([]apiv1.VolumeMount, 0)
	initContainerMounts := make([]apiv1.VolumeMount, 0)

	for _, volume := range public.Volumes {
		if volume.Init {
			initContainerMounts = append(initContainerMounts, apiv1.VolumeMount{
				Name:      volume.Name,
				MountPath: volume.MountPath,
			})
		} else {
			normalContainerMounts = append(normalContainerMounts, apiv1.VolumeMount{
				Name:      volume.Name,
				MountPath: volume.MountPath,
			})
		}
	}

	initContainers := make([]apiv1.Container, len(public.InitContainers))
	for i, initContainer := range public.InitContainers {
		initContainers[i] = apiv1.Container{
			Name:         initContainer.Name,
			Image:        initContainer.Image,
			Command:      initContainer.Command,
			Args:         initContainer.Args,
			VolumeMounts: initContainerMounts,
		}
	}

	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      public.Name,
			Namespace: public.Namespace,
			Labels: map[string]string{
				keys.ManifestLabelID:   public.ID,
				keys.ManifestLabelName: public.Name,
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					keys.ManifestLabelID: public.ID,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						keys.ManifestLabelID:   public.ID,
						keys.ManifestLabelName: public.Name,
					},
				},
				Spec: apiv1.PodSpec{
					Volumes: volumes,
					Containers: []apiv1.Container{
						{
							Name:    public.Name,
							Image:   public.DockerImage,
							Command: public.Command,
							Args:    public.Args,
							Env:     envs,
							Resources: apiv1.ResourceRequirements{
								Limits:   limits,
								Requests: requests,
							},
							Lifecycle:    lifecycle,
							VolumeMounts: normalContainerMounts,
						},
					},
					InitContainers: initContainers,
				},
			},
		},
	}
}
func CreateServiceManifest(public *models.ServicePublic) *apiv1.Service {
	return &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      public.Name,
			Namespace: public.Namespace,
			Labels: map[string]string{
				keys.ManifestLabelID:   public.ID,
				keys.ManifestLabelName: public.Name,
			},
		},
		Spec: apiv1.ServiceSpec{
			Ports: []apiv1.ServicePort{
				{
					Name:       "app-port",
					Protocol:   "TCP",
					Port:       int32(public.Port),
					TargetPort: intstr.FromInt(public.TargetPort),
				},
			},
			Selector: map[string]string{
				keys.ManifestLabelName: public.Name,
			},
		},
		Status: apiv1.ServiceStatus{},
	}
}

func CreateIngressManifest(public *models.IngressPublic) *networkingv1.Ingress {
	rules := make([]networkingv1.IngressRule, len(public.Hosts))
	for idx, host := range public.Hosts {
		rules[idx] = networkingv1.IngressRule{
			Host: host,
			IngressRuleValue: networkingv1.IngressRuleValue{
				HTTP: &networkingv1.HTTPIngressRuleValue{
					Paths: []networkingv1.HTTPIngressPath{
						{
							Path:     "/",
							PathType: pathTypeAddr("Prefix"),
							Backend: networkingv1.IngressBackend{
								Service: &networkingv1.IngressServiceBackend{
									Name: public.ServiceName,
									Port: networkingv1.ServiceBackendPort{
										Number: int32(public.ServicePort),
									},
								},
							},
						},
					},
				},
			},
		}
	}

	return &networkingv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      public.Name,
			Namespace: public.Namespace,
			Labels: map[string]string{
				keys.ManifestLabelID:   public.ID,
				keys.ManifestLabelName: public.Name,
			},
			Annotations: map[string]string{
				"kubernetes.io/ingress.class": public.IngressClass,
			},
		},
		Spec: networkingv1.IngressSpec{
			Rules: rules,
		},
	}
}

func CreatePvManifest(public *models.PvPublic) *apiv1.PersistentVolume {
	return &apiv1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name: public.Name,
			Labels: map[string]string{
				keys.ManifestLabelID:   public.ID,
				keys.ManifestLabelName: public.Name,
			},
		},
		Spec: apiv1.PersistentVolumeSpec{
			AccessModes: []apiv1.PersistentVolumeAccessMode{
				apiv1.ReadWriteMany,
			},
			Capacity: apiv1.ResourceList{
				apiv1.ResourceStorage: resource.MustParse(public.Capacity),
			},
			PersistentVolumeSource: apiv1.PersistentVolumeSource{
				NFS: &apiv1.NFSVolumeSource{
					Server:   public.NfsServer,
					Path:     public.NfsPath,
					ReadOnly: false,
				},
			},
		},
	}
}

func CreatePvcManifest(public *models.PvcPublic) *apiv1.PersistentVolumeClaim {
	return &apiv1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      public.Name,
			Namespace: public.Namespace,
			Labels: map[string]string{
				keys.ManifestLabelID:   public.ID,
				keys.ManifestLabelName: public.Name,
			},
		},
		Spec: apiv1.PersistentVolumeClaimSpec{
			AccessModes: []apiv1.PersistentVolumeAccessMode{
				apiv1.ReadWriteMany,
			},
			Resources: apiv1.ResourceRequirements{
				Requests: apiv1.ResourceList{
					apiv1.ResourceStorage: resource.MustParse(public.Capacity),
				},
				Limits: apiv1.ResourceList{
					apiv1.ResourceStorage: resource.MustParse(public.Capacity),
				},
			},
			VolumeName: public.PvName,
		},
	}
}

func CreateJobManifest(public *models.JobPublic) *v1.Job {
	volumes := make([]apiv1.Volume, len(public.Volumes))
	for i, volume := range public.Volumes {
		volumes[i] = apiv1.Volume{
			Name: volume.Name,
			VolumeSource: apiv1.VolumeSource{
				PersistentVolumeClaim: &apiv1.PersistentVolumeClaimVolumeSource{
					ClaimName: volume.PvcName,
				},
			},
		}
	}

	volumeMounts := make([]apiv1.VolumeMount, 0)
	for _, volume := range public.Volumes {
		if !volume.Init {
			volumeMounts = append(volumeMounts, apiv1.VolumeMount{
				Name:      volume.Name,
				MountPath: volume.MountPath,
			})
		}
	}

	return &v1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      public.Name,
			Namespace: public.Namespace,
			Labels: map[string]string{
				keys.ManifestLabelID:   public.ID,
				keys.ManifestLabelName: public.Name,
			},
		},
		Spec: v1.JobSpec{
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						keys.ManifestLabelID:   public.ID,
						keys.ManifestLabelName: public.Name,
					},
				},
				Spec: apiv1.PodSpec{
					RestartPolicy: apiv1.RestartPolicyNever,
					Volumes:       volumes,
					Containers: []apiv1.Container{
						{
							Name:            public.Name,
							Image:           public.Image,
							ImagePullPolicy: apiv1.PullIfNotPresent,
							Command:         public.Command,
							Args:            public.Args,
							VolumeMounts:    volumeMounts,
						},
					},
				},
			},
		},
	}
}

func pathTypeAddr(s string) *networkingv1.PathType {
	return (*networkingv1.PathType)(&s)
}

func createResourceList(cpu, memory string) apiv1.ResourceList {
	limits := apiv1.ResourceList{}

	cpuQuantity, err := resource.ParseQuantity(cpu)
	if err == nil {
		limits[apiv1.ResourceCPU] = cpuQuantity
	}

	memoryQuantity, err := resource.ParseQuantity(memory)
	if err == nil {
		limits[apiv1.ResourceMemory] = memoryQuantity
	}

	return limits
}
