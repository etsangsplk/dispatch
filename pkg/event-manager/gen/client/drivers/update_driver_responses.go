///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package drivers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// UpdateDriverReader is a Reader for the UpdateDriver structure.
type UpdateDriverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDriverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateDriverOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateDriverBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateDriverNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUpdateDriverInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUpdateDriverDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateDriverOK creates a UpdateDriverOK with default headers values
func NewUpdateDriverOK() *UpdateDriverOK {
	return &UpdateDriverOK{}
}

/*UpdateDriverOK handles this case with default header values.

Successful operation
*/
type UpdateDriverOK struct {
	Payload *v1.EventDriver
}

func (o *UpdateDriverOK) Error() string {
	return fmt.Sprintf("[PUT /drivers/{driverName}][%d] updateDriverOK  %+v", 200, o.Payload)
}

func (o *UpdateDriverOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.EventDriver)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDriverBadRequest creates a UpdateDriverBadRequest with default headers values
func NewUpdateDriverBadRequest() *UpdateDriverBadRequest {
	return &UpdateDriverBadRequest{}
}

/*UpdateDriverBadRequest handles this case with default header values.

Invalid Name supplied
*/
type UpdateDriverBadRequest struct {
	Payload *v1.Error
}

func (o *UpdateDriverBadRequest) Error() string {
	return fmt.Sprintf("[PUT /drivers/{driverName}][%d] updateDriverBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateDriverBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDriverNotFound creates a UpdateDriverNotFound with default headers values
func NewUpdateDriverNotFound() *UpdateDriverNotFound {
	return &UpdateDriverNotFound{}
}

/*UpdateDriverNotFound handles this case with default header values.

Driver not found
*/
type UpdateDriverNotFound struct {
	Payload *v1.Error
}

func (o *UpdateDriverNotFound) Error() string {
	return fmt.Sprintf("[PUT /drivers/{driverName}][%d] updateDriverNotFound  %+v", 404, o.Payload)
}

func (o *UpdateDriverNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDriverInternalServerError creates a UpdateDriverInternalServerError with default headers values
func NewUpdateDriverInternalServerError() *UpdateDriverInternalServerError {
	return &UpdateDriverInternalServerError{}
}

/*UpdateDriverInternalServerError handles this case with default header values.

Internal server error
*/
type UpdateDriverInternalServerError struct {
	Payload *v1.Error
}

func (o *UpdateDriverInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /drivers/{driverName}][%d] updateDriverInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateDriverInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDriverDefault creates a UpdateDriverDefault with default headers values
func NewUpdateDriverDefault(code int) *UpdateDriverDefault {
	return &UpdateDriverDefault{
		_statusCode: code,
	}
}

/*UpdateDriverDefault handles this case with default header values.

Unknown error
*/
type UpdateDriverDefault struct {
	_statusCode int

	Payload *v1.Error
}

// Code gets the status code for the update driver default response
func (o *UpdateDriverDefault) Code() int {
	return o._statusCode
}

func (o *UpdateDriverDefault) Error() string {
	return fmt.Sprintf("[PUT /drivers/{driverName}][%d] updateDriver default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateDriverDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
