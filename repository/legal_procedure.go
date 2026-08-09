package repository

import (
	"amox/lex_engine_lib/domain"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type LegalProcedureRepository interface {
	SaveLegalProcedure(legalProcedure domain.LegalProcedure) (*domain.LegalProcedure, error)
	GetLegalProcedureByID(id string) (*domain.LegalProcedure, error)
	GetLegalProcedures(offset, limit int) ([]domain.LegalProcedure, error)
	CountLegalProcedures() (int, error)
}

type LegalProcedureRepositoryHandler struct {
	Ctx  context.Context
	pool *pgxpool.Pool
}

func NewLegalProcedureRepositoryHandler(ctx context.Context, pool *pgxpool.Pool) *LegalProcedureRepositoryHandler {
	return &LegalProcedureRepositoryHandler{
		Ctx:  ctx,
		pool: pool,
	}
}

// This function will look for the ID of the legal procedure, if it doesn't
// exists, it will create a new one. If it exists, it will update the label and
// description of the legal procedure.
func (h *LegalProcedureRepositoryHandler) SaveLegalProcedure(legalProcedure domain.LegalProcedure) (*domain.LegalProcedure, error) {
	query := `SELECT id FROM lexon.legal_procedures WHERE id = $1`
	var id string
	err := h.pool.QueryRow(h.Ctx, query, legalProcedure.ID).Scan(&id)

	var saved domain.LegalProcedure
	if err != nil {
		// If the legal procedure doesn't exist, create a new one
		insertQuery := `INSERT INTO lexon.legal_procedures (id, label, description) VALUES ($1, $2, $3) RETURNING id, label, description`
		err = h.pool.QueryRow(h.Ctx, insertQuery, legalProcedure.ID, legalProcedure.Label, legalProcedure.Description).Scan(&saved.ID, &saved.Label, &saved.Description)
		if err != nil {
			return nil, err
		}
	} else {
		// If the legal procedure exists, update its label and description
		updateQuery := `UPDATE lexon.legal_procedures SET label = $1, description = $2 WHERE id = $3 RETURNING id, label, description`
		err = h.pool.QueryRow(h.Ctx, updateQuery, legalProcedure.Label, legalProcedure.Description, legalProcedure.ID).Scan(&saved.ID, &saved.Label, &saved.Description)
		if err != nil {
			return nil, err
		}
	}

	return &saved, nil
}

// This function is used to get a LegalProcedure from the database by its
// id.
func (h *LegalProcedureRepositoryHandler) GetLegalProcedureByID(id string) (*domain.LegalProcedure, error) {
	query := `SELECT id, label, description FROM lexon.legal_procedures WHERE id = $1`
	var legalProcedure domain.LegalProcedure
	err := h.pool.QueryRow(h.Ctx, query, id).Scan(&legalProcedure.ID, &legalProcedure.Label, &legalProcedure.Description)
	if err != nil {
		return nil, err
	}

	return &legalProcedure, nil
}

func (h *LegalProcedureRepositoryHandler) GetLegalProcedures(offset, limit int) ([]domain.LegalProcedure, error) {
	query := `SELECT id, label, description FROM lexon.legal_procedures ORDER BY id LIMIT $1 OFFSET $2`
	rows, err := h.pool.Query(h.Ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var legalProcedures []domain.LegalProcedure
	for rows.Next() {
		var legalProcedure domain.LegalProcedure
		err := rows.Scan(&legalProcedure.ID, &legalProcedure.Label, &legalProcedure.Description)
		if err != nil {
			return nil, err
		}
		legalProcedures = append(legalProcedures, legalProcedure)
	}

	return legalProcedures, nil
}

func (h *LegalProcedureRepositoryHandler) CountLegalProcedures() (int, error) {
	query := `SELECT COUNT(*) FROM lexon.legal_procedures`
	var count int
	err := h.pool.QueryRow(h.Ctx, query).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
