package main

import (
	"database/sql"
	"fmt"
	"muazwzxv/Micheal-stack-book/internal/config"
	"muazwzxv/Micheal-stack-book/internal/logger"
	"muazwzxv/Micheal-stack-book/internal/monolith"
	"muazwzxv/Micheal-stack-book/internal/rpc"
	"muazwzxv/Micheal-stack-book/internal/waiter"
	"muazwzxv/Micheal-stack-book/internal/web"
	"muazwzxv/Micheal-stack-book/stores"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func run() (err error) {
	var cfg config.AppConfig

	// parse config/env/...
	cfg, err = config.InitConfig()
	if err != nil {
		return err
	}

	m := app{cfg: cfg}

	m.db, err = sql.Open("pgx", cfg.PG.Conn)
	if err != nil {
		return err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(m.db)
	m.logger = logger.New(logger.LogConfig{
		Environment: cfg.Environment,
		LogLevel:    logger.Level(cfg.Loglevel),
	})
  m.rpc = initRpc(cfg.Rpc)
  m.mux = initMux(cfg.Web)
  m.waiter = waiter.New(waiter.CatchSignals())

  // init modules
  m.modules = []monolith.Module{
    &stores.Module{},
  }

  if err = m.startupModules(); err != nil {
    return err
  }

  //TODO: Mount general web resources
  return
}


func initRpc(_ rpc.RpcConfig) *grpc.Server {
  server := grpc.NewServer()
  reflection.Register(server)

  return server
}

func initMux(_ web.WebConfig) *chi.Mux {
  return chi.NewMux()
}
