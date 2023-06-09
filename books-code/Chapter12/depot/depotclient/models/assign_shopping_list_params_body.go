// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AssignShoppingListParamsBody assign shopping list params body
//
// swagger:model assignShoppingListParamsBody
type AssignShoppingListParamsBody struct {

	// bot Id
	BotID string `json:"botId,omitempty"`
}

// Validate validates this assign shopping list params body
func (m *AssignShoppingListParamsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this assign shopping list params body based on context it is used
func (m *AssignShoppingListParamsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AssignShoppingListParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssignShoppingListParamsBody) UnmarshalBinary(b []byte) error {
	var res AssignShoppingListParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
