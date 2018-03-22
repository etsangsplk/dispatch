///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/dispatch/pkg/function-manager/gen/models"
)

// GetFunctionReader is a Reader for the GetFunction structure.
type GetFunctionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFunctionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFunctionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetFunctionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetFunctionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetFunctionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFunctionOK creates a GetFunctionOK with default headers values
func NewGetFunctionOK() *GetFunctionOK {
	return &GetFunctionOK{}
}

/*GetFunctionOK handles this case with default header values.

Successful operation
*/
type GetFunctionOK struct {
	Payload *models.Function
}

func (o *GetFunctionOK) Error() string {
	return fmt.Sprintf("[GET /function/{functionName}][%d] getFunctionOK  %+v", 200, o.Payload)
}

func (o *GetFunctionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Function)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFunctionBadRequest creates a GetFunctionBadRequest with default headers values
func NewGetFunctionBadRequest() *GetFunctionBadRequest {
	return &GetFunctionBadRequest{}
}

/*GetFunctionBadRequest handles this case with default header values.

Invalid Name supplied
*/
type GetFunctionBadRequest struct {
	Payload *models.Error
}

func (o *GetFunctionBadRequest) Error() string {
	return fmt.Sprintf("[GET /function/{functionName}][%d] getFunctionBadRequest  %+v", 400, o.Payload)
}

func (o *GetFunctionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFunctionNotFound creates a GetFunctionNotFound with default headers values
func NewGetFunctionNotFound() *GetFunctionNotFound {
	return &GetFunctionNotFound{}
}

/*GetFunctionNotFound handles this case with default header values.

Function not found
*/
type GetFunctionNotFound struct {
	Payload *models.Error
}

func (o *GetFunctionNotFound) Error() string {
	return fmt.Sprintf("[GET /function/{functionName}][%d] getFunctionNotFound  %+v", 404, o.Payload)
}

func (o *GetFunctionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFunctionInternalServerError creates a GetFunctionInternalServerError with default headers values
func NewGetFunctionInternalServerError() *GetFunctionInternalServerError {
	return &GetFunctionInternalServerError{}
}

/*GetFunctionInternalServerError handles this case with default header values.

Internal error
*/
type GetFunctionInternalServerError struct {
	Payload *models.Error
}

func (o *GetFunctionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /function/{functionName}][%d] getFunctionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFunctionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
