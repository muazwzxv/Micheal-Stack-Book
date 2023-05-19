package application

import (
	"context"
	"muazwzxv/Micheal-stack-book/stores/internal/application/commands"
	"muazwzxv/Micheal-stack-book/stores/internal/domain"
)

type (
	App interface {
		ICommands
		IQueries
	}
	ICommands interface {
		CreateStore(ctx context.Context, cmd commands.CreateStore) error
		EnableParticipation(ctx context.Context, cmd commands.EnableParticipation) error
		DisableParticipation(ctx context.Context, cmd commands.DisableParticipation) error
		AddProduct(ctx context.Context, cmd commands.AddProduct) error
		RemoveProduct(ctx context.Context, cmd commands.RemoveProduct) error
	}
	IQueries interface{}

	Application struct {
		appCommands
		appQueries
	}
	appCommands struct {
		commands.CreateStoreHandler
		commands.EnableParticipationHandler
		commands.DisableParticipationHandler
		commands.AddProductHandler
		commands.RemoveProductHandler
	}
	appQueries struct{}
)

func New(stores domain.IStoreRepository, participatingStores domain.IParticipatingStoreRepository, products domain.IProductRepository) *Application {
	return &Application{
		appCommands: appCommands{
			CreateStoreHandler: commands.NewCreateStoreHandler(stores),
      EnableParticipationHandler: commands.NewEnableParticipationHandler(stores),
      DisableParticipationHandler: commands.NewDisableParticipationHandler(stores),
      AddProductHandler: commands.NewAddProductHandler(stores, products),
      RemoveProductHandler: commands.NewRemoveProductHandler(stores, products),
		},
		appQueries: appQueries{},
	}
}
