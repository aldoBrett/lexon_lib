CREATE TABLE
    IF NOT EXISTS lexon.machine_state_transitions (
        id TEXT PRIMARY KEY NOT NULL,
        source_state_id TEXT NOT NULL REFERENCES lexon.machine_states (id) ON DELETE CASCADE,
        target_state_id TEXT NOT NULL REFERENCES lexon.machine_states (id) ON DELETE CASCADE,
        condition TEXT NOT NULL,
        actions TEXT NOT NULL,
        risk TEXT NOT NULL,
        note TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );