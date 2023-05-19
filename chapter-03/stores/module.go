package stores

import (
	"context"
	"muazwzxv/Micheal-stack-book/internal/monolith"
	"muazwzxv/Micheal-stack-book/stores/internal/repository"
)

type Module struct {}

func (m *Module) Startup(ctx context.Context, mono monolith.Monolith) error {
  // setup driver adapters
  stores := repository.NewStoreRepository("stores.stores", mono.DB())

}
