package queries

import (
	"context"
	"muazwzxv/Micheal-stack-book/stores/internal/domain"
)

type GetStore struct {
  ID string
}

type GetStoreHandler struct {
  stores domain.IStoreRepository
}

func NewGetStoreHandler(stores domain.IStoreRepository) GetStoreHandler {
  return GetStoreHandler{
    stores: stores,
  }
}

func (h GetStoreHandler) GetStore(ctx context.Context, query GetStore) (*domain.Store, error) {
  return h.stores.Find(ctx, query.ID)
}

