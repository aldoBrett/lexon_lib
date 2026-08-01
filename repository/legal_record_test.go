package repository

import (
	"context"
	"testing"

	"amox/lex_engine_lib/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestRepositoryLegalRecord(t *testing.T) {
	dsn := "host=db user=admin password=secretpassword dbname=lexon_engine_test_db port=5432 sslmode=disable"
	pool, _ := pgxpool.New(context.Background(), dsn)

	// Clean up the legal_records and legal_procedures tables before running tests
	_, err := pool.Exec(context.Background(), "DELETE FROM lexon.legal_records")
	if err != nil {
		t.Fatalf("Failed to clean up legal_records table: %v", err)
	}
	_, err = pool.Exec(context.Background(), "DELETE FROM lexon.legal_procedures")
	if err != nil {
		t.Fatalf("Failed to clean up legal_procedures table: %v", err)
	}

	repository := NewLegalRecordRepositoryHandler(context.Background(), pool)

	t.Run("Save LegalRecord", func(t *testing.T) {
		legalProcedureId := uuid.New().String()
		legalProcedure := domain.LegalProcedure{
			ID:          legalProcedureId,
			Label:       "Test Legal Procedure",
			Description: "Test Legal Procedure Description",
		}
		legalProcedureRepository := NewLegalProcedureRepositoryHandler(context.Background(), pool)
		if err := legalProcedureRepository.SaveLegalProcedure(legalProcedure); err != nil {
			t.Fatalf("Failed to save legal procedure: %v", err)
		}

		legalRecord := domain.LegalRecord{
			LegalProcedureID: &legalProcedureId,
			Actor:            "Test Actor",
			Defendant:        "Test Defendant",
		}

		savedLegalRecord, err := repository.SaveLegalRecord(legalRecord)
		if err != nil {
			t.Fatalf("Failed to save legal record: %v", err)
		}

		// Verify that the legal record was saved correctly
		if savedLegalRecord.LegalProcedureID == nil || *savedLegalRecord.LegalProcedureID != legalProcedureId {
			t.Fatalf("Saved legal record does not match the expected legal procedure ID")
		}
		if savedLegalRecord.Actor != "Test Actor" || savedLegalRecord.Defendant != "Test Defendant" {
			t.Fatalf("Saved legal record does not match the expected values")
		}
		if savedLegalRecord.ID == "" {
			t.Fatalf("Saved legal record ID should not be empty")
		}
	})

	t.Run("Save and edit LegalRecord", func(t *testing.T) {
		legalProcedureId := uuid.New().String()
		legalProcedure := domain.LegalProcedure{
			ID:          legalProcedureId,
			Label:       "Test Legal Procedure",
			Description: "Test Legal Procedure Description",
		}
		legalProcedureRepository := NewLegalProcedureRepositoryHandler(context.Background(), pool)
		if err := legalProcedureRepository.SaveLegalProcedure(legalProcedure); err != nil {
			t.Fatalf("Failed to save legal procedure: %v", err)
		}

		legalRecord := domain.LegalRecord{
			LegalProcedureID: &legalProcedureId,
			Actor:            "Test Actor",
			Defendant:        "Test Defendant",
		}

		savedLegalRecord, err := repository.SaveLegalRecord(legalRecord)
		if err != nil {
			t.Fatalf("Failed to save legal record: %v", err)
		}

		// Edit the legal record
		savedLegalRecord.Actor = "Updated Actor"
		savedLegalRecord.Defendant = "Updated Defendant"
		updatedLegalRecord, err := repository.SaveLegalRecord(*savedLegalRecord)
		if err != nil {
			t.Fatalf("Failed to edit legal record: %v", err)
		}

		// Verify that the legal record was updated correctly
		if updatedLegalRecord.Actor != "Updated Actor" || updatedLegalRecord.Defendant != "Updated Defendant" {
			t.Fatalf("Updated legal record does not match the expected values")
		}
	})
}
