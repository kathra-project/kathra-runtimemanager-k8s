// Code generated by go-swagger; DO NOT EDIT.

package logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetLogsParams creates a new GetLogsParams object
// no default values defined in spec.
func NewGetLogsParams() GetLogsParams {

	return GetLogsParams{}
}

// GetLogsParams contains all the bound params for the get logs operation
// typically these are obtained from a http.Request
//
// swagger:parameters getLogs
type GetLogsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Regex filters
	  In: formData
	  Collection Format: multi
	*/
	Filters []string
	/*N latest lines
	  In: formData
	*/
	MaxLines *string
	/*RuntimeAppInstance providerId
	  Required: true
	  In: path
	*/
	ProviderIDAI string
	/*RuntimeEnvironment providerId
	  Required: true
	  In: path
	*/
	ProviderIDRE string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetLogsParams() beforehand.
func (o *GetLogsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		if err != http.ErrNotMultipart {
			return errors.New(400, "%v", err)
		} else if err := r.ParseForm(); err != nil {
			return errors.New(400, "%v", err)
		}
	}
	fds := runtime.Values(r.Form)

	fdFilters, fdhkFilters, _ := fds.GetOK("filters")
	if err := o.bindFilters(fdFilters, fdhkFilters, route.Formats); err != nil {
		res = append(res, err)
	}

	fdMaxLines, fdhkMaxLines, _ := fds.GetOK("maxLines")
	if err := o.bindMaxLines(fdMaxLines, fdhkMaxLines, route.Formats); err != nil {
		res = append(res, err)
	}

	rProviderIDAI, rhkProviderIDAI, _ := route.Params.GetOK("providerIdAI")
	if err := o.bindProviderIDAI(rProviderIDAI, rhkProviderIDAI, route.Formats); err != nil {
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

// bindFilters binds and validates array parameter Filters from formData.
//
// Arrays are parsed according to CollectionFormat: "multi" (defaults to "csv" when empty).
func (o *GetLogsParams) bindFilters(rawData []string, hasKey bool, formats strfmt.Registry) error {

	// CollectionFormat: multi
	filtersIC := rawData

	if len(filtersIC) == 0 {
		return nil
	}

	var filtersIR []string
	for _, filtersIV := range filtersIC {
		filtersI := filtersIV

		filtersIR = append(filtersIR, filtersI)
	}

	o.Filters = filtersIR

	return nil
}

// bindMaxLines binds and validates parameter MaxLines from formData.
func (o *GetLogsParams) bindMaxLines(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.MaxLines = &raw

	return nil
}

// bindProviderIDAI binds and validates parameter ProviderIDAI from path.
func (o *GetLogsParams) bindProviderIDAI(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ProviderIDAI = raw

	return nil
}

// bindProviderIDRE binds and validates parameter ProviderIDRE from path.
func (o *GetLogsParams) bindProviderIDRE(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ProviderIDRE = raw

	return nil
}
