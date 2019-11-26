// Code generated by go-swagger; DO NOT EDIT.

package container_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteRuntimeContainerParams creates a new DeleteRuntimeContainerParams object
// no default values defined in spec.
func NewDeleteRuntimeContainerParams() DeleteRuntimeContainerParams {

	return DeleteRuntimeContainerParams{}
}

// DeleteRuntimeContainerParams contains all the bound params for the delete runtime container operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteRuntimeContainer
type DeleteRuntimeContainerParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*RuntimeAppInstance providerId
	  Required: true
	  In: path
	*/
	ProviderIDAI string
	/*RuntimeAppInstance providerId
	  Required: true
	  In: path
	*/
	ProviderIDContainer string
	/*RuntimeEnvironment providerId
	  Required: true
	  In: path
	*/
	ProviderIDRE string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteRuntimeContainerParams() beforehand.
func (o *DeleteRuntimeContainerParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rProviderIDAI, rhkProviderIDAI, _ := route.Params.GetOK("providerIdAI")
	if err := o.bindProviderIDAI(rProviderIDAI, rhkProviderIDAI, route.Formats); err != nil {
		res = append(res, err)
	}

	rProviderIDContainer, rhkProviderIDContainer, _ := route.Params.GetOK("providerIdContainer")
	if err := o.bindProviderIDContainer(rProviderIDContainer, rhkProviderIDContainer, route.Formats); err != nil {
		res = append(res, err)
	}

	rProviderIDRE, rhkProviderIDRE, _ := route.Params.GetOK("providerIdRE")
	if err := o.bindProviderIDRE(rProviderIDRE, rhkProviderIDRE, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindProviderIDAI binds and validates parameter ProviderIDAI from path.
func (o *DeleteRuntimeContainerParams) bindProviderIDAI(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ProviderIDAI = raw

	return nil
}

// bindProviderIDContainer binds and validates parameter ProviderIDContainer from path.
func (o *DeleteRuntimeContainerParams) bindProviderIDContainer(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ProviderIDContainer = raw

	return nil
}

// bindProviderIDRE binds and validates parameter ProviderIDRE from path.
func (o *DeleteRuntimeContainerParams) bindProviderIDRE(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ProviderIDRE = raw

	return nil
}
