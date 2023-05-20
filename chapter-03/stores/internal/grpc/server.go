package grpc

import (
	"context"
	"muazwzxv/Micheal-stack-book/stores/internal/application"
	"muazwzxv/Micheal-stack-book/stores/storespb"

	"google.golang.org/grpc"
)

type server struct {
  app application.App
  storespb.UnimplementedStoresServiceServer
}

func RegisterServer(_ context.Context, app application.App, registrar grpc.ServiceRegistrar) error {
  storespb.RegisterStoresServiceServer(registrar, server{app: app})
  return nil
}

// TODO: Implement gRCP methods
