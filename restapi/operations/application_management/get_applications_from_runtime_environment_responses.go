// Code generated by go-swagger; DO NOT EDIT.

package application_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/models"
)

// GetApplicationsFromRuntimeEnvironmentOKCode is the HTTP code returned for type GetApplicationsFromRuntimeEnvironmentOK
const GetApplicationsFromRuntimeEnvironmentOKCode int = 200

/*GetApplicationsFromRuntimeEnvironmentOK Runtimes environments

swagger:response getApplicationsFromRuntimeEnvironmentOK
*/
type GetApplicationsFromRuntimeEnvironmentOK struct {

	/*
	  In: Body
	*/
	Payload []*models.RuntimeAppInstance `json:"body,omitempty"`
}

// NewGetApplicationsFromRuntimeEnvironmentOK creates GetApplicationsFromRuntimeEnvironmentOK with default headers values
func NewGetApplicationsFromRuntimeEnvironmentOK() *GetApplicationsFromRuntimeEnvironmentOK {

	return &GetApplicationsFromRuntimeEnvironmentOK{}
}

// WithPayload adds the payload to the get applications from runtime environment o k response
func (o *GetApplicationsFromRuntimeEnvironmentOK) WithPayload(payload []*models.RuntimeAppInstance) *GetApplicationsFromRuntimeEnvironmentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get applications from runtime environment o k response
func (o *GetApplicationsFromRuntimeEnvironmentOK) SetPayload(payload []*models.RuntimeAppInstance) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetApplicationsFromRuntimeEnvironmentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.RuntimeAppInstance, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetApplicationsFromRuntimeEnvironmentBadRequestCode is the HTTP code returned for type GetApplicationsFromRuntimeEnvironmentBadRequest
const GetApplicationsFromRuntimeEnvironmentBadRequestCode int = 400

/*GetApplicationsFromRuntimeEnvironmentBadRequest Bad request

swagger:response getApplicationsFromRuntimeEnvironmentBadRequest
*/
type GetApplicationsFromRuntimeEnvironmentBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetApplicationsFromRuntimeEnvironmentBadRequest creates GetApplicationsFromRuntimeEnvironmentBadRequest with default headers values
func NewGetApplicationsFromRuntimeEnvironmentBadRequest() *GetApplicationsFromRuntimeEnvironmentBadRequest {

	return &GetApplicationsFromRuntimeEnvironmentBadRequest{}
}

// WithPayload adds the payload to the get applications from runtime environment bad request response
func (o *GetApplicationsFromRuntimeEnvironmentBadRequest) WithPayload(payload string) *GetApplicationsFromRuntimeEnvironmentBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get applications from runtime environment bad request response
func (o *GetApplicationsFromRuntimeEnvironmentBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetApplicationsFromRuntimeEnvironmentBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetApplicationsFromRuntimeEnvironmentNotFoundCode is the HTTP code returned for type GetApplicationsFromRuntimeEnvironmentNotFound
const GetApplicationsFromRuntimeEnvironmentNotFoundCode int = 404

/*GetApplicationsFromRuntimeEnvironmentNotFound Not found

swagger:response getApplicationsFromRuntimeEnvironmentNotFound
*/
type GetApplicationsFromRuntimeEnvironmentNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetApplicationsFromRuntimeEnvironmentNotFound creates GetApplicationsFromRuntimeEnvironmentNotFound with default headers values
func NewGetApplicationsFromRuntimeEnvironmentNotFound() *GetApplicationsFromRuntimeEnvironmentNotFound {

	return &GetApplicationsFromRuntimeEnvironmentNotFound{}
}

// WithPayload adds the payload to the get applications from runtime environment not found response
func (o *GetApplicationsFromRuntimeEnvironmentNotFound) WithPayload(payload string) *GetApplicationsFromRuntimeEnvironmentNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get applications from runtime environment not found response
func (o *GetApplicationsFromRuntimeEnvironmentNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetApplicationsFromRuntimeEnvironmentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetApplicationsFromRuntimeEnvironmentInternalServerErrorCode is the HTTP code returned for type GetApplicationsFromRuntimeEnvironmentInternalServerError
const GetApplicationsFromRuntimeEnvironmentInternalServerErrorCode int = 500

/*GetApplicationsFromRuntimeEnvironmentInternalServerError Internal error

swagger:response getApplicationsFromRuntimeEnvironmentInternalServerError
*/
type GetApplicationsFromRuntimeEnvironmentInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetApplicationsFromRuntimeEnvironmentInternalServerError creates GetApplicationsFromRuntimeEnvironmentInternalServerError with default headers values
func NewGetApplicationsFromRuntimeEnvironmentInternalServerError() *GetApplicationsFromRuntimeEnvironmentInternalServerError {

	return &GetApplicationsFromRuntimeEnvironmentInternalServerError{}
}

// WithPayload adds the payload to the get applications from runtime environment internal server error response
func (o *GetApplicationsFromRuntimeEnvironmentInternalServerError) WithPayload(payload string) *GetApplicationsFromRuntimeEnvironmentInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get applications from runtime environment internal server error response
func (o *GetApplicationsFromRuntimeEnvironmentInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetApplicationsFromRuntimeEnvironmentInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}