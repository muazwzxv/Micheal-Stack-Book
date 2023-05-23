package domain

import (
	"github.com/stackus/errors"
)

var (
	ErrProductNameIsBlank = errors.Wrap(errors.ErrBadRequest, "the product name cannot be blank")
	ErrProductIsNegative  = errors.Wrap(errors.ErrBadRequest, "the product price cannot be negative")
)

type Product struct {
	ID          string
	StoreID     string
	Name        string
	Description string
	SKU         string
	Price       float64
}

func CreateProduct(
	id, storeId, name, description, sku string,
	price float64,
) (*Product, error) {
	if name == "" {
		return nil, ErrProductNameIsBlank
	}

	if price < 0 {
		return nil, ErrProductIsNegative
	}

	product := &Product{
		ID:          id,
		StoreID:     storeId,
		Name:        name,
		Description: description,
		SKU:         sku,
		Price:       price,
	}

	return product, nil
}
