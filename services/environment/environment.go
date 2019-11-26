/*
 * Kathra Runtime Environment Manager
 *
 * KATHRA Runtime Environment Management API
 *
 * Environment functions
 */
package environment

import (
    "errors"
	"log"
    "context"
    k8sClt "github.com/ericchiang/k8s"
    corev1 "github.com/ericchiang/k8s/apis/core/v1"
	metav1 "github.com/ericchiang/k8s/apis/meta/v1"
	svc "gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/services"
	api "gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/models"
)

type EnvironmentService struct {  
    client *k8sClt.Client
}

func NewEnvironmentService() EnvironmentService {  
    return EnvironmentService {client: svc.KubernetesClientServiceInstance.Client}
}

var Instance = NewEnvironmentService()

func (service EnvironmentService) List() ([]*api.RuntimeEnvironment, error) {
    var nsList corev1.NamespaceList
    var list = []*api.RuntimeEnvironment{}
    if err := service.client.List(context.Background(), "", &nsList); err != nil {
        log.Fatal(err)
    }
    for _, ns := range nsList.Items {
        list = append(list, mapEnvironmentFromNamespace(*ns))
    }
    return list, nil
}

func (service EnvironmentService) Get(name string) (*api.RuntimeEnvironment, error) {
    var nsList corev1.NamespaceList
    if err := service.client.List(context.Background(), "", &nsList); err != nil {
        log.Fatal(err)
    }
    for _, ns := range nsList.Items {
        if (*ns.Metadata.Name == name) {
            return mapEnvironmentFromNamespace(*ns), nil
        }
    }
    return nil, nil
}

func (service EnvironmentService) Add(env *api.RuntimeEnvironment) (*api.RuntimeEnvironment, error) {
    
    var namespace, errMapping = mapNamespaceFromEnv(env)
    if (errMapping != nil) {
        return nil, errMapping
    }
    var ns corev1.Namespace
    if err := service.client.Get(context.Background(), "", *namespace.Metadata.Name, &ns); err != nil {
        if err := service.client.Create(context.Background(), namespace); err != nil {
            return &api.RuntimeEnvironment{}, err
        }
        return Instance.Get(*namespace.Metadata.Name)
    } else {
        // already exists
        return nil, errors.New("Environment '"+env.ProviderID+"' already exists")
    }
}

func (service EnvironmentService) Delete(name string) (*api.RuntimeEnvironment, error) {
    var ns corev1.Namespace
    if err := service.client.Get(context.Background(), "", name, &ns); err != nil {
        return &api.RuntimeEnvironment{}, err
    }
    if err := service.client.Delete(context.Background(), &ns); err != nil {
        return &api.RuntimeEnvironment{}, err
    }
    return mapEnvironmentFromNamespace(ns), nil
}

// Convert RuntimeEnvironment From Namespace
func mapEnvironmentFromNamespace(namespace corev1.Namespace) (*api.RuntimeEnvironment) {
    var labels = namespace.Metadata.Labels
    var name = labels["kathraName"]
    var description = labels["kathraDescription"]
    if (len(name) == 0) {
        name = *namespace.Metadata.Name
    }
    return &api.RuntimeEnvironment{ProviderID: *namespace.Metadata.Name, Name: name, Description: description}
}

// Convert Namespace From RuntimeEnvironment
func mapNamespaceFromEnv(env *api.RuntimeEnvironment) (*corev1.Namespace, error) {
    if (len(env.Name) == 0) {
        return nil, errors.New("Environment's name is empty")
    }
    var providerID = env.ProviderID
    if (len(env.ProviderID) == 0) {
        // generate ProviderID from name
        providerID = env.Name
    }
    var labels map[string]string
    labels = make(map[string]string)
    labels["kathraName"] =  env.Name
    labels["kathraDescription"] =  env.Description
    var namespace = corev1.Namespace {
        Metadata: &metav1.ObjectMeta{
            Name:    &providerID,
            Labels:  labels,
        },
    }
    return &namespace, nil
}