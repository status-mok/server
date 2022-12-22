// Code generated by go-swagger; DO NOT EDIT.

package expectation_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/status-mok/server/e2e/app/http-client/models"
)

// ExpectationServiceCreateReader is a Reader for the ExpectationServiceCreate structure.
type ExpectationServiceCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExpectationServiceCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExpectationServiceCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExpectationServiceCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExpectationServiceCreateOK creates a ExpectationServiceCreateOK with default headers values
func NewExpectationServiceCreateOK() *ExpectationServiceCreateOK {
	return &ExpectationServiceCreateOK{}
}

/*
ExpectationServiceCreateOK describes a response with status code 200, with default header values.

A successful response.
*/
type ExpectationServiceCreateOK struct {
	Payload *models.ExpectationServiceCreateResponse
}

// IsSuccess returns true when this expectation service create o k response has a 2xx status code
func (o *ExpectationServiceCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this expectation service create o k response has a 3xx status code
func (o *ExpectationServiceCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this expectation service create o k response has a 4xx status code
func (o *ExpectationServiceCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this expectation service create o k response has a 5xx status code
func (o *ExpectationServiceCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this expectation service create o k response a status code equal to that given
func (o *ExpectationServiceCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExpectationServiceCreateOK) Error() string {
	return fmt.Sprintf("[POST /expectation/create][%d] expectationServiceCreateOK  %+v", 200, o.Payload)
}

func (o *ExpectationServiceCreateOK) String() string {
	return fmt.Sprintf("[POST /expectation/create][%d] expectationServiceCreateOK  %+v", 200, o.Payload)
}

func (o *ExpectationServiceCreateOK) GetPayload() *models.ExpectationServiceCreateResponse {
	return o.Payload
}

func (o *ExpectationServiceCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExpectationServiceCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExpectationServiceCreateDefault creates a ExpectationServiceCreateDefault with default headers values
func NewExpectationServiceCreateDefault(code int) *ExpectationServiceCreateDefault {
	return &ExpectationServiceCreateDefault{
		_statusCode: code,
	}
}

/*
ExpectationServiceCreateDefault describes a response with status code -1, with default header values.

ExpectationServiceCreateDefault expectation service create default
*/
type ExpectationServiceCreateDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the expectation service create default response
func (o *ExpectationServiceCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this expectation service create default response has a 2xx status code
func (o *ExpectationServiceCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this expectation service create default response has a 3xx status code
func (o *ExpectationServiceCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this expectation service create default response has a 4xx status code
func (o *ExpectationServiceCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this expectation service create default response has a 5xx status code
func (o *ExpectationServiceCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this expectation service create default response a status code equal to that given
func (o *ExpectationServiceCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExpectationServiceCreateDefault) Error() string {
	return fmt.Sprintf("[POST /expectation/create][%d] ExpectationService_Create default  %+v", o._statusCode, o.Payload)
}

func (o *ExpectationServiceCreateDefault) String() string {
	return fmt.Sprintf("[POST /expectation/create][%d] ExpectationService_Create default  %+v", o._statusCode, o.Payload)
}

func (o *ExpectationServiceCreateDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *ExpectationServiceCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
