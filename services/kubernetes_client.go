package services

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	k8sClt "github.com/ericchiang/k8s"
	"github.com/ghodss/yaml"
)

// KubernetesClientService struct
type KubernetesClientService struct {
	Client *k8sClt.Client
}

// NewKubernetesClientService constructor
func NewKubernetesClientService() KubernetesClientService {
	client, err := loadClient()
	if err != nil {
		log.Fatal(err)
	}
	return KubernetesClientService{Client: client}
}

// KubernetesClientServiceInstance instance
var KubernetesClientServiceInstance = NewKubernetesClientService()

// loadClient parses a kubeconfig from a file and returns a Kubernetes
// client. It does not support extensions or client auth providers.
func loadClient() (*k8sClt.Client, error) {
	data, err := ioutil.ReadFile(os.Getenv("KUBERNETES_CLIENT_CONFIG"))
	if err != nil {
		return nil, fmt.Errorf("read kubeconfig: %v", err)
	}

	// Unmarshal YAML into a Kubernetes config object.
	var config k8sClt.Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("unmarshal kubeconfig: %v", err)
	}
	return k8sClt.NewClient(&config)
}
