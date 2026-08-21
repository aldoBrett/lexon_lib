package seeds

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// LegalDocumentField mirrors a row of lexon.legal_document_fields.
type LegalDocumentField struct {
	ID              string
	Name            string
	Description     string
	LegalDocumentID string
}

// LegalDocumentFields is the seed data for the civil legal document fields catalog.
var LegalDocumentFields = []LegalDocumentField{
	{ID: "DOC.0001.F001", Name: "Actor", Description: "Nombre completo de la persona que demanda", LegalDocumentID: "DOC.0001"},
	{ID: "DOC.0001.F002", Name: "Demandado", Description: "Nombre completo de la persona que ha sido demandada", LegalDocumentID: "DOC.0001"},
	{ID: "DOC.0001.F003", Name: "Domicilio actor", Description: "Domicilio de la persona que demanda", LegalDocumentID: "DOC.0001"},
	{ID: "DOC.0001.F004", Name: "Domicilio demandado", Description: "Domicilio de la persona que ha sido demandada", LegalDocumentID: "DOC.0001"},
	{ID: "DOC.0001.F005", Name: "Abogado actor", Description: "Nombre completo del abogado de la persona que demanda", LegalDocumentID: "DOC.0001"},
	{ID: "DOC.0001.F006", Name: "Autoridad a la que va dirigida", Description: "", LegalDocumentID: "DOC.0001"},
	{ID: "DOC.0001.F007", Name: "Acción principal", Description: "Lo que estás demandando", LegalDocumentID: "DOC.0001"},
	{ID: "DOC.0002.F001", Name: "Fecha ingreso", Description: "Fecha en que la demanda fue ingresada / aceptada", LegalDocumentID: "DOC.0002"},
	{ID: "DOC.0002.F002", Name: "Folio", Description: "Folio asignado a la demanda", LegalDocumentID: "DOC.0002"},
	{ID: "DOC.0002.F003", Name: "Autoridad que recibe la demanda", Description: "", LegalDocumentID: "DOC.0002"},
	{ID: "DOC.0003.F001", Name: "Datos de emplazamiento", Description: "Información del desahogo del emplazamiento", LegalDocumentID: "DOC.0003"},
}

// SeedLegalDocumentFields upserts LegalDocumentFields into lexon.legal_document_fields.
func SeedLegalDocumentFields(ctx context.Context, pool *pgxpool.Pool) error {
	batch := &pgx.Batch{}
	for _, f := range LegalDocumentFields {
		batch.Queue(
			`INSERT INTO lexon.legal_document_fields (id, name, description, legal_document_id)
			 VALUES ($1, $2, $3, $4)
			 ON CONFLICT (id) DO UPDATE SET
			   name = EXCLUDED.name,
			   description = EXCLUDED.description,
			   legal_document_id = EXCLUDED.legal_document_id`,
			f.ID, f.Name, f.Description, f.LegalDocumentID,
		)
	}

	br := pool.SendBatch(ctx, batch)
	defer br.Close()

	for range LegalDocumentFields {
		if _, err := br.Exec(); err != nil {
			return fmt.Errorf("seeds: insert legal document field: %w", err)
		}
	}

	return nil
}
