CREATE TABLE
    IF NOT EXISTS lexon.legal_records (
        id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid (),
        legal_procedure_id uuid NOT NULL REFERENCES lexon.legal_procedures (id) ON DELETE CASCADE,
        -- TODO: I think this has to be not nullable.
        -- TODO: We need the catalogs
        -- trial_kind TEXT NOT NULL,
        trial_kind TEXT,
        -- record_number TEXT NOT NULL,
        record_number TEXT,
        actor TEXT NOT NULL,
        defendant TEXT NOT NULL,
        -- defendant_address TEXT NOT NULL,
        defendant_address TEXT,
        -- court TEXT NOT NULL,
        -- jurisdiction TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );