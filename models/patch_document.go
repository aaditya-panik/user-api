// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PatchDocument patch document
// swagger:model patchDocument
type PatchDocument struct {

	// email
	// Min Length: 6
	// Format: email
	Email strfmt.Email `json:"email,omitempty"`

	// first name
	// Min Length: 3
	FirstName string `json:"first_name,omitempty"`

	// last name
	// Min Length: 3
	LastName string `json:"last_name,omitempty"`
}

// Validate validates this patch document
func (m *PatchDocument) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchDocument) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(m.Email) { // not required
		return nil
	}

	if err := validate.MinLength("email", "body", string(m.Email), 6); err != nil {
		return err
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PatchDocument) validateFirstName(formats strfmt.Registry) error {

	if swag.IsZero(m.FirstName) { // not required
		return nil
	}

	if err := validate.MinLength("first_name", "body", string(m.FirstName), 3); err != nil {
		return err
	}

	return nil
}

func (m *PatchDocument) validateLastName(formats strfmt.Registry) error {

	if swag.IsZero(m.LastName) { // not required
		return nil
	}

	if err := validate.MinLength("last_name", "body", string(m.LastName), 3); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchDocument) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchDocument) UnmarshalBinary(b []byte) error {
	var res PatchDocument
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
