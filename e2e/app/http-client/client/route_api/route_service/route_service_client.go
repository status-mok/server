// Code generated by go-swagger; DO NOT EDIT.

package route_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new route service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for route service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	RouteServiceCreate(params *RouteServiceCreateParams, opts ...ClientOption) (*RouteServiceCreateOK, error)

	RouteServiceDelete(params *RouteServiceDeleteParams, opts ...ClientOption) (*RouteServiceDeleteOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
RouteServiceCreate creates
*/
func (a *Client) RouteServiceCreate(params *RouteServiceCreateParams, opts ...ClientOption) (*RouteServiceCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRouteServiceCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RouteService_Create",
		Method:             "POST",
		PathPattern:        "/route/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RouteServiceCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RouteServiceCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RouteServiceCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RouteServiceDelete deletes
*/
func (a *Client) RouteServiceDelete(params *RouteServiceDeleteParams, opts ...ClientOption) (*RouteServiceDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRouteServiceDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RouteService_Delete",
		Method:             "POST",
		PathPattern:        "/route/delete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RouteServiceDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RouteServiceDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RouteServiceDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
