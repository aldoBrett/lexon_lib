package repository

import (
	"amox/lex_engine_lib/domain"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LegalClaimsRepository interface {
	SaveLegalClaim(legalClaim domain.LegalClaim) (*domain.LegalClaim, error)
	GetLegalClaimsByLegalRecordID(legalRecordId string) ([]domain.LegalClaim, error)
}

type LegalClaimsRepositoryHandler struct {
	Ctx  context.Context
	pool *pgxpool.Pool
}

func NewLegalClaimsRepositoryHandler(ctx context.Context, pool *pgxpool.Pool) *LegalClaimsRepositoryHandler {
	return &LegalClaimsRepositoryHandler{
		Ctx:  ctx,
		pool: pool,
	}
}

func (h *LegalClaimsRepositoryHandler) SaveLegalClaim(legalClaim domain.LegalClaim) (*domain.LegalClaim, error) {
	if legalClaim.ID == "" {
		legalClaim.ID = uuid.New().String()
		insertQuery := `INSERT INTO lexon.legal_claims (id, legal_record_id, description) VALUES ($1, $2, $3)`
		_, err := h.pool.Exec(h.Ctx, insertQuery,
			legalClaim.ID,
			legalClaim.LegalRecordID,
			legalClaim.Description,
		)
		if err != nil {
			return nil, err
		}
		return &legalClaim, nil
	}

	updateQuery := `UPDATE lexon.legal_claims SET legal_record_id = $1, description = $2 WHERE id = $3`
	_, err := h.pool.Exec(h.Ctx, updateQuery,
		legalClaim.LegalRecordID,
		legalClaim.Description,
		legalClaim.ID,
	)
	if err != nil {
		return nil, err
	}
	return &legalClaim, nil
}

func (h *LegalClaimsRepositoryHandler) GetLegalClaimsByLegalRecordID(legalRecordId string) ([]domain.LegalClaim, error) {
	query := `SELECT id, legal_record_id, description FROM lexon.legal_claims WHERE legal_record_id = $1`
	rows, err := h.pool.Query(h.Ctx, query, legalRecordId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var legalClaims []domain.LegalClaim
	for rows.Next() {
		var legalClaim domain.LegalClaim
		err := rows.Scan(&legalClaim.ID, &legalClaim.LegalRecordID, &legalClaim.Description)
		if err != nil {
			return nil, err
		}
		legalClaims = append(legalClaims, legalClaim)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return legalClaims, nil
}
