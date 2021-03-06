// Code generated by go-swagger; DO NOT EDIT.

package application_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/models"
)

// GetApplicationOKCode is the HTTP code returned for type GetApplicationOK
const GetApplicationOKCode int = 200

/*GetApplicationOK RuntimeAppInstance

swagger:response getApplicationOK
*/
type GetApplicationOK struct {

	/*
	  In: Body
	*/
	Payload *models.RuntimeAppInstance `json:"body,omitempty"`
}

// NewGetApplicationOK creates GetApplicationOK with default headers values
func NewGetApplicationOK() *GetApplicationOK {

	return &GetApplicationOK{}
}

// WithPayload adds the payload to the get application o k response
func (o *GetApplicationOK) WithPayload(payload *models.RuntimeAppInstance) *GetApplicationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get application o k response
func (o *GetApplicationOK) SetPayload(payload *models.RuntimeAppInstance) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetApplicationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetApplicationBadRequestCode is the HTTP code returned for type GetApplicationBadRequest
const GetApplicationBadRequestCode int = 400

/*GetApplicationBadRequest Bad request

swagger:response getApplicationBadRequest
*/
type GetApplicationBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetApplicationBadRequest creates GetApplicationBadRequest with default headers values
func NewGetApplicationBadRequest() *GetApplicationBadRequest {

	return &GetApplicationBadRequest{}
}

// WithPayload adds the payload to the get application bad request response
func (o *GetApplicationBadRequest) WithPayload(payload string) *GetApplicationBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get application bad request response
func (o *GetApplicationBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetApplicationBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetApplicationNotFoundCode is the HTTP code returned for type GetApplicationNotFound
const GetApplicationNotFoundCode int = 404

/*GetApplicationNotFound Not found

swagger:response getApplicationNotFound
*/
type GetApplicationNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetApplicationNotFound creates GetApplicationNotFound with default headers values
func NewGetApplicationNotFound() *GetApplicationNotFound {

	return &GetApplicationNotFound{}
}

// WithPayload adds the payload to the get application not found response
func (o *GetApplicationNotFound) WithPayload(payload string) *GetApplicationNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get application not found response
func (o *GetApplicationNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetApplicationNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetApplicationInternalServerErrorCode is the HTTP code returned for type GetApplicationInternalServerError
const GetApplicationInternalServerErrorCode int = 500

/*GetApplicationInternalServerError Internal error

swagger:response getApplicationInternalServerError
*/
type GetApplicationInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetApplicationInternalServerError creates GetApplicationInternalServerError with default headers values
func NewGetApplicationInternalServerError() *GetApplicationInternalServerError {

	return &GetApplicationInternalServerError{}
}

// WithPayload adds the payload to the get application internal server error response
func (o *GetApplicationInternalServerError) WithPayload(payload string) *GetApplicationInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get application internal server error response
func (o *GetApplicationInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetApplicationInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
