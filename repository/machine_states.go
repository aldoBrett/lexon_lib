package repository

import (
	"amox/lex_engine_lib/domain"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MachineStatesRepository interface {
	GetMachineStates(offset, limit int) ([]domain.MachineState, error)
	GetMachineStateByID(id string) (*domain.MachineState, error)
	CountMachineStates() (int, error)
}

type MachineStatesRepositoryHandler struct {
	Ctx  context.Context
	pool *pgxpool.Pool
}

func NewMachineStatesRepositoryHandler(ctx context.Context, pool *pgxpool.Pool) *MachineStatesRepositoryHandler {
	return &MachineStatesRepositoryHandler{
		Ctx:  ctx,
		pool: pool,
	}
}

func (h *MachineStatesRepositoryHandler) GetMachineStates(offset, limit int) ([]domain.MachineState, error) {
	query := `SELECT id, name, description, stage_id, kind FROM lexon.machine_states ORDER BY id LIMIT $1 OFFSET $2`
	rows, err := h.pool.Query(h.Ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var states []domain.MachineState
	for rows.Next() {
		var state domain.MachineState
		if err := rows.Scan(&state.ID, &state.Name, &state.Description, &state.StageID, &state.Kind); err != nil {
			return nil, err
		}
		states = append(states, state)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return states, nil
}

func (h *MachineStatesRepositoryHandler) GetMachineStateByID(id string) (*domain.MachineState, error) {
	query := `SELECT id, name, description, stage_id, kind FROM lexon.machine_states WHERE id = $1`
	var state domain.MachineState
	err := h.pool.QueryRow(h.Ctx, query, id).Scan(&state.ID, &state.Name, &state.Description, &state.StageID, &state.Kind)
	if err != nil {
		return nil, err
	}
	return &state, nil
}

func (h *MachineStatesRepositoryHandler) CountMachineStates() (int, error) {
	query := `SELECT COUNT(*) FROM lexon.machine_states`
	var count int
	err := h.pool.QueryRow(h.Ctx, query).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
