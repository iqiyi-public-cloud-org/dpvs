// Code generated by go-swagger; DO NOT EDIT.

package ipset

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// IsInOKCode is the HTTP code returned for type IsInOK
const IsInOKCode int = 200

/*
IsInOK Succeed

swagger:response isInOK
*/
type IsInOK struct {

	/*
	  In: Body
	*/
	Payload *IsInOKBody `json:"body,omitempty"`
}

// NewIsInOK creates IsInOK with default headers values
func NewIsInOK() *IsInOK {

	return &IsInOK{}
}

// WithPayload adds the payload to the is in o k response
func (o *IsInOK) WithPayload(payload *IsInOKBody) *IsInOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the is in o k response
func (o *IsInOK) SetPayload(payload *IsInOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *IsInOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// IsInBadRequestCode is the HTTP code returned for type IsInBadRequest
const IsInBadRequestCode int = 400

/*
IsInBadRequest Invalid ipset parameter

swagger:response isInBadRequest
*/
type IsInBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewIsInBadRequest creates IsInBadRequest with default headers values
func NewIsInBadRequest() *IsInBadRequest {

	return &IsInBadRequest{}
}

// WithPayload adds the payload to the is in bad request response
func (o *IsInBadRequest) WithPayload(payload string) *IsInBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the is in bad request response
func (o *IsInBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *IsInBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// IsInNotFoundCode is the HTTP code returned for type IsInNotFound
const IsInNotFoundCode int = 404

/*
IsInNotFound Ipset not found

swagger:response isInNotFound
*/
type IsInNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewIsInNotFound creates IsInNotFound with default headers values
func NewIsInNotFound() *IsInNotFound {

	return &IsInNotFound{}
}

// WithPayload adds the payload to the is in not found response
func (o *IsInNotFound) WithPayload(payload string) *IsInNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the is in not found response
func (o *IsInNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *IsInNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// IsInFailureCode is the HTTP code returned for type IsInFailure
const IsInFailureCode int = 500

/*
IsInFailure Service not available

swagger:response isInFailure
*/
type IsInFailure struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewIsInFailure creates IsInFailure with default headers values
func NewIsInFailure() *IsInFailure {

	return &IsInFailure{}
}

// WithPayload adds the payload to the is in failure response
func (o *IsInFailure) WithPayload(payload string) *IsInFailure {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the is in failure response
func (o *IsInFailure) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *IsInFailure) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
