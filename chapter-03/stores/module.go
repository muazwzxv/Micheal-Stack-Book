package stores

import (
	"context"
	"muazwzxv/Micheal-stack-book/internal/monolith"
	"muazwzxv/Micheal-stack-book/stores/internal/application"
	"muazwzxv/Micheal-stack-book/stores/internal/logging"
	"muazwzxv/Micheal-stack-book/stores/internal/repository"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	stores := repository.NewStoreRepository("stores.stores", mono.DB())
	participatingStores := repository.NewParticipatingStoreRepository("stores.stores", mono.DB())
	products := repository.NewProductRepository("stores.products", mono.DB())

	// setup application
	var app application.App
	app = application.New(stores, participatingStores, products)
	app = logging.LogApplicationAccess(app, mono.Logger())


	// TODO: Setup drivers adapters

	return nil
}
