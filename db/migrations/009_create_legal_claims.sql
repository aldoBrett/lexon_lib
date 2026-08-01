CREATE TABLE
    IF NOT EXISTS lexon.legal_claims (
        id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid (),
        legal_record_id uuid NOT NULL REFERENCES lexon.legal_records (id) ON DELETE CASCADE,
        description TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );