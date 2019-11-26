/*
 * Kathra Runtime Environment Manager
 *
 * KATHRA Runtime Environment Management API
 *
 * AppService functions
 */
package app_service

import (
	"context"
	k8sClt "github.com/ericchiang/k8s"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
	api "gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/models"
	svc "gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/services"
	"log"
)

type AppServiceService struct {
	client *k8sClt.Client
}

func NewAppServiceService() AppServiceService {
	return AppServiceService{client: svc.KubernetesClientServiceInstance.Client}
}

var Instance = NewAppServiceService()

func (service AppServiceService) List(namespaceProviderID string, applicationProviderID string) ([]*api.RuntimeAppService, error) {
	var svcList corev1.ServiceList
	var list = []*api.RuntimeAppService{}
	labels := new(k8sClt.LabelSelector)
	labels.Eq("release", applicationProviderID)
	if err := service.client.List(context.Background(), namespaceProviderID, &svcList, labels.Selector()); err != nil {
		log.Fatal(err)
	}
	for _, svc := range svcList.Items {
		list = append(list, mapAppServiceFromService(*svc))
	}
	return list, nil
}

func (service AppServiceService) Get(namespaceProviderID string, applicationProviderID string, serviceProviderID string) (*api.RuntimeAppService, error) {
	var svcList corev1.ServiceList
	labels := new(k8sClt.LabelSelector)
	labels.Eq("release", applicationProviderID)
	if err := service.client.List(context.Background(), namespaceProviderID, &svcList, labels.Selector()); err != nil {
		log.Fatal(err)
	}
	for _, svc := range svcList.Items {
		if *svc.Metadata.Name == serviceProviderID {
			return mapAppServiceFromService(*svc), nil
		}
	}
	return nil, nil
}

// Convert RuntimeAppService From Service
func mapAppServiceFromService(service corev1.Service) *api.RuntimeAppService {
	ports := []*api.RuntimeAppServicePort{}
	for _, port := range service.Spec.Ports {
		ports = append(ports, &api.RuntimeAppServicePort{Port: int64(*port.Port), Protocol: *port.Protocol, TargetPort: *port.TargetPort.StrVal})
	}
	selectors := []*api.RuntimeAppServiceSelector{}
	for k, v := range service.Spec.Selector {
		selectors = append(selectors, &api.RuntimeAppServiceSelector{Key: k, Value: v})
	}
	return &api.RuntimeAppService{ProviderID: *service.Metadata.Name, Name: *service.Metadata.Name, Ports: ports, Selectors: selectors}
}
