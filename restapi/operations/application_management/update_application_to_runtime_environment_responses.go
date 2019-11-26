// Code generated by go-swagger; DO NOT EDIT.

package application_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/models"
)

// UpdateApplicationToRuntimeEnvironmentOKCode is the HTTP code returned for type UpdateApplicationToRuntimeEnvironmentOK
const UpdateApplicationToRuntimeEnvironmentOKCode int = 200

/*UpdateApplicationToRuntimeEnvironmentOK RuntimeAppInstance deployed with success

swagger:response updateApplicationToRuntimeEnvironmentOK
*/
type UpdateApplicationToRuntimeEnvironmentOK struct {

	/*
	  In: Body
	*/
	Payload *models.RuntimeAppInstance `json:"body,omitempty"`
}

// NewUpdateApplicationToRuntimeEnvironmentOK creates UpdateApplicationToRuntimeEnvironmentOK with default headers values
func NewUpdateApplicationToRuntimeEnvironmentOK() *UpdateApplicationToRuntimeEnvironmentOK {

	return &UpdateApplicationToRuntimeEnvironmentOK{}
}

// WithPayload adds the payload to the update application to runtime environment o k response
func (o *UpdateApplicationToRuntimeEnvironmentOK) WithPayload(payload *models.RuntimeAppInstance) *UpdateApplicationToRuntimeEnvironmentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update application to runtime environment o k response
func (o *UpdateApplicationToRuntimeEnvironmentOK) SetPayload(payload *models.RuntimeAppInstance) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateApplicationToRuntimeEnvironmentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateApplicationToRuntimeEnvironmentBadRequestCode is the HTTP code returned for type UpdateApplicationToRuntimeEnvironmentBadRequest
const UpdateApplicationToRuntimeEnvironmentBadRequestCode int = 400

/*UpdateApplicationToRuntimeEnvironmentBadRequest Bad request

swagger:response updateApplicationToRuntimeEnvironmentBadRequest
*/
type UpdateApplicationToRuntimeEnvironmentBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewUpdateApplicationToRuntimeEnvironmentBadRequest creates UpdateApplicationToRuntimeEnvironmentBadRequest with default headers values
func NewUpdateApplicationToRuntimeEnvironmentBadRequest() *UpdateApplicationToRuntimeEnvironmentBadRequest {

	return &UpdateApplicationToRuntimeEnvironmentBadRequest{}
}

// WithPayload adds the payload to the update application to runtime environment bad request response
func (o *UpdateApplicationToRuntimeEnvironmentBadRequest) WithPayload(payload string) *UpdateApplicationToRuntimeEnvironmentBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update application to runtime environment bad request response
func (o *UpdateApplicationToRuntimeEnvironmentBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateApplicationToRuntimeEnvironmentBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateApplicationToRuntimeEnvironmentNotFoundCode is the HTTP code returned for type UpdateApplicationToRuntimeEnvironmentNotFound
const UpdateApplicationToRuntimeEnvironmentNotFoundCode int = 404

/*UpdateApplicationToRuntimeEnvironmentNotFound Not found

swagger:response updateApplicationToRuntimeEnvironmentNotFound
*/
type UpdateApplicationToRuntimeEnvironmentNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewUpdateApplicationToRuntimeEnvironmentNotFound creates UpdateApplicationToRuntimeEnvironmentNotFound with default headers values
func NewUpdateApplicationToRuntimeEnvironmentNotFound() *UpdateApplicationToRuntimeEnvironmentNotFound {

	return &UpdateApplicationToRuntimeEnvironmentNotFound{}
}

// WithPayload adds the payload to the update application to runtime environment not found response
func (o *UpdateApplicationToRuntimeEnvironmentNotFound) WithPayload(payload string) *UpdateApplicationToRuntimeEnvironmentNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update application to runtime environment not found response
func (o *UpdateApplicationToRuntimeEnvironmentNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateApplicationToRuntimeEnvironmentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateApplicationToRuntimeEnvironmentInternalServerErrorCode is the HTTP code returned for type UpdateApplicationToRuntimeEnvironmentInternalServerError
const UpdateApplicationToRuntimeEnvironmentInternalServerErrorCode int = 500

/*UpdateApplicationToRuntimeEnvironmentInternalServerError Internal error

swagger:response updateApplicationToRuntimeEnvironmentInternalServerError
*/
type UpdateApplicationToRuntimeEnvironmentInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewUpdateApplicationToRuntimeEnvironmentInternalServerError creates UpdateApplicationToRuntimeEnvironmentInternalServerError with default headers values
func NewUpdateApplicationToRuntimeEnvironmentInternalServerError() *UpdateApplicationToRuntimeEnvironmentInternalServerError {

	return &UpdateApplicationToRuntimeEnvironmentInternalServerError{}
}

// WithPayload adds the payload to the update application to runtime environment internal server error response
func (o *UpdateApplicationToRuntimeEnvironmentInternalServerError) WithPayload(payload string) *UpdateApplicationToRuntimeEnvironmentInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update application to runtime environment internal server error response
func (o *UpdateApplicationToRuntimeEnvironmentInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateApplicationToRuntimeEnvironmentInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
