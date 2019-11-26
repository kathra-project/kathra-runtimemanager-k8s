// Code generated by go-swagger; DO NOT EDIT.

package application_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetRuntimeAppServicesHandlerFunc turns a function with the right signature into a get runtime app services handler
type GetRuntimeAppServicesHandlerFunc func(GetRuntimeAppServicesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRuntimeAppServicesHandlerFunc) Handle(params GetRuntimeAppServicesParams) middleware.Responder {
	return fn(params)
}

// GetRuntimeAppServicesHandler interface for that can handle valid get runtime app services params
type GetRuntimeAppServicesHandler interface {
	Handle(GetRuntimeAppServicesParams) middleware.Responder
}

// NewGetRuntimeAppServices creates a new http.Handler for the get runtime app services operation
func NewGetRuntimeAppServices(ctx *middleware.Context, handler GetRuntimeAppServicesHandler) *GetRuntimeAppServices {
	return &GetRuntimeAppServices{Context: ctx, Handler: handler}
}

/*GetRuntimeAppServices swagger:route GET /runtimeEnvironments/{providerIdRE}/applications/{providerIdAI}/services Application management getRuntimeAppServices

Get RuntimeAppServices

*/
type GetRuntimeAppServices struct {
	Context *middleware.Context
	Handler GetRuntimeAppServicesHandler
}

func (o *GetRuntimeAppServices) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRuntimeAppServicesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}