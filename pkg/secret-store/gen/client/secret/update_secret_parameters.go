///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package secret

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

	models "github.com/vmware/dispatch/pkg/secret-store/gen/models"
)

// NewUpdateSecretParams creates a new UpdateSecretParams object
// with the default values initialized.
func NewUpdateSecretParams() *UpdateSecretParams {
	var ()
	return &UpdateSecretParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSecretParamsWithTimeout creates a new UpdateSecretParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSecretParamsWithTimeout(timeout time.Duration) *UpdateSecretParams {
	var ()
	return &UpdateSecretParams{

		timeout: timeout,
	}
}

// NewUpdateSecretParamsWithContext creates a new UpdateSecretParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSecretParamsWithContext(ctx context.Context) *UpdateSecretParams {
	var ()
	return &UpdateSecretParams{

		Context: ctx,
	}
}

// NewUpdateSecretParamsWithHTTPClient creates a new UpdateSecretParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSecretParamsWithHTTPClient(client *http.Client) *UpdateSecretParams {
	var ()
	return &UpdateSecretParams{
		HTTPClient: client,
	}
}

/*UpdateSecretParams contains all the parameters to send to the API endpoint
for the update secret operation typically these are written to a http.Request
*/
type UpdateSecretParams struct {

	/*Secret*/
	Secret *models.Secret
	/*SecretName*/
	SecretName string
	/*Tags
	  Filter based on tags

	*/
	Tags []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update secret params
func (o *UpdateSecretParams) WithTimeout(timeout time.Duration) *UpdateSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update secret params
func (o *UpdateSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update secret params
func (o *UpdateSecretParams) WithContext(ctx context.Context) *UpdateSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update secret params
func (o *UpdateSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update secret params
func (o *UpdateSecretParams) WithHTTPClient(client *http.Client) *UpdateSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update secret params
func (o *UpdateSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSecret adds the secret to the update secret params
func (o *UpdateSecretParams) WithSecret(secret *models.Secret) *UpdateSecretParams {
	o.SetSecret(secret)
	return o
}

// SetSecret adds the secret to the update secret params
func (o *UpdateSecretParams) SetSecret(secret *models.Secret) {
	o.Secret = secret
}

// WithSecretName adds the secretName to the update secret params
func (o *UpdateSecretParams) WithSecretName(secretName string) *UpdateSecretParams {
	o.SetSecretName(secretName)
	return o
}

// SetSecretName adds the secretName to the update secret params
func (o *UpdateSecretParams) SetSecretName(secretName string) {
	o.SecretName = secretName
}

// WithTags adds the tags to the update secret params
func (o *UpdateSecretParams) WithTags(tags []string) *UpdateSecretParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the update secret params
func (o *UpdateSecretParams) SetTags(tags []string) {
	o.Tags = tags
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Secret != nil {
		if err := r.SetBodyParam(o.Secret); err != nil {
			return err
		}
	}

	// path param secretName
	if err := r.SetPathParam("secretName", o.SecretName); err != nil {
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
