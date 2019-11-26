// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RuntimeVolume runtime volume
// swagger:model RuntimeVolume
type RuntimeVolume struct {

	// Volume name
	Name string `json:"name,omitempty"`

	// RuntimeVolume providerId
	ProviderID string `json:"providerId,omitempty"`

	// Volume space
	Space string `json:"space,omitempty"`

	// RuntimeVolume status
	// Enum: [PROVISIONING AVAILABLE TERMINATING]
	Status string `json:"status,omitempty"`
}

// Validate validates this runtime volume
func (m *RuntimeVolume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var runtimeVolumeTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PROVISIONING","AVAILABLE","TERMINATING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		runtimeVolumeTypeStatusPropEnum = append(runtimeVolumeTypeStatusPropEnum, v)
	}
}

const (

	// RuntimeVolumeStatusPROVISIONING captures enum value "PROVISIONING"
	RuntimeVolumeStatusPROVISIONING string = "PROVISIONING"

	// RuntimeVolumeStatusAVAILABLE captures enum value "AVAILABLE"
	RuntimeVolumeStatusAVAILABLE string = "AVAILABLE"

	// RuntimeVolumeStatusTERMINATING captures enum value "TERMINATING"
	RuntimeVolumeStatusTERMINATING string = "TERMINATING"
)

// prop value enum
func (m *RuntimeVolume) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, runtimeVolumeTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *RuntimeVolume) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RuntimeVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RuntimeVolume) UnmarshalBinary(b []byte) error {
	var res RuntimeVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
