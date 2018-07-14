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

// User user
// swagger:model user
type User struct {

	// first name
	// Required: true
	// Min Length: 8
	FirstName *string `json:"first_name"`

	// id
	// Read Only: true
	ID string `json:"id,omitempty"`

	// last name
	// Required: true
	// Min Length: 8
	LastName *string `json:"last_name"`

	// username
	// Required: true
	// Min Length: 6
	Username *string `json:"username"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFirstName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("first_name", "body", m.FirstName); err != nil {
		return err
	}

	if err := validate.MinLength("first_name", "body", string(*m.FirstName), 8); err != nil {
		return err
	}

	return nil
}

func (m *User) validateLastName(formats strfmt.Registry) error {

	if err := validate.Required("last_name", "body", m.LastName); err != nil {
		return err
	}

	if err := validate.MinLength("last_name", "body", string(*m.LastName), 8); err != nil {
		return err
	}

	return nil
}

func (m *User) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	if err := validate.MinLength("username", "body", string(*m.Username), 6); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
