// Copyright 2018-2020 Polyaxon, Inc.
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

package projects_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// PatchProjectTeamsReader is a Reader for the PatchProjectTeams structure.
type PatchProjectTeamsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchProjectTeamsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchProjectTeamsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewPatchProjectTeamsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPatchProjectTeamsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchProjectTeamsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchProjectTeamsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchProjectTeamsOK creates a PatchProjectTeamsOK with default headers values
func NewPatchProjectTeamsOK() *PatchProjectTeamsOK {
	return &PatchProjectTeamsOK{}
}

/*PatchProjectTeamsOK handles this case with default header values.

A successful response.
*/
type PatchProjectTeamsOK struct {
	Payload *service_model.V1ProjectTeams
}

func (o *PatchProjectTeamsOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/{project}/teams][%d] patchProjectTeamsOK  %+v", 200, o.Payload)
}

func (o *PatchProjectTeamsOK) GetPayload() *service_model.V1ProjectTeams {
	return o.Payload
}

func (o *PatchProjectTeamsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ProjectTeams)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchProjectTeamsNoContent creates a PatchProjectTeamsNoContent with default headers values
func NewPatchProjectTeamsNoContent() *PatchProjectTeamsNoContent {
	return &PatchProjectTeamsNoContent{}
}

/*PatchProjectTeamsNoContent handles this case with default header values.

No content.
*/
type PatchProjectTeamsNoContent struct {
	Payload interface{}
}

func (o *PatchProjectTeamsNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/{project}/teams][%d] patchProjectTeamsNoContent  %+v", 204, o.Payload)
}

func (o *PatchProjectTeamsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchProjectTeamsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchProjectTeamsForbidden creates a PatchProjectTeamsForbidden with default headers values
func NewPatchProjectTeamsForbidden() *PatchProjectTeamsForbidden {
	return &PatchProjectTeamsForbidden{}
}

/*PatchProjectTeamsForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type PatchProjectTeamsForbidden struct {
	Payload interface{}
}

func (o *PatchProjectTeamsForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/{project}/teams][%d] patchProjectTeamsForbidden  %+v", 403, o.Payload)
}

func (o *PatchProjectTeamsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchProjectTeamsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchProjectTeamsNotFound creates a PatchProjectTeamsNotFound with default headers values
func NewPatchProjectTeamsNotFound() *PatchProjectTeamsNotFound {
	return &PatchProjectTeamsNotFound{}
}

/*PatchProjectTeamsNotFound handles this case with default header values.

Resource does not exist.
*/
type PatchProjectTeamsNotFound struct {
	Payload interface{}
}

func (o *PatchProjectTeamsNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/{project}/teams][%d] patchProjectTeamsNotFound  %+v", 404, o.Payload)
}

func (o *PatchProjectTeamsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchProjectTeamsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchProjectTeamsDefault creates a PatchProjectTeamsDefault with default headers values
func NewPatchProjectTeamsDefault(code int) *PatchProjectTeamsDefault {
	return &PatchProjectTeamsDefault{
		_statusCode: code,
	}
}

/*PatchProjectTeamsDefault handles this case with default header values.

An unexpected error response
*/
type PatchProjectTeamsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the patch project teams default response
func (o *PatchProjectTeamsDefault) Code() int {
	return o._statusCode
}

func (o *PatchProjectTeamsDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/{project}/teams][%d] PatchProjectTeams default  %+v", o._statusCode, o.Payload)
}

func (o *PatchProjectTeamsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *PatchProjectTeamsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}