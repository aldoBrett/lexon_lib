package engine

import (
	"fmt"

	"github.com/google/uuid"
)

// The association from a legal procedure to a machine instance is one-to-one.
// So we're going to check if there's already a machine instance with the
// legalProcedureID, if there is, we will return an error. If there isn't,
// we will create a new machine instance with the legalProcedureID and the
// initial state of the machine.
func (e *EngineHandler) EngineCreate(legalProcedureID string) error {
	query := `SELECT id FROM lexon.machine_instances WHERE legal_procedure_id = $1`
	var id string
	err := e.pool.QueryRow(e.ctx, query, legalProcedureID).Scan(&id)
	if err == nil {
		// If the legal procedure already has a machine instance, return an error
		return fmt.Errorf("legal procedure %s already has a machine instance with id %s", legalProcedureID, id)
	}

	// If the legal procedure doesn't have a machine instance, create a new one
	insertQuery := `INSERT INTO lexon.machine_instances (id, legal_procedure_id) VALUES ($1, $2)`
	_, err = e.pool.Exec(e.ctx, insertQuery, uuid.New().String(), legalProcedureID)
	if err != nil {
		return err
	}

	return nil
}
