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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon-sdks/go/http_client/v1/service_model"
)

// LoginReader is a Reader for the Login structure.
type LoginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLoginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewLoginForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewLoginNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLoginOK creates a LoginOK with default headers values
func NewLoginOK() *LoginOK {
	return &LoginOK{}
}

/*LoginOK handles this case with default header values.

A successful response.
*/
type LoginOK struct {
	Payload *service_model.V1Auth
}

func (o *LoginOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/users/token][%d] loginOK  %+v", 200, o.Payload)
}

func (o *LoginOK) GetPayload() *service_model.V1Auth {
	return o.Payload
}

func (o *LoginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Auth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoginForbidden creates a LoginForbidden with default headers values
func NewLoginForbidden() *LoginForbidden {
	return &LoginForbidden{}
}

/*LoginForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type LoginForbidden struct {
	Payload interface{}
}

func (o *LoginForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/users/token][%d] loginForbidden  %+v", 403, o.Payload)
}

func (o *LoginForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *LoginForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoginNotFound creates a LoginNotFound with default headers values
func NewLoginNotFound() *LoginNotFound {
	return &LoginNotFound{}
}

/*LoginNotFound handles this case with default header values.

Resource does not exist.
*/
type LoginNotFound struct {
	Payload string
}

func (o *LoginNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/users/token][%d] loginNotFound  %+v", 404, o.Payload)
}

func (o *LoginNotFound) GetPayload() string {
	return o.Payload
}

func (o *LoginNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
