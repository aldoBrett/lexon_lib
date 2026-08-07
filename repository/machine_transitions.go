package repository

import (
	"amox/lex_engine_lib/domain"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MachineTransitionsRepository interface {
	GetMachineTransitionByID(transitionID string) (*domain.MachineStateTransition, error)
}

type MachineTransitionsRepositoryHandler struct {
	pool *pgxpool.Pool
	Ctx  context.Context
}

func NewMachineTransitionsRepositoryHandler(ctx context.Context, pool *pgxpool.Pool) *MachineTransitionsRepositoryHandler {
	return &MachineTransitionsRepositoryHandler{
		pool: pool,
		Ctx:  ctx,
	}
}

func (h *MachineTransitionsRepositoryHandler) GetMachineTransitionByID(transitionID string) (*domain.MachineStateTransition, error) {
	machineTransition := &domain.MachineStateTransition{}
	query := `SELECT id, source_state_id, target_state_id, condition, actions, risk, note, created_at, updated_at FROM lexon.machine_state_transitions WHERE id = $1`
	err := h.pool.QueryRow(h.Ctx, query, transitionID).Scan(
		&machineTransition.ID,
		&machineTransition.SourceStateID,
		&machineTransition.TargetStateID,
		&machineTransition.Condition,
		&machineTransition.Actions,
		&machineTransition.Risk,
		&machineTransition.Note,
		&machineTransition.CreatedAt,
		&machineTransition.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return machineTransition, nil
}
