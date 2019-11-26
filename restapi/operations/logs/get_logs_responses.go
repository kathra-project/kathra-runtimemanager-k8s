// Code generated by go-swagger; DO NOT EDIT.

package logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetLogsOKCode is the HTTP code returned for type GetLogsOK
const GetLogsOKCode int = 200

/*GetLogsOK RuntimeAppService

swagger:response getLogsOK
*/
type GetLogsOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetLogsOK creates GetLogsOK with default headers values
func NewGetLogsOK() *GetLogsOK {

	return &GetLogsOK{}
}

// WithPayload adds the payload to the get logs o k response
func (o *GetLogsOK) WithPayload(payload string) *GetLogsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get logs o k response
func (o *GetLogsOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLogsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetLogsBadRequestCode is the HTTP code returned for type GetLogsBadRequest
const GetLogsBadRequestCode int = 400

/*GetLogsBadRequest Bad request

swagger:response getLogsBadRequest
*/
type GetLogsBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetLogsBadRequest creates GetLogsBadRequest with default headers values
func NewGetLogsBadRequest() *GetLogsBadRequest {

	return &GetLogsBadRequest{}
}

// WithPayload adds the payload to the get logs bad request response
func (o *GetLogsBadRequest) WithPayload(payload string) *GetLogsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get logs bad request response
func (o *GetLogsBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLogsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetLogsNotFoundCode is the HTTP code returned for type GetLogsNotFound
const GetLogsNotFoundCode int = 404

/*GetLogsNotFound Not found

swagger:response getLogsNotFound
*/
type GetLogsNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetLogsNotFound creates GetLogsNotFound with default headers values
func NewGetLogsNotFound() *GetLogsNotFound {

	return &GetLogsNotFound{}
}

// WithPayload adds the payload to the get logs not found response
func (o *GetLogsNotFound) WithPayload(payload string) *GetLogsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get logs not found response
func (o *GetLogsNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLogsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetLogsInternalServerErrorCode is the HTTP code returned for type GetLogsInternalServerError
const GetLogsInternalServerErrorCode int = 500

/*GetLogsInternalServerError Internal error

swagger:response getLogsInternalServerError
*/
type GetLogsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetLogsInternalServerError creates GetLogsInternalServerError with default headers values
func NewGetLogsInternalServerError() *GetLogsInternalServerError {

	return &GetLogsInternalServerError{}
}

// WithPayload adds the payload to the get logs internal server error response
func (o *GetLogsInternalServerError) WithPayload(payload string) *GetLogsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get logs internal server error response
func (o *GetLogsInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLogsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
