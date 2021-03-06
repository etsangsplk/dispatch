///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package base_image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// UpdateBaseImageByNameReader is a Reader for the UpdateBaseImageByName structure.
type UpdateBaseImageByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateBaseImageByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateBaseImageByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateBaseImageByNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateBaseImageByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUpdateBaseImageByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateBaseImageByNameOK creates a UpdateBaseImageByNameOK with default headers values
func NewUpdateBaseImageByNameOK() *UpdateBaseImageByNameOK {
	return &UpdateBaseImageByNameOK{}
}

/*UpdateBaseImageByNameOK handles this case with default header values.

successful operation
*/
type UpdateBaseImageByNameOK struct {
	Payload *v1.BaseImage
}

func (o *UpdateBaseImageByNameOK) Error() string {
	return fmt.Sprintf("[PUT /baseimage/{baseImageName}][%d] updateBaseImageByNameOK  %+v", 200, o.Payload)
}

func (o *UpdateBaseImageByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.BaseImage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBaseImageByNameBadRequest creates a UpdateBaseImageByNameBadRequest with default headers values
func NewUpdateBaseImageByNameBadRequest() *UpdateBaseImageByNameBadRequest {
	return &UpdateBaseImageByNameBadRequest{}
}

/*UpdateBaseImageByNameBadRequest handles this case with default header values.

Invalid input
*/
type UpdateBaseImageByNameBadRequest struct {
	Payload *v1.Error
}

func (o *UpdateBaseImageByNameBadRequest) Error() string {
	return fmt.Sprintf("[PUT /baseimage/{baseImageName}][%d] updateBaseImageByNameBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateBaseImageByNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBaseImageByNameNotFound creates a UpdateBaseImageByNameNotFound with default headers values
func NewUpdateBaseImageByNameNotFound() *UpdateBaseImageByNameNotFound {
	return &UpdateBaseImageByNameNotFound{}
}

/*UpdateBaseImageByNameNotFound handles this case with default header values.

Image not found
*/
type UpdateBaseImageByNameNotFound struct {
	Payload *v1.Error
}

func (o *UpdateBaseImageByNameNotFound) Error() string {
	return fmt.Sprintf("[PUT /baseimage/{baseImageName}][%d] updateBaseImageByNameNotFound  %+v", 404, o.Payload)
}

func (o *UpdateBaseImageByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBaseImageByNameDefault creates a UpdateBaseImageByNameDefault with default headers values
func NewUpdateBaseImageByNameDefault(code int) *UpdateBaseImageByNameDefault {
	return &UpdateBaseImageByNameDefault{
		_statusCode: code,
	}
}

/*UpdateBaseImageByNameDefault handles this case with default header values.

Generic error response
*/
type UpdateBaseImageByNameDefault struct {
	_statusCode int

	Payload *v1.Error
}

// Code gets the status code for the update base image by name default response
func (o *UpdateBaseImageByNameDefault) Code() int {
	return o._statusCode
}

func (o *UpdateBaseImageByNameDefault) Error() string {
	return fmt.Sprintf("[PUT /baseimage/{baseImageName}][%d] updateBaseImageByName default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateBaseImageByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
