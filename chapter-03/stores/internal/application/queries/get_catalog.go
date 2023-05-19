package queries

import (
	"context"
	"muazwzxv/Micheal-stack-book/stores/internal/domain"
)

type GetCatalog struct {
  StoreID string
}

type GetCatalogHandler struct {
  products domain.IProductRepository
}

func NewCatalogHandler(products domain.IProductRepository) GetCatalogHandler {
  return GetCatalogHandler{
    products: products,
  }
}

func (h GetCatalogHandler) GetCatalog(ctx context.Context, query GetCatalog) ([]*domain.Product, error) {
  return h.products.GetCatalog(ctx, query.StoreID)
}
