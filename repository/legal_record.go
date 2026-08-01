package repository

import (
	"amox/lex_engine_lib/domain"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LegalRecordRepository interface {
	SaveLegalRecord(legalRecord domain.LegalRecord) (*domain.LegalRecord, error)
	// GetLegalRecordByID(id string) (*domain.LegalRecord, error)
}

type LegalRecordRepositoryHandler struct {
	Ctx  context.Context
	pool *pgxpool.Pool
}

func NewLegalRecordRepositoryHandler(ctx context.Context, pool *pgxpool.Pool) *LegalRecordRepositoryHandler {
	return &LegalRecordRepositoryHandler{
		Ctx:  ctx,
		pool: pool,
	}
}

// The LegalRecord instance can come withouth id (when it's first created), after that, when the user
// is editing the legal record it will come with and id. So we need to check if it has an id, if it doesn't
// have one, we're going to create a new LegalRecord. When the id arrives, we're going to update the existing
// LegalRecord with the new data.
func (h *LegalRecordRepositoryHandler) SaveLegalRecord(legalRecord domain.LegalRecord) (*domain.LegalRecord, error) {
	if legalRecord.ID == "" {
		legalRecord.ID = uuid.New().String()
		insertQuery := `INSERT INTO lexon.legal_records (id, legal_procedure_id, trial_kind, record_number, actor, defendant, defendant_address) VALUES ($1, $2, $3, $4, $5, $6, $7)`
		_, err := h.pool.Exec(h.Ctx, insertQuery,
			legalRecord.ID,
			legalRecord.LegalProcedureID,
			legalRecord.TrialKind,
			legalRecord.RecordNumber,
			legalRecord.Actor,
			legalRecord.Defendant,
			legalRecord.DefendantAddress,
		)
		if err != nil {
			return nil, err
		}
		return &legalRecord, nil
	}

	updateQuery := `UPDATE lexon.legal_records SET legal_procedure_id = $1, trial_kind = $2, record_number = $3, actor = $4, defendant = $5, defendant_address = $6 WHERE id = $7`
	_, err := h.pool.Exec(h.Ctx, updateQuery,
		legalRecord.LegalProcedureID,
		legalRecord.TrialKind,
		legalRecord.RecordNumber,
		legalRecord.Actor,
		legalRecord.Defendant,
		legalRecord.DefendantAddress,
		legalRecord.ID,
	)
	if err != nil {
		return nil, err
	}

	return &legalRecord, nil
}
