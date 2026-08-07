package repository

import (
	"context"
	"testing"

	"amox/lex_engine_lib/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestRepositoryLegalClaim(t *testing.T) {
	dsn := "host=db user=admin password=secretpassword dbname=lexon_engine_test_db port=5432 sslmode=disable"
	pool, _ := pgxpool.New(context.Background(), dsn)

	// Clean up the legal_claims table before running tests
	_, err := pool.Exec(context.Background(), "DELETE FROM lexon.legal_claims")
	if err != nil {
		t.Fatalf("Failed to clean up legal_claims table: %v", err)
	}

	repository := NewLegalClaimsRepositoryHandler(context.Background(), pool)

	legalProcedureId := uuid.New().String()
	legalProcedure := domain.LegalProcedure{
		ID:          legalProcedureId,
		Label:       "Test Legal Procedure",
		Description: "Test Legal Procedure Description",
	}
	legalProcedureRepository := NewLegalProcedureRepositoryHandler(context.Background(), pool)
	if _, err := legalProcedureRepository.SaveLegalProcedure(legalProcedure); err != nil {
		t.Fatalf("Failed to save legal procedure: %v", err)
	}
	legalRecord := domain.LegalRecord{
		LegalProcedureID: &legalProcedureId,
		Actor:            "Test Actor",
		Defendant:        "Test Defendant",
	}
	legalRecordRepository := NewLegalRecordRepositoryHandler(context.Background(), pool)
	savedLegalRecord, err := legalRecordRepository.SaveLegalRecord(legalRecord)
	if err != nil {
		t.Fatalf("Failed to save legal record: %v", err)
	}

	t.Run("Save LegalClaim and get it by LegalRecordID", func(t *testing.T) {
		legalClaim := domain.LegalClaim{
			LegalRecordID: savedLegalRecord.ID,
			Description:   "Test Legal Claim Description",
		}
		savedLegalClaim, err := repository.SaveLegalClaim(legalClaim)
		if err != nil {
			t.Fatalf("Failed to save legal claim: %v", err)
		}
		if savedLegalClaim.ID == "" {
			t.Fatalf("Saved legal claim ID should not be empty")
		}

		// Verify that the legal claim can be retrieved by legal record ID
		legalClaims, err := repository.GetLegalClaimsByLegalRecordID(savedLegalRecord.ID)
		if err != nil {
			t.Fatalf("Failed to get legal claims by legal record ID: %v", err)
		}
		if len(legalClaims) != 1 {
			t.Fatalf("Expected 1 legal claim, got %d", len(legalClaims))
		}
		retrievedLegalClaim := legalClaims[0]
		if retrievedLegalClaim.ID != savedLegalClaim.ID {
			t.Fatalf("Retrieved legal claim does not match the expected ID")
		}
		if retrievedLegalClaim.Description != "Test Legal Claim Description" {
			t.Fatalf("Retrieved legal claim description does not match the expected value")
		}
	})

	t.Run("Save and update LegalClaim", func(t *testing.T) {
		legalClaim := domain.LegalClaim{
			LegalRecordID: savedLegalRecord.ID,
			Description:   "Test Legal Claim Description",
		}
		savedLegalClaim, err := repository.SaveLegalClaim(legalClaim)
		if err != nil {
			t.Fatalf("Failed to save legal claim: %v", err)
		}

		// Update the description of the saved legal claim
		savedLegalClaim.Description = "Updated Legal Claim Description"
		updatedLegalClaim, err := repository.SaveLegalClaim(*savedLegalClaim)
		if err != nil {
			t.Fatalf("Failed to update legal claim: %v", err)
		}

		// Verify that the legal claim was updated correctly
		if updatedLegalClaim.Description != "Updated Legal Claim Description" {
			t.Fatalf("Updated legal claim description does not match the expected value")
		}
	})
}
