package commands

import (
	"context"
	"muazwzxv/Micheal-stack-book/stores/internal/domain"
)

type DisableParticipation struct {
  ID string
}

type DisableParticipationHandler struct {
  stores domain.IStoreRepository
}

func NewDisableParticipationHandler(stores domain.IStoreRepository) DisableParticipationHandler {
  return DisableParticipationHandler{stores: stores}
}

func (h DisableParticipationHandler) DisableParticipation(ctx context.Context, cmd DisableParticipation) error {
  store, err := h.stores.Find(ctx, cmd.ID)
  if err != nil {
    return err
  }

  err = store.DisableParticipation()
  if err != nil {
    return err
  }

  return h.stores.Update(ctx, store)
}
