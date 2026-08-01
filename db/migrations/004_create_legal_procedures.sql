CREATE TABLE
    IF NOT EXISTS lexon.legal_procedures (
        id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid (),
        label TEXT NOT NULL,
        description TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );