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

// DeleteBaseImageByNameOKCode is the HTTP code returned for type DeleteBaseImageByNameOK
const DeleteBaseImageByNameOKCode int = 200

/*DeleteBaseImageByNameOK successful operation

swagger:response deleteBaseImageByNameOK
*/
type DeleteBaseImageByNameOK struct {

	/*
	  In: Body
	*/
	Payload *v1.BaseImage `json:"body,omitempty"`
}

// NewDeleteBaseImageByNameOK creates DeleteBaseImageByNameOK with default headers values
func NewDeleteBaseImageByNameOK() *DeleteBaseImageByNameOK {

	return &DeleteBaseImageByNameOK{}
}

// WithPayload adds the payload to the delete base image by name o k response
func (o *DeleteBaseImageByNameOK) WithPayload(payload *v1.BaseImage) *DeleteBaseImageByNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete base image by name o k response
func (o *DeleteBaseImageByNameOK) SetPayload(payload *v1.BaseImage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteBaseImageByNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteBaseImageByNameBadRequestCode is the HTTP code returned for type DeleteBaseImageByNameBadRequest
const DeleteBaseImageByNameBadRequestCode int = 400

/*DeleteBaseImageByNameBadRequest Invalid ID supplied

swagger:response deleteBaseImageByNameBadRequest
*/
type DeleteBaseImageByNameBadRequest struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewDeleteBaseImageByNameBadRequest creates DeleteBaseImageByNameBadRequest with default headers values
func NewDeleteBaseImageByNameBadRequest() *DeleteBaseImageByNameBadRequest {

	return &DeleteBaseImageByNameBadRequest{}
}

// WithPayload adds the payload to the delete base image by name bad request response
func (o *DeleteBaseImageByNameBadRequest) WithPayload(payload *v1.Error) *DeleteBaseImageByNameBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete base image by name bad request response
func (o *DeleteBaseImageByNameBadRequest) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteBaseImageByNameBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteBaseImageByNameNotFoundCode is the HTTP code returned for type DeleteBaseImageByNameNotFound
const DeleteBaseImageByNameNotFoundCode int = 404

/*DeleteBaseImageByNameNotFound Base image not found

swagger:response deleteBaseImageByNameNotFound
*/
type DeleteBaseImageByNameNotFound struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewDeleteBaseImageByNameNotFound creates DeleteBaseImageByNameNotFound with default headers values
func NewDeleteBaseImageByNameNotFound() *DeleteBaseImageByNameNotFound {

	return &DeleteBaseImageByNameNotFound{}
}

// WithPayload adds the payload to the delete base image by name not found response
func (o *DeleteBaseImageByNameNotFound) WithPayload(payload *v1.Error) *DeleteBaseImageByNameNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete base image by name not found response
func (o *DeleteBaseImageByNameNotFound) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteBaseImageByNameNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteBaseImageByNameDefault Generic error response

swagger:response deleteBaseImageByNameDefault
*/
type DeleteBaseImageByNameDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewDeleteBaseImageByNameDefault creates DeleteBaseImageByNameDefault with default headers values
func NewDeleteBaseImageByNameDefault(code int) *DeleteBaseImageByNameDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteBaseImageByNameDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete base image by name default response
func (o *DeleteBaseImageByNameDefault) WithStatusCode(code int) *DeleteBaseImageByNameDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete base image by name default response
func (o *DeleteBaseImageByNameDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete base image by name default response
func (o *DeleteBaseImageByNameDefault) WithPayload(payload *v1.Error) *DeleteBaseImageByNameDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete base image by name default response
func (o *DeleteBaseImageByNameDefault) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteBaseImageByNameDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
