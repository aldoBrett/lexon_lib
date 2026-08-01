package repository

import (
	"context"
	"testing"

	"amox/lex_engine_lib/engine"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestRepositoryMachineInstance(t *testing.T) {
	dsn := "host=db user=admin password=secretpassword dbname=lexon_engine_test_db port=5432 sslmode=disable"
	pool, _ := pgxpool.New(context.Background(), dsn)

	t.Run("Create MachineInstance for legal procedure", func(t *testing.T) {
		legalProcedureId := uuid.New().String()
		legalProcedure := engine.LegalProcedure{
			ID:          legalProcedureId,
			Label:       "Test Legal Procedure",
			Description: "Test Legal Procedure Description",
		}

		//? Is this necessary? Maybe if we check on the creationg of the machine instance that the
		//? legalProcedure exists...
		legalProcedureRepository := NewLegalProcedureRepositoryHandler(context.Background(), pool)
		err := legalProcedureRepository.SaveLegalProcedure(legalProcedure)
		if err != nil {
			t.Fatalf("Failed to save legal procedure: %v", err)
		}

		repository := NewMachineInstanceRepositoryHandler(context.Background(), pool)
		err = repository.CreateMachineInstanceForLegalProcedure(legalProcedureId)
		if err != nil {
			t.Fatalf("Failed to create machine instance for legal procedure: %v", err)
		}

		// Verify that the machine instance was created correctly
		query := `SELECT id, legal_procedure_id FROM lexon.machine_instances WHERE legal_procedure_id = $1`
		var savedMachineInstance engine.MachineInstance
		err = pool.QueryRow(context.Background(), query, legalProcedureId).Scan(&savedMachineInstance.ID, &savedMachineInstance.LegalProcedureID)
		if err != nil {
			t.Fatalf("Failed to retrieve saved machine instance: %v", err)
		}

		if savedMachineInstance.LegalProcedureID != legalProcedureId {
			t.Fatalf("Saved machine instance does not match the expected legal procedure ID")
		}
	})

	t.Run("Create MachineInstance for legal procedure that already has one", func(t *testing.T) {
		legalProcedureId := uuid.New().String()
		legalProcedure := engine.LegalProcedure{
			ID:          legalProcedureId,
			Label:       "Test Legal Procedure",
			Description: "Test Legal Procedure Description",
		}
		legalProcedureRepository := NewLegalProcedureRepositoryHandler(context.Background(), pool)
		err := legalProcedureRepository.SaveLegalProcedure(legalProcedure)
		if err != nil {
			t.Fatalf("Failed to save legal procedure: %v", err)
		}

		repository := NewMachineInstanceRepositoryHandler(context.Background(), pool)
		err = repository.CreateMachineInstanceForLegalProcedure(legalProcedureId)
		if err != nil {
			t.Fatalf("Failed to create machine instance for legal procedure: %v", err)
		}

		// Try to create another machine instance for the same legal procedure
		err = repository.CreateMachineInstanceForLegalProcedure(legalProcedureId)
		if err == nil {
			t.Fatalf("Expected error when creating a second machine instance for the same legal procedure, but got none")
		}
	})

	t.Run("Create MachineInstance for non-existent legal procedure", func(t *testing.T) {
		nonExistentLegalProcedureId := uuid.New().String()
		repository := NewMachineInstanceRepositoryHandler(context.Background(), pool)
		err := repository.CreateMachineInstanceForLegalProcedure(nonExistentLegalProcedureId)
		if err == nil {
			t.Fatalf("Expected error when creating a machine instance for a non-existent legal procedure, but got none")
		}
	})

	t.Run("Get MachineInstance by legal procedure ID", func(t *testing.T) {
		legalProcedureId := uuid.New().String()
		legalProcedure := engine.LegalProcedure{
			ID:          legalProcedureId,
			Label:       "Test Legal Procedure",
			Description: "Test Legal Procedure Description",
		}
		legalProcedureRepository := NewLegalProcedureRepositoryHandler(context.Background(), pool)
		err := legalProcedureRepository.SaveLegalProcedure(legalProcedure)
		if err != nil {
			t.Fatalf("Failed to save legal procedure: %v", err)
		}

		repository := NewMachineInstanceRepositoryHandler(context.Background(), pool)
		err = repository.CreateMachineInstanceForLegalProcedure(legalProcedureId)
		if err != nil {
			t.Fatalf("Failed to create machine instance for legal procedure: %v", err)
		}

		machineInstance, err := repository.GetMachineInstanceByLegalProcedureID(legalProcedureId)
		if err != nil {
			t.Fatalf("Failed to get machine instance by legal procedure ID: %v", err)
		}

		if machineInstance.LegalProcedureID != legalProcedureId {
			t.Fatalf("Retrieved machine instance does not match the expected legal procedure ID")
		}
	})

	t.Run("Get MachineInstance by legal procedure ID that doesn't exist", func(t *testing.T) {
		nonExistentLegalProcedureId := uuid.New().String()
		repository := NewMachineInstanceRepositoryHandler(context.Background(), pool)
		_, err := repository.GetMachineInstanceByLegalProcedureID(nonExistentLegalProcedureId)
		if err == nil {
			t.Fatalf("Expected error when retrieving machine instance for non-existent legal procedure ID, but got none")
		}
	})
}
