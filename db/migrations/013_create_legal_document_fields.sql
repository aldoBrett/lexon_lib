CREATE TABLE
    IF NOT EXISTS lexon.legal_document_fields (
        id TEXT PRIMARY KEY NOT NULL,
        legal_document_id TEXT NOT NULL REFERENCES lexon.legal_documents (id) ON DELETE CASCADE,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );