package repository

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TestMachineStatesRepository(t *testing.T) {
	dsn := "host=db user=admin password=secretpassword dbname=lexon_engine_test_db port=5432 sslmode=disable"
	pool, _ := pgxpool.New(context.Background(), dsn)

	repository := NewMachineStatesRepositoryHandler(context.Background(), pool)

	t.Run("Get all machine states with pagination", func(t *testing.T) {
		offset := 0
		limit := 10
		states, err := repository.GetMachineStates(offset, limit)
		if err != nil {
			t.Fatalf("Failed to get machine states: %v", err)
		}

		if len(states) == 0 {
			t.Fatalf("Expected to find machine states, but got none")
		}
	})

	t.Run("Get machine state by ID", func(t *testing.T) {
		stateID := "CIV.ORD.S00" // Use a valid state ID from your seed data
		state, err := repository.GetMachineStateByID(stateID)
		if err != nil {
			t.Fatalf("Failed to get machine state by ID: %v", err)
		}

		if state == nil || state.ID != stateID {
			t.Fatalf("Expected to find machine state with ID %s, but got %+v", stateID, state)
		}
	})

	t.Run("Count machine states", func(t *testing.T) {
		count, err := repository.CountMachineStates()
		if err != nil {
			t.Fatalf("Failed to count machine states: %v", err)
		}

		if count <= 0 {
			t.Fatalf("Expected to find at least one machine state, but got count: %d", count)
		}
	})
}
