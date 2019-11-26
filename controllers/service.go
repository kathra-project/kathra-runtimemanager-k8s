package controllers


import (
	"fmt"
	middleware "github.com/go-openapi/runtime/middleware"
	svc "gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/services/app_service"
	api "gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/restapi/operations/application_management"
)


func GetAppServices() (api.GetRuntimeAppServicesHandlerFunc) {
	return api.GetRuntimeAppServicesHandlerFunc(func(params api.GetRuntimeAppServicesParams) middleware.Responder {
		var services, err = svc.Instance.List(params.ProviderIDRE, params.ProviderIDAI)
		if (err != nil ) {
			fmt.Println(err)
			return api.NewGetRuntimeAppServicesInternalServerError().WithPayload("RuntimeAppServices with ProviderID '"+params.ProviderIDRE+"' generates internal error")
		} else if (services == nil) {
			return api.NewGetRuntimeAppServicesNotFound().WithPayload("RuntimeAppServices with ProviderID '"+params.ProviderIDRE+"' not found")
		} else {
			return api.NewGetRuntimeAppServicesOK().WithPayload(services)
		}
    })
}
