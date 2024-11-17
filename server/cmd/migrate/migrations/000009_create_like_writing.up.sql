CREATE TABLE IF NOT EXISTS like_writing (
    like_id uuid NOT NULL,
    writing_id uuid NOT NULL,
    PRIMARY KEY (like_id, writing_id),
    FOREIGN KEY (like_id) REFERENCES likes (id) ON DELETE CASCADE,
    FOREIGN KEY (writing_id) REFERENCES writings (id) ON DELETE CASCADE
)