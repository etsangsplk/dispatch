///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package base_image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// UpdateBaseImageByNameOKCode is the HTTP code returned for type UpdateBaseImageByNameOK
const UpdateBaseImageByNameOKCode int = 200

/*UpdateBaseImageByNameOK successful operation

swagger:response updateBaseImageByNameOK
*/
type UpdateBaseImageByNameOK struct {

	/*
	  In: Body
	*/
	Payload *v1.BaseImage `json:"body,omitempty"`
}

// NewUpdateBaseImageByNameOK creates UpdateBaseImageByNameOK with default headers values
func NewUpdateBaseImageByNameOK() *UpdateBaseImageByNameOK {

	return &UpdateBaseImageByNameOK{}
}

// WithPayload adds the payload to the update base image by name o k response
func (o *UpdateBaseImageByNameOK) WithPayload(payload *v1.BaseImage) *UpdateBaseImageByNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update base image by name o k response
func (o *UpdateBaseImageByNameOK) SetPayload(payload *v1.BaseImage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateBaseImageByNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateBaseImageByNameBadRequestCode is the HTTP code returned for type UpdateBaseImageByNameBadRequest
const UpdateBaseImageByNameBadRequestCode int = 400

/*UpdateBaseImageByNameBadRequest Invalid input

swagger:response updateBaseImageByNameBadRequest
*/
type UpdateBaseImageByNameBadRequest struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewUpdateBaseImageByNameBadRequest creates UpdateBaseImageByNameBadRequest with default headers values
func NewUpdateBaseImageByNameBadRequest() *UpdateBaseImageByNameBadRequest {

	return &UpdateBaseImageByNameBadRequest{}
}

// WithPayload adds the payload to the update base image by name bad request response
func (o *UpdateBaseImageByNameBadRequest) WithPayload(payload *v1.Error) *UpdateBaseImageByNameBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update base image by name bad request response
func (o *UpdateBaseImageByNameBadRequest) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateBaseImageByNameBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateBaseImageByNameNotFoundCode is the HTTP code returned for type UpdateBaseImageByNameNotFound
const UpdateBaseImageByNameNotFoundCode int = 404

/*UpdateBaseImageByNameNotFound Image not found

swagger:response updateBaseImageByNameNotFound
*/
type UpdateBaseImageByNameNotFound struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewUpdateBaseImageByNameNotFound creates UpdateBaseImageByNameNotFound with default headers values
func NewUpdateBaseImageByNameNotFound() *UpdateBaseImageByNameNotFound {

	return &UpdateBaseImageByNameNotFound{}
}

// WithPayload adds the payload to the update base image by name not found response
func (o *UpdateBaseImageByNameNotFound) WithPayload(payload *v1.Error) *UpdateBaseImageByNameNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update base image by name not found response
func (o *UpdateBaseImageByNameNotFound) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateBaseImageByNameNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UpdateBaseImageByNameDefault Generic error response

swagger:response updateBaseImageByNameDefault
*/
type UpdateBaseImageByNameDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewUpdateBaseImageByNameDefault creates UpdateBaseImageByNameDefault with default headers values
func NewUpdateBaseImageByNameDefault(code int) *UpdateBaseImageByNameDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateBaseImageByNameDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update base image by name default response
func (o *UpdateBaseImageByNameDefault) WithStatusCode(code int) *UpdateBaseImageByNameDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update base image by name default response
func (o *UpdateBaseImageByNameDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update base image by name default response
func (o *UpdateBaseImageByNameDefault) WithPayload(payload *v1.Error) *UpdateBaseImageByNameDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update base image by name default response
func (o *UpdateBaseImageByNameDefault) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateBaseImageByNameDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
