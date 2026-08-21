package seeds

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// LegalDocument mirrors a row of lexon.legal_documents.
type LegalDocument struct {
	ID             string
	Name           string
	Description    string
	MachineEventID string
}

// LegalDocuments is the seed data for the civil legal documents catalog.
var LegalDocuments = []LegalDocument{
	{ID: "DOC.0001", Name: "Demanda", Description: "", MachineEventID: "CIV.ORD.E002"},
	{ID: "DOC.0002", Name: "Demanda ingresada", Description: "Es la demanda con el sello de ingresado", MachineEventID: "CIV.ORD.E005"},
	{ID: "DOC.0003", Name: "Demanda admitida", Description: "Demanda con acuerdo de admisión", MachineEventID: "CIV.ORD.E010"},
}

// SeedLegalDocuments upserts LegalDocuments into lexon.legal_documents.
func SeedLegalDocuments(ctx context.Context, pool *pgxpool.Pool) error {
	batch := &pgx.Batch{}
	for _, d := range LegalDocuments {
		batch.Queue(
			`INSERT INTO lexon.legal_documents (id, name, description, machine_event_id)
			 VALUES ($1, $2, $3, $4)
			 ON CONFLICT (id) DO UPDATE SET
			   name = EXCLUDED.name,
			   description = EXCLUDED.description,
			   machine_event_id = EXCLUDED.machine_event_id`,
			d.ID, d.Name, d.Description, d.MachineEventID,
		)
	}

	br := pool.SendBatch(ctx, batch)
	defer br.Close()

	for range LegalDocuments {
		if _, err := br.Exec(); err != nil {
			return fmt.Errorf("seeds: insert legal document: %w", err)
		}
	}

	return nil
}
