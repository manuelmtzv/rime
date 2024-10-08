CREATE TABLE
  IF NOT EXISTS users (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
    username text NOT NULL UNIQUE,
    name text NOT NULL,
    lastname text NOT NULL,
    email text NOT NULL UNIQUE,
    hashedPassword text NOT NULL,
    created_at timestamp
    with
      time zone DEFAULT now (),
      updated_at timestamp
    with
      time zone DEFAULT now ()
  );