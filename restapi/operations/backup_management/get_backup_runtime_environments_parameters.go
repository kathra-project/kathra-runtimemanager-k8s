// Code generated by go-swagger; DO NOT EDIT.

package backup_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetBackupRuntimeEnvironmentsParams creates a new GetBackupRuntimeEnvironmentsParams object
// no default values defined in spec.
func NewGetBackupRuntimeEnvironmentsParams() GetBackupRuntimeEnvironmentsParams {

	return GetBackupRuntimeEnvironmentsParams{}
}

// GetBackupRuntimeEnvironmentsParams contains all the bound params for the get backup runtime environments operation
// typically these are obtained from a http.Request
//
// swagger:parameters getBackupRuntimeEnvironments
type GetBackupRuntimeEnvironmentsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*RuntimeEnvironment providerId
	  Required: true
	  In: path
	*/
	ProviderID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetBackupRuntimeEnvironmentsParams() beforehand.
func (o *GetBackupRuntimeEnvironmentsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rProviderID, rhkProviderID, _ := route.Params.GetOK("providerId")
	if err := o.bindProviderID(rProviderID, rhkProviderID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindProviderID binds and validates parameter ProviderID from path.
func (o *GetBackupRuntimeEnvironmentsParams) bindProviderID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ProviderID = raw

	return nil
}
