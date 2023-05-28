package queries

import (
	"context"
	"muazwzxv/Micheal-stack-book/stores/internal/domain"
)

type GetProduct struct {
  ID string
}

type GetProductHandler struct {
  products domain.IProductRepository
}

func NewGetProductHandler(products domain.IProductRepository) GetProductHandler {
  return GetProductHandler{
    products: products,
  }
}

func (h GetProductHandler) GetProduct(ctx context.Context, query GetProduct) (*domain.Product, error) {
  return h.products.FindProduct(ctx, query.ID)
}
