package repository

import (
	"context"
	"database/sql"
	"fmt"
	"muazwzxv/Micheal-stack-book/stores/internal/domain"

	"github.com/stackus/errors"
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
func (r StoreRepository) Find(ctx context.Context, storeID string) (*domain.Store, error) {
	const query = `
    SELECT
      name, location, participating
    FROM %s
    WHERE 
      id = $1
    LIMIT 1`

	store := &domain.Store{
		ID: storeID,
	}

	err := r.db.QueryRowContext(ctx, r.table(query), store.ID).
		Scan(&store.Name, &store.Location, &store.Participating)
	if err != nil {
		return nil, errors.Wrap(err, "scanning stores")
	}

	return store, nil
}

// TODO: offset and limit
func (r StoreRepository) FindAll(ctx context.Context) ([]*domain.Store, error) {
	const query = `
    SELECT 
      id, name, location, participating
    FROM %s
    `

	rows, err := r.db.QueryContext(ctx, r.table(query))
	if err != nil {
		return nil, errors.Wrap(err, "querying stores")
	}

	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	stores := make([]*domain.Store, 0)

	for rows.Next() {
		store := &domain.Store{}
		err := rows.Scan(&store.ID, &store.Name, &store.Location, &store.Participating)
		if err != nil {
			return nil, errors.Wrap(err, "scanning store")
		}
		stores = append(stores, store)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "finishing store rows")
	}

	return stores, nil
}

func (r StoreRepository) Save(ctx context.Context, store *domain.Store) error {
	const query = `
    INSERT INTO %s
      (id, name, location, participating)
    VALUES
      ($1, $2, $3, $4)
    `
	_, err := r.db.ExecContext(ctx, r.table(query),
		store.ID,
		store.Name,
		store.Location,
		store.Participating,
	)
	if err != nil {
		return errors.Wrap(err, "inserting store")
	}

	return nil
}

func (r StoreRepository) Update(ctx context.Context, store *domain.Store) error {
	const query = `
    UPDATE %s 
    SET 
      name = $2, location = $3, participating = $4 
    WHERE 
      id = $1`

	_, err := r.db.ExecContext(ctx, r.table(query), store.ID, store.Name, store.Location, store.Participating)
	if err != nil {
		return errors.Wrap(err, "updating store")
	}

	return nil
}

func (r StoreRepository) Delete(ctx context.Context, storeID string) error {
	const query = "DELETE FROM %s WHERE id = $1 LIMIT 1"

	_, err := r.db.ExecContext(ctx, r.table(query), storeID)
	if err != nil {
		return errors.Wrap(err, "deleting store")
	}

	return nil
}

func (r StoreRepository) FindParticipatingStore(ctx context.Context) ([]*domain.Store, error) {
	const query = `
    SELECT 
      id, name, location, participating 
    FROM %s 
    WHERE 
      participating is true`

	rows, err := r.db.QueryContext(ctx, r.table(query))
	if err != nil {
		return nil, errors.Wrap(err, "querying participating stores")
	}

	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

  stores := make([]*domain.Store, 0)

	for rows.Next() {
		store := &domain.Store{}
		err := rows.Scan(&store.ID, &store.Name, &store.Location, &store.Participating)
		if err != nil {
			return nil, errors.Wrap(err, "scanning participating store")
		}
		stores = append(stores, store)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "finishing participating store rows")
	}

	return stores, nil
}

func (r StoreRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
