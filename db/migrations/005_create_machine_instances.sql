CREATE TABLE
    IF NOT EXISTS lexon.machine_instances (
        id TEXT PRIMARY KEY NOT NULL,
        current_state_id TEXT NOT NULL REFERENCES lexon.machine_states (id) ON DELETE CASCADE,
        legal_procedure_id uuid NOT NULL REFERENCES lexon.legal_procedures (id) ON DELETE CASCADE,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
