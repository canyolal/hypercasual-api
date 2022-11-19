CREATE TABLE maillist (
    id bigserial PRIMARY KEY,
    email citext UNIQUE NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

INSERT INTO maillist (email,created_at) VALUES (
    'selami@sahin.com',
    '2022-12-12 12:12:12'
);