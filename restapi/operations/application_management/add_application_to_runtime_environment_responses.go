// Code generated by go-swagger; DO NOT EDIT.

package application_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/models"
)

// AddApplicationToRuntimeEnvironmentOKCode is the HTTP code returned for type AddApplicationToRuntimeEnvironmentOK
const AddApplicationToRuntimeEnvironmentOKCode int = 200

/*AddApplicationToRuntimeEnvironmentOK RuntimeAppInstance deployed with success

swagger:response addApplicationToRuntimeEnvironmentOK
*/
type AddApplicationToRuntimeEnvironmentOK struct {

	/*
	  In: Body
	*/
	Payload *models.RuntimeAppInstance `json:"body,omitempty"`
}

// NewAddApplicationToRuntimeEnvironmentOK creates AddApplicationToRuntimeEnvironmentOK with default headers values
func NewAddApplicationToRuntimeEnvironmentOK() *AddApplicationToRuntimeEnvironmentOK {

	return &AddApplicationToRuntimeEnvironmentOK{}
}

// WithPayload adds the payload to the add application to runtime environment o k response
func (o *AddApplicationToRuntimeEnvironmentOK) WithPayload(payload *models.RuntimeAppInstance) *AddApplicationToRuntimeEnvironmentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add application to runtime environment o k response
func (o *AddApplicationToRuntimeEnvironmentOK) SetPayload(payload *models.RuntimeAppInstance) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddApplicationToRuntimeEnvironmentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddApplicationToRuntimeEnvironmentBadRequestCode is the HTTP code returned for type AddApplicationToRuntimeEnvironmentBadRequest
const AddApplicationToRuntimeEnvironmentBadRequestCode int = 400

/*AddApplicationToRuntimeEnvironmentBadRequest Bad request

swagger:response addApplicationToRuntimeEnvironmentBadRequest
*/
type AddApplicationToRuntimeEnvironmentBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewAddApplicationToRuntimeEnvironmentBadRequest creates AddApplicationToRuntimeEnvironmentBadRequest with default headers values
func NewAddApplicationToRuntimeEnvironmentBadRequest() *AddApplicationToRuntimeEnvironmentBadRequest {

	return &AddApplicationToRuntimeEnvironmentBadRequest{}
}

// WithPayload adds the payload to the add application to runtime environment bad request response
func (o *AddApplicationToRuntimeEnvironmentBadRequest) WithPayload(payload string) *AddApplicationToRuntimeEnvironmentBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add application to runtime environment bad request response
func (o *AddApplicationToRuntimeEnvironmentBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddApplicationToRuntimeEnvironmentBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// AddApplicationToRuntimeEnvironmentNotFoundCode is the HTTP code returned for type AddApplicationToRuntimeEnvironmentNotFound
const AddApplicationToRuntimeEnvironmentNotFoundCode int = 404

/*AddApplicationToRuntimeEnvironmentNotFound Not found

swagger:response addApplicationToRuntimeEnvironmentNotFound
*/
type AddApplicationToRuntimeEnvironmentNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewAddApplicationToRuntimeEnvironmentNotFound creates AddApplicationToRuntimeEnvironmentNotFound with default headers values
func NewAddApplicationToRuntimeEnvironmentNotFound() *AddApplicationToRuntimeEnvironmentNotFound {

	return &AddApplicationToRuntimeEnvironmentNotFound{}
}

// WithPayload adds the payload to the add application to runtime environment not found response
func (o *AddApplicationToRuntimeEnvironmentNotFound) WithPayload(payload string) *AddApplicationToRuntimeEnvironmentNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add application to runtime environment not found response
func (o *AddApplicationToRuntimeEnvironmentNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddApplicationToRuntimeEnvironmentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// AddApplicationToRuntimeEnvironmentInternalServerErrorCode is the HTTP code returned for type AddApplicationToRuntimeEnvironmentInternalServerError
const AddApplicationToRuntimeEnvironmentInternalServerErrorCode int = 500

/*AddApplicationToRuntimeEnvironmentInternalServerError Internal error

swagger:response addApplicationToRuntimeEnvironmentInternalServerError
*/
type AddApplicationToRuntimeEnvironmentInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewAddApplicationToRuntimeEnvironmentInternalServerError creates AddApplicationToRuntimeEnvironmentInternalServerError with default headers values
func NewAddApplicationToRuntimeEnvironmentInternalServerError() *AddApplicationToRuntimeEnvironmentInternalServerError {

	return &AddApplicationToRuntimeEnvironmentInternalServerError{}
}

// WithPayload adds the payload to the add application to runtime environment internal server error response
func (o *AddApplicationToRuntimeEnvironmentInternalServerError) WithPayload(payload string) *AddApplicationToRuntimeEnvironmentInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add application to runtime environment internal server error response
func (o *AddApplicationToRuntimeEnvironmentInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddApplicationToRuntimeEnvironmentInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
