package repository

import (
	"testing"

	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TestMachineStateStagesRepository(t *testing.T) {
	dsn := "host=db user=admin password=secretpassword dbname=lexon_engine_test_db port=5432 sslmode=disable"
	pool, _ := pgxpool.New(context.Background(), dsn)

	repository := NewMachineStateStagesRepositoryHandler(context.Background(), pool)

	t.Run("Get all machine state stages", func(t *testing.T) {
		stages, err := repository.GetMachineStateStages()
		if err != nil {
			t.Fatalf("Failed to get machine state stages: %v", err)
		}

		if len(stages) == 0 {
			t.Fatalf("Expected to find machine state stages, but got none")
		}
	})

	t.Run("Get machine state stage by ID", func(t *testing.T) {
		stageID := "CIV.M01" // Use a valid stage ID from your seed data
		stage, err := repository.GetMachineStateStageByID(stageID)
		if err != nil {
			t.Fatalf("Failed to get machine state stage by ID: %v", err)
		}

		if stage == nil || stage.ID != stageID {
			t.Fatalf("Expected to find machine state stage with ID %s, but got %+v", stageID, stage)
		}
	})

}
