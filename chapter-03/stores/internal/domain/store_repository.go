package domain

import "golang.org/x/net/context"

type IStoreRepository interface {
  Save(ctx context.Context, store *Store) error
  Update(ctx context.Context, store *Store) error
  Delete(ctx context.Context, storeID string) error
  Find(ctx context.Context, storeID string) (*Store, error)
  FindAll(ctx context.Context) ([]*Store, error)
}
