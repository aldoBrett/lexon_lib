package repository

import (
	"amox/lex_engine_lib/engine"
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestRepositoryLegalProcedure(t *testing.T) {
	dsn := "host=db user=admin password=secretpassword dbname=lexon_engine_test_db port=5432 sslmode=disable"
	pool, _ := pgxpool.New(context.Background(), dsn)

	t.Run("Save LegalProcedure", func(t *testing.T) {
		legalProcedureId := uuid.New().String()
		legalProcedure := engine.LegalProcedure{
			ID:          legalProcedureId,
			Label:       "Test Legal Procedure",
			Description: "Test Legal Procedure Description",
		}

		repository := NewLegalProcedureRepositoryHandler(context.Background(), pool)
		err := repository.SaveLegalProcedure(legalProcedure)
		if err != nil {
			t.Fatalf("Failed to save legal procedure: %v", err)
		}

		// Verify that the legal procedure was saved correctly
		query := `SELECT id, label, description FROM lexon.legal_procedures WHERE id = $1`
		var savedLegalProcedure engine.LegalProcedure
		err = pool.QueryRow(context.Background(), query, legalProcedureId).Scan(&savedLegalProcedure.ID, &savedLegalProcedure.Label, &savedLegalProcedure.Description)
		if err != nil {
			t.Fatalf("Failed to retrieve saved legal procedure: %v", err)
		}

		if savedLegalProcedure.ID != legalProcedure.ID || savedLegalProcedure.Label != legalProcedure.Label || savedLegalProcedure.Description != legalProcedure.Description {
			t.Fatalf("Saved legal procedure does not match the expected values")
		}
	})

	t.Run("Save and get LegalProcedure", func(t *testing.T) {
		legalProcedureId := uuid.New().String()
		legalProcedure := engine.LegalProcedure{
			ID:          legalProcedureId,
			Label:       "Test Legal Procedure",
			Description: "Test Legal Procedure Description",
		}

		repository := NewLegalProcedureRepositoryHandler(context.Background(), pool)
		err := repository.SaveLegalProcedure(legalProcedure)
		if err != nil {
			t.Fatalf("Failed to save legal procedure: %v", err)
		}

		// Now retrieve the legal procedure by ID
		retrievedLegalProcedure, err := repository.GetLegalProcedureByID(legalProcedureId)
		if err != nil {
			t.Fatalf("Failed to retrieve legal procedure by ID: %v", err)
		}

		if retrievedLegalProcedure.ID != legalProcedure.ID || retrievedLegalProcedure.Label != legalProcedure.Label || retrievedLegalProcedure.Description != legalProcedure.Description {
			t.Fatalf("Retrieved legal procedure does not match the expected values")
		}
	})

	t.Run("Get LegalProcedure not found", func(t *testing.T) {
		nonExistentId := uuid.New().String()
		repository := NewLegalProcedureRepositoryHandler(context.Background(), pool)
		_, err := repository.GetLegalProcedureByID(nonExistentId)
		if err == nil {
			t.Fatalf("Expected error when retrieving non-existent legal procedure, but got none")
		}
	})
}
