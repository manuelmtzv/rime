CREATE TABLE
  IF NOT EXISTS tags (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
    name TEXT NOT NULL
  );