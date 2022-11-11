CREATE TABLE IF NOT EXISTS publishers (
    id bigserial PRIMARY KEY,
    name text UNIQUE NOT NULL,
    link text NOT NULL,
    version integer NOT NULL DEFAULT 1
);