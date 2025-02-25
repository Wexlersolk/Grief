CREATE TABLE board_events (
    id BIGSERIAL PRIMARY KEY,
    board_id INTEGER NOT NULL REFERENCES boards(id),
    event_type TEXT NOT NULL, -- e.g., 'CardMoved', 'CardUpdated'
    event_data JSONB NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
