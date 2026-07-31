CREATE TABLE
    IF NOT EXISTS lexon.machine_state_transitions_history (
        id TEXT PRIMARY KEY NOT NULL,
        machine_instance_id TEXT NOT NULL REFERENCES lexon.machine_instances (id) ON DELETE CASCADE,
        from_state_id TEXT NOT NULL REFERENCES lexon.machine_states (id) ON DELETE CASCADE,
        to_state_id TEXT NOT NULL REFERENCES lexon.machine_states (id) ON DELETE CASCADE,
        event_id TEXT NOT NULL REFERENCES lexon.machine_events (id) ON DELETE CASCADE,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP

        -- TODO: Add a column for the user who triggered the transition, if applicable
        -- TODO: Add a column for the reason for the transition, if applicable
        -- TODO: Add a column for any additional metadata related to the transition, if applicable
    );