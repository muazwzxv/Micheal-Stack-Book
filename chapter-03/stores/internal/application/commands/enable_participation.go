package commands

import (
	"context"
	"muazwzxv/Micheal-stack-book/stores/internal/domain"
)

type EnableParticipation struct {
  ID string
}

type EnableParticipationHandler struct {
  stores domain.IStoreRepository
} 

func NewEnableParticipationHandler(stores domain.IStoreRepository) EnableParticipationHandler {
  return EnableParticipationHandler{stores: stores}
}

func (h EnableParticipationHandler) EnableParticipation(ctx context.Context, cmd EnableParticipation) error {
  store, err := h.stores.Find(ctx, cmd.ID)
  if err != nil {
    return err
  }

  err = store.EnableParticipation()
  if err != nil {
    return err
  }

  return h.stores.Update(ctx, store)
}
