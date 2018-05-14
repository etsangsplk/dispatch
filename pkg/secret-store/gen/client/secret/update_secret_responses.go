///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package secret

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// UpdateSecretReader is a Reader for the UpdateSecret structure.
type UpdateSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewUpdateSecretCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateSecretBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateSecretNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUpdateSecretDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSecretCreated creates a UpdateSecretCreated with default headers values
func NewUpdateSecretCreated() *UpdateSecretCreated {
	return &UpdateSecretCreated{}
}

/*UpdateSecretCreated handles this case with default header values.

The updated secret
*/
type UpdateSecretCreated struct {
	Payload *v1.Secret
}

func (o *UpdateSecretCreated) Error() string {
	return fmt.Sprintf("[PUT /{secretName}][%d] updateSecretCreated  %+v", 201, o.Payload)
}

func (o *UpdateSecretCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Secret)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSecretBadRequest creates a UpdateSecretBadRequest with default headers values
func NewUpdateSecretBadRequest() *UpdateSecretBadRequest {
	return &UpdateSecretBadRequest{}
}

/*UpdateSecretBadRequest handles this case with default header values.

Bad Request
*/
type UpdateSecretBadRequest struct {
	Payload *v1.Error
}

func (o *UpdateSecretBadRequest) Error() string {
	return fmt.Sprintf("[PUT /{secretName}][%d] updateSecretBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateSecretBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSecretNotFound creates a UpdateSecretNotFound with default headers values
func NewUpdateSecretNotFound() *UpdateSecretNotFound {
	return &UpdateSecretNotFound{}
}

/*UpdateSecretNotFound handles this case with default header values.

Resource Not Found if no secret exists with the given name
*/
type UpdateSecretNotFound struct {
	Payload *v1.Error
}

func (o *UpdateSecretNotFound) Error() string {
	return fmt.Sprintf("[PUT /{secretName}][%d] updateSecretNotFound  %+v", 404, o.Payload)
}

func (o *UpdateSecretNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSecretDefault creates a UpdateSecretDefault with default headers values
func NewUpdateSecretDefault(code int) *UpdateSecretDefault {
	return &UpdateSecretDefault{
		_statusCode: code,
	}
}

/*UpdateSecretDefault handles this case with default header values.

generic error
*/
type UpdateSecretDefault struct {
	_statusCode int

	Payload *v1.Error
}

// Code gets the status code for the update secret default response
func (o *UpdateSecretDefault) Code() int {
	return o._statusCode
}

func (o *UpdateSecretDefault) Error() string {
	return fmt.Sprintf("[PUT /{secretName}][%d] updateSecret default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSecretDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
