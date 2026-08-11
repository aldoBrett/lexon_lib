package repository

import (
	"amox/lex_engine_lib/domain"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type LegalActionsRepository interface {
	GetLegalActions() ([]domain.LegalAction, error)
	GetLegalActionByID(id string) (*domain.LegalAction, error)
}

type LegalActionsRepositoryHandler struct {
	Ctx  context.Context
	pool *pgxpool.Pool
}

func NewLegalActionsRepositoryHandler(ctx context.Context, pool *pgxpool.Pool) *LegalActionsRepositoryHandler {
	return &LegalActionsRepositoryHandler{
		Ctx:  ctx,
		pool: pool,
	}
}

func (h *LegalActionsRepositoryHandler) GetLegalActions() ([]domain.LegalAction, error) {
	query := `SELECT id, category, sub_category, action_name, via FROM lexon.legal_actions`
	rows, err := h.pool.Query(h.Ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var actions []domain.LegalAction
	for rows.Next() {
		var action domain.LegalAction
		if err := rows.Scan(&action.ID, &action.Category, &action.SubCategory, &action.ActionName, &action.Via); err != nil {
			return nil, err
		}
		actions = append(actions, action)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return actions, nil
}

func (h *LegalActionsRepositoryHandler) GetLegalActionByID(id string) (*domain.LegalAction, error) {
	query := `SELECT id, category, sub_category, action_name, via FROM lexon.legal_actions WHERE id = $1`
	row := h.pool.QueryRow(h.Ctx, query, id)

	var action domain.LegalAction
	if err := row.Scan(&action.ID, &action.Category, &action.SubCategory, &action.ActionName, &action.Via); err != nil {
		return nil, err
	}

	return &action, nil
}
