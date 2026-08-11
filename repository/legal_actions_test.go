package repository

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TestLegalActionsRepository(t *testing.T) {
	dsn := "host=db user=admin password=secretpassword dbname=lexon_engine_test_db port=5432 sslmode=disable"
	pool, _ := pgxpool.New(context.Background(), dsn)

	repository := NewLegalActionsRepositoryHandler(context.Background(), pool)

	t.Run("Get all legal actions", func(t *testing.T) {
		actions, err := repository.GetLegalActions()
		if err != nil {
			t.Fatalf("Failed to get legal actions: %v", err)
		}

		if len(actions) == 0 {
			t.Fatalf("Expected to find legal actions, but got none")
		}
	})

	t.Run("Get legal action by ID", func(t *testing.T) {
		actionID := "ACT.CIV.0001"
		action, err := repository.GetLegalActionByID(actionID)
		if err != nil {
			t.Fatalf("Failed to get legal action by ID: %v", err)
		}

		if action == nil || action.ID != actionID {
			t.Fatalf("Expected to find legal action with ID %s, but got %+v", actionID, action)
		}
	})
}
