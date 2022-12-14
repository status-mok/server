// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ServerServiceServerType server service server type
//
// swagger:model ServerServiceServerType
type ServerServiceServerType string

func NewServerServiceServerType(value ServerServiceServerType) *ServerServiceServerType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ServerServiceServerType.
func (m ServerServiceServerType) Pointer() *ServerServiceServerType {
	return &m
}

const (

	// ServerServiceServerTypeSERVERTYPEUNSPECIFIED captures enum value "SERVER_TYPE_UNSPECIFIED"
	ServerServiceServerTypeSERVERTYPEUNSPECIFIED ServerServiceServerType = "SERVER_TYPE_UNSPECIFIED"

	// ServerServiceServerTypeSERVERTYPEHTTP captures enum value "SERVER_TYPE_HTTP"
	ServerServiceServerTypeSERVERTYPEHTTP ServerServiceServerType = "SERVER_TYPE_HTTP"

	// ServerServiceServerTypeSERVERTYPEGRPC captures enum value "SERVER_TYPE_GRPC"
	ServerServiceServerTypeSERVERTYPEGRPC ServerServiceServerType = "SERVER_TYPE_GRPC"

	// ServerServiceServerTypeSERVERTYPETHRIFT captures enum value "SERVER_TYPE_THRIFT"
	ServerServiceServerTypeSERVERTYPETHRIFT ServerServiceServerType = "SERVER_TYPE_THRIFT"

	// ServerServiceServerTypeSERVERTYPETCP captures enum value "SERVER_TYPE_TCP"
	ServerServiceServerTypeSERVERTYPETCP ServerServiceServerType = "SERVER_TYPE_TCP"

	// ServerServiceServerTypeSERVERTYPEUDP captures enum value "SERVER_TYPE_UDP"
	ServerServiceServerTypeSERVERTYPEUDP ServerServiceServerType = "SERVER_TYPE_UDP"
)

// for schema
var serverServiceServerTypeEnum []interface{}

func init() {
	var res []ServerServiceServerType
	if err := json.Unmarshal([]byte(`["SERVER_TYPE_UNSPECIFIED","SERVER_TYPE_HTTP","SERVER_TYPE_GRPC","SERVER_TYPE_THRIFT","SERVER_TYPE_TCP","SERVER_TYPE_UDP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverServiceServerTypeEnum = append(serverServiceServerTypeEnum, v)
	}
}

func (m ServerServiceServerType) validateServerServiceServerTypeEnum(path, location string, value ServerServiceServerType) error {
	if err := validate.EnumCase(path, location, value, serverServiceServerTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this server service server type
func (m ServerServiceServerType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateServerServiceServerTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this server service server type based on context it is used
func (m ServerServiceServerType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
