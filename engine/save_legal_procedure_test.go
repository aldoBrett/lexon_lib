package engine

import (
	"context"
	"os"
	"testing"

	"github.com/google/uuid"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TestEngineSaveLegalProcedure(t *testing.T) {
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

	t.Run("SaveLegalProcedure", func(t *testing.T) {
		id := uuid.New().String()

		legalProcedure := LegalProcedure{
			ID:          id,
			Name:        "Test Legal Procedure",
			Description: "Test Legal Procedure Description",
		}

		engine.EngineSaveLegalProcedure(legalProcedure)

		// Verify that the legal procedure was saved correctly
		query := `SELECT id, name, description FROM lexon.legal_procedures WHERE id = $1`
		var savedLegalProcedure LegalProcedure
		err := engine.pool.QueryRow(engine.ctx, query, id).Scan(&savedLegalProcedure.ID, &savedLegalProcedure.Name, &savedLegalProcedure.Description)
		if err != nil {
			t.Fatalf("Failed to retrieve saved legal procedure: %v", err)
		}

		if savedLegalProcedure.ID != legalProcedure.ID || savedLegalProcedure.Name != legalProcedure.Name || savedLegalProcedure.Description != legalProcedure.Description {
			t.Fatalf("Saved legal procedure does not match the expected values")
		}
	})
}
