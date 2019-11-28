// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/darwinnn/TestTask/src/models"
)

// PostStateOKCode is the HTTP code returned for type PostStateOK
const PostStateOKCode int = 200

/*PostStateOK state updates

swagger:response postStateOK
*/
type PostStateOK struct {

	/*
	  In: Body
	*/
	Payload *models.StateObj `json:"body,omitempty"`
}

// NewPostStateOK creates PostStateOK with default headers values
func NewPostStateOK() *PostStateOK {

	return &PostStateOK{}
}

// WithPayload adds the payload to the post state o k response
func (o *PostStateOK) WithPayload(payload *models.StateObj) *PostStateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post state o k response
func (o *PostStateOK) SetPayload(payload *models.StateObj) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostStateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PostStateDefault error message

swagger:response postStateDefault
*/
type PostStateDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostStateDefault creates PostStateDefault with default headers values
func NewPostStateDefault(code int) *PostStateDefault {
	if code <= 0 {
		code = 500
	}

	return &PostStateDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post state default response
func (o *PostStateDefault) WithStatusCode(code int) *PostStateDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post state default response
func (o *PostStateDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post state default response
func (o *PostStateDefault) WithPayload(payload *models.Error) *PostStateDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post state default response
func (o *PostStateDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostStateDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
