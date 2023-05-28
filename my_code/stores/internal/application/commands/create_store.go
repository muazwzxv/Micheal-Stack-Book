package commands

import (
	"context"
	"muazwzxv/Micheal-stack-book/stores/internal/domain"
)

type (
	CreateStore struct {
		ID       string
		Name     string
		Location string
	}

  CreateStoreHandler struct {
    stores domain.IStoreRepository
  }
)

func NewCreateStoreHandler(stores domain.IStoreRepository) CreateStoreHandler {
  return CreateStoreHandler{
    stores: stores,
  }
}

func (h CreateStoreHandler) CreateStore(ctx context.Context, cmd CreateStore) error {
  store, err := domain.CreateStore(cmd.ID, cmd.Name, cmd.Location)
  if err != nil {
    return err
  }

  err = h.stores.Save(ctx, store)
  
  return err
}

