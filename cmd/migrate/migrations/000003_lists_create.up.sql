CREATE TABLE lists (
    id SERIAL PRIMARY KEY,
    board_id INTEGER REFERENCES boards(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    position INTEGER NOT NULL
);

