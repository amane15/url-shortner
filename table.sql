CREATE TABLE IF NOT EXISTS "urls" (
    original_url TEXT PRIMARY KEY NOT NULL,
    shortened_url TEXT UNIQUE NOT NULL,
    clicks INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE index idx_shortened ON urls (shortened_url); 
