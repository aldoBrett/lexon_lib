package repository

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TestLegalDocumentFieldsRepository(t *testing.T) {
	dsn := "host=db user=admin password=secretpassword dbname=lexon_engine_test_db port=5432 sslmode=disable"
	pool, _ := pgxpool.New(context.Background(), dsn)

	repository := NewLegalDocumentFieldsRepositoryHandler(context.Background(), pool)

	t.Run("Get all legal document fields with pagination", func(t *testing.T) {
		offset := 0
		limit := 10
		fields, err := repository.GetLegalDocumentFields(offset, limit)
		if err != nil {
			t.Fatalf("Failed to get legal document fields: %v", err)
		}

		if len(fields) == 0 {
			t.Fatalf("Expected to find legal document fields, but got none")
		}
	})

	t.Run("Count legal document fields", func(t *testing.T) {
		count, err := repository.CountLegalDocumentFields()
		if err != nil {
			t.Fatalf("Failed to count legal document fields: %v", err)
		}

		if count <= 0 {
			t.Fatalf("Expected to find at least one legal document field, but got count: %d", count)
		}
	})

	t.Run("Get legal document fields by document ID", func(t *testing.T) {
		documentID := "DOC.0001"
		fields, err := repository.GetLegalDocumentFieldsByDocumentID(documentID)
		if err != nil {
			t.Fatalf("Failed to get legal document fields by document ID: %v", err)
		}

		if len(fields) == 0 {
			t.Fatalf("Expected to find legal document fields for document ID %s, but got none", documentID)
		}
	})
}
