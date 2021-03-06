///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package serviceaccount

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// GetServiceAccountOKCode is the HTTP code returned for type GetServiceAccountOK
const GetServiceAccountOKCode int = 200

/*GetServiceAccountOK Successful operation

swagger:response getServiceAccountOK
*/
type GetServiceAccountOK struct {

	/*
	  In: Body
	*/
	Payload *v1.ServiceAccount `json:"body,omitempty"`
}

// NewGetServiceAccountOK creates GetServiceAccountOK with default headers values
func NewGetServiceAccountOK() *GetServiceAccountOK {

	return &GetServiceAccountOK{}
}

// WithPayload adds the payload to the get service account o k response
func (o *GetServiceAccountOK) WithPayload(payload *v1.ServiceAccount) *GetServiceAccountOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get service account o k response
func (o *GetServiceAccountOK) SetPayload(payload *v1.ServiceAccount) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServiceAccountOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetServiceAccountBadRequestCode is the HTTP code returned for type GetServiceAccountBadRequest
const GetServiceAccountBadRequestCode int = 400

/*GetServiceAccountBadRequest Invalid Name supplied

swagger:response getServiceAccountBadRequest
*/
type GetServiceAccountBadRequest struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetServiceAccountBadRequest creates GetServiceAccountBadRequest with default headers values
func NewGetServiceAccountBadRequest() *GetServiceAccountBadRequest {

	return &GetServiceAccountBadRequest{}
}

// WithPayload adds the payload to the get service account bad request response
func (o *GetServiceAccountBadRequest) WithPayload(payload *v1.Error) *GetServiceAccountBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get service account bad request response
func (o *GetServiceAccountBadRequest) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServiceAccountBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetServiceAccountNotFoundCode is the HTTP code returned for type GetServiceAccountNotFound
const GetServiceAccountNotFoundCode int = 404

/*GetServiceAccountNotFound Service Account not found

swagger:response getServiceAccountNotFound
*/
type GetServiceAccountNotFound struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetServiceAccountNotFound creates GetServiceAccountNotFound with default headers values
func NewGetServiceAccountNotFound() *GetServiceAccountNotFound {

	return &GetServiceAccountNotFound{}
}

// WithPayload adds the payload to the get service account not found response
func (o *GetServiceAccountNotFound) WithPayload(payload *v1.Error) *GetServiceAccountNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get service account not found response
func (o *GetServiceAccountNotFound) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServiceAccountNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetServiceAccountInternalServerErrorCode is the HTTP code returned for type GetServiceAccountInternalServerError
const GetServiceAccountInternalServerErrorCode int = 500

/*GetServiceAccountInternalServerError Internal error

swagger:response getServiceAccountInternalServerError
*/
type GetServiceAccountInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetServiceAccountInternalServerError creates GetServiceAccountInternalServerError with default headers values
func NewGetServiceAccountInternalServerError() *GetServiceAccountInternalServerError {

	return &GetServiceAccountInternalServerError{}
}

// WithPayload adds the payload to the get service account internal server error response
func (o *GetServiceAccountInternalServerError) WithPayload(payload *v1.Error) *GetServiceAccountInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get service account internal server error response
func (o *GetServiceAccountInternalServerError) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServiceAccountInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
