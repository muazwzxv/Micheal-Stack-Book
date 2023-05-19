package commands

import (
	"context"
	"muazwzxv/Micheal-stack-book/stores/internal/domain"
)

type RemoveProduct struct {
	ID string
}

type RemoveProductHandler struct {
	stores   domain.IStoreRepository
	products domain.IProductRepository
}

func NewRemoveProductHandler(stores domain.IStoreRepository, products domain.IProductRepository) RemoveProductHandler {
  return RemoveProductHandler{
    stores: stores,
    products: products,
  }
}

func (h RemoveProductHandler) RemoveProduct(ctx context.Context, cmd RemoveProduct) error {
  return h.products.RemoveProduct(ctx, cmd.ID)
}
