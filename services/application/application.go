/*
 * Kathra Runtime Environment Manager
 *
 * KATHRA Runtime Environment Management API
 *
 * Envirnonment functions
 */
package application

import (
	"log"
	k8sClt "github.com/ericchiang/k8s"
	svc "gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/services"
	api "gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/models"
)

type ApplicationService struct {  
	k8sClt *k8sClt.Client
	helmClt *svc.HelmClientService
}

func NewApplicationService() ApplicationService {  
	return ApplicationService {
		k8sClt: svc.KubernetesClientServiceInstance.Client,
		helmClt: &svc.HelmClientServiceInstance,
	}
}

var Instance = NewApplicationService()

func (self ApplicationService) List(namespace string) ([]*api.RuntimeAppInstance, error) {
	var helmReleases, err = self.helmClt.GetReleasesFromNamespace(namespace)
	if (err != nil) {
		log.Println(err)
        return nil, err
    }
	var list = []*api.RuntimeAppInstance{}
	for _, ns := range helmReleases {
        list = append(list, mapRuntimeAppInstanceFromHelmRelease(ns))
    }
    return list, nil
}
func (self ApplicationService) Add(namespace string, app *api.RuntimeAppInstance) (*api.RuntimeAppInstance, error) {
	self.helmClt.InstallRelease(namespace, app.Name, app.CatalogEntry.ProviderID, app.CatalogEntry.Version, convertCatalogEntryArgumentsToStringMap(app.Parameters))
    return self.Get(namespace, app.Name)
}
func (self ApplicationService) Update(namespace string, app *api.RuntimeAppInstance) (*api.RuntimeAppInstance, error) {
    return nil, nil
}
func (self ApplicationService) Delete(namespace string, ProviderID string) (*api.RuntimeAppInstance, error) {
	var app, err = self.Get(namespace, ProviderID)
	if (err != nil) {
		log.Println(err)
        return nil, err
	} else if (app == nil) {
		return nil, nil
	}
	err = self.helmClt.DeleteReleaseFromNamespace(namespace, ProviderID)
	if (err != nil) {
		log.Println(err)
        return nil, err
	}
	app.Status = "DELETED"
    return app, nil
}
func (self ApplicationService) Get(namespace string, ProviderID string) (*api.RuntimeAppInstance, error) {
	var helmRelease, err = self.helmClt.GetReleaseFromNamespace(namespace, ProviderID)
	if (err != nil) {
		log.Println(err)
        return nil, err
	}
	if (helmRelease == nil) {
        return nil, nil
	}
	var app = mapRuntimeAppInstanceFromHelmRelease(helmRelease)
	var values, errValues = self.helmClt.GetReleaseValueFromNamespace(namespace, ProviderID)
	if (errValues != nil) {
		log.Println(errValues)
        return nil, errValues
    }
	app.Parameters = convertStringMapToCatalogEntryArguments(values)
    return app, nil
}

func mapRuntimeAppInstanceFromHelmRelease(release *svc.HelmRelease)(*api.RuntimeAppInstance) {
	var providerId = release.Name
	var catalogEntry = api.CatalogEntry {
		ProviderID: release.Chart,
		Version: release.AppVersion,
	}
	return &api.RuntimeAppInstance{ProviderID: providerId, Name: release.Name, Status: release.Status, CatalogEntry: &catalogEntry}
}

func convertStringMapToCatalogEntryArguments(entries map[string]string)([]*api.CatalogEntryArgument) {
    var list = []*api.CatalogEntryArgument{}
	for k, v := range entries { 
		list = append(list, &api.CatalogEntryArgument{Key: k, Value: v})
	}
	return list
}
func convertCatalogEntryArgumentsToStringMap(entries []*api.CatalogEntryArgument)(map[string]string) {
    var values map[string]string
    values = make(map[string]string)
	for _, entry := range entries { 
		values[entry.Key] = entry.Value
	}
	return values
}


