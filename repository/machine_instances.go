package repository

import (
	"context"
	"fmt"

	"amox/lex_engine_lib/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MachineInstancesRepository interface {
	CreateMachineInstanceForLegalProcedure(legalProcedureID string) (*domain.MachineInstance, error)
	GetMachineInstanceByLegalProcedureID(legalProcedureID string) (*domain.MachineInstance, error)
	UpdateMachineInstanceCurrentState(machineInstanceID string, nextState string) (*domain.MachineInstance, error)
}

type MachineInstancesRepositoryHandler struct {
	pool *pgxpool.Pool
	Ctx  context.Context
}

func NewMachineInstancesRepositoryHandler(ctx context.Context, pool *pgxpool.Pool) *MachineInstancesRepositoryHandler {
	return &MachineInstancesRepositoryHandler{
		pool: pool,
		Ctx:  ctx,
	}
}

// Create a new MachineInstance for a give LegalProcedure. First we're going to check if there exists
// already a MachineInstance for the given LegalProcedure, if it does, we're going to return an error.
// If it doesn't exist, we're going to create a new MachineInstance and return it.
func (h *MachineInstancesRepositoryHandler) CreateMachineInstanceForLegalProcedure(legalProcedureID string) (*domain.MachineInstance, error) {
	// Check if a MachineInstance already exists for the given LegalProcedure
	existingInstance, err := h.GetMachineInstanceByLegalProcedureID(legalProcedureID)
	if err == nil && existingInstance != nil {
		return nil, fmt.Errorf("machine instance already exists for legal procedure ID: %s", legalProcedureID)
	}

	machineInstanceID := uuid.New().String()
	insertQuery := `INSERT INTO lexon.machine_instances (id, legal_procedure_id, created_at, updated_at) VALUES ($1, $2, NOW(), NOW())`
	_, err = h.pool.Exec(h.Ctx, insertQuery, machineInstanceID, legalProcedureID)
	if err != nil {
		return nil, err
	}
	return &domain.MachineInstance{
		ID:               machineInstanceID,
		LegalProcedureID: legalProcedureID,
	}, nil
}

// Since we have only one MachineInstance per LegalProcedure, we can get the MachineInstance by the LegalProcedureID.
// If it doesn't exist, we're going to return an error.
func (h *MachineInstancesRepositoryHandler) GetMachineInstanceByLegalProcedureID(legalProcedureID string) (*domain.MachineInstance, error) {
	machineInstance := &domain.MachineInstance{}
	query := `SELECT id, current_state_id, legal_procedure_id, created_at, updated_at FROM lexon.machine_instances WHERE legal_procedure_id = $1`
	err := h.pool.QueryRow(h.Ctx, query, legalProcedureID).Scan(
		&machineInstance.ID,
		&machineInstance.CurrentStateID,
		&machineInstance.LegalProcedureID,
		&machineInstance.CreatedAt,
		&machineInstance.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return machineInstance, nil
}

func (h *MachineInstancesRepositoryHandler) UpdateMachineInstanceCurrentState(machineInstanceID string, nextState string) (*domain.MachineInstance, error) {
	machineInstance := &domain.MachineInstance{}
	updateQuery := `UPDATE lexon.machine_instances SET current_state_id = $1, updated_at = NOW() WHERE id = $2
		RETURNING id, current_state_id, legal_procedure_id, created_at, updated_at`
	err := h.pool.QueryRow(h.Ctx, updateQuery, nextState, machineInstanceID).Scan(
		&machineInstance.ID,
		&machineInstance.CurrentStateID,
		&machineInstance.LegalProcedureID,
		&machineInstance.CreatedAt,
		&machineInstance.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return machineInstance, nil
}
