package engine

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TestEngineProcessSignal(t *testing.T) {
	dsn := "host=localhost user=admin password=secretpassword dbname=lexon_engine_test_db port=5432 sslmode=disable"
	pool, _ := pgxpool.New(context.Background(), dsn)

	engine := NewEngineHandler(EngineParams{
		Ctx:  context.Background(),
		Pool: pool,
	})
	signal := MachineSignal{
		Origin: "document",
		// DocumentType: "test_document_type",
	}

	engine.EngineProcessSignal(signal)
}
