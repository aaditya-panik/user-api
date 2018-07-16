// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "userapi/models"
)

// PatchOneOKCode is the HTTP code returned for type PatchOneOK
const PatchOneOKCode int = 200

/*PatchOneOK Patch user with specific id

swagger:response patchOneOK
*/
type PatchOneOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewPatchOneOK creates PatchOneOK with default headers values
func NewPatchOneOK() *PatchOneOK {

	return &PatchOneOK{}
}

// WithPayload adds the payload to the patch one o k response
func (o *PatchOneOK) WithPayload(payload *models.User) *PatchOneOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch one o k response
func (o *PatchOneOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchOneOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchOneBadRequestCode is the HTTP code returned for type PatchOneBadRequest
const PatchOneBadRequestCode int = 400

/*PatchOneBadRequest Invalid Patch Form - Bad Request

swagger:response patchOneBadRequest
*/
type PatchOneBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPatchOneBadRequest creates PatchOneBadRequest with default headers values
func NewPatchOneBadRequest() *PatchOneBadRequest {

	return &PatchOneBadRequest{}
}

// WithPayload adds the payload to the patch one bad request response
func (o *PatchOneBadRequest) WithPayload(payload *models.Error) *PatchOneBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch one bad request response
func (o *PatchOneBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchOneBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchOneNotFoundCode is the HTTP code returned for type PatchOneNotFound
const PatchOneNotFoundCode int = 404

/*PatchOneNotFound User Not Found

swagger:response patchOneNotFound
*/
type PatchOneNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPatchOneNotFound creates PatchOneNotFound with default headers values
func NewPatchOneNotFound() *PatchOneNotFound {

	return &PatchOneNotFound{}
}

// WithPayload adds the payload to the patch one not found response
func (o *PatchOneNotFound) WithPayload(payload *models.Error) *PatchOneNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch one not found response
func (o *PatchOneNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchOneNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PatchOneDefault Internal Server Error

swagger:response patchOneDefault
*/
type PatchOneDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPatchOneDefault creates PatchOneDefault with default headers values
func NewPatchOneDefault(code int) *PatchOneDefault {
	if code <= 0 {
		code = 500
	}

	return &PatchOneDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the patch one default response
func (o *PatchOneDefault) WithStatusCode(code int) *PatchOneDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the patch one default response
func (o *PatchOneDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the patch one default response
func (o *PatchOneDefault) WithPayload(payload *models.Error) *PatchOneDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch one default response
func (o *PatchOneDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchOneDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
