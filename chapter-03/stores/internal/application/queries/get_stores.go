package queries

import (
	"context"
	"muazwzxv/Micheal-stack-book/stores/internal/domain"
)

type GetStores struct{}

type GetStoresHandler struct {
	stores domain.IStoreRepository
}

func NewGetStoresHandler(stores domain.IStoreRepository) GetStoresHandler {
	return GetStoresHandler{stores: stores}
}

func (h GetStoresHandler) GetStores(ctx context.Context, _ GetStores) ([]*domain.Store, error) {
	return h.stores.FindAll(ctx)
}
