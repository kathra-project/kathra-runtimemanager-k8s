// Code generated by go-swagger; DO NOT EDIT.

package backup_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/models"
)

// AddBackupRuntimeEnvironmentsOKCode is the HTTP code returned for type AddBackupRuntimeEnvironmentsOK
const AddBackupRuntimeEnvironmentsOKCode int = 200

/*AddBackupRuntimeEnvironmentsOK RuntimeAppInstance deployed with success

swagger:response addBackupRuntimeEnvironmentsOK
*/
type AddBackupRuntimeEnvironmentsOK struct {

	/*
	  In: Body
	*/
	Payload *models.RuntimeEnvironmentBackup `json:"body,omitempty"`
}

// NewAddBackupRuntimeEnvironmentsOK creates AddBackupRuntimeEnvironmentsOK with default headers values
func NewAddBackupRuntimeEnvironmentsOK() *AddBackupRuntimeEnvironmentsOK {

	return &AddBackupRuntimeEnvironmentsOK{}
}

// WithPayload adds the payload to the add backup runtime environments o k response
func (o *AddBackupRuntimeEnvironmentsOK) WithPayload(payload *models.RuntimeEnvironmentBackup) *AddBackupRuntimeEnvironmentsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add backup runtime environments o k response
func (o *AddBackupRuntimeEnvironmentsOK) SetPayload(payload *models.RuntimeEnvironmentBackup) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddBackupRuntimeEnvironmentsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddBackupRuntimeEnvironmentsBadRequestCode is the HTTP code returned for type AddBackupRuntimeEnvironmentsBadRequest
const AddBackupRuntimeEnvironmentsBadRequestCode int = 400

/*AddBackupRuntimeEnvironmentsBadRequest Bad request

swagger:response addBackupRuntimeEnvironmentsBadRequest
*/
type AddBackupRuntimeEnvironmentsBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewAddBackupRuntimeEnvironmentsBadRequest creates AddBackupRuntimeEnvironmentsBadRequest with default headers values
func NewAddBackupRuntimeEnvironmentsBadRequest() *AddBackupRuntimeEnvironmentsBadRequest {

	return &AddBackupRuntimeEnvironmentsBadRequest{}
}

// WithPayload adds the payload to the add backup runtime environments bad request response
func (o *AddBackupRuntimeEnvironmentsBadRequest) WithPayload(payload string) *AddBackupRuntimeEnvironmentsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add backup runtime environments bad request response
func (o *AddBackupRuntimeEnvironmentsBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddBackupRuntimeEnvironmentsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// AddBackupRuntimeEnvironmentsInternalServerErrorCode is the HTTP code returned for type AddBackupRuntimeEnvironmentsInternalServerError
const AddBackupRuntimeEnvironmentsInternalServerErrorCode int = 500

/*AddBackupRuntimeEnvironmentsInternalServerError Internal error

swagger:response addBackupRuntimeEnvironmentsInternalServerError
*/
type AddBackupRuntimeEnvironmentsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewAddBackupRuntimeEnvironmentsInternalServerError creates AddBackupRuntimeEnvironmentsInternalServerError with default headers values
func NewAddBackupRuntimeEnvironmentsInternalServerError() *AddBackupRuntimeEnvironmentsInternalServerError {

	return &AddBackupRuntimeEnvironmentsInternalServerError{}
}

// WithPayload adds the payload to the add backup runtime environments internal server error response
func (o *AddBackupRuntimeEnvironmentsInternalServerError) WithPayload(payload string) *AddBackupRuntimeEnvironmentsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add backup runtime environments internal server error response
func (o *AddBackupRuntimeEnvironmentsInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddBackupRuntimeEnvironmentsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
