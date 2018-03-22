///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/dispatch/pkg/api-manager/gen/models"
)

// GetApisReader is a Reader for the GetApis structure.
type GetApisReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApisReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetApisOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetApisInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetApisDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetApisOK creates a GetApisOK with default headers values
func NewGetApisOK() *GetApisOK {
	return &GetApisOK{}
}

/*GetApisOK handles this case with default header values.

Successful operation
*/
type GetApisOK struct {
	Payload []*models.API
}

func (o *GetApisOK) Error() string {
	return fmt.Sprintf("[GET /][%d] getApisOK  %+v", 200, o.Payload)
}

func (o *GetApisOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApisInternalServerError creates a GetApisInternalServerError with default headers values
func NewGetApisInternalServerError() *GetApisInternalServerError {
	return &GetApisInternalServerError{}
}

/*GetApisInternalServerError handles this case with default header values.

Internal Error
*/
type GetApisInternalServerError struct {
	Payload *models.Error
}

func (o *GetApisInternalServerError) Error() string {
	return fmt.Sprintf("[GET /][%d] getApisInternalServerError  %+v", 500, o.Payload)
}

func (o *GetApisInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApisDefault creates a GetApisDefault with default headers values
func NewGetApisDefault(code int) *GetApisDefault {
	return &GetApisDefault{
		_statusCode: code,
	}
}

/*GetApisDefault handles this case with default header values.

Unexpected Error
*/
type GetApisDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get apis default response
func (o *GetApisDefault) Code() int {
	return o._statusCode
}

func (o *GetApisDefault) Error() string {
	return fmt.Sprintf("[GET /][%d] getAPIs default  %+v", o._statusCode, o.Payload)
}

func (o *GetApisDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
