CREATE TABLE
    IF NOT EXISTS lexon.machine_states (
        id TEXT PRIMARY KEY NOT NULL,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        stage_id TEXT NOT NULL REFERENCES lexon.machine_state_stages (id),
        kind TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

-- TODO: look for the naming convention with Kain