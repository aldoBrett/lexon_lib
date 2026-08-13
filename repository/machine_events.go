package repository

import (
	"amox/lex_engine_lib/domain"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MachineEventsRepository interface {
	GetMachineEventsByIDs(ids []string) ([]*domain.MachineEvent, error)
}

type MachineEventsRepositoryHandler struct {
	pool *pgxpool.Pool
	Ctx  context.Context
}

func NewMachineEventsRepositoryHandler(ctx context.Context, pool *pgxpool.Pool) *MachineEventsRepositoryHandler {
	return &MachineEventsRepositoryHandler{
		pool: pool,
		Ctx:  ctx,
	}
}

func (h *MachineEventsRepositoryHandler) GetMachineEventsByIDs(ids []string) ([]*domain.MachineEvent, error) {
	events := []*domain.MachineEvent{}
	query := `SELECT id, name, description, kind FROM lexon.machine_events WHERE id = ANY($1)`
	rows, err := h.pool.Query(h.Ctx, query, ids)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		event := &domain.MachineEvent{}
		err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.Kind,
		)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}
