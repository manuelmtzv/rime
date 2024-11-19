CREATE TABLE IF NOT EXISTS writing_likes (
    author_id uuid NOT NULL,
    writing_id uuid NOT NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    PRIMARY KEY (author_id, writing_id),
    FOREIGN KEY (author_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (writing_id) REFERENCES comments (id) ON DELETE CASCADE
);