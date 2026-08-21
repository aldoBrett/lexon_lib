package repository

import (
	"amox/lex_engine_lib/domain"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type LegalDocumentFieldsRepository interface {
	GetLegalDocumentFields(offset, limit int) ([]domain.LegalDocumentField, error)
	CountLegalDocumentFields() (int, error)
	GetLegalDocumentFieldsByDocumentID(documentID string) ([]domain.LegalDocumentField, error)
}

type LegalDocumentFieldsRepositoryHandler struct {
	Ctx  context.Context
	pool *pgxpool.Pool
}

func NewLegalDocumentFieldsRepositoryHandler(ctx context.Context, pool *pgxpool.Pool) *LegalDocumentFieldsRepositoryHandler {
	return &LegalDocumentFieldsRepositoryHandler{
		Ctx:  ctx,
		pool: pool,
	}
}

func (h *LegalDocumentFieldsRepositoryHandler) GetLegalDocumentFields(offset, limit int) ([]domain.LegalDocumentField, error) {
	query := `SELECT id, name, description FROM lexon.legal_document_fields ORDER BY id LIMIT $1 OFFSET $2`
	rows, err := h.pool.Query(h.Ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fields []domain.LegalDocumentField
	for rows.Next() {
		var field domain.LegalDocumentField
		if err := rows.Scan(&field.ID, &field.Name, &field.Description); err != nil {
			return nil, err
		}
		fields = append(fields, field)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return fields, nil
}

func (h *LegalDocumentFieldsRepositoryHandler) CountLegalDocumentFields() (int, error) {
	query := `SELECT COUNT(*) FROM lexon.legal_document_fields`
	var count int
	err := h.pool.QueryRow(h.Ctx, query).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (h *LegalDocumentFieldsRepositoryHandler) GetLegalDocumentFieldsByDocumentID(documentID string) ([]domain.LegalDocumentField, error) {
	query := `SELECT id, name, description FROM lexon.legal_document_fields WHERE legal_document_id = $1`
	rows, err := h.pool.Query(h.Ctx, query, documentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fields []domain.LegalDocumentField
	for rows.Next() {
		var field domain.LegalDocumentField
		if err := rows.Scan(&field.ID, &field.Name, &field.Description); err != nil {
			return nil, err
		}
		fields = append(fields, field)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return fields, nil
}
