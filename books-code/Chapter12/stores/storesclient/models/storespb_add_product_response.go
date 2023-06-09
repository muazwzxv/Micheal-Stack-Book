// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorespbAddProductResponse storespb add product response
//
// swagger:model storespbAddProductResponse
type StorespbAddProductResponse struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this storespb add product response
func (m *StorespbAddProductResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this storespb add product response based on context it is used
func (m *StorespbAddProductResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorespbAddProductResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorespbAddProductResponse) UnmarshalBinary(b []byte) error {
	var res StorespbAddProductResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
