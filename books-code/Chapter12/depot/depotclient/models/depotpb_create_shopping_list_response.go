// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DepotpbCreateShoppingListResponse depotpb create shopping list response
//
// swagger:model depotpbCreateShoppingListResponse
type DepotpbCreateShoppingListResponse struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this depotpb create shopping list response
func (m *DepotpbCreateShoppingListResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this depotpb create shopping list response based on context it is used
func (m *DepotpbCreateShoppingListResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DepotpbCreateShoppingListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DepotpbCreateShoppingListResponse) UnmarshalBinary(b []byte) error {
	var res DepotpbCreateShoppingListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
