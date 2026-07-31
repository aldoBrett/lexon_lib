// Command migrate applies all pending lexon schema migrations.
package main

import (
	"context"
	"log"
	"os"

	"amox/lex_engine_lib/db"
	"amox/lex_engine_lib/db/migrations"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("migrate: DATABASE_URL is required")
	}

	ctx := context.Background()

	pool, err := db.New(ctx, dsn)
	if err != nil {
		log.Fatalf("migrate: %v", err)
	}
	defer pool.Close()

	if err := migrations.Run(ctx, pool); err != nil {
		log.Fatalf("migrate: %v", err)
	}

	log.Println("migrate: all migrations applied")
}
