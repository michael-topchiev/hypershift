// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudNetworksGetReader is a Reader for the PcloudNetworksGet structure.
type PcloudNetworksGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudNetworksGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudNetworksGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudNetworksGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudNetworksGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudNetworksGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudNetworksGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudNetworksGetOK creates a PcloudNetworksGetOK with default headers values
func NewPcloudNetworksGetOK() *PcloudNetworksGetOK {
	return &PcloudNetworksGetOK{}
}

/* PcloudNetworksGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudNetworksGetOK struct {
	Payload *models.Network
}

func (o *PcloudNetworksGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksGetOK  %+v", 200, o.Payload)
}
func (o *PcloudNetworksGetOK) GetPayload() *models.Network {
	return o.Payload
}

func (o *PcloudNetworksGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Network)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksGetBadRequest creates a PcloudNetworksGetBadRequest with default headers values
func NewPcloudNetworksGetBadRequest() *PcloudNetworksGetBadRequest {
	return &PcloudNetworksGetBadRequest{}
}

/* PcloudNetworksGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudNetworksGetBadRequest struct {
	Payload *models.Error
}

func (o *PcloudNetworksGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksGetBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudNetworksGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksGetUnauthorized creates a PcloudNetworksGetUnauthorized with default headers values
func NewPcloudNetworksGetUnauthorized() *PcloudNetworksGetUnauthorized {
	return &PcloudNetworksGetUnauthorized{}
}

/* PcloudNetworksGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudNetworksGetUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudNetworksGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksGetUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudNetworksGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksGetNotFound creates a PcloudNetworksGetNotFound with default headers values
func NewPcloudNetworksGetNotFound() *PcloudNetworksGetNotFound {
	return &PcloudNetworksGetNotFound{}
}

/* PcloudNetworksGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudNetworksGetNotFound struct {
	Payload *models.Error
}

func (o *PcloudNetworksGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksGetNotFound  %+v", 404, o.Payload)
}
func (o *PcloudNetworksGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksGetInternalServerError creates a PcloudNetworksGetInternalServerError with default headers values
func NewPcloudNetworksGetInternalServerError() *PcloudNetworksGetInternalServerError {
	return &PcloudNetworksGetInternalServerError{}
}

/* PcloudNetworksGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudNetworksGetInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudNetworksGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksGetInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudNetworksGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
