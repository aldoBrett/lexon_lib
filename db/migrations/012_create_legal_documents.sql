CREATE TABLE
    IF NOT EXISTS lexon.legal_documents (
        id TEXT PRIMARY KEY NOT NULL,
        machine_event_id TEXT NOT NULL REFERENCES lexon.machine_events (id) ON DELETE CASCADE,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );