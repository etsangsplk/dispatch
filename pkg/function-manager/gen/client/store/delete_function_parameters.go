///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteFunctionParams creates a new DeleteFunctionParams object
// with the default values initialized.
func NewDeleteFunctionParams() *DeleteFunctionParams {
	var ()
	return &DeleteFunctionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteFunctionParamsWithTimeout creates a new DeleteFunctionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteFunctionParamsWithTimeout(timeout time.Duration) *DeleteFunctionParams {
	var ()
	return &DeleteFunctionParams{

		timeout: timeout,
	}
}

// NewDeleteFunctionParamsWithContext creates a new DeleteFunctionParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteFunctionParamsWithContext(ctx context.Context) *DeleteFunctionParams {
	var ()
	return &DeleteFunctionParams{

		Context: ctx,
	}
}

// NewDeleteFunctionParamsWithHTTPClient creates a new DeleteFunctionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteFunctionParamsWithHTTPClient(client *http.Client) *DeleteFunctionParams {
	var ()
	return &DeleteFunctionParams{
		HTTPClient: client,
	}
}

/*DeleteFunctionParams contains all the parameters to send to the API endpoint
for the delete function operation typically these are written to a http.Request
*/
type DeleteFunctionParams struct {

	/*XDispatchOrg*/
	XDispatchOrg string
	/*FunctionName
	  Name of function to work on

	*/
	FunctionName string
	/*Tags
	  Filter based on tags

	*/
	Tags []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete function params
func (o *DeleteFunctionParams) WithTimeout(timeout time.Duration) *DeleteFunctionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete function params
func (o *DeleteFunctionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete function params
func (o *DeleteFunctionParams) WithContext(ctx context.Context) *DeleteFunctionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete function params
func (o *DeleteFunctionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete function params
func (o *DeleteFunctionParams) WithHTTPClient(client *http.Client) *DeleteFunctionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete function params
func (o *DeleteFunctionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDispatchOrg adds the xDispatchOrg to the delete function params
func (o *DeleteFunctionParams) WithXDispatchOrg(xDispatchOrg string) *DeleteFunctionParams {
	o.SetXDispatchOrg(xDispatchOrg)
	return o
}

// SetXDispatchOrg adds the xDispatchOrg to the delete function params
func (o *DeleteFunctionParams) SetXDispatchOrg(xDispatchOrg string) {
	o.XDispatchOrg = xDispatchOrg
}

// WithFunctionName adds the functionName to the delete function params
func (o *DeleteFunctionParams) WithFunctionName(functionName string) *DeleteFunctionParams {
	o.SetFunctionName(functionName)
	return o
}

// SetFunctionName adds the functionName to the delete function params
func (o *DeleteFunctionParams) SetFunctionName(functionName string) {
	o.FunctionName = functionName
}

// WithTags adds the tags to the delete function params
func (o *DeleteFunctionParams) WithTags(tags []string) *DeleteFunctionParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the delete function params
func (o *DeleteFunctionParams) SetTags(tags []string) {
	o.Tags = tags
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFunctionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Dispatch-Org
	if err := r.SetHeaderParam("X-Dispatch-Org", o.XDispatchOrg); err != nil {
		return err
	}

	// path param functionName
	if err := r.SetPathParam("functionName", o.FunctionName); err != nil {
		return err
	}

	valuesTags := o.Tags

	joinedTags := swag.JoinByFormat(valuesTags, "multi")
	// query array param tags
	if err := r.SetQueryParam("tags", joinedTags...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
