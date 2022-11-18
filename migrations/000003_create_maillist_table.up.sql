CREATE TABLE IF NOT EXISTS maillist (
    id bigserial PRIMARY KEY,
    email citext UNIQUE NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);