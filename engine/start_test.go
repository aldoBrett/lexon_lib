package engine

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestEngineStart(t *testing.T) {
	dsn := "host=db user=admin password=secretpassword dbname=lexon_engine_test_db port=5432 sslmode=disable"
	pool, _ := pgxpool.New(context.Background(), dsn)

	// Clean database before running tests
	_, err := pool.Exec(context.Background(), "DELETE FROM lexon.legal_procedures")
	if err != nil {
		t.Fatalf("Failed to clean up legal_procedures table: %v", err)
	}
	_, err = pool.Exec(context.Background(), "DELETE FROM lexon.machine_instances")
	if err != nil {
		t.Fatalf("Failed to clean up machine_instances table: %v", err)
	}
	_, err = pool.Exec(context.Background(), "DELETE FROM lexon.legal_records")
	if err != nil {
		t.Fatalf("Failed to clean up legal_records table: %v", err)
	}
	_, err = pool.Exec(context.Background(), "DELETE FROM lexon.legal_claims")
	if err != nil {
		t.Fatalf("Failed to clean up legal_claims table: %v", err)
	}

	engine := NewEngineHandler(EngineParams{
		Ctx:  context.Background(),
		Pool: pool,
	})

	legalProcedureId := uuid.New().String()
	legalProcedure := LegalProcedure{
		ID:          legalProcedureId,
		Label:       "Test Legal Procedure",
		Description: "Test Legal Procedure Description",
	}
	legalRecord := LegalRecord{
		// ID:          uuid.New().String(),
		// LegalProcedureID: &legalProcedureId,
		Actor:     "Test Actor",
		Defendant: "Test Defendant",
	}
	legalClaims := []LegalClaim{
		{
			Description: "Legal claim description",
		},
	}

	t.Run("Start an engine", func(t *testing.T) {
		engine.Start(StartParams{
			legalProcedure: legalProcedure,
			legalRecord:    legalRecord,
			legalClaims:    legalClaims,
		})
	})

}
