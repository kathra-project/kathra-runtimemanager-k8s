// Code generated by go-swagger; DO NOT EDIT.

package application_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetApplicationsFromRuntimeEnvironmentHandlerFunc turns a function with the right signature into a get applications from runtime environment handler
type GetApplicationsFromRuntimeEnvironmentHandlerFunc func(GetApplicationsFromRuntimeEnvironmentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetApplicationsFromRuntimeEnvironmentHandlerFunc) Handle(params GetApplicationsFromRuntimeEnvironmentParams) middleware.Responder {
	return fn(params)
}

// GetApplicationsFromRuntimeEnvironmentHandler interface for that can handle valid get applications from runtime environment params
type GetApplicationsFromRuntimeEnvironmentHandler interface {
	Handle(GetApplicationsFromRuntimeEnvironmentParams) middleware.Responder
}

// NewGetApplicationsFromRuntimeEnvironment creates a new http.Handler for the get applications from runtime environment operation
func NewGetApplicationsFromRuntimeEnvironment(ctx *middleware.Context, handler GetApplicationsFromRuntimeEnvironmentHandler) *GetApplicationsFromRuntimeEnvironment {
	return &GetApplicationsFromRuntimeEnvironment{Context: ctx, Handler: handler}
}

/*GetApplicationsFromRuntimeEnvironment swagger:route GET /runtimeEnvironments/{providerId}/applications Application management getApplicationsFromRuntimeEnvironment

Get application instances from RuntimeEnvironment

*/
type GetApplicationsFromRuntimeEnvironment struct {
	Context *middleware.Context
	Handler GetApplicationsFromRuntimeEnvironmentHandler
}

func (o *GetApplicationsFromRuntimeEnvironment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetApplicationsFromRuntimeEnvironmentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
