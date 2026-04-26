package main

import (
    `context`
    `log`
    `os`
    `os/signal`
    `syscall`

    `fintechledger/internal/platform/redis`
    `fintechledger/internal/checker`
)

func main() {
    redisAddr := os.Getenv(`REDIS_ADDR`)
    
    cache := redis.NewClient(redisAddr)
    fraudChecker := checker.NewFraudService(cache)

    ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
    defer stop()

    log.Println(`Fraud Detection Worker started`)

    err := fraudChecker.ListenForTransactions(ctx)
    if err != nil && err != context.Canceled {
        log.Fatalf(`Worker failure: %v`, err)
    }

    log.Println(`Worker shut down gracefully`)
}
