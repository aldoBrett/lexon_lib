package repository

import (
	"amox/lex_engine_lib/domain"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type LegalDocumentsRepository interface {
	GetLegalDocuments(offset, limit int) ([]domain.LegalDocument, error)
	CountLegalDocuments() (int, error)
	GetLegalDocumentByMachineEventID(eventID string) (domain.LegalDocument, error)
}

type LegalDocumentsRepositoryHandler struct {
	Ctx  context.Context
	pool *pgxpool.Pool
}

func NewLegalDocumentsRepositoryHandler(ctx context.Context, pool *pgxpool.Pool) *LegalDocumentsRepositoryHandler {
	return &LegalDocumentsRepositoryHandler{
		Ctx:  ctx,
		pool: pool,
	}
}

func (h *LegalDocumentsRepositoryHandler) GetLegalDocuments(offset, limit int) ([]domain.LegalDocument, error) {
	query := `SELECT id, name, description FROM lexon.legal_documents ORDER BY id LIMIT $1 OFFSET $2`
	rows, err := h.pool.Query(h.Ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var documents []domain.LegalDocument
	for rows.Next() {
		var document domain.LegalDocument
		if err := rows.Scan(&document.ID, &document.Name, &document.Description); err != nil {
			return nil, err
		}
		documents = append(documents, document)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return documents, nil
}

func (h *LegalDocumentsRepositoryHandler) CountLegalDocuments() (int, error) {
	query := `SELECT COUNT(*) FROM lexon.legal_documents`
	var count int
	err := h.pool.QueryRow(h.Ctx, query).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (h *LegalDocumentsRepositoryHandler) GetLegalDocumentByMachineEventID(machineEventID string) (domain.LegalDocument, error) {
	query := `SELECT id, name, description FROM lexon.legal_documents WHERE machine_event_id = $1`
	var document domain.LegalDocument
	err := h.pool.QueryRow(h.Ctx, query, machineEventID).Scan(&document.ID, &document.Name, &document.Description)
	if err != nil {
		return domain.LegalDocument{}, err
	}

	return document, nil
}
