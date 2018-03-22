///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package drivers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/vmware/dispatch/pkg/event-manager/gen/models"
)

// GetDriverTypeOKCode is the HTTP code returned for type GetDriverTypeOK
const GetDriverTypeOKCode int = 200

/*GetDriverTypeOK Successful operation

swagger:response getDriverTypeOK
*/
type GetDriverTypeOK struct {

	/*
	  In: Body
	*/
	Payload *models.DriverType `json:"body,omitempty"`
}

// NewGetDriverTypeOK creates GetDriverTypeOK with default headers values
func NewGetDriverTypeOK() *GetDriverTypeOK {

	return &GetDriverTypeOK{}
}

// WithPayload adds the payload to the get driver type o k response
func (o *GetDriverTypeOK) WithPayload(payload *models.DriverType) *GetDriverTypeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get driver type o k response
func (o *GetDriverTypeOK) SetPayload(payload *models.DriverType) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDriverTypeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetDriverTypeBadRequestCode is the HTTP code returned for type GetDriverTypeBadRequest
const GetDriverTypeBadRequestCode int = 400

/*GetDriverTypeBadRequest Invalid Name supplied

swagger:response getDriverTypeBadRequest
*/
type GetDriverTypeBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetDriverTypeBadRequest creates GetDriverTypeBadRequest with default headers values
func NewGetDriverTypeBadRequest() *GetDriverTypeBadRequest {

	return &GetDriverTypeBadRequest{}
}

// WithPayload adds the payload to the get driver type bad request response
func (o *GetDriverTypeBadRequest) WithPayload(payload *models.Error) *GetDriverTypeBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get driver type bad request response
func (o *GetDriverTypeBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDriverTypeBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetDriverTypeNotFoundCode is the HTTP code returned for type GetDriverTypeNotFound
const GetDriverTypeNotFoundCode int = 404

/*GetDriverTypeNotFound Driver not found

swagger:response getDriverTypeNotFound
*/
type GetDriverTypeNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetDriverTypeNotFound creates GetDriverTypeNotFound with default headers values
func NewGetDriverTypeNotFound() *GetDriverTypeNotFound {

	return &GetDriverTypeNotFound{}
}

// WithPayload adds the payload to the get driver type not found response
func (o *GetDriverTypeNotFound) WithPayload(payload *models.Error) *GetDriverTypeNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get driver type not found response
func (o *GetDriverTypeNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDriverTypeNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetDriverTypeInternalServerErrorCode is the HTTP code returned for type GetDriverTypeInternalServerError
const GetDriverTypeInternalServerErrorCode int = 500

/*GetDriverTypeInternalServerError Internal server error

swagger:response getDriverTypeInternalServerError
*/
type GetDriverTypeInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetDriverTypeInternalServerError creates GetDriverTypeInternalServerError with default headers values
func NewGetDriverTypeInternalServerError() *GetDriverTypeInternalServerError {

	return &GetDriverTypeInternalServerError{}
}

// WithPayload adds the payload to the get driver type internal server error response
func (o *GetDriverTypeInternalServerError) WithPayload(payload *models.Error) *GetDriverTypeInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get driver type internal server error response
func (o *GetDriverTypeInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDriverTypeInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetDriverTypeDefault Unknown error

swagger:response getDriverTypeDefault
*/
type GetDriverTypeDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetDriverTypeDefault creates GetDriverTypeDefault with default headers values
func NewGetDriverTypeDefault(code int) *GetDriverTypeDefault {
	if code <= 0 {
		code = 500
	}

	return &GetDriverTypeDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get driver type default response
func (o *GetDriverTypeDefault) WithStatusCode(code int) *GetDriverTypeDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get driver type default response
func (o *GetDriverTypeDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get driver type default response
func (o *GetDriverTypeDefault) WithPayload(payload *models.Error) *GetDriverTypeDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get driver type default response
func (o *GetDriverTypeDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDriverTypeDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
