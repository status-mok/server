// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RouteServiceCreateRequest route service create request
//
// swagger:model RouteServiceCreateRequest
type RouteServiceCreateRequest struct {

	// Server name is a unique identifier of the mock server.
	// Example: http-server-8080
	// Required: true
	// Min Length: 1
	ServerName *string `json:"server_name"`

	// 1: ReqResp
	// 2: WebSocket
	// Required: true
	Type *RouteServiceRouteType `json:"type"`

	// URL is a relative URI of the route.
	//
	// The route type.
	// Example: /some-url
	// Required: true
	// Min Length: 1
	URL *string `json:"url"`
}

// Validate validates this route service create request
func (m *RouteServiceCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServerName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RouteServiceCreateRequest) validateServerName(formats strfmt.Registry) error {

	if err := validate.Required("server_name", "body", m.ServerName); err != nil {
		return err
	}

	if err := validate.MinLength("server_name", "body", *m.ServerName, 1); err != nil {
		return err
	}

	return nil
}

func (m *RouteServiceCreateRequest) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *RouteServiceCreateRequest) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	if err := validate.MinLength("url", "body", *m.URL, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this route service create request based on the context it is used
func (m *RouteServiceCreateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RouteServiceCreateRequest) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RouteServiceCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RouteServiceCreateRequest) UnmarshalBinary(b []byte) error {
	var res RouteServiceCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
