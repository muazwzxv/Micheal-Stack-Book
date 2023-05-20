package queries

import (
	"context"
	"muazwzxv/Micheal-stack-book/stores/internal/domain"
)

type GetParticipatingStores struct {}

type GetParticipatingStoresHandler struct {
  participatingStores domain.IParticipatingStoreRepository
}

func NewParticipatingStoresHandler(participatingStores domain.IParticipatingStoreRepository) GetParticipatingStoresHandler {
  return GetParticipatingStoresHandler{
    participatingStores: participatingStores,
  }
}

func (h GetParticipatingStoresHandler) GetParticipatingStores(ctx context.Context, _ GetParticipatingStores) ([]*domain.Store, error) {
  return h.participatingStores.FindAll(ctx)
}
