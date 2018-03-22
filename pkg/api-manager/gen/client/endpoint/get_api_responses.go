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

// GetAPIReader is a Reader for the GetAPI structure.
type GetAPIReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAPIOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetAPIBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAPINotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetAPIInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAPIOK creates a GetAPIOK with default headers values
func NewGetAPIOK() *GetAPIOK {
	return &GetAPIOK{}
}

/*GetAPIOK handles this case with default header values.

Successful operation
*/
type GetAPIOK struct {
	Payload *models.API
}

func (o *GetAPIOK) Error() string {
	return fmt.Sprintf("[GET /{api}][%d] getApiOK  %+v", 200, o.Payload)
}

func (o *GetAPIOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.API)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIBadRequest creates a GetAPIBadRequest with default headers values
func NewGetAPIBadRequest() *GetAPIBadRequest {
	return &GetAPIBadRequest{}
}

/*GetAPIBadRequest handles this case with default header values.

Invalid Name supplied
*/
type GetAPIBadRequest struct {
	Payload *models.Error
}

func (o *GetAPIBadRequest) Error() string {
	return fmt.Sprintf("[GET /{api}][%d] getApiBadRequest  %+v", 400, o.Payload)
}

func (o *GetAPIBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPINotFound creates a GetAPINotFound with default headers values
func NewGetAPINotFound() *GetAPINotFound {
	return &GetAPINotFound{}
}

/*GetAPINotFound handles this case with default header values.

API not found
*/
type GetAPINotFound struct {
	Payload *models.Error
}

func (o *GetAPINotFound) Error() string {
	return fmt.Sprintf("[GET /{api}][%d] getApiNotFound  %+v", 404, o.Payload)
}

func (o *GetAPINotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIInternalServerError creates a GetAPIInternalServerError with default headers values
func NewGetAPIInternalServerError() *GetAPIInternalServerError {
	return &GetAPIInternalServerError{}
}

/*GetAPIInternalServerError handles this case with default header values.

Internal error
*/
type GetAPIInternalServerError struct {
	Payload *models.Error
}

func (o *GetAPIInternalServerError) Error() string {
	return fmt.Sprintf("[GET /{api}][%d] getApiInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAPIInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
