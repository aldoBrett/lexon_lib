// Command seed loads all lexon reference-table seed data.
package main

import (
	"context"
	"log"
	"os"

	"amox/lex_engine_lib/db"
	"amox/lex_engine_lib/db/seeds"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("seed: DATABASE_URL is required")
	}

	ctx := context.Background()

	pool, err := db.New(ctx, dsn)
	if err != nil {
		log.Fatalf("seed: %v", err)
	}
	defer pool.Close()

	if err := seeds.Load(ctx, pool); err != nil {
		log.Fatalf("seed: %v", err)
	}

	log.Println("seed: all seeds applied")
}
