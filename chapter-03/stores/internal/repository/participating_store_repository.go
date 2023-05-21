package repository

import (
	"context"
	"database/sql"
	"fmt"
	"muazwzxv/Micheal-stack-book/stores/internal/domain"

	"github.com/stackus/errors"
)

type ParticipatingStoreRepository struct {
	tableName string
	db        *sql.DB
}

var _ domain.IParticipatingStoreRepository = (*ParticipatingStoreRepository)(nil)

func NewParticipatingStoreRepository(tableName string, db *sql.DB) ParticipatingStoreRepository {
	return ParticipatingStoreRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r ParticipatingStoreRepository) FindAll(ctx context.Context) ([]*domain.Store, error) {
	const query = `
    SELECT 
      id, name, location, participating 
    FROM 
      %s
    WHERE participating is true`

	rows, err := r.db.QueryContext(ctx, r.table(query))
	if err != nil {
		return nil, errors.Wrap(err, "query participating stores")
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	stores := make([]*domain.Store, 0)

	for rows.Next() {
		store := &domain.Store{}
		err := rows.Scan(
			&store.ID,
			&store.Name,
			&store.Location,
			&store.Participating,
		)
		if err != nil {
			return nil, errors.Wrap(err, "scanning participatings tore")
		}

		stores = append(stores, store)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "finishing participating store rows")
	}

	return stores, nil
}

func (r ParticipatingStoreRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
