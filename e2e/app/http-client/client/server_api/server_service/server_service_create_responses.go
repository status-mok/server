// Code generated by go-swagger; DO NOT EDIT.

package server_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/status-mok/server/e2e/app/http-client/models"
)

// ServerServiceCreateReader is a Reader for the ServerServiceCreate structure.
type ServerServiceCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServerServiceCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServerServiceCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServerServiceCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServerServiceCreateOK creates a ServerServiceCreateOK with default headers values
func NewServerServiceCreateOK() *ServerServiceCreateOK {
	return &ServerServiceCreateOK{}
}

/*
ServerServiceCreateOK describes a response with status code 200, with default header values.

A successful response.
*/
type ServerServiceCreateOK struct {
	Payload *models.ServerServiceCreateResponse
}

// IsSuccess returns true when this server service create o k response has a 2xx status code
func (o *ServerServiceCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this server service create o k response has a 3xx status code
func (o *ServerServiceCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this server service create o k response has a 4xx status code
func (o *ServerServiceCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this server service create o k response has a 5xx status code
func (o *ServerServiceCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this server service create o k response a status code equal to that given
func (o *ServerServiceCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ServerServiceCreateOK) Error() string {
	return fmt.Sprintf("[POST /server/create][%d] serverServiceCreateOK  %+v", 200, o.Payload)
}

func (o *ServerServiceCreateOK) String() string {
	return fmt.Sprintf("[POST /server/create][%d] serverServiceCreateOK  %+v", 200, o.Payload)
}

func (o *ServerServiceCreateOK) GetPayload() *models.ServerServiceCreateResponse {
	return o.Payload
}

func (o *ServerServiceCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerServiceCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServerServiceCreateDefault creates a ServerServiceCreateDefault with default headers values
func NewServerServiceCreateDefault(code int) *ServerServiceCreateDefault {
	return &ServerServiceCreateDefault{
		_statusCode: code,
	}
}

/*
ServerServiceCreateDefault describes a response with status code -1, with default header values.

ServerServiceCreateDefault server service create default
*/
type ServerServiceCreateDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the server service create default response
func (o *ServerServiceCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this server service create default response has a 2xx status code
func (o *ServerServiceCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this server service create default response has a 3xx status code
func (o *ServerServiceCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this server service create default response has a 4xx status code
func (o *ServerServiceCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this server service create default response has a 5xx status code
func (o *ServerServiceCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this server service create default response a status code equal to that given
func (o *ServerServiceCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ServerServiceCreateDefault) Error() string {
	return fmt.Sprintf("[POST /server/create][%d] ServerService_Create default  %+v", o._statusCode, o.Payload)
}

func (o *ServerServiceCreateDefault) String() string {
	return fmt.Sprintf("[POST /server/create][%d] ServerService_Create default  %+v", o._statusCode, o.Payload)
}

func (o *ServerServiceCreateDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *ServerServiceCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
