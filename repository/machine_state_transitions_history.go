package repository

import (
	"amox/lex_engine_lib/domain"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MachineStateTransitionsHistoryRepository interface {
	CreateMachineStateTransitionHistory(machineInstanceID string, transitionID string) (*domain.MachineStateTransitionHistory, error)
	GetMachineStateTransitionsHistoryByMachineStateStageID(machineInstanceID string, machineStateStageID string) ([]*domain.MachineStateTransitionHistory, error)
	GetMachineStateTransitionsHistoryByMachineInstanceID(machineInstanceID string) ([]*domain.MachineStateTransitionHistory, error)
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

func (h *MachineStateTransitionsHistoryRepositoryHandler) GetMachineStateTransitionsHistoryByMachineStateStageID(machineInstanceID string, machineStateStageID string) ([]*domain.MachineStateTransitionHistory, error) {
	query := `SELECT h.id, h.machine_instance_id, h.transition_id, h.created_at
		FROM lexon.machine_state_transitions_history h
		JOIN lexon.machine_state_transitions t ON t.id = h.transition_id
		JOIN lexon.machine_states s ON s.id = t.target_state_id
		WHERE h.machine_instance_id = $1 AND s.stage_id = $2
		ORDER BY h.created_at ASC`
	rows, err := h.pool.Query(h.Ctx, query, machineInstanceID, machineStateStageID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []*domain.MachineStateTransitionHistory
	for rows.Next() {
		h := &domain.MachineStateTransitionHistory{}
		err := rows.Scan(
			&h.ID,
			&h.MachineInstanceID,
			&h.TransitionID,
			&h.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		history = append(history, h)
	}

	return history, nil
}

func (h *MachineStateTransitionsHistoryRepositoryHandler) GetMachineStateTransitionsHistoryByMachineInstanceID(machineInstanceID string) ([]*domain.MachineStateTransitionHistory, error) {
	query := `SELECT id, machine_instance_id, transition_id, created_at
		FROM lexon.machine_state_transitions_history
		WHERE machine_instance_id = $1
		ORDER BY created_at ASC`
	rows, err := h.pool.Query(h.Ctx, query, machineInstanceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []*domain.MachineStateTransitionHistory
	for rows.Next() {
		h := &domain.MachineStateTransitionHistory{}
		err := rows.Scan(
			&h.ID,
			&h.MachineInstanceID,
			&h.TransitionID,
			&h.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		history = append(history, h)
	}

	return history, nil
}
