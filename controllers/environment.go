package controllers


import (
	"fmt"
	middleware "github.com/go-openapi/runtime/middleware"
	svc "gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/services/environment"
	api "gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/restapi/operations/environment_management"
)


func GetRuntimeEnvironment() (api.GetRuntimeEnvironmentHandlerFunc) {
	return api.GetRuntimeEnvironmentHandlerFunc(func(params api.GetRuntimeEnvironmentParams) middleware.Responder {
		var environment, err = svc.Instance.Get(params.ProviderID)
		if (err != nil ) {
			fmt.Println(err)
			return api.NewGetRuntimeEnvironmentInternalServerError().WithPayload("RuntimeEnvironment with ProviderID '"+params.ProviderID+"' generates internal error")
		} else if (environment == nil) {
			return api.NewGetRuntimeEnvironmentNotFound().WithPayload("RuntimeEnvironment with ProviderID '"+params.ProviderID+"' not found")
		} else {
			return api.NewGetRuntimeEnvironmentOK().WithPayload(environment)
		}
    })
}

func GetRuntimeEnvironments() (api.GetRuntimeEnvironmentsHandlerFunc) {
	return api.GetRuntimeEnvironmentsHandlerFunc(func(params api.GetRuntimeEnvironmentsParams) middleware.Responder {
		var environments, err = svc.Instance.List()
		if (err != nil ) {
			fmt.Println(err)
			return api.NewGetRuntimeEnvironmentsInternalServerError().WithPayload("Get RuntimeEnvironments generates internal error")
		} else {
			return api.NewGetRuntimeEnvironmentsOK().WithPayload(environments)
		}
    })
}

func AddRuntimeEnvironment() (api.AddRuntimeEnvironmentHandlerFunc) {
	return api.AddRuntimeEnvironmentHandlerFunc(func(params api.AddRuntimeEnvironmentParams) middleware.Responder {
		var environment, err = svc.Instance.Add(params.RuntimeEnvironment)
		if (err != nil ) {
			fmt.Println(err)
			return api.NewAddRuntimeEnvironmentInternalServerError().WithPayload("Add RuntimeEnvironment '"+params.RuntimeEnvironment.Name+"' generates internal error")
		} else {
			return api.NewAddRuntimeEnvironmentOK().WithPayload(environment)
		}
    })
}

func DeleteRuntimeEnvironment() (api.DeleteRuntimeEnvironmentHandlerFunc) {
	return api.DeleteRuntimeEnvironmentHandlerFunc(func(params api.DeleteRuntimeEnvironmentParams) middleware.Responder {
		environment, err := svc.Instance.Delete(params.ProviderID)
		if (err != nil ) {
			fmt.Println(err)
			return api.NewDeleteRuntimeEnvironmentInternalServerError().WithPayload("Delete RuntimeEnvironment '"+params.ProviderID+"' generates internal error")
		} else if (environment == nil) {
			return api.NewGetRuntimeEnvironmentNotFound().WithPayload("RuntimeEnvironment with ProviderID '"+params.ProviderID+"' not found")
		} else {
			return api.NewDeleteRuntimeEnvironmentOK().WithPayload("RuntimeEnvironment "+environment.ProviderID+" deleted")
		}
    })
}