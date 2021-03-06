// Code generated by go-swagger; DO NOT EDIT.

package container_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteRuntimeContainerHandlerFunc turns a function with the right signature into a delete runtime container handler
type DeleteRuntimeContainerHandlerFunc func(DeleteRuntimeContainerParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteRuntimeContainerHandlerFunc) Handle(params DeleteRuntimeContainerParams) middleware.Responder {
	return fn(params)
}

// DeleteRuntimeContainerHandler interface for that can handle valid delete runtime container params
type DeleteRuntimeContainerHandler interface {
	Handle(DeleteRuntimeContainerParams) middleware.Responder
}

// NewDeleteRuntimeContainer creates a new http.Handler for the delete runtime container operation
func NewDeleteRuntimeContainer(ctx *middleware.Context, handler DeleteRuntimeContainerHandler) *DeleteRuntimeContainer {
	return &DeleteRuntimeContainer{Context: ctx, Handler: handler}
}

/*DeleteRuntimeContainer swagger:route DELETE /runtimeEnvironments/{providerIdRE}/applications/{providerIdAI}/containers/{providerIdContainer} Container management deleteRuntimeContainer

Kill RuntimeContainer

*/
type DeleteRuntimeContainer struct {
	Context *middleware.Context
	Handler DeleteRuntimeContainerHandler
}

func (o *DeleteRuntimeContainer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteRuntimeContainerParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
