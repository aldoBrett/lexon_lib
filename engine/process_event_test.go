package engine

import (
	"fmt"
	"testing"
)

func TestEngineProcessEvent(t *testing.T) {
	// dsn := "host=localhost user=admin password=secretpassword dbname=lexon_engine_test_db port=5432 sslmode=disable"
	// pool, _ := pgxpool.New(context.Background(), dsn)

	fmt.Println("---->TestEngineProcessEvent<----")
	// engine := NewEngineHandler(EngineParams{
	// 	Pool: pool,
	// })
	// event := MachineEvent{
	// 	Origin:       "test_origin",
	// 	DocumentType: "test_document_type",
	// }

	// engine.EngineProcessEvent(event)

	fmt.Println("----END----")
}
