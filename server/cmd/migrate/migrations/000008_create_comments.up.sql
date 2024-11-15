CREATE TABLE IF EXISTS comments (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    content text NOT NULL,
    author_id uuid NOT NULL REFERENCES users (id),
    writing_id uuid NOT NULL REFERENCES writings (id),
    reply_to uuid REFERENCES comments (id),
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now()
)