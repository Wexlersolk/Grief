CREATE TABLE boards (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    user_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
