package conf

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"regexp"

	"go-deploy/pkg/imp/cloudstack"
)

type Environment struct {
	Port          int    `yaml:"port"`
	ExternalUrl   string `yaml:"externalUrl"`
	Manager       string `yaml:"manager"`
	SessionSecret string `yaml:"sessionSecret"`

	GPU struct {
		PrivilegedGPUs []string `yaml:"privilegedGpus"`
		ExcludedHosts  []string `yaml:"excludedHosts"`
	} `yaml:"gpu"`

	DockerRegistry struct {
		Url         string `yaml:"url"`
		Placeholder struct {
			Project    string `yaml:"project"`
			Repository string `yaml:"repository"`
		} `yaml:"placeholder"`
	} `yaml:"dockerRegistry"`

	App struct {
		ParentDomain string `yaml:"parentDomain"`
		Port         int    `yaml:"port"`
		Prefix       string `yaml:"prefix"`
		DefaultQuota int    `yaml:"defaultQuota"`
	} `yaml:"app"`

	VM struct {
		ParentDomain      string `yaml:"parentDomain"`
		DefaultQuota      int    `yaml:"defaultQuota"`
		AdminSshPublicKey string `yaml:"adminSshPublicKey"`
	} `yaml:"vm"`

	Keycloak struct {
		Url        string `yaml:"url"`
		Realm      string `yaml:"realm"`
		AdminGroup string `yaml:"adminGroup"`
		GpuGroup   string `yaml:"gpuGroup"`
	} `yaml:"keycloak"`

	DB struct {
		Url  string `yaml:"url"`
		Name string `yaml:"name"`
	} `yaml:"db"`

	CS struct {
		Url    string `yaml:"url"`
		ApiKey string `yaml:"apiKey"`
		Secret string `yaml:"secret"`
	} `yaml:"cs"`

	PfSense struct {
		User      string `yaml:"user"`
		Password  string `yaml:"password"`
		Url       string `yaml:"url"`
		PublicIP  string `yaml:"publicIp"`
		PortRange struct {
			Start int `yaml:"start"`
			End   int `yaml:"end"`
		} `yaml:"portRange"`
	} `yaml:"pfSense"`

	Landing struct {
		Url      string `yaml:"url"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		ClientID string `yaml:"clientId"`
	} `yaml:"landing"`

	K8s struct {
		Name   string `yaml:"name"`
		URL    string `yaml:"url"`
		Client *kubernetes.Clientset
	} `yaml:"k8s"`

	NPM struct {
		Url      string `yaml:"url"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"npm"`

	Harbor struct {
		Url           string `yaml:"url"`
		User          string `yaml:"user"`
		Password      string `yaml:"password"`
		WebhookSecret string `yaml:"webhookSecret"`
	} `yaml:"harbor"`
}

var Env Environment

func SetupEnvironment() {
	makeError := func(err error) error {
		return fmt.Errorf("failed to setup environment. details: %s", err)
	}

	filepath, found := os.LookupEnv("DEPLOY_CONFIG_FILE")
	if !found {
		log.Fatalln(makeError(fmt.Errorf("config file not found. please set DEPLOY_CONFIG_FILE environment variable")))
	}

	log.Println("reading config from", filepath)
	yamlFile, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf(makeError(err).Error())
	}

	err = yaml.Unmarshal(yamlFile, &Env)
	if err != nil {
		log.Fatalf(makeError(err).Error())
	}

	err = setupK8sClusters()
	if err != nil {
		log.Fatalln(makeError(err))
	}

	log.Println("config loaded")
}

func setupK8sClusters() error {
	makeError := func(err error) error {
		return fmt.Errorf("failed to setup k8s clusters. details: %s", err)
	}

	log.Println("fetching available k8s clusters")

	csClient := cloudstack.NewAsyncClient(
		Env.CS.Url,
		Env.CS.ApiKey,
		Env.CS.Secret,
		true,
	)

	listClusterParams := csClient.Kubernetes.NewListKubernetesClustersParams()
	listClusterParams.SetListall(true)
	clusters, err := csClient.Kubernetes.ListKubernetesClusters(listClusterParams)
	if err != nil {
		log.Fatalln(makeError(err))
	}

	fetchConfig := func(name string, publicUrl string) string {

		log.Println("fetching k8s cluster config for", name)

		clusterIdx := -1
		for idx, cluster := range clusters.KubernetesClusters {
			if cluster.Name == name {
				clusterIdx = idx
				break
			}
		}

		if clusterIdx == -1 {
			log.Println("cluster", name, "not found")
			return ""
		}

		params := csClient.Kubernetes.NewGetKubernetesClusterConfigParams()
		params.SetId(clusters.KubernetesClusters[clusterIdx].Id)

		config, err := csClient.Kubernetes.GetKubernetesClusterConfig(params)
		if err != nil {
			log.Fatalln(makeError(err))
		}

		// use regex to replace the private ip in config.ConffigData 172.31.1.* with the public ip
		regex := regexp.MustCompile(`https://172.31.1.[0-9]+:6443`)
		config.Configdata = regex.ReplaceAllString(config.Configdata, publicUrl)

		return config.Configdata
	}

	configData := fetchConfig(Env.K8s.Name, Env.K8s.URL)
	if configData == "" {
		return makeError(fmt.Errorf("failed to fetch k8s cluster config"))
	}

	Env.K8s.Client, err = createClient([]byte(configData))
	if err != nil {
		return makeError(err)
	}

	log.Println("k8s clusters setup done")
	return nil
}

func createClient(configData []byte) (*kubernetes.Clientset, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to create k8s client. details: %s", err)
	}

	kubeConfig, err := clientcmd.RESTConfigFromKubeConfig(configData)
	if err != nil {
		return nil, makeError(err)
	}

	k8sClient, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		return nil, makeError(err)
	}

	return k8sClient, nil
}
