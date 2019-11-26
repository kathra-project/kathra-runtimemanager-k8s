// Code generated by go-swagger; DO NOT EDIT.

package environment_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/models"
)

// AddRuntimeEnvironmentOKCode is the HTTP code returned for type AddRuntimeEnvironmentOK
const AddRuntimeEnvironmentOKCode int = 200

/*AddRuntimeEnvironmentOK RuntimeEnvironment deployed with success

swagger:response addRuntimeEnvironmentOK
*/
type AddRuntimeEnvironmentOK struct {

	/*
	  In: Body
	*/
	Payload *models.RuntimeEnvironment `json:"body,omitempty"`
}

// NewAddRuntimeEnvironmentOK creates AddRuntimeEnvironmentOK with default headers values
func NewAddRuntimeEnvironmentOK() *AddRuntimeEnvironmentOK {

	return &AddRuntimeEnvironmentOK{}
}

// WithPayload adds the payload to the add runtime environment o k response
func (o *AddRuntimeEnvironmentOK) WithPayload(payload *models.RuntimeEnvironment) *AddRuntimeEnvironmentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add runtime environment o k response
func (o *AddRuntimeEnvironmentOK) SetPayload(payload *models.RuntimeEnvironment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddRuntimeEnvironmentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddRuntimeEnvironmentBadRequestCode is the HTTP code returned for type AddRuntimeEnvironmentBadRequest
const AddRuntimeEnvironmentBadRequestCode int = 400

/*AddRuntimeEnvironmentBadRequest Bad request

swagger:response addRuntimeEnvironmentBadRequest
*/
type AddRuntimeEnvironmentBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewAddRuntimeEnvironmentBadRequest creates AddRuntimeEnvironmentBadRequest with default headers values
func NewAddRuntimeEnvironmentBadRequest() *AddRuntimeEnvironmentBadRequest {

	return &AddRuntimeEnvironmentBadRequest{}
}

// WithPayload adds the payload to the add runtime environment bad request response
func (o *AddRuntimeEnvironmentBadRequest) WithPayload(payload string) *AddRuntimeEnvironmentBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add runtime environment bad request response
func (o *AddRuntimeEnvironmentBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddRuntimeEnvironmentBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// AddRuntimeEnvironmentNotFoundCode is the HTTP code returned for type AddRuntimeEnvironmentNotFound
const AddRuntimeEnvironmentNotFoundCode int = 404

/*AddRuntimeEnvironmentNotFound Not found

swagger:response addRuntimeEnvironmentNotFound
*/
type AddRuntimeEnvironmentNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewAddRuntimeEnvironmentNotFound creates AddRuntimeEnvironmentNotFound with default headers values
func NewAddRuntimeEnvironmentNotFound() *AddRuntimeEnvironmentNotFound {

	return &AddRuntimeEnvironmentNotFound{}
}

// WithPayload adds the payload to the add runtime environment not found response
func (o *AddRuntimeEnvironmentNotFound) WithPayload(payload string) *AddRuntimeEnvironmentNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add runtime environment not found response
func (o *AddRuntimeEnvironmentNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddRuntimeEnvironmentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// AddRuntimeEnvironmentInternalServerErrorCode is the HTTP code returned for type AddRuntimeEnvironmentInternalServerError
const AddRuntimeEnvironmentInternalServerErrorCode int = 500

/*AddRuntimeEnvironmentInternalServerError Internal error

swagger:response addRuntimeEnvironmentInternalServerError
*/
type AddRuntimeEnvironmentInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewAddRuntimeEnvironmentInternalServerError creates AddRuntimeEnvironmentInternalServerError with default headers values
func NewAddRuntimeEnvironmentInternalServerError() *AddRuntimeEnvironmentInternalServerError {

	return &AddRuntimeEnvironmentInternalServerError{}
}

// WithPayload adds the payload to the add runtime environment internal server error response
func (o *AddRuntimeEnvironmentInternalServerError) WithPayload(payload string) *AddRuntimeEnvironmentInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add runtime environment internal server error response
func (o *AddRuntimeEnvironmentInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddRuntimeEnvironmentInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
