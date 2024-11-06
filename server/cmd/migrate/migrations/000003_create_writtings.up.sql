CREATE TABLE
    IF NOT EXISTS writings (
        id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
        type varchar(255) NOT NULL,
        title varchar(255) NOT NULL,
        content text NOT NULL,
        author_id uuid NOT NULL REFERENCES users (id),
        created_at timestamp NOT NULL DEFAULT now (),
        updated_at timestamp NOT NULL DEFAULT now ()
    )