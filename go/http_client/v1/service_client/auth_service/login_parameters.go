// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package auth_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon-sdks/go/http_client/v1/service_model"
)

// NewLoginParams creates a new LoginParams object
// with the default values initialized.
func NewLoginParams() *LoginParams {
	var ()
	return &LoginParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLoginParamsWithTimeout creates a new LoginParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLoginParamsWithTimeout(timeout time.Duration) *LoginParams {
	var ()
	return &LoginParams{

		timeout: timeout,
	}
}

// NewLoginParamsWithContext creates a new LoginParams object
// with the default values initialized, and the ability to set a context for a request
func NewLoginParamsWithContext(ctx context.Context) *LoginParams {
	var ()
	return &LoginParams{

		Context: ctx,
	}
}

// NewLoginParamsWithHTTPClient creates a new LoginParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLoginParamsWithHTTPClient(client *http.Client) *LoginParams {
	var ()
	return &LoginParams{
		HTTPClient: client,
	}
}

/*LoginParams contains all the parameters to send to the API endpoint
for the login operation typically these are written to a http.Request
*/
type LoginParams struct {

	/*Body*/
	Body *service_model.V1CredsBodyRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the login params
func (o *LoginParams) WithTimeout(timeout time.Duration) *LoginParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the login params
func (o *LoginParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the login params
func (o *LoginParams) WithContext(ctx context.Context) *LoginParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the login params
func (o *LoginParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the login params
func (o *LoginParams) WithHTTPClient(client *http.Client) *LoginParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the login params
func (o *LoginParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the login params
func (o *LoginParams) WithBody(body *service_model.V1CredsBodyRequest) *LoginParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the login params
func (o *LoginParams) SetBody(body *service_model.V1CredsBodyRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *LoginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
