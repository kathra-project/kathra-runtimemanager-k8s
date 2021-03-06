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

// RuntimeJob runtime job
// swagger:model RuntimeJob
type RuntimeJob struct {

	// RuntimeJob name
	Name string `json:"name,omitempty"`

	// RuntimeJob providerId
	ProviderID string `json:"providerId,omitempty"`

	// RuntimeEnvironment status
	// Enum: [SCHEDULED RUNNING SUCCESS FAILED]
	Status string `json:"status,omitempty"`
}

// Validate validates this runtime job
func (m *RuntimeJob) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var runtimeJobTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SCHEDULED","RUNNING","SUCCESS","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		runtimeJobTypeStatusPropEnum = append(runtimeJobTypeStatusPropEnum, v)
	}
}

const (

	// RuntimeJobStatusSCHEDULED captures enum value "SCHEDULED"
	RuntimeJobStatusSCHEDULED string = "SCHEDULED"

	// RuntimeJobStatusRUNNING captures enum value "RUNNING"
	RuntimeJobStatusRUNNING string = "RUNNING"

	// RuntimeJobStatusSUCCESS captures enum value "SUCCESS"
	RuntimeJobStatusSUCCESS string = "SUCCESS"

	// RuntimeJobStatusFAILED captures enum value "FAILED"
	RuntimeJobStatusFAILED string = "FAILED"
)

// prop value enum
func (m *RuntimeJob) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, runtimeJobTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *RuntimeJob) validateStatus(formats strfmt.Registry) error {

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
func (m *RuntimeJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RuntimeJob) UnmarshalBinary(b []byte) error {
	var res RuntimeJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
