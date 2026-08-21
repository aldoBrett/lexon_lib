package repository

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TestLegalDocumentsRepository(t *testing.T) {
	dsn := "host=db user=admin password=secretpassword dbname=lexon_engine_test_db port=5432 sslmode=disable"
	pool, _ := pgxpool.New(context.Background(), dsn)

	repository := NewLegalDocumentsRepositoryHandler(context.Background(), pool)

	t.Run("Get all legal documents with pagination", func(t *testing.T) {
		offset := 0
		limit := 10
		documents, err := repository.GetLegalDocuments(offset, limit)
		if err != nil {
			t.Fatalf("Failed to get legal documents: %v", err)
		}

		if len(documents) == 0 {
			t.Fatalf("Expected to find legal documents, but got none")
		}
	})

	t.Run("Count legal documents", func(t *testing.T) {
		count, err := repository.CountLegalDocuments()
		if err != nil {
			t.Fatalf("Failed to count legal documents: %v", err)
		}

		if count <= 0 {
			t.Fatalf("Expected to find at least one legal document, but got count: %d", count)
		}
	})

	t.Run("Get legal document by event ID", func(t *testing.T) {
		eventID := "CIV.ORD.E002"
		document, err := repository.GetLegalDocumentByMachineEventID(eventID)
		if err != nil {
			t.Fatalf("Failed to get legal document by event ID: %v", err)
		}

		if document.ID == "" {
			t.Fatalf("Expected to find legal document for event ID %s, but got none", eventID)
		}
	})
}
