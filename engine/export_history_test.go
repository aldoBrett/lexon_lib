package engine

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestExportHistory(t *testing.T) {
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
	_, err = pool.Exec(context.Background(), "DELETE FROM lexon.machine_state_transitions_history")
	if err != nil {
		t.Fatalf("Failed to clean up machine_state_transitions_history table: %v", err)
	}

	t.Run("Create an engine with initialize set to true and arrive to CIV.ORD.E001, and retrieve the history", func(t *testing.T) {
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
		legalActionId := "ACT.CIV.0001"
		legalRecord := LegalRecord{
			LegalProcedureID: &legalProcedureId,
			Actor:            "Test Actor",
			Defendant:        "Test Defendant",
			LegalActionID:    &legalActionId,
		}
		legalClaims := []LegalClaim{
			{
				Description: "Legal claim description",
			},
		}

		initialize := true
		engine.Create(CreateParams{
			LegalProcedure: legalProcedure,
			LegalRecord:    legalRecord,
			LegalClaims:    legalClaims,
			Initialize:     &initialize,
		})

		// Get the current state of the machine instance
		machineInstance, err := engine.repos.MachineInstances.GetMachineInstanceByLegalProcedureID(legalProcedureId)
		if err != nil {
			t.Fatalf("Failed to retrieve machine instance: %v", err)
		}

		// TODO: we have to check this final state when we add the other initialization signal
		currentStateID := "CIV.ORD.S01"
		if machineInstance.CurrentStateID == nil || *machineInstance.CurrentStateID != currentStateID {
			t.Errorf("Expected current state ID to be %s, but got %v", currentStateID, machineInstance.CurrentStateID)
		}

		machineStateStateID := "CIV.M01"
		history, err := engine.ExportHistory(ExportHistoryParams{
			LegalProcedureID:    legalProcedureId,
			MachineStateStageID: &machineStateStateID,
		})
		if err != nil {
			t.Fatalf("Failed to export history: %v", err)
		}
		if len(history) == 0 {
			t.Fatalf("Expected at least one history item, but got none")
		}
		for _, item := range history {
			if item.ID == "" || item.From == "" || item.To == "" {
				t.Errorf("Expected history item to have id, from and to set, got %+v", item)
			}
			if item.Event.ID == "" || item.Event.Name == "" {
				t.Errorf("Expected history item event to have id and name set, got %+v", item.Event)
			}
		}
	})
}
