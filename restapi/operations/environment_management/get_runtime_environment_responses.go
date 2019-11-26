// Code generated by go-swagger; DO NOT EDIT.

package environment_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/models"
)

// GetRuntimeEnvironmentOKCode is the HTTP code returned for type GetRuntimeEnvironmentOK
const GetRuntimeEnvironmentOKCode int = 200

/*GetRuntimeEnvironmentOK RuntimeEnvironments

swagger:response getRuntimeEnvironmentOK
*/
type GetRuntimeEnvironmentOK struct {

	/*
	  In: Body
	*/
	Payload *models.RuntimeEnvironment `json:"body,omitempty"`
}

// NewGetRuntimeEnvironmentOK creates GetRuntimeEnvironmentOK with default headers values
func NewGetRuntimeEnvironmentOK() *GetRuntimeEnvironmentOK {

	return &GetRuntimeEnvironmentOK{}
}

// WithPayload adds the payload to the get runtime environment o k response
func (o *GetRuntimeEnvironmentOK) WithPayload(payload *models.RuntimeEnvironment) *GetRuntimeEnvironmentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get runtime environment o k response
func (o *GetRuntimeEnvironmentOK) SetPayload(payload *models.RuntimeEnvironment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRuntimeEnvironmentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRuntimeEnvironmentBadRequestCode is the HTTP code returned for type GetRuntimeEnvironmentBadRequest
const GetRuntimeEnvironmentBadRequestCode int = 400

/*GetRuntimeEnvironmentBadRequest Bad request

swagger:response getRuntimeEnvironmentBadRequest
*/
type GetRuntimeEnvironmentBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetRuntimeEnvironmentBadRequest creates GetRuntimeEnvironmentBadRequest with default headers values
func NewGetRuntimeEnvironmentBadRequest() *GetRuntimeEnvironmentBadRequest {

	return &GetRuntimeEnvironmentBadRequest{}
}

// WithPayload adds the payload to the get runtime environment bad request response
func (o *GetRuntimeEnvironmentBadRequest) WithPayload(payload string) *GetRuntimeEnvironmentBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get runtime environment bad request response
func (o *GetRuntimeEnvironmentBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRuntimeEnvironmentBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetRuntimeEnvironmentNotFoundCode is the HTTP code returned for type GetRuntimeEnvironmentNotFound
const GetRuntimeEnvironmentNotFoundCode int = 404

/*GetRuntimeEnvironmentNotFound Not found

swagger:response getRuntimeEnvironmentNotFound
*/
type GetRuntimeEnvironmentNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetRuntimeEnvironmentNotFound creates GetRuntimeEnvironmentNotFound with default headers values
func NewGetRuntimeEnvironmentNotFound() *GetRuntimeEnvironmentNotFound {

	return &GetRuntimeEnvironmentNotFound{}
}

// WithPayload adds the payload to the get runtime environment not found response
func (o *GetRuntimeEnvironmentNotFound) WithPayload(payload string) *GetRuntimeEnvironmentNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get runtime environment not found response
func (o *GetRuntimeEnvironmentNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRuntimeEnvironmentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetRuntimeEnvironmentInternalServerErrorCode is the HTTP code returned for type GetRuntimeEnvironmentInternalServerError
const GetRuntimeEnvironmentInternalServerErrorCode int = 500

/*GetRuntimeEnvironmentInternalServerError Internal error

swagger:response getRuntimeEnvironmentInternalServerError
*/
type GetRuntimeEnvironmentInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetRuntimeEnvironmentInternalServerError creates GetRuntimeEnvironmentInternalServerError with default headers values
func NewGetRuntimeEnvironmentInternalServerError() *GetRuntimeEnvironmentInternalServerError {

	return &GetRuntimeEnvironmentInternalServerError{}
}

// WithPayload adds the payload to the get runtime environment internal server error response
func (o *GetRuntimeEnvironmentInternalServerError) WithPayload(payload string) *GetRuntimeEnvironmentInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get runtime environment internal server error response
func (o *GetRuntimeEnvironmentInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRuntimeEnvironmentInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}