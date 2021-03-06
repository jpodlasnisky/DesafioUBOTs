// Code generated by go-swagger; DO NOT EDIT.

package cliente

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"std/github.com/jpodlasnisky/DesafioUBOTs/models"
)

// ClientesMaisFieisOKCode is the HTTP code returned for type ClientesMaisFieisOK
const ClientesMaisFieisOKCode int = 200

/*ClientesMaisFieisOK Cliente com maior compra única

swagger:response clientesMaisFieisOK
*/
type ClientesMaisFieisOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Cliente `json:"body,omitempty"`
}

// NewClientesMaisFieisOK creates ClientesMaisFieisOK with default headers values
func NewClientesMaisFieisOK() *ClientesMaisFieisOK {

	return &ClientesMaisFieisOK{}
}

// WithPayload adds the payload to the clientes mais fieis o k response
func (o *ClientesMaisFieisOK) WithPayload(payload []*models.Cliente) *ClientesMaisFieisOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the clientes mais fieis o k response
func (o *ClientesMaisFieisOK) SetPayload(payload []*models.Cliente) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ClientesMaisFieisOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Cliente, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ClientesMaisFieisInternalServerErrorCode is the HTTP code returned for type ClientesMaisFieisInternalServerError
const ClientesMaisFieisInternalServerErrorCode int = 500

/*ClientesMaisFieisInternalServerError Status 500

swagger:response clientesMaisFieisInternalServerError
*/
type ClientesMaisFieisInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Erro `json:"body,omitempty"`
}

// NewClientesMaisFieisInternalServerError creates ClientesMaisFieisInternalServerError with default headers values
func NewClientesMaisFieisInternalServerError() *ClientesMaisFieisInternalServerError {

	return &ClientesMaisFieisInternalServerError{}
}

// WithPayload adds the payload to the clientes mais fieis internal server error response
func (o *ClientesMaisFieisInternalServerError) WithPayload(payload *models.Erro) *ClientesMaisFieisInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the clientes mais fieis internal server error response
func (o *ClientesMaisFieisInternalServerError) SetPayload(payload *models.Erro) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ClientesMaisFieisInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
