// Code generated by go-swagger; DO NOT EDIT.

package environment_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetRuntimeEnvironmentsHandlerFunc turns a function with the right signature into a get runtime environments handler
type GetRuntimeEnvironmentsHandlerFunc func(GetRuntimeEnvironmentsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRuntimeEnvironmentsHandlerFunc) Handle(params GetRuntimeEnvironmentsParams) middleware.Responder {
	return fn(params)
}

// GetRuntimeEnvironmentsHandler interface for that can handle valid get runtime environments params
type GetRuntimeEnvironmentsHandler interface {
	Handle(GetRuntimeEnvironmentsParams) middleware.Responder
}

// NewGetRuntimeEnvironments creates a new http.Handler for the get runtime environments operation
func NewGetRuntimeEnvironments(ctx *middleware.Context, handler GetRuntimeEnvironmentsHandler) *GetRuntimeEnvironments {
	return &GetRuntimeEnvironments{Context: ctx, Handler: handler}
}

/*GetRuntimeEnvironments swagger:route GET /runtimeEnvironments Environment management getRuntimeEnvironments

Get all RuntimeEnvironments

*/
type GetRuntimeEnvironments struct {
	Context *middleware.Context
	Handler GetRuntimeEnvironmentsHandler
}

func (o *GetRuntimeEnvironments) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRuntimeEnvironmentsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
