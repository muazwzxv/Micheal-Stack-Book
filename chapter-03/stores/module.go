package stores

import (
	"context"
	"muazwzxv/Micheal-stack-book/internal/monolith"
	"muazwzxv/Micheal-stack-book/stores/internal/application"
	"muazwzxv/Micheal-stack-book/stores/internal/grpc"
	"muazwzxv/Micheal-stack-book/stores/internal/logging"
	"muazwzxv/Micheal-stack-book/stores/internal/repository"
	"muazwzxv/Micheal-stack-book/stores/internal/rest"
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

  if err := grpc.RegisterServer(ctx, app, mono.Rpc()); err != nil{
    return err
  }

  if err := rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
    return err
  }

  if err := rest.RegisterSwagger(mono.Mux()); err != nil {
    return err
  }

	return nil
}
