CREATE TABLE IF NOT EXISTS like_comment (
    like_id uuid NOT NULL,
    comment_id uuid NOT NULL,
    PRIMARY KEY (like_id, comment_id),
    FOREIGN KEY (like_id) REFERENCES likes (id) ON DELETE CASCADE,
    FOREIGN KEY (comment_id) REFERENCES comments (id) ON DELETE CASCADE
);