package stores

import (
	"context"
	"muazwzxv/Micheal-stack-book/internal/monolith"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	// setup driver adapters
	return nil
}
