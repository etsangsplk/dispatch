///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetOrganizationsParams creates a new GetOrganizationsParams object
// with the default values initialized.
func NewGetOrganizationsParams() *GetOrganizationsParams {

	return &GetOrganizationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationsParamsWithTimeout creates a new GetOrganizationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOrganizationsParamsWithTimeout(timeout time.Duration) *GetOrganizationsParams {

	return &GetOrganizationsParams{

		timeout: timeout,
	}
}

// NewGetOrganizationsParamsWithContext creates a new GetOrganizationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOrganizationsParamsWithContext(ctx context.Context) *GetOrganizationsParams {

	return &GetOrganizationsParams{

		Context: ctx,
	}
}

// NewGetOrganizationsParamsWithHTTPClient creates a new GetOrganizationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOrganizationsParamsWithHTTPClient(client *http.Client) *GetOrganizationsParams {

	return &GetOrganizationsParams{
		HTTPClient: client,
	}
}

/*GetOrganizationsParams contains all the parameters to send to the API endpoint
for the get organizations operation typically these are written to a http.Request
*/
type GetOrganizationsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get organizations params
func (o *GetOrganizationsParams) WithTimeout(timeout time.Duration) *GetOrganizationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organizations params
func (o *GetOrganizationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organizations params
func (o *GetOrganizationsParams) WithContext(ctx context.Context) *GetOrganizationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organizations params
func (o *GetOrganizationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organizations params
func (o *GetOrganizationsParams) WithHTTPClient(client *http.Client) *GetOrganizationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organizations params
func (o *GetOrganizationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
