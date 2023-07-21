CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY,
  name text NOT NULL,
  email text NOT NULL,
  mobile text NOT NULL,
  created_at timestamptz NOT NULL,
  updated_at timestamptz NOT NULL
);