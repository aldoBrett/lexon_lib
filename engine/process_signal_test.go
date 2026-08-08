package engine

import (
	"amox/lex_engine_lib/domain"
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestEngineProcessSignal(t *testing.T) {
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

	t.Run("Start an engine and move by EventID", func(t *testing.T) {
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
			Actor:     "Test Actor",
			Defendant: "Test Defendant",
		}
		legalClaims := []LegalClaim{
			{
				Description: "Legal claim description",
			},
		}

		initialize := false
		engine.Create(CreateParams{
			legalProcedure: legalProcedure,
			legalRecord:    legalRecord,
			legalClaims:    legalClaims,
			initialize:     &initialize,
		})

		eventID := "CIV.ORD.E001"
		updatedStateMachine, err := engine.ProcessSignal(ProcessSignalParams{
			signal: domain.MachineSignal{
				ID:      uuid.New().String(),
				Code:    "",
				Origin:  "system",
				EventID: &eventID,
			},
		})
		if err != nil {
			t.Fatalf("Failed to process signal: %v", err)
		}

		if updatedStateMachine == nil || updatedStateMachine.CurrentStateID == nil {
			t.Fatalf("Expected updated state machine with a current state, got nil")
		}
		if *updatedStateMachine.CurrentStateID != "CIV.ORD.S00" {
			t.Fatalf("Expected current state ID to be 'CIV.ORD.S00', got '%s'", *updatedStateMachine.CurrentStateID)
		}

		var historyCount int
		err = pool.QueryRow(
			context.Background(),
			"SELECT COUNT(*) FROM lexon.machine_state_transitions_history WHERE machine_instance_id = $1",
			updatedStateMachine.ID,
		).Scan(&historyCount)
		if err != nil {
			t.Fatalf("Failed to query machine_state_transitions_history: %v", err)
		}
		if historyCount != 1 {
			t.Fatalf("Expected 1 machine_state_transitions_history row for machine instance, got %d", historyCount)
		}
	})

	t.Run("Start engine and move by EventID to CIV.ORD.E002", func(t *testing.T) {
		t.Skip("")
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
			Actor:     "Test Actor",
			Defendant: "Test Defendant",
		}
		legalClaims := []LegalClaim{
			{
				Description: "Legal claim description",
			},
		}

		initialize := false
		engine.Create(CreateParams{
			legalProcedure: legalProcedure,
			legalRecord:    legalRecord,
			legalClaims:    legalClaims,
			initialize:     &initialize,
		})

		eventID := "CIV.ORD.E001"
		engine.ProcessSignal(ProcessSignalParams{
			signal: domain.MachineSignal{
				ID:      uuid.New().String(),
				Code:    "",
				Origin:  "system",
				EventID: &eventID,
			},
		})

		eventID = "CIV.ORD.E002"
		updatedStateMachine, err := engine.ProcessSignal(ProcessSignalParams{
			signal: domain.MachineSignal{
				ID:      uuid.New().String(),
				Code:    "",
				Origin:  "system",
				EventID: &eventID,
			},
		})
		if err != nil {
			t.Fatalf("Failed to process signal: %v", err)
		}

		if updatedStateMachine == nil || updatedStateMachine.CurrentStateID == nil {
			t.Fatalf("Expected updated state machine with a current state, got nil")
		}
		if *updatedStateMachine.CurrentStateID != "CIV.ORD.S01" {
			t.Fatalf("Expected current state ID to be 'CIV.ORD.S01', got '%s'", *updatedStateMachine.CurrentStateID)
		}

		eventID = "CIV.ORD.E004"
		updatedStateMachine, err = engine.ProcessSignal(ProcessSignalParams{
			signal: domain.MachineSignal{
				ID:      uuid.New().String(),
				Code:    "",
				Origin:  "system",
				EventID: &eventID,
			},
		})
		if err != nil {
			t.Fatalf("Failed to process signal: %v", err)
		}

		if updatedStateMachine == nil || updatedStateMachine.CurrentStateID == nil {
			t.Fatalf("Expected updated state machine with a current state, got nil")
		}
		if *updatedStateMachine.CurrentStateID != "CIV.ORD.S02" {
			t.Fatalf("Expected current state ID to be 'CIV.ORD.S02', got '%s'", *updatedStateMachine.CurrentStateID)
		}
	})
}
