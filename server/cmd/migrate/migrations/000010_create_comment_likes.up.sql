CREATE TABLE IF NOT EXISTS comment_likes (
    author_id uuid NOT NULL,
    comment_id uuid NOT NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    PRIMARY KEY (author_id, comment_id),
    FOREIGN KEY (author_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (comment_id) REFERENCES comments (id) ON DELETE CASCADE
);