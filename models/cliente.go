// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Cliente Dados do cliente
//
// swagger:model Cliente
type Cliente struct {

	// CPF do Cliente
	// Min Length: 1
	CpfClientes string `json:"cpfClientes,omitempty"`

	// Maior compra única do cliente
	MaiorCompraUnicaDoCliente float64 `json:"maiorCompraUnicaDoCliente,omitempty"`

	// Nome do Cliente
	// Min Length: 1
	NomeCliente string `json:"nomeCliente,omitempty"`

	// Quantidade de compras realizadas (fidelidade do cliente)
	QuantidadeDeComprasDoCliente int64 `json:"quantidadeDeComprasDoCliente,omitempty"`

	// Total em compras realizadas pelo cliente
	TotalEmComprasDoCliente float64 `json:"totalEmComprasDoCliente,omitempty"`
}

// Validate validates this cliente
func (m *Cliente) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCpfClientes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNomeCliente(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Cliente) validateCpfClientes(formats strfmt.Registry) error {

	if swag.IsZero(m.CpfClientes) { // not required
		return nil
	}

	if err := validate.MinLength("cpfClientes", "body", string(m.CpfClientes), 1); err != nil {
		return err
	}

	return nil
}

func (m *Cliente) validateNomeCliente(formats strfmt.Registry) error {

	if swag.IsZero(m.NomeCliente) { // not required
		return nil
	}

	if err := validate.MinLength("nomeCliente", "body", string(m.NomeCliente), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Cliente) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cliente) UnmarshalBinary(b []byte) error {
	var res Cliente
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}