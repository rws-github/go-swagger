// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/rws-github/go-swagger/examples/task-tracker/models"
)

// GetTaskCommentsReader is a Reader for the GetTaskComments structure.
type GetTaskCommentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTaskCommentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTaskCommentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetTaskCommentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTaskCommentsOK creates a GetTaskCommentsOK with default headers values
func NewGetTaskCommentsOK() *GetTaskCommentsOK {
	return &GetTaskCommentsOK{}
}

/*GetTaskCommentsOK handles this case with default header values.

The list of comments
*/
type GetTaskCommentsOK struct {
	Payload []*models.Comment
}

func (o *GetTaskCommentsOK) Error() string {
	return fmt.Sprintf("[GET /tasks/{id}/comments][%d] getTaskCommentsOK  %+v", 200, o.Payload)
}

func (o *GetTaskCommentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTaskCommentsDefault creates a GetTaskCommentsDefault with default headers values
func NewGetTaskCommentsDefault(code int) *GetTaskCommentsDefault {
	return &GetTaskCommentsDefault{
		_statusCode: code,
	}
}

/*GetTaskCommentsDefault handles this case with default header values.

Error response
*/
type GetTaskCommentsDefault struct {
	_statusCode int

	XErrorCode string

	Payload *models.Error
}

// Code gets the status code for the get task comments default response
func (o *GetTaskCommentsDefault) Code() int {
	return o._statusCode
}

func (o *GetTaskCommentsDefault) Error() string {
	return fmt.Sprintf("[GET /tasks/{id}/comments][%d] getTaskComments default  %+v", o._statusCode, o.Payload)
}

func (o *GetTaskCommentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Error-Code
	o.XErrorCode = response.GetHeader("X-Error-Code")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
