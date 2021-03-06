///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package serviceaccount

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// DeleteServiceAccountReader is a Reader for the DeleteServiceAccount structure.
type DeleteServiceAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteServiceAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteServiceAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteServiceAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteServiceAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteServiceAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteServiceAccountOK creates a DeleteServiceAccountOK with default headers values
func NewDeleteServiceAccountOK() *DeleteServiceAccountOK {
	return &DeleteServiceAccountOK{}
}

/*DeleteServiceAccountOK handles this case with default header values.

Successful operation
*/
type DeleteServiceAccountOK struct {
	Payload *v1.ServiceAccount
}

func (o *DeleteServiceAccountOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/iam/serviceaccount/{serviceAccountName}][%d] deleteServiceAccountOK  %+v", 200, o.Payload)
}

func (o *DeleteServiceAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.ServiceAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteServiceAccountBadRequest creates a DeleteServiceAccountBadRequest with default headers values
func NewDeleteServiceAccountBadRequest() *DeleteServiceAccountBadRequest {
	return &DeleteServiceAccountBadRequest{}
}

/*DeleteServiceAccountBadRequest handles this case with default header values.

Invalid Name supplied
*/
type DeleteServiceAccountBadRequest struct {
	Payload *v1.Error
}

func (o *DeleteServiceAccountBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v1/iam/serviceaccount/{serviceAccountName}][%d] deleteServiceAccountBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteServiceAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteServiceAccountNotFound creates a DeleteServiceAccountNotFound with default headers values
func NewDeleteServiceAccountNotFound() *DeleteServiceAccountNotFound {
	return &DeleteServiceAccountNotFound{}
}

/*DeleteServiceAccountNotFound handles this case with default header values.

Service Account not found
*/
type DeleteServiceAccountNotFound struct {
	Payload *v1.Error
}

func (o *DeleteServiceAccountNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/iam/serviceaccount/{serviceAccountName}][%d] deleteServiceAccountNotFound  %+v", 404, o.Payload)
}

func (o *DeleteServiceAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteServiceAccountInternalServerError creates a DeleteServiceAccountInternalServerError with default headers values
func NewDeleteServiceAccountInternalServerError() *DeleteServiceAccountInternalServerError {
	return &DeleteServiceAccountInternalServerError{}
}

/*DeleteServiceAccountInternalServerError handles this case with default header values.

Internal error
*/
type DeleteServiceAccountInternalServerError struct {
	Payload *v1.Error
}

func (o *DeleteServiceAccountInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/iam/serviceaccount/{serviceAccountName}][%d] deleteServiceAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteServiceAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
