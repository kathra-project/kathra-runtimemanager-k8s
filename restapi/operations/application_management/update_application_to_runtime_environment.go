// Code generated by go-swagger; DO NOT EDIT.

package application_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UpdateApplicationToRuntimeEnvironmentHandlerFunc turns a function with the right signature into a update application to runtime environment handler
type UpdateApplicationToRuntimeEnvironmentHandlerFunc func(UpdateApplicationToRuntimeEnvironmentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateApplicationToRuntimeEnvironmentHandlerFunc) Handle(params UpdateApplicationToRuntimeEnvironmentParams) middleware.Responder {
	return fn(params)
}

// UpdateApplicationToRuntimeEnvironmentHandler interface for that can handle valid update application to runtime environment params
type UpdateApplicationToRuntimeEnvironmentHandler interface {
	Handle(UpdateApplicationToRuntimeEnvironmentParams) middleware.Responder
}

// NewUpdateApplicationToRuntimeEnvironment creates a new http.Handler for the update application to runtime environment operation
func NewUpdateApplicationToRuntimeEnvironment(ctx *middleware.Context, handler UpdateApplicationToRuntimeEnvironmentHandler) *UpdateApplicationToRuntimeEnvironment {
	return &UpdateApplicationToRuntimeEnvironment{Context: ctx, Handler: handler}
}

/*UpdateApplicationToRuntimeEnvironment swagger:route PUT /runtimeEnvironments/{providerId}/applications Application management updateApplicationToRuntimeEnvironment

Update application to RuntimeEnvironment

*/
type UpdateApplicationToRuntimeEnvironment struct {
	Context *middleware.Context
	Handler UpdateApplicationToRuntimeEnvironmentHandler
}

func (o *UpdateApplicationToRuntimeEnvironment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateApplicationToRuntimeEnvironmentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
