package main

import (
	"context"
	"log"
	"os"

	"github.com/SaranHiruthikM/ecomm-go/internal/env"
	"github.com/jackc/pgx/v5"
	"github.com/kaptinlin/messageformat-go/pkg/logger"
)

func main() {
	ctx := context.Background()

	cfg := config{addr: ":8080", db: dbconfig{
		dsn: env.GetString("GOOSE_DBSTRING", "host=localhost user=postgres password=postgres dbname=ecom sslmode=disable"),
	}}

	// Database
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)
	}

	defer conn.Close(ctx)

	logger.Info("connected to database", "dsn", cfg.db.dsn)

	api := application{
		config: cfg,
	}

	if err := api.run(api.mount()); err != nil {
		log.Printf("Server has failed to start, err: %s", err)
		os.Exit(1)
	}
}
