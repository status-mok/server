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

// ServerServiceStartReader is a Reader for the ServerServiceStart structure.
type ServerServiceStartReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServerServiceStartReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServerServiceStartOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServerServiceStartDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServerServiceStartOK creates a ServerServiceStartOK with default headers values
func NewServerServiceStartOK() *ServerServiceStartOK {
	return &ServerServiceStartOK{}
}

/*
ServerServiceStartOK describes a response with status code 200, with default header values.

A successful response.
*/
type ServerServiceStartOK struct {
	Payload *models.ServerServiceStartResponse
}

// IsSuccess returns true when this server service start o k response has a 2xx status code
func (o *ServerServiceStartOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this server service start o k response has a 3xx status code
func (o *ServerServiceStartOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this server service start o k response has a 4xx status code
func (o *ServerServiceStartOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this server service start o k response has a 5xx status code
func (o *ServerServiceStartOK) IsServerError() bool {
	return false
}

// IsCode returns true when this server service start o k response a status code equal to that given
func (o *ServerServiceStartOK) IsCode(code int) bool {
	return code == 200
}

func (o *ServerServiceStartOK) Error() string {
	return fmt.Sprintf("[POST /server/start][%d] serverServiceStartOK  %+v", 200, o.Payload)
}

func (o *ServerServiceStartOK) String() string {
	return fmt.Sprintf("[POST /server/start][%d] serverServiceStartOK  %+v", 200, o.Payload)
}

func (o *ServerServiceStartOK) GetPayload() *models.ServerServiceStartResponse {
	return o.Payload
}

func (o *ServerServiceStartOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerServiceStartResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServerServiceStartDefault creates a ServerServiceStartDefault with default headers values
func NewServerServiceStartDefault(code int) *ServerServiceStartDefault {
	return &ServerServiceStartDefault{
		_statusCode: code,
	}
}

/*
ServerServiceStartDefault describes a response with status code -1, with default header values.

ServerServiceStartDefault server service start default
*/
type ServerServiceStartDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the server service start default response
func (o *ServerServiceStartDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this server service start default response has a 2xx status code
func (o *ServerServiceStartDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this server service start default response has a 3xx status code
func (o *ServerServiceStartDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this server service start default response has a 4xx status code
func (o *ServerServiceStartDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this server service start default response has a 5xx status code
func (o *ServerServiceStartDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this server service start default response a status code equal to that given
func (o *ServerServiceStartDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ServerServiceStartDefault) Error() string {
	return fmt.Sprintf("[POST /server/start][%d] ServerService_Start default  %+v", o._statusCode, o.Payload)
}

func (o *ServerServiceStartDefault) String() string {
	return fmt.Sprintf("[POST /server/start][%d] ServerService_Start default  %+v", o._statusCode, o.Payload)
}

func (o *ServerServiceStartDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *ServerServiceStartDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}