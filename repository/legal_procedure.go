package repository

import (
	"amox/lex_engine_lib/engine"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type LegalProcedureRepository interface {
	SaveLegalProcedure(legalProcedure engine.LegalProcedure) error
	GetLegalProcedureByID(id string) (*engine.LegalProcedure, error)
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
func (h *LegalProcedureRepositoryHandler) SaveLegalProcedure(legalProcedure engine.LegalProcedure) error {
	query := `SELECT id FROM lexon.legal_procedures WHERE id = $1`
	var id string
	err := h.pool.QueryRow(h.Ctx, query, legalProcedure.ID).Scan(&id)
	if err != nil {
		// If the legal procedure doesn't exist, create a new one
		insertQuery := `INSERT INTO lexon.legal_procedures (id, label, description) VALUES ($1, $2, $3)`
		_, err = h.pool.Exec(h.Ctx, insertQuery, legalProcedure.ID, legalProcedure.Label, legalProcedure.Description)
		if err != nil {
			return err
		}
	} else {
		// If the legal procedure exists, update its label and description
		updateQuery := `UPDATE lexon.legal_procedures SET label = $1, description = $2 WHERE id = $3`
		_, err = h.pool.Exec(h.Ctx, updateQuery, legalProcedure.Label, legalProcedure.Description, legalProcedure.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

// This function is used to get a LegalProcedure from the database by its
// id.
func (h *LegalProcedureRepositoryHandler) GetLegalProcedureByID(id string) (*engine.LegalProcedure, error) {
	query := `SELECT id, label, description FROM lexon.legal_procedures WHERE id = $1`
	var legalProcedure engine.LegalProcedure
	err := h.pool.QueryRow(h.Ctx, query, id).Scan(&legalProcedure.ID, &legalProcedure.Label, &legalProcedure.Description)
	if err != nil {
		return nil, err
	}

	return &legalProcedure, nil
}
