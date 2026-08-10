package repository

import (
	"amox/lex_engine_lib/domain"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MachineStateStagesRepository interface {
	GetMachineStateStages() ([]domain.MachineStateStage, error)
	GetMachineStateStageByID(id string) (*domain.MachineStateStage, error)
}

type MachineStateStagesRepositoryHandler struct {
	Ctx  context.Context
	pool *pgxpool.Pool
}

func NewMachineStateStagesRepositoryHandler(ctx context.Context, pool *pgxpool.Pool) *MachineStateStagesRepositoryHandler {
	return &MachineStateStagesRepositoryHandler{
		Ctx:  ctx,
		pool: pool,
	}
}

func (h *MachineStateStagesRepositoryHandler) GetMachineStateStages() ([]domain.MachineStateStage, error) {
	query := `SELECT id, code, name, description, content, stage_order FROM lexon.machine_state_stages ORDER BY stage_order`
	rows, err := h.pool.Query(h.Ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stages []domain.MachineStateStage
	for rows.Next() {
		var stage domain.MachineStateStage
		if err := rows.Scan(&stage.ID, &stage.Code, &stage.Name, &stage.Description, &stage.Content, &stage.StageOrder); err != nil {
			return nil, err
		}
		stages = append(stages, stage)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return stages, nil
}

func (h *MachineStateStagesRepositoryHandler) GetMachineStateStageByID(id string) (*domain.MachineStateStage, error) {
	query := `SELECT id, code, name, description, content, stage_order FROM lexon.machine_state_stages WHERE id = $1`
	var stage domain.MachineStateStage
	err := h.pool.QueryRow(h.Ctx, query, id).Scan(&stage.ID, &stage.Code, &stage.Name, &stage.Description, &stage.Content, &stage.StageOrder)
	if err != nil {
		return nil, err
	}
	return &stage, nil
}
