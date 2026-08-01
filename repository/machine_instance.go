package repository

import (
	"amox/lex_engine_lib/engine"
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MachineInstanceRepository interface {
	SaveMachineInstance(machineInstance engine.MachineInstance) error
	GetMachineInstanceByID(id string) (*engine.MachineInstance, error)
}

type MachineInstanceHandler struct {
	Ctx  context.Context
	pool *pgxpool.Pool
}

func NewMachineInstanceHandler(ctx context.Context, pool *pgxpool.Pool) *MachineInstanceHandler {
	return &MachineInstanceHandler{
		Ctx:  ctx,
		pool: pool,
	}
}

func (h *MachineInstanceHandler) CreateMachineInstanceForLegalProcedure(legalProcedureId string) error {
	query := `SELECT id FROM lexon.machine_instances WHERE legal_procedure_id = $1`
	var id string
	err := h.pool.QueryRow(h.Ctx, query, legalProcedureId).Scan(&id)
	if err == nil {
		// If the legal procedure already has a machine instance, return an error
		return fmt.Errorf("legal procedure %s already has a machine instance with id %s", legalProcedureId, id)
	}

	// If the legal procedure doesn't have a machine instance, create a new one
	insertQuery := `INSERT INTO lexon.machine_instances (id, legal_procedure_id) VALUES ($1, $2)`
	_, err = h.pool.Exec(h.Ctx, insertQuery, uuid.New().String(), legalProcedureId)
	if err != nil {
		return err
	}

	return nil
}

// func (h *MachineInstanceHandler) SaveMachineInstance(machineInstance engine.MachineInstance) error {
// 	query := `SELECT id FROM lexon.machine_instances WHERE id = $1`
// 	var id string
// 	err := h.pool.QueryRow(h.Ctx, query, machineInstance.ID).Scan(&id)
// 	if err != nil {
// 		// If the machine instance doesn't exist, create a new one
// 		insertQuery := `INSERT INTO lexon.machine_instances (id, name, description) VALUES ($1, $2, $3)`
// 		_, err = h.pool.Exec(h.Ctx, insertQuery, machineInstance.ID, machineInstance.Name, machineInstance.Description)
// 		if err != nil {
// 			return err
// 		}
// 	} else {
// 		// If the machine instance exists, update its name and description
// 		updateQuery := `UPDATE lexon.machine_instances SET name = $1, description = $2 WHERE id = $3`
// 		_, err = h.pool.Exec(h.Ctx, updateQuery, machineInstance.Name, machineInstance.Description, machineInstance.ID)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

func (h *MachineInstanceHandler) GetMachineInstanceByID(id string) (*engine.MachineInstance, error) {
	query := `SELECT id, name, current_state_id, legal_procedure_id FROM lexon.machine_instances WHERE id = $1`
	var machineInstance engine.MachineInstance
	err := h.pool.QueryRow(h.Ctx, query, id).Scan(
		&machineInstance.ID,
		&machineInstance.CurrentStateID,
		&machineInstance.LegalProcedureID,
	)
	if err != nil {
		return nil, err
	}
	return &machineInstance, nil
}
