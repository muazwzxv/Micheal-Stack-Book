package repository

import (
	"context"
	"database/sql"
	"muazwzxv/Micheal-stack-book/stores/internal/domain"
)

type ProductRepository struct {
	tableName string
	db        *sql.DB
}

var _ domain.IProductRepository = (*ProductRepository)(nil)

func NewProductRepository(tableName string, db *sql.DB) ProductRepository {
	return ProductRepository{
		tableName: tableName,
		db:        db,
	}
}

/**
  Implements IProductRepository
*/

// TODO: implements all these methods
func (r ProductRepository) FindProduct(ctx context.Context, id string) (*domain.Product, error) {
	return nil, nil
}

func (r ProductRepository) AddProduct(ctx context.Context, product *domain.Product) error {
	return nil
}

func (r ProductRepository) RemoveProduct(ctx context.Context, id string) error {
	return nil
}

func (r ProductRepository) GetCatalog(ctx context.Context, storeID string) ([]*domain.Product, error) {
	return nil, nil
}

// func (r ProductRepository) table(query string) string {
//   return fmt.Sprintf(query, r.tableName)
// }
