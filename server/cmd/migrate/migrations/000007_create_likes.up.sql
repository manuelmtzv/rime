CREATE TABLE IF NOT EXISTS likes (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    author_id uuid NOT NULL REFERENCES users (id),
    writing_id uuid NOT NULL REFERENCES writings (id),
    created_at timestamp NOT NULL DEFAULT now()
);