///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// UpdateSubscriptionOKCode is the HTTP code returned for type UpdateSubscriptionOK
const UpdateSubscriptionOKCode int = 200

/*UpdateSubscriptionOK Successful operation

swagger:response updateSubscriptionOK
*/
type UpdateSubscriptionOK struct {

	/*
	  In: Body
	*/
	Payload *v1.Subscription `json:"body,omitempty"`
}

// NewUpdateSubscriptionOK creates UpdateSubscriptionOK with default headers values
func NewUpdateSubscriptionOK() *UpdateSubscriptionOK {

	return &UpdateSubscriptionOK{}
}

// WithPayload adds the payload to the update subscription o k response
func (o *UpdateSubscriptionOK) WithPayload(payload *v1.Subscription) *UpdateSubscriptionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update subscription o k response
func (o *UpdateSubscriptionOK) SetPayload(payload *v1.Subscription) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateSubscriptionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateSubscriptionBadRequestCode is the HTTP code returned for type UpdateSubscriptionBadRequest
const UpdateSubscriptionBadRequestCode int = 400

/*UpdateSubscriptionBadRequest Invalid Name supplied

swagger:response updateSubscriptionBadRequest
*/
type UpdateSubscriptionBadRequest struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewUpdateSubscriptionBadRequest creates UpdateSubscriptionBadRequest with default headers values
func NewUpdateSubscriptionBadRequest() *UpdateSubscriptionBadRequest {

	return &UpdateSubscriptionBadRequest{}
}

// WithPayload adds the payload to the update subscription bad request response
func (o *UpdateSubscriptionBadRequest) WithPayload(payload *v1.Error) *UpdateSubscriptionBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update subscription bad request response
func (o *UpdateSubscriptionBadRequest) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateSubscriptionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateSubscriptionNotFoundCode is the HTTP code returned for type UpdateSubscriptionNotFound
const UpdateSubscriptionNotFoundCode int = 404

/*UpdateSubscriptionNotFound Subscription not found

swagger:response updateSubscriptionNotFound
*/
type UpdateSubscriptionNotFound struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewUpdateSubscriptionNotFound creates UpdateSubscriptionNotFound with default headers values
func NewUpdateSubscriptionNotFound() *UpdateSubscriptionNotFound {

	return &UpdateSubscriptionNotFound{}
}

// WithPayload adds the payload to the update subscription not found response
func (o *UpdateSubscriptionNotFound) WithPayload(payload *v1.Error) *UpdateSubscriptionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update subscription not found response
func (o *UpdateSubscriptionNotFound) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateSubscriptionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateSubscriptionInternalServerErrorCode is the HTTP code returned for type UpdateSubscriptionInternalServerError
const UpdateSubscriptionInternalServerErrorCode int = 500

/*UpdateSubscriptionInternalServerError Internal server error

swagger:response updateSubscriptionInternalServerError
*/
type UpdateSubscriptionInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewUpdateSubscriptionInternalServerError creates UpdateSubscriptionInternalServerError with default headers values
func NewUpdateSubscriptionInternalServerError() *UpdateSubscriptionInternalServerError {

	return &UpdateSubscriptionInternalServerError{}
}

// WithPayload adds the payload to the update subscription internal server error response
func (o *UpdateSubscriptionInternalServerError) WithPayload(payload *v1.Error) *UpdateSubscriptionInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update subscription internal server error response
func (o *UpdateSubscriptionInternalServerError) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateSubscriptionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UpdateSubscriptionDefault Unknown error

swagger:response updateSubscriptionDefault
*/
type UpdateSubscriptionDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewUpdateSubscriptionDefault creates UpdateSubscriptionDefault with default headers values
func NewUpdateSubscriptionDefault(code int) *UpdateSubscriptionDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateSubscriptionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update subscription default response
func (o *UpdateSubscriptionDefault) WithStatusCode(code int) *UpdateSubscriptionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update subscription default response
func (o *UpdateSubscriptionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update subscription default response
func (o *UpdateSubscriptionDefault) WithPayload(payload *v1.Error) *UpdateSubscriptionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update subscription default response
func (o *UpdateSubscriptionDefault) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateSubscriptionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
