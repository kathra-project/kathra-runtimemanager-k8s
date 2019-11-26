// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/restapi/operations"
	"gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/restapi/operations/application_management"
	"gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/restapi/operations/backup_management"
	"gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/restapi/operations/container_management"
	"gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/restapi/operations/environment_management"
	"gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/restapi/operations/job_management"
	"gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/restapi/operations/load_balancer_management"
	"gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/restapi/operations/logs"
	"gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/restapi/operations/volume_storage_management"
)

//go:generate swagger generate server --target ..\..\kathra-runtimemanager-api --name KathraRuntimeEnvironmentManager --spec ..\swagger.yaml

func configureFlags(api *operations.KathraRuntimeEnvironmentManagerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.KathraRuntimeEnvironmentManagerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.MultipartformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	if api.ApplicationManagementAddApplicationToRuntimeEnvironmentHandler == nil {
		api.ApplicationManagementAddApplicationToRuntimeEnvironmentHandler = application_management.AddApplicationToRuntimeEnvironmentHandlerFunc(func(params application_management.AddApplicationToRuntimeEnvironmentParams) middleware.Responder {
			return middleware.NotImplemented("operation application_management.AddApplicationToRuntimeEnvironment has not yet been implemented")
		})
	}
	if api.BackupManagementAddBackupRuntimeEnvironmentsHandler == nil {
		api.BackupManagementAddBackupRuntimeEnvironmentsHandler = backup_management.AddBackupRuntimeEnvironmentsHandlerFunc(func(params backup_management.AddBackupRuntimeEnvironmentsParams) middleware.Responder {
			return middleware.NotImplemented("operation backup_management.AddBackupRuntimeEnvironments has not yet been implemented")
		})
	}
	if api.EnvironmentManagementAddRuntimeEnvironmentHandler == nil {
		api.EnvironmentManagementAddRuntimeEnvironmentHandler = environment_management.AddRuntimeEnvironmentHandlerFunc(func(params environment_management.AddRuntimeEnvironmentParams) middleware.Responder {
			return middleware.NotImplemented("operation environment_management.AddRuntimeEnvironment has not yet been implemented")
		})
	}
	if api.ApplicationManagementDeleteApplicationHandler == nil {
		api.ApplicationManagementDeleteApplicationHandler = application_management.DeleteApplicationHandlerFunc(func(params application_management.DeleteApplicationParams) middleware.Responder {
			return middleware.NotImplemented("operation application_management.DeleteApplication has not yet been implemented")
		})
	}
	if api.BackupManagementDeleteBackupRuntimeEnvironmentsHandler == nil {
		api.BackupManagementDeleteBackupRuntimeEnvironmentsHandler = backup_management.DeleteBackupRuntimeEnvironmentsHandlerFunc(func(params backup_management.DeleteBackupRuntimeEnvironmentsParams) middleware.Responder {
			return middleware.NotImplemented("operation backup_management.DeleteBackupRuntimeEnvironments has not yet been implemented")
		})
	}
	if api.ContainerManagementDeleteRuntimeContainerHandler == nil {
		api.ContainerManagementDeleteRuntimeContainerHandler = container_management.DeleteRuntimeContainerHandlerFunc(func(params container_management.DeleteRuntimeContainerParams) middleware.Responder {
			return middleware.NotImplemented("operation container_management.DeleteRuntimeContainer has not yet been implemented")
		})
	}
	if api.EnvironmentManagementDeleteRuntimeEnvironmentHandler == nil {
		api.EnvironmentManagementDeleteRuntimeEnvironmentHandler = environment_management.DeleteRuntimeEnvironmentHandlerFunc(func(params environment_management.DeleteRuntimeEnvironmentParams) middleware.Responder {
			return middleware.NotImplemented("operation environment_management.DeleteRuntimeEnvironment has not yet been implemented")
		})
	}
	if api.BackupManagementDownloadBackupRuntimeEnvironmentsHandler == nil {
		api.BackupManagementDownloadBackupRuntimeEnvironmentsHandler = backup_management.DownloadBackupRuntimeEnvironmentsHandlerFunc(func(params backup_management.DownloadBackupRuntimeEnvironmentsParams) middleware.Responder {
			return middleware.NotImplemented("operation backup_management.DownloadBackupRuntimeEnvironments has not yet been implemented")
		})
	}
	if api.ApplicationManagementGetApplicationHandler == nil {
		api.ApplicationManagementGetApplicationHandler = application_management.GetApplicationHandlerFunc(func(params application_management.GetApplicationParams) middleware.Responder {
			return middleware.NotImplemented("operation application_management.GetApplication has not yet been implemented")
		})
	}
	if api.ApplicationManagementGetApplicationsFromRuntimeEnvironmentHandler == nil {
		api.ApplicationManagementGetApplicationsFromRuntimeEnvironmentHandler = application_management.GetApplicationsFromRuntimeEnvironmentHandlerFunc(func(params application_management.GetApplicationsFromRuntimeEnvironmentParams) middleware.Responder {
			return middleware.NotImplemented("operation application_management.GetApplicationsFromRuntimeEnvironment has not yet been implemented")
		})
	}
	if api.BackupManagementGetBackupRuntimeEnvironmentsHandler == nil {
		api.BackupManagementGetBackupRuntimeEnvironmentsHandler = backup_management.GetBackupRuntimeEnvironmentsHandlerFunc(func(params backup_management.GetBackupRuntimeEnvironmentsParams) middleware.Responder {
			return middleware.NotImplemented("operation backup_management.GetBackupRuntimeEnvironments has not yet been implemented")
		})
	}
	if api.LoadBalancerManagementGetLoadBalancersFromRuntimeHandler == nil {
		api.LoadBalancerManagementGetLoadBalancersFromRuntimeHandler = load_balancer_management.GetLoadBalancersFromRuntimeHandlerFunc(func(params load_balancer_management.GetLoadBalancersFromRuntimeParams) middleware.Responder {
			return middleware.NotImplemented("operation load_balancer_management.GetLoadBalancersFromRuntime has not yet been implemented")
		})
	}
	if api.LogsGetLogsHandler == nil {
		api.LogsGetLogsHandler = logs.GetLogsHandlerFunc(func(params logs.GetLogsParams) middleware.Responder {
			return middleware.NotImplemented("operation logs.GetLogs has not yet been implemented")
		})
	}
	if api.ApplicationManagementGetRuntimeAppServicesHandler == nil {
		api.ApplicationManagementGetRuntimeAppServicesHandler = application_management.GetRuntimeAppServicesHandlerFunc(func(params application_management.GetRuntimeAppServicesParams) middleware.Responder {
			return middleware.NotImplemented("operation application_management.GetRuntimeAppServices has not yet been implemented")
		})
	}
	if api.ContainerManagementGetRuntimeContainerHandler == nil {
		api.ContainerManagementGetRuntimeContainerHandler = container_management.GetRuntimeContainerHandlerFunc(func(params container_management.GetRuntimeContainerParams) middleware.Responder {
			return middleware.NotImplemented("operation container_management.GetRuntimeContainer has not yet been implemented")
		})
	}
	if api.EnvironmentManagementGetRuntimeEnvironmentHandler == nil {
		api.EnvironmentManagementGetRuntimeEnvironmentHandler = environment_management.GetRuntimeEnvironmentHandlerFunc(func(params environment_management.GetRuntimeEnvironmentParams) middleware.Responder {
			return middleware.NotImplemented("operation environment_management.GetRuntimeEnvironment has not yet been implemented")
		})
	}
	if api.EnvironmentManagementGetRuntimeEnvironmentsHandler == nil {
		api.EnvironmentManagementGetRuntimeEnvironmentsHandler = environment_management.GetRuntimeEnvironmentsHandlerFunc(func(params environment_management.GetRuntimeEnvironmentsParams) middleware.Responder {
			return middleware.NotImplemented("operation environment_management.GetRuntimeEnvironments has not yet been implemented")
		})
	}
	if api.JobManagementGetRuntimeJobsHandler == nil {
		api.JobManagementGetRuntimeJobsHandler = job_management.GetRuntimeJobsHandlerFunc(func(params job_management.GetRuntimeJobsParams) middleware.Responder {
			return middleware.NotImplemented("operation job_management.GetRuntimeJobs has not yet been implemented")
		})
	}
	if api.VolumeStorageManagementGetRuntimeVolumeHandler == nil {
		api.VolumeStorageManagementGetRuntimeVolumeHandler = volume_storage_management.GetRuntimeVolumeHandlerFunc(func(params volume_storage_management.GetRuntimeVolumeParams) middleware.Responder {
			return middleware.NotImplemented("operation volume_storage_management.GetRuntimeVolume has not yet been implemented")
		})
	}
	if api.BackupManagementImportBackupRuntimeEnvironmentsHandler == nil {
		api.BackupManagementImportBackupRuntimeEnvironmentsHandler = backup_management.ImportBackupRuntimeEnvironmentsHandlerFunc(func(params backup_management.ImportBackupRuntimeEnvironmentsParams) middleware.Responder {
			return middleware.NotImplemented("operation backup_management.ImportBackupRuntimeEnvironments has not yet been implemented")
		})
	}
	if api.BackupManagementRestoreBackupRuntimeEnvironmentsHandler == nil {
		api.BackupManagementRestoreBackupRuntimeEnvironmentsHandler = backup_management.RestoreBackupRuntimeEnvironmentsHandlerFunc(func(params backup_management.RestoreBackupRuntimeEnvironmentsParams) middleware.Responder {
			return middleware.NotImplemented("operation backup_management.RestoreBackupRuntimeEnvironments has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
