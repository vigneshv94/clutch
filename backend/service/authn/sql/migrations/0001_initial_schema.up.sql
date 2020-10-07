CREATE TABLE authn(
    id text PRIMARY KEY,
    updated_at TIMESTAMPTZ DEFAULT NOW()
);
