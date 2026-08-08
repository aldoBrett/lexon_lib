package repository

import (
	"amox/lex_engine_lib/domain"
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestMachineInstancesRepository(t *testing.T) {
	dsn := "host=db user=admin password=secretpassword dbname=lexon_engine_test_db port=5432 sslmode=disable"
	pool, _ := pgxpool.New(context.Background(), dsn)

	// Clean up the machine_instances table before running tests
	_, err := pool.Exec(context.Background(), "DELETE FROM lexon.machine_instances")
	if err != nil {
		t.Fatalf("Failed to clean up machine_instances table: %v", err)
	}

	repository := NewMachineInstancesRepositoryHandler(context.Background(), pool)

	t.Run("Create MachineInstance for legal procedure", func(t *testing.T) {
		legalProcedureId := uuid.New().String()
		legalProcedure := domain.LegalProcedure{
			ID:          legalProcedureId,
			Label:       "Test Legal Procedure",
			Description: "Test Legal Procedure Description",
		}
		legalProcedureRepository := NewLegalProcedureRepositoryHandler(context.Background(), pool)
		_, err := legalProcedureRepository.SaveLegalProcedure(legalProcedure)
		if err != nil {
			t.Fatalf("Failed to save legal procedure: %v", err)
		}

		savedMachineInstance, err := repository.CreateMachineInstanceForLegalProcedure(legalProcedureId)
		if err != nil {
			t.Fatalf("Failed to create machine instance for legal procedure: %v", err)
		}

		// Verify that the machine instance was created correctly
		query := `SELECT id, legal_procedure_id FROM lexon.machine_instances WHERE legal_procedure_id = $1`
		var retrievedMachineInstance domain.MachineInstance
		err = pool.QueryRow(context.Background(), query, legalProcedureId).Scan(&retrievedMachineInstance.ID, &retrievedMachineInstance.LegalProcedureID)
		if err != nil {
			t.Fatalf("Failed to retrieve saved machine instance: %v", err)
		}

		if retrievedMachineInstance.LegalProcedureID != legalProcedureId {
			t.Fatalf("Saved machine instance does not match the expected legal procedure ID")
		}

		// Also verify by the savedMachineInstance id
		if retrievedMachineInstance.ID != savedMachineInstance.ID {
			t.Fatalf("Retrieved machine instance ID does not match the saved machine instance ID")
		}
	})

	t.Run("Create machine instance and try to create for same legal procedure again", func(t *testing.T) {
		legalProcedureId := uuid.New().String()
		legalProcedure := domain.LegalProcedure{
			ID:          legalProcedureId,
			Label:       "Test Legal Procedure",
			Description: "Test Legal Procedure Description",
		}
		legalProcedureRepository := NewLegalProcedureRepositoryHandler(context.Background(), pool)
		_, err := legalProcedureRepository.SaveLegalProcedure(legalProcedure)
		if err != nil {
			t.Fatalf("Failed to save legal procedure: %v", err)
		}

		savedMachineInstance, err := repository.CreateMachineInstanceForLegalProcedure(legalProcedureId)
		if err != nil {
			t.Fatalf("Failed to create machine instance for legal procedure: %v", err)
		}

		// Verify that the machine instance was created correctly
		query := `SELECT id, legal_procedure_id FROM lexon.machine_instances WHERE legal_procedure_id = $1`
		var retrievedMachineInstance domain.MachineInstance
		err = pool.QueryRow(context.Background(), query, legalProcedureId).Scan(&retrievedMachineInstance.ID, &retrievedMachineInstance.LegalProcedureID)
		if err != nil {
			t.Fatalf("Failed to retrieve saved machine instance: %v", err)
		}

		if retrievedMachineInstance.LegalProcedureID != legalProcedureId {
			t.Fatalf("Saved machine instance does not match the expected legal procedure ID")
		}

		// Also verify by the savedMachineInstance id
		if retrievedMachineInstance.ID != savedMachineInstance.ID {
			t.Fatalf("Retrieved machine instance ID does not match the saved machine instance ID")
		}

		// Try to create another machine instance for the same legal procedure
		_, err = repository.CreateMachineInstanceForLegalProcedure(legalProcedureId)
		if err == nil {
			t.Fatalf("Expected error when creating a second machine instance for the same legal procedure, but got none")
		}
	})

	t.Run("Get machine instance by legal procedure ID", func(t *testing.T) {
		legalProcedureId := uuid.New().String()
		legalProcedure := domain.LegalProcedure{
			ID:          legalProcedureId,
			Label:       "Test Legal Procedure",
			Description: "Test Legal Procedure Description",
		}
		legalProcedureRepository := NewLegalProcedureRepositoryHandler(context.Background(), pool)
		_, err := legalProcedureRepository.SaveLegalProcedure(legalProcedure)
		if err != nil {
			t.Fatalf("Failed to save legal procedure: %v", err)
		}

		savedMachineInstance, err := repository.CreateMachineInstanceForLegalProcedure(legalProcedureId)
		if err != nil {
			t.Fatalf("Failed to create machine instance for legal procedure: %v", err)
		}

		retrievedMachineInstance, err := repository.GetMachineInstanceByLegalProcedureID(legalProcedureId)
		if err != nil {
			t.Fatalf("Failed to retrieve machine instance by legal procedure ID: %v", err)
		}

		if retrievedMachineInstance.ID != savedMachineInstance.ID {
			t.Fatalf("Retrieved machine instance ID does not match the saved machine instance ID")
		}
	})
}
