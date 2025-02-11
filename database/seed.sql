CREATE TABLE IF NOT EXISTS articles (
    route TEXT PRIMARY KEY,
    title TEXT,
    description TEXT,
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
