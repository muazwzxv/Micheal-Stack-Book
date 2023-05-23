package domain

import "context"

type IParticipatingStoreRepository interface {
  FindAll(ctx context.Context) ([]*Store, error)
}
