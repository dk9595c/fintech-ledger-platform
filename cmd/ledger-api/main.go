package main

import (
    `log`
    `net/http`
    `os`

    `fintechledger/internal/platform/postgres`
    `fintechledger/internal/platform/redis`
    `fintechledger/internal/transaction`
)

func main() {
    dbConn := os.Getenv(`DATABASE_URL`)
    redisAddr := os.Getenv(`REDIS_ADDR`)

    db, err := postgres.NewPool(dbConn)
    if err != nil {
        log.Fatalf(`Database connection failed: %v`, err)
    }
    defer db.Close()

    cache := redis.NewClient(redisAddr)

    txHandler := transaction.NewHandler(db, cache)

    mux := http.NewServeMux()
    mux.HandleFunc(`/api/v1/transfer`, txHandler.HandleTransfer)

    log.Println(`Enterprise Ledger API starting on port 8080`)
    if err := http.ListenAndServe(`:8080`, mux); err != nil {
        log.Fatalf(`Server error: %v`, err)
    }
}
