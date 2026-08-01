package engine

import (
	"context"
	"os"
	"testing"

	"github.com/google/uuid"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TestEngineGetLegalProcedure(t *testing.T) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=db user=admin password=secretpassword dbname=lexon_engine_test_db port=5432 sslmode=disable"
	}
	pool, _ := pgxpool.New(context.Background(), dsn)

	// Remove the legal_procedures in the database
	_, err := pool.Exec(context.Background(), `DELETE FROM lexon.legal_procedures`)
	if err != nil {
		t.Fatalf("Failed to clean up legal procedures: %v", err)
	}

	engine := NewEngineHandler(EngineParams{
		Ctx:  context.Background(),
		Pool: pool,
	})

	t.Run("GetLegalProcedure", func(t *testing.T) {
		id := uuid.New().String()

		legalProcedure := LegalProcedure{
			ID:          id,
			Label:       "Test Legal Procedure",
			Description: "Test Legal Procedure Description",
		}

		engine.EngineSaveLegalProcedure(legalProcedure)

		// Verify that the legal procedure was retrieved correctly
		retrievedLegalProcedure := engine.EngineGetLegalProcedure(id)
		if retrievedLegalProcedure.ID != legalProcedure.ID || retrievedLegalProcedure.Label != legalProcedure.Label || retrievedLegalProcedure.Description != legalProcedure.Description {
			t.Fatalf("Retrieved legal procedure does not match the expected values")
		}
	})

	t.Run("GetLegalProcedureNotFound", func(t *testing.T) {
		id := uuid.New().String()

		// Verify that retrieving a non-existent legal procedure panics
		defer func() {
			if r := recover(); r == nil {
				t.Fatalf("Expected panic when retrieving non-existent legal procedure, but did not panic")
			}
		}()
		engine.EngineGetLegalProcedure(id)
	})
}
