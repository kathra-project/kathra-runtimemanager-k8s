// Code generated by go-swagger; DO NOT EDIT.

package backup_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/models"
)

// RestoreBackupRuntimeEnvironmentsOKCode is the HTTP code returned for type RestoreBackupRuntimeEnvironmentsOK
const RestoreBackupRuntimeEnvironmentsOKCode int = 200

/*RestoreBackupRuntimeEnvironmentsOK RuntimeAppInstance restored with success

swagger:response restoreBackupRuntimeEnvironmentsOK
*/
type RestoreBackupRuntimeEnvironmentsOK struct {

	/*
	  In: Body
	*/
	Payload *models.RuntimeEnvironmentBackup `json:"body,omitempty"`
}

// NewRestoreBackupRuntimeEnvironmentsOK creates RestoreBackupRuntimeEnvironmentsOK with default headers values
func NewRestoreBackupRuntimeEnvironmentsOK() *RestoreBackupRuntimeEnvironmentsOK {

	return &RestoreBackupRuntimeEnvironmentsOK{}
}

// WithPayload adds the payload to the restore backup runtime environments o k response
func (o *RestoreBackupRuntimeEnvironmentsOK) WithPayload(payload *models.RuntimeEnvironmentBackup) *RestoreBackupRuntimeEnvironmentsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the restore backup runtime environments o k response
func (o *RestoreBackupRuntimeEnvironmentsOK) SetPayload(payload *models.RuntimeEnvironmentBackup) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RestoreBackupRuntimeEnvironmentsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RestoreBackupRuntimeEnvironmentsBadRequestCode is the HTTP code returned for type RestoreBackupRuntimeEnvironmentsBadRequest
const RestoreBackupRuntimeEnvironmentsBadRequestCode int = 400

/*RestoreBackupRuntimeEnvironmentsBadRequest Bad request

swagger:response restoreBackupRuntimeEnvironmentsBadRequest
*/
type RestoreBackupRuntimeEnvironmentsBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewRestoreBackupRuntimeEnvironmentsBadRequest creates RestoreBackupRuntimeEnvironmentsBadRequest with default headers values
func NewRestoreBackupRuntimeEnvironmentsBadRequest() *RestoreBackupRuntimeEnvironmentsBadRequest {

	return &RestoreBackupRuntimeEnvironmentsBadRequest{}
}

// WithPayload adds the payload to the restore backup runtime environments bad request response
func (o *RestoreBackupRuntimeEnvironmentsBadRequest) WithPayload(payload string) *RestoreBackupRuntimeEnvironmentsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the restore backup runtime environments bad request response
func (o *RestoreBackupRuntimeEnvironmentsBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RestoreBackupRuntimeEnvironmentsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// RestoreBackupRuntimeEnvironmentsInternalServerErrorCode is the HTTP code returned for type RestoreBackupRuntimeEnvironmentsInternalServerError
const RestoreBackupRuntimeEnvironmentsInternalServerErrorCode int = 500

/*RestoreBackupRuntimeEnvironmentsInternalServerError Internal error

swagger:response restoreBackupRuntimeEnvironmentsInternalServerError
*/
type RestoreBackupRuntimeEnvironmentsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewRestoreBackupRuntimeEnvironmentsInternalServerError creates RestoreBackupRuntimeEnvironmentsInternalServerError with default headers values
func NewRestoreBackupRuntimeEnvironmentsInternalServerError() *RestoreBackupRuntimeEnvironmentsInternalServerError {

	return &RestoreBackupRuntimeEnvironmentsInternalServerError{}
}

// WithPayload adds the payload to the restore backup runtime environments internal server error response
func (o *RestoreBackupRuntimeEnvironmentsInternalServerError) WithPayload(payload string) *RestoreBackupRuntimeEnvironmentsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the restore backup runtime environments internal server error response
func (o *RestoreBackupRuntimeEnvironmentsInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RestoreBackupRuntimeEnvironmentsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}