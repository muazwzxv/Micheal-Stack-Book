package config

import (
	"muazwzxv/Micheal-stack-book/internal/rpc"
	"muazwzxv/Micheal-stack-book/internal/web"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/stackus/dotenv"
)

 
type (
  PGConfig struct {
    Conn string `required:"true"`
  }

  AppConfig struct {
    Environment string
    Loglevel string  `envconfig:"LOG_LEVEL" default:"DEBUG"`
    PG PGConfig
    Rpc rpc.RpcConfig
    Web web.WebConfig
    ShutdownTimeout time.Duration   `envconfig:"SHUTDOWN_TIMEOUT" default:"30s"`
  }
)

func InitConfig () (cfg AppConfig, err error) {
  if err = dotenv.Load(dotenv.EnvironmentFiles(os.Getenv("ENVIRONMENT"))); err != nil {
    return
  }

  err = envconfig.Process("", &cfg)
  return 
}
