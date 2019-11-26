package controllers


import (
	"fmt"
	middleware "github.com/go-openapi/runtime/middleware"
	svc "gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/services/application"
	svc_env "gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/services/environment"
	api "gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/restapi/operations/application_management"
)


func GetApplications() (api.GetApplicationsFromRuntimeEnvironmentHandlerFunc) {
	return api.GetApplicationsFromRuntimeEnvironmentHandlerFunc(func(params api.GetApplicationsFromRuntimeEnvironmentParams) middleware.Responder {
		env, err := svc_env.Instance.Get(params.ProviderID)
		if (err != nil ) {
			fmt.Println(err)
			return api.NewGetApplicationsFromRuntimeEnvironmentInternalServerError().WithPayload("Internal error")
		} else if (env == nil) {
			return api.NewGetApplicationNotFound().WithPayload("Namespace not found")
		}
		
		applications, err := svc.Instance.List(params.ProviderID)
		if (err != nil ) {
			fmt.Println(err)
			return api.NewGetApplicationsFromRuntimeEnvironmentInternalServerError().WithPayload("Get Applications from RuntimeEnv '"+params.ProviderID+"' generates internal error")
		} else {
			return api.NewGetApplicationsFromRuntimeEnvironmentOK().WithPayload(applications)
		}
    })
}

func GetApplication() (api.GetApplicationHandlerFunc) {
	return api.GetApplicationHandlerFunc(func(params api.GetApplicationParams) middleware.Responder {
		env, err := svc_env.Instance.Get(params.ProviderIDRE)
		if (err != nil ) {
			fmt.Println(err)
			return api.NewGetApplicationsFromRuntimeEnvironmentInternalServerError().WithPayload("Internal error")
		} else if (env == nil) {
			return api.NewGetApplicationNotFound().WithPayload("Namespace not found")
		}

		application, err := svc.Instance.Get(params.ProviderIDRE, params.ProviderIDAI)
		if (err != nil ) {
			fmt.Println(err)
			return api.NewGetApplicationInternalServerError().WithPayload("Get Application generates internal error")
		} else if (application == nil) {
			return api.NewGetApplicationNotFound()
		} else {
			return api.NewGetApplicationOK().WithPayload(application)
		} 
    })
}

func AddAplication() (api.AddApplicationToRuntimeEnvironmentHandlerFunc) {
	return api.AddApplicationToRuntimeEnvironmentHandlerFunc(func(params api.AddApplicationToRuntimeEnvironmentParams) middleware.Responder {
		env, err := svc_env.Instance.Get(params.ProviderID)
		if (err != nil ) {
			fmt.Println(err)
			return api.NewGetApplicationsFromRuntimeEnvironmentInternalServerError().WithPayload("Internal error")
		} else if (env == nil) {
			return api.NewGetApplicationNotFound().WithPayload("Namespace not found")
		}

		application, err := svc.Instance.Add(params.ProviderID, params.RuntimeAppInstance)
		if (err != nil ) {
			fmt.Println(err)
			return api.NewAddApplicationToRuntimeEnvironmentInternalServerError().WithPayload("Add Application '"+params.RuntimeAppInstance.Name+"' generates internal error")
		} else {
			return api.NewAddApplicationToRuntimeEnvironmentOK().WithPayload(application)
		}
    })
}

func DeleteAplication() (api.DeleteApplicationHandlerFunc) {
	return api.DeleteApplicationHandlerFunc(func(params api.DeleteApplicationParams) middleware.Responder {
		env, err := svc_env.Instance.Get(params.ProviderIDRE)
		if (err != nil ) {
			fmt.Println(err)
			return api.NewGetApplicationsFromRuntimeEnvironmentInternalServerError().WithPayload("Internal error")
		} else if (env == nil) {
			return api.NewGetApplicationNotFound().WithPayload("Namespace not found")
		}

		app, err := svc.Instance.Delete(params.ProviderIDRE, params.ProviderIDAI)
		if (err != nil ) {
			fmt.Println(err)
			return api.NewDeleteApplicationInternalServerError().WithPayload("Delete Application '"+params.ProviderIDAI+"' generates internal error")
		} else if (app == nil) {
			return api.NewGetApplicationNotFound().WithPayload("Application with ProviderID '"+params.ProviderIDAI+"' not found")
		} else {
			return api.NewDeleteApplicationOK().WithPayload("Application "+app.ProviderID+" deleted")
		}
    })
}