package commands

import (
	"context"
	"muazwzxv/Micheal-stack-book/stores/internal/domain"

	"github.com/stackus/errors"
)

type AddProduct struct {
	ID          string
	StoreID     string
	Name        string
	Description string
	SKU         string
	Price       float64
}

type AddProductHandler struct {
	stores   domain.IStoreRepository
	products domain.IProductRepository
}

func NewAddProductHandler(stores domain.IStoreRepository, products domain.IProductRepository) AddProductHandler {
	return AddProductHandler{
		stores:   stores,
		products: products,
	}
}

func (h AddProductHandler) AddProduct(ctx context.Context, cmd AddProduct) error {
  _, err := h.stores.Find(ctx, cmd.StoreID)
  if err != nil {
    return errors.Wrap(err, "error adding product")
  }

  product, err := domain.CreateProduct(cmd.ID, cmd.StoreID, cmd.Name, cmd.Description, cmd.SKU, cmd.Price)
  if err != nil {
    return errors.Wrap(err, "error adding product")
  }

  return errors.Wrap(h.products.AddProduct(ctx, product), "error adding product")
}
