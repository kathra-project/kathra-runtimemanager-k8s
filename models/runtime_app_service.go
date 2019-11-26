// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RuntimeAppService runtime app service
// swagger:model RuntimeAppService
type RuntimeAppService struct {

	// RuntimeAppService name
	Name string `json:"name,omitempty"`

	// Ports exposed
	Ports []*RuntimeAppServicePort `json:"ports"`

	// RuntimeAppService providerId
	ProviderID string `json:"providerId,omitempty"`

	// RuntimeAppServicePort selectors
	Selectors []*RuntimeAppServiceSelector `json:"selectors"`
}

// Validate validates this runtime app service
func (m *RuntimeAppService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelectors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RuntimeAppService) validatePorts(formats strfmt.Registry) error {

	if swag.IsZero(m.Ports) { // not required
		return nil
	}

	for i := 0; i < len(m.Ports); i++ {
		if swag.IsZero(m.Ports[i]) { // not required
			continue
		}

		if m.Ports[i] != nil {
			if err := m.Ports[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RuntimeAppService) validateSelectors(formats strfmt.Registry) error {

	if swag.IsZero(m.Selectors) { // not required
		return nil
	}

	for i := 0; i < len(m.Selectors); i++ {
		if swag.IsZero(m.Selectors[i]) { // not required
			continue
		}

		if m.Selectors[i] != nil {
			if err := m.Selectors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("selectors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RuntimeAppService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RuntimeAppService) UnmarshalBinary(b []byte) error {
	var res RuntimeAppService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}