package rest

import (
	"context"
	"muazwzxv/Micheal-stack-book/stores/storespb"

	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RegisterGateway(ctx context.Context, mux *chi.Mux, grpcAddr string) error {
	const api = "/api/stores"

	gateway := runtime.NewServeMux()
  err := storespb.RegisterStoresServiceHandlerFromEndpoint(ctx, gateway, grpcAddr, []grpc.DialOption{
    grpc.WithTransportCredentials(insecure.NewCredentials()),
  })
  if err != nil {
    return err
  }

  // mount the GRPC Gateway
  mux.Mount(api, gateway)

  return nil
}
