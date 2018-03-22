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

	models "github.com/vmware/dispatch/pkg/secret-store/gen/models"
)

// GetSecretsReader is a Reader for the GetSecrets structure.
type GetSecretsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSecretsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSecretsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetSecretsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetSecretsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSecretsOK creates a GetSecretsOK with default headers values
func NewGetSecretsOK() *GetSecretsOK {
	return &GetSecretsOK{}
}

/*GetSecretsOK handles this case with default header values.

An array of registered secrets
*/
type GetSecretsOK struct {
	Payload []*models.Secret
}

func (o *GetSecretsOK) Error() string {
	return fmt.Sprintf("[GET /][%d] getSecretsOK  %+v", 200, o.Payload)
}

func (o *GetSecretsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSecretsBadRequest creates a GetSecretsBadRequest with default headers values
func NewGetSecretsBadRequest() *GetSecretsBadRequest {
	return &GetSecretsBadRequest{}
}

/*GetSecretsBadRequest handles this case with default header values.

Bad Request
*/
type GetSecretsBadRequest struct {
	Payload *models.Error
}

func (o *GetSecretsBadRequest) Error() string {
	return fmt.Sprintf("[GET /][%d] getSecretsBadRequest  %+v", 400, o.Payload)
}

func (o *GetSecretsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSecretsDefault creates a GetSecretsDefault with default headers values
func NewGetSecretsDefault(code int) *GetSecretsDefault {
	return &GetSecretsDefault{
		_statusCode: code,
	}
}

/*GetSecretsDefault handles this case with default header values.

Standard error
*/
type GetSecretsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get secrets default response
func (o *GetSecretsDefault) Code() int {
	return o._statusCode
}

func (o *GetSecretsDefault) Error() string {
	return fmt.Sprintf("[GET /][%d] getSecrets default  %+v", o._statusCode, o.Payload)
}

func (o *GetSecretsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
