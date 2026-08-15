package repository

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TestMachineTransitionsRepository(t *testing.T) {
	dsn := "host=db user=admin password=secretpassword dbname=lexon_engine_test_db port=5432 sslmode=disable"
	pool, _ := pgxpool.New(context.Background(), dsn)

	repository := NewMachineTransitionsRepositoryHandler(context.Background(), pool)

	t.Run("Get machine transition by ID", func(t *testing.T) {
		// Insert a test machine transition into the database
		transitionID := "T001"

		// Retrieve the machine transition by ID
		machineTransition, err := repository.GetMachineTransitionByID(transitionID)
		if err != nil {
			t.Fatalf("Failed to get machine transition by ID: %v", err)
		}

		if machineTransition == nil {
			t.Fatalf("Expected machine transition, but got nil")
		}

		if machineTransition.ID != transitionID {
			t.Fatalf("Expected transition ID %s, but got %s", transitionID, machineTransition.ID)
		}
		if machineTransition.SourceStateID == "" {
			t.Fatalf("Expected non-empty source state ID, but got empty")
		}
	})

	t.Run("Get machine transition by source state ID and event ID", func(t *testing.T) {
		sourceStateID := "CIV.ORD.S00"
		eventID := "CIV.ORD.E001"

		machineTransition, err := repository.GetMachineTransitionBySourceAndEvent(sourceStateID, eventID)
		if err != nil {
			t.Fatalf("Failed to get machine transition by source state ID and event ID: %v", err)
		}

		if machineTransition == nil {
			t.Fatalf("Expected machine transition, but got nil")
		}

		if machineTransition.SourceStateID != sourceStateID {
			t.Fatalf("Expected source state ID %s, but got %s", sourceStateID, machineTransition.SourceStateID)
		}
		if machineTransition.EventID != eventID {
			t.Fatalf("Expected event ID %s, but got %s", eventID, machineTransition.EventID)
		}
	})

	t.Run("Get machine transitions by source state ID", func(t *testing.T) {
		sourceStateID := "CIV.ORD.S00"

		machineTransitions, err := repository.GetMachineStateTransitionsBySourceStateID(sourceStateID)
		if err != nil {
			t.Fatalf("Failed to get machine transitions by source state ID: %v", err)
		}

		if len(machineTransitions) == 0 {
			t.Fatalf("Expected at least one machine transition, but got none")
		}

		for _, transition := range machineTransitions {
			if transition.SourceStateID != sourceStateID {
				t.Fatalf("Expected source state ID %s, but got %s", sourceStateID, transition.SourceStateID)
			}
		}
	})
}
