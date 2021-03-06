///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// DeleteOrganizationOKCode is the HTTP code returned for type DeleteOrganizationOK
const DeleteOrganizationOKCode int = 200

/*DeleteOrganizationOK Successful operation

swagger:response deleteOrganizationOK
*/
type DeleteOrganizationOK struct {

	/*
	  In: Body
	*/
	Payload *v1.Organization `json:"body,omitempty"`
}

// NewDeleteOrganizationOK creates DeleteOrganizationOK with default headers values
func NewDeleteOrganizationOK() *DeleteOrganizationOK {

	return &DeleteOrganizationOK{}
}

// WithPayload adds the payload to the delete organization o k response
func (o *DeleteOrganizationOK) WithPayload(payload *v1.Organization) *DeleteOrganizationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete organization o k response
func (o *DeleteOrganizationOK) SetPayload(payload *v1.Organization) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteOrganizationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteOrganizationBadRequestCode is the HTTP code returned for type DeleteOrganizationBadRequest
const DeleteOrganizationBadRequestCode int = 400

/*DeleteOrganizationBadRequest Invalid Name supplied

swagger:response deleteOrganizationBadRequest
*/
type DeleteOrganizationBadRequest struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewDeleteOrganizationBadRequest creates DeleteOrganizationBadRequest with default headers values
func NewDeleteOrganizationBadRequest() *DeleteOrganizationBadRequest {

	return &DeleteOrganizationBadRequest{}
}

// WithPayload adds the payload to the delete organization bad request response
func (o *DeleteOrganizationBadRequest) WithPayload(payload *v1.Error) *DeleteOrganizationBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete organization bad request response
func (o *DeleteOrganizationBadRequest) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteOrganizationBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteOrganizationNotFoundCode is the HTTP code returned for type DeleteOrganizationNotFound
const DeleteOrganizationNotFoundCode int = 404

/*DeleteOrganizationNotFound Organization not found

swagger:response deleteOrganizationNotFound
*/
type DeleteOrganizationNotFound struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewDeleteOrganizationNotFound creates DeleteOrganizationNotFound with default headers values
func NewDeleteOrganizationNotFound() *DeleteOrganizationNotFound {

	return &DeleteOrganizationNotFound{}
}

// WithPayload adds the payload to the delete organization not found response
func (o *DeleteOrganizationNotFound) WithPayload(payload *v1.Error) *DeleteOrganizationNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete organization not found response
func (o *DeleteOrganizationNotFound) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteOrganizationNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteOrganizationInternalServerErrorCode is the HTTP code returned for type DeleteOrganizationInternalServerError
const DeleteOrganizationInternalServerErrorCode int = 500

/*DeleteOrganizationInternalServerError Internal error

swagger:response deleteOrganizationInternalServerError
*/
type DeleteOrganizationInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewDeleteOrganizationInternalServerError creates DeleteOrganizationInternalServerError with default headers values
func NewDeleteOrganizationInternalServerError() *DeleteOrganizationInternalServerError {

	return &DeleteOrganizationInternalServerError{}
}

// WithPayload adds the payload to the delete organization internal server error response
func (o *DeleteOrganizationInternalServerError) WithPayload(payload *v1.Error) *DeleteOrganizationInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete organization internal server error response
func (o *DeleteOrganizationInternalServerError) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteOrganizationInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
