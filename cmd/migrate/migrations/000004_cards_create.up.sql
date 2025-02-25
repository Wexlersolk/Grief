CREATE TABLE cards (
    id SERIAL PRIMARY KEY,
    list_id INTEGER REFERENCES lists(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    description TEXT,
    color TEXT DEFAULT 'white',
    done BOOLEAN DEFAULT FALSE,
    position INTEGER NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW()
);
