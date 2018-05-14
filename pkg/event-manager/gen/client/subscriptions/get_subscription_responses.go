///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// GetSubscriptionReader is a Reader for the GetSubscription structure.
type GetSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetSubscriptionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetSubscriptionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetSubscriptionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetSubscriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSubscriptionOK creates a GetSubscriptionOK with default headers values
func NewGetSubscriptionOK() *GetSubscriptionOK {
	return &GetSubscriptionOK{}
}

/*GetSubscriptionOK handles this case with default header values.

Successful operation
*/
type GetSubscriptionOK struct {
	Payload *v1.Subscription
}

func (o *GetSubscriptionOK) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{subscriptionName}][%d] getSubscriptionOK  %+v", 200, o.Payload)
}

func (o *GetSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Subscription)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionBadRequest creates a GetSubscriptionBadRequest with default headers values
func NewGetSubscriptionBadRequest() *GetSubscriptionBadRequest {
	return &GetSubscriptionBadRequest{}
}

/*GetSubscriptionBadRequest handles this case with default header values.

Invalid Name supplied
*/
type GetSubscriptionBadRequest struct {
	Payload *v1.Error
}

func (o *GetSubscriptionBadRequest) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{subscriptionName}][%d] getSubscriptionBadRequest  %+v", 400, o.Payload)
}

func (o *GetSubscriptionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionNotFound creates a GetSubscriptionNotFound with default headers values
func NewGetSubscriptionNotFound() *GetSubscriptionNotFound {
	return &GetSubscriptionNotFound{}
}

/*GetSubscriptionNotFound handles this case with default header values.

Subscription not found
*/
type GetSubscriptionNotFound struct {
	Payload *v1.Error
}

func (o *GetSubscriptionNotFound) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{subscriptionName}][%d] getSubscriptionNotFound  %+v", 404, o.Payload)
}

func (o *GetSubscriptionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionInternalServerError creates a GetSubscriptionInternalServerError with default headers values
func NewGetSubscriptionInternalServerError() *GetSubscriptionInternalServerError {
	return &GetSubscriptionInternalServerError{}
}

/*GetSubscriptionInternalServerError handles this case with default header values.

Internal server error
*/
type GetSubscriptionInternalServerError struct {
	Payload *v1.Error
}

func (o *GetSubscriptionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{subscriptionName}][%d] getSubscriptionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSubscriptionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionDefault creates a GetSubscriptionDefault with default headers values
func NewGetSubscriptionDefault(code int) *GetSubscriptionDefault {
	return &GetSubscriptionDefault{
		_statusCode: code,
	}
}

/*GetSubscriptionDefault handles this case with default header values.

Unknown error
*/
type GetSubscriptionDefault struct {
	_statusCode int

	Payload *v1.Error
}

// Code gets the status code for the get subscription default response
func (o *GetSubscriptionDefault) Code() int {
	return o._statusCode
}

func (o *GetSubscriptionDefault) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{subscriptionName}][%d] getSubscription default  %+v", o._statusCode, o.Payload)
}

func (o *GetSubscriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
