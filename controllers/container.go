package controllers

import (
	"fmt"
	middleware "github.com/go-openapi/runtime/middleware"
	api "gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/restapi/operations/container_management"
	svc "gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/services/container_service"
)

func GetContainers() api.GetRuntimeContainerHandlerFunc {
	return api.GetRuntimeContainerHandlerFunc(func(params api.GetRuntimeContainerParams) middleware.Responder {
		var containers, err = svc.Instance.List(params.ProviderIDRE, params.ProviderIDAI)
		if err != nil {
			fmt.Println(err)
			return api.NewGetRuntimeContainerInternalServerError().WithPayload("RuntimeAppServices with ProviderID '" + params.ProviderIDRE + "' generates internal error")
		} else if containers == nil {
			return api.NewGetRuntimeContainerNotFound().WithPayload("RuntimeAppServices with ProviderID '" + params.ProviderIDRE + "' not found")
		} else {
			return api.NewGetRuntimeContainerOK().WithPayload(containers)
		}
	})
}
