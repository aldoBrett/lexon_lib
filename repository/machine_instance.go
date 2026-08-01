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

type MachineInstanceRepositoryHandler struct {
	Ctx  context.Context
	pool *pgxpool.Pool
}

func NewMachineInstanceRepositoryHandler(ctx context.Context, pool *pgxpool.Pool) *MachineInstanceRepositoryHandler {
	return &MachineInstanceRepositoryHandler{
		Ctx:  ctx,
		pool: pool,
	}
}

func (h *MachineInstanceRepositoryHandler) CreateMachineInstanceForLegalProcedure(legalProcedureId string) error {
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

//? Do we need this one (._. )?
// func (h *MachineInstanceRepositoryHandler) GetMachineInstanceByID(id string) (*engine.MachineInstance, error) {
// 	query := `SELECT id, name, current_state_id, legal_procedure_id FROM lexon.machine_instances WHERE id = $1`
// 	var machineInstance engine.MachineInstance
// 	err := h.pool.QueryRow(h.Ctx, query, id).Scan(
// 		&machineInstance.ID,
// 		&machineInstance.CurrentStateID,
// 		&machineInstance.LegalProcedureID,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &machineInstance, nil
// }

func (h *MachineInstanceRepositoryHandler) GetMachineInstanceByLegalProcedureID(legalProcedureId string) (*engine.MachineInstance, error) {
	query := `SELECT id, name, current_state_id, legal_procedure_id FROM lexon.machine_instances WHERE legal_procedure_id = $1`
	var machineInstance engine.MachineInstance
	err := h.pool.QueryRow(h.Ctx, query, legalProcedureId).Scan(
		&machineInstance.ID,
		&machineInstance.CurrentStateID,
		&machineInstance.LegalProcedureID,
	)
	if err != nil {
		return nil, err
	}
	return &machineInstance, nil
}
