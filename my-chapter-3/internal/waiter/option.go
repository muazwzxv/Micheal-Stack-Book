package waiter

import "context"

type WaiterOption func(c *waiterConfig)

func ParentContext(ctx context.Context) WaiterOption {
  return func(c *waiterConfig) {
    c.parentCtx = ctx
  }
}

func CatchSignals() WaiterOption {
  return func(c *waiterConfig) {
    c.catchSignals = true
  }
}
