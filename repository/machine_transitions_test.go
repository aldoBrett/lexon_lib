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
}
