package repository

import (
	"context"
	"database/sql"
	"muazwzxv/Micheal-stack-book/stores/internal/domain"
)

type StoreRepository struct {
	tableName string
	db        *sql.DB
}

var _ domain.IStoreRepository = (*StoreRepository)(nil)

func NewStoreRepository(tableName string, db *sql.DB) StoreRepository {
	return StoreRepository{
		tableName: tableName,
		db:        db,
	}
}

/**
  Implements IStoreRepository
*/

// TODO: implements all these methods
func (r StoreRepository) Find(ctx context.Context, storeID string) (*domain.Store, error) {
	return nil, nil
}

func (r StoreRepository) FindAll(ctx context.Context) ([]*domain.Store, error) {
	return nil, nil
}

func (r StoreRepository) Save(ctx context.Context, store *domain.Store) error {
	return nil
}

func (r StoreRepository) Update(ctx context.Context, store *domain.Store) error {
	return nil
}

func (r StoreRepository) Delete(ctx context.Context, storeID string) error {
	return nil
}

func (r StoreRepository) FindParticipatingStore(ctx context.Context) ([]*domain.Store, error) {
	return nil, nil
}

// func (r StoreRepository) table(query string) string {
//   return fmt.Sprintf(query, r.tableName)
// }
