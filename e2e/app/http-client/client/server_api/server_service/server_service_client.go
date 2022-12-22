// Code generated by go-swagger; DO NOT EDIT.

package server_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new server service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for server service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ServerServiceCreate(params *ServerServiceCreateParams, opts ...ClientOption) (*ServerServiceCreateOK, error)

	ServerServiceDelete(params *ServerServiceDeleteParams, opts ...ClientOption) (*ServerServiceDeleteOK, error)

	ServerServiceStart(params *ServerServiceStartParams, opts ...ClientOption) (*ServerServiceStartOK, error)

	ServerServiceStop(params *ServerServiceStopParams, opts ...ClientOption) (*ServerServiceStopOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ServerServiceCreate creates
*/
func (a *Client) ServerServiceCreate(params *ServerServiceCreateParams, opts ...ClientOption) (*ServerServiceCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServerServiceCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ServerService_Create",
		Method:             "POST",
		PathPattern:        "/server/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ServerServiceCreateReader{formats: a.formats},
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
	success, ok := result.(*ServerServiceCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ServerServiceCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ServerServiceDelete deletes
*/
func (a *Client) ServerServiceDelete(params *ServerServiceDeleteParams, opts ...ClientOption) (*ServerServiceDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServerServiceDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ServerService_Delete",
		Method:             "POST",
		PathPattern:        "/server/delete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ServerServiceDeleteReader{formats: a.formats},
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
	success, ok := result.(*ServerServiceDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ServerServiceDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ServerServiceStart starts
*/
func (a *Client) ServerServiceStart(params *ServerServiceStartParams, opts ...ClientOption) (*ServerServiceStartOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServerServiceStartParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ServerService_Start",
		Method:             "POST",
		PathPattern:        "/server/start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ServerServiceStartReader{formats: a.formats},
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
	success, ok := result.(*ServerServiceStartOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ServerServiceStartDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ServerServiceStop stops
*/
func (a *Client) ServerServiceStop(params *ServerServiceStopParams, opts ...ClientOption) (*ServerServiceStopOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServerServiceStopParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ServerService_Stop",
		Method:             "POST",
		PathPattern:        "/server/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ServerServiceStopReader{formats: a.formats},
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
	success, ok := result.(*ServerServiceStopOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ServerServiceStopDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
