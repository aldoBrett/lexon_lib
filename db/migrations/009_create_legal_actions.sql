CREATE TABLE
    IF NOT EXISTS lexon.legal_actions (
        id TEXT PRIMARY KEY NOT NULL,
        category TEXT NOT NULL,
        sub_category TEXT NOT NULL,
        action_name TEXT NOT NULL,
        via TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );