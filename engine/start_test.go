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

		// Verify that the legal procedure was saved correctly
		savedLegalProcedure, err := engine.repos.LegalProcedure.GetLegalProcedureByID(legalProcedureId)
		if err != nil {
			t.Fatalf("Failed to get legal procedure: %v", err)
		}
		if savedLegalProcedure.ID != legalProcedureId {
			t.Fatalf("Saved legal procedure does not match the expected ID")
		}

		// Verify that the legal record was saved correctly
		legalRecords, err := engine.repos.LegalRecord.GetLegalRecordsByLegalProcedureID(legalProcedureId)
		if err != nil {
			t.Fatalf("Failed to get legal records: %v", err)
		}
		if len(legalRecords) != 1 {
			t.Fatalf("Expected 1 legal record, got %d", len(legalRecords))
		}
		retrievedLegalRecord := legalRecords[0]
		if retrievedLegalRecord.Actor != "Test Actor" || retrievedLegalRecord.Defendant != "Test Defendant" {
			t.Fatalf("Retrieved legal record does not match the expected values")
		}

		// Verify that the legal claim was saved correctly
		legalClaims, err := engine.repos.LegalClaims.GetLegalClaimsByLegalRecordID(retrievedLegalRecord.ID)
		if err != nil {
			t.Fatalf("Failed to get legal claims: %v", err)
		}
		if len(legalClaims) != 1 {
			t.Fatalf("Expected 1 legal claim, got %d", len(legalClaims))
		}
		retrievedLegalClaim := legalClaims[0]
		if retrievedLegalClaim.Description != "Legal claim description" {
			t.Fatalf("Retrieved legal claim does not match the expected values")
		}
	})

}
