// Code generated by go-swagger; DO NOT EDIT.

package basket

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
)

// NewCancelBasketParams creates a new CancelBasketParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCancelBasketParams() *CancelBasketParams {
	return &CancelBasketParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCancelBasketParamsWithTimeout creates a new CancelBasketParams object
// with the ability to set a timeout on a request.
func NewCancelBasketParamsWithTimeout(timeout time.Duration) *CancelBasketParams {
	return &CancelBasketParams{
		timeout: timeout,
	}
}

// NewCancelBasketParamsWithContext creates a new CancelBasketParams object
// with the ability to set a context for a request.
func NewCancelBasketParamsWithContext(ctx context.Context) *CancelBasketParams {
	return &CancelBasketParams{
		Context: ctx,
	}
}

// NewCancelBasketParamsWithHTTPClient creates a new CancelBasketParams object
// with the ability to set a custom HTTPClient for a request.
func NewCancelBasketParamsWithHTTPClient(client *http.Client) *CancelBasketParams {
	return &CancelBasketParams{
		HTTPClient: client,
	}
}

/* CancelBasketParams contains all the parameters to send to the API endpoint
   for the cancel basket operation.

   Typically these are written to a http.Request.
*/
type CancelBasketParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cancel basket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelBasketParams) WithDefaults() *CancelBasketParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cancel basket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelBasketParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cancel basket params
func (o *CancelBasketParams) WithTimeout(timeout time.Duration) *CancelBasketParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel basket params
func (o *CancelBasketParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel basket params
func (o *CancelBasketParams) WithContext(ctx context.Context) *CancelBasketParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel basket params
func (o *CancelBasketParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel basket params
func (o *CancelBasketParams) WithHTTPClient(client *http.Client) *CancelBasketParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel basket params
func (o *CancelBasketParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the cancel basket params
func (o *CancelBasketParams) WithID(id string) *CancelBasketParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the cancel basket params
func (o *CancelBasketParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CancelBasketParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
