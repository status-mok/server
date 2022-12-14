// Code generated by go-swagger; DO NOT EDIT.

package expectation_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/status-mok/server/e2e/app/http-client/models"
)

// NewExpectationServiceCreateParams creates a new ExpectationServiceCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExpectationServiceCreateParams() *ExpectationServiceCreateParams {
	return &ExpectationServiceCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExpectationServiceCreateParamsWithTimeout creates a new ExpectationServiceCreateParams object
// with the ability to set a timeout on a request.
func NewExpectationServiceCreateParamsWithTimeout(timeout time.Duration) *ExpectationServiceCreateParams {
	return &ExpectationServiceCreateParams{
		timeout: timeout,
	}
}

// NewExpectationServiceCreateParamsWithContext creates a new ExpectationServiceCreateParams object
// with the ability to set a context for a request.
func NewExpectationServiceCreateParamsWithContext(ctx context.Context) *ExpectationServiceCreateParams {
	return &ExpectationServiceCreateParams{
		Context: ctx,
	}
}

// NewExpectationServiceCreateParamsWithHTTPClient creates a new ExpectationServiceCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExpectationServiceCreateParamsWithHTTPClient(client *http.Client) *ExpectationServiceCreateParams {
	return &ExpectationServiceCreateParams{
		HTTPClient: client,
	}
}

/*
ExpectationServiceCreateParams contains all the parameters to send to the API endpoint

	for the expectation service create operation.

	Typically these are written to a http.Request.
*/
type ExpectationServiceCreateParams struct {

	// Body.
	Body *models.ExpectationServiceCreateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the expectation service create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExpectationServiceCreateParams) WithDefaults() *ExpectationServiceCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the expectation service create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExpectationServiceCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the expectation service create params
func (o *ExpectationServiceCreateParams) WithTimeout(timeout time.Duration) *ExpectationServiceCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the expectation service create params
func (o *ExpectationServiceCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the expectation service create params
func (o *ExpectationServiceCreateParams) WithContext(ctx context.Context) *ExpectationServiceCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the expectation service create params
func (o *ExpectationServiceCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the expectation service create params
func (o *ExpectationServiceCreateParams) WithHTTPClient(client *http.Client) *ExpectationServiceCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the expectation service create params
func (o *ExpectationServiceCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the expectation service create params
func (o *ExpectationServiceCreateParams) WithBody(body *models.ExpectationServiceCreateRequest) *ExpectationServiceCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the expectation service create params
func (o *ExpectationServiceCreateParams) SetBody(body *models.ExpectationServiceCreateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ExpectationServiceCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
