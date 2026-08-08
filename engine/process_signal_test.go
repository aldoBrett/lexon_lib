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

	// legalProcedureId := uuid.New().String()
	// legalProcedure := LegalProcedure{
	// 	ID:          legalProcedureId,
	// 	Label:       "Test Legal Procedure",
	// 	Description: "Test Legal Procedure Description",
	// }
	// // También vamos a agregar la info del expediente base.
	// // Y también vienen las pretenciones.
	// engine.EngineSaveLegalProcedure(legalProcedure)

	// // signal := MachineSignal{
	// // 	Origin: "document",
	// // 	// DocumentType: "test_document_type",
	// // }

	// // engine.EngineProcessSignal(signal)
	// t.Run("Move until CIV.ORD.E003", func(t *testing.T) {
	// 	singalZero := MachineSignal{
	// 		// Origin: "document",
	// 		Origin: "ui",
	// 	}

	// 	engine.EngineProcessSignal(ProcessSignalParams{
	// 		signal:           singalZero,
	// 		legalProcedureID: &legalProcedureId,
	// 	})
	// })

	t.Run("Start and engine and move by TransitionID", func(t *testing.T) {
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

		initialize := false
		engine.Create(CreateParams{
			legalProcedure: legalProcedure,
			legalRecord:    legalRecord,
			legalClaims:    legalClaims,
			initialize:     &initialize,
		})

		transitionID := "T001"
		updatedStateMachine, err := engine.ProcessSignal(ProcessSignalParams{
			signal: domain.MachineSignal{
				ID:           uuid.New().String(),
				Code:         "",
				Origin:       "system",
				TransitionID: &transitionID,
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
	})
}
