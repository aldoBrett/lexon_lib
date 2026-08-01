package engine

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TestEngineProcessSignal(t *testing.T) {
	dsn := "host=db user=admin password=secretpassword dbname=lexon_engine_test_db port=5432 sslmode=disable"
	pool, _ := pgxpool.New(context.Background(), dsn)

	engine := NewEngineHandler(EngineParams{
		Ctx:  context.Background(),
		Pool: pool,
	})

	fmt.Println("engine: ", engine)
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
}
