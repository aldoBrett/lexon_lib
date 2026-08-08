package repository

import (
	"amox/lex_engine_lib/domain"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MachineStateTransitionsHistoryRepository interface {
	CreateMachineStateTransitionHistory(machineInstanceID string, transitionID string) (*domain.MachineStateTransitionHistory, error)
}

type MachineStateTransitionsHistoryRepositoryHandler struct {
	pool *pgxpool.Pool
	Ctx  context.Context
}

func NewMachineStateTransitionsHistoryRepositoryHandler(ctx context.Context, pool *pgxpool.Pool) *MachineStateTransitionsHistoryRepositoryHandler {
	return &MachineStateTransitionsHistoryRepositoryHandler{
		pool: pool,
		Ctx:  ctx,
	}
}

// Record that a MachineInstance moved through a MachineStateTransition. This is an append-only
// audit trail: one row per transition applied to a machine instance.
func (h *MachineStateTransitionsHistoryRepositoryHandler) CreateMachineStateTransitionHistory(machineInstanceID string, transitionID string) (*domain.MachineStateTransitionHistory, error) {
	historyID := uuid.New().String()
	history := &domain.MachineStateTransitionHistory{}
	insertQuery := `INSERT INTO lexon.machine_state_transitions_history (id, machine_instance_id, transition_id, created_at)
		VALUES ($1, $2, $3, NOW())
		RETURNING id, machine_instance_id, transition_id, created_at`
	err := h.pool.QueryRow(h.Ctx, insertQuery, historyID, machineInstanceID, transitionID).Scan(
		&history.ID,
		&history.MachineInstanceID,
		&history.TransitionID,
		&history.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return history, nil
}
