// Code generated by go-swagger; DO NOT EDIT.

package apimodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Resource resource
// swagger:model resource
type Resource struct {

	// Resource content
	// Required: true
	ResourceContent *string `json:"resourceContent"`

	// Resource URI
	// Required: true
	ResourceURI *string `json:"resourceURI"`
}

// Validate validates this resource
func (m *Resource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Resource) validateResourceContent(formats strfmt.Registry) error {

	if err := validate.Required("resourceContent", "body", m.ResourceContent); err != nil {
		return err
	}

	return nil
}

func (m *Resource) validateResourceURI(formats strfmt.Registry) error {

	if err := validate.Required("resourceURI", "body", m.ResourceURI); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Resource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Resource) UnmarshalBinary(b []byte) error {
	var res Resource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
