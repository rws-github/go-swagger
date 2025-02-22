// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/rws-github/go-swagger/examples/contributed-templates/stratoscale/models"
)

// PetCreateReader is a Reader for the PetCreate structure.
type PetCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PetCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPetCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 405:
		result := NewPetCreateMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPetCreateCreated creates a PetCreateCreated with default headers values
func NewPetCreateCreated() *PetCreateCreated {
	return &PetCreateCreated{}
}

/*PetCreateCreated handles this case with default header values.

Created
*/
type PetCreateCreated struct {
	Payload *models.Pet
}

func (o *PetCreateCreated) Error() string {
	return fmt.Sprintf("[POST /pet][%d] petCreateCreated  %+v", 201, o.Payload)
}

func (o *PetCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Pet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPetCreateMethodNotAllowed creates a PetCreateMethodNotAllowed with default headers values
func NewPetCreateMethodNotAllowed() *PetCreateMethodNotAllowed {
	return &PetCreateMethodNotAllowed{}
}

/*PetCreateMethodNotAllowed handles this case with default header values.

Invalid input
*/
type PetCreateMethodNotAllowed struct {
}

func (o *PetCreateMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /pet][%d] petCreateMethodNotAllowed ", 405)
}

func (o *PetCreateMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
