CREATE TABLE IF NOT EXISTS tags (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  name TEXT NOT NULL,
  created_at timestamp NOT NULL DEFAULT now (),
  updated_at timestamp NOT NULL DEFAULT now ()
);