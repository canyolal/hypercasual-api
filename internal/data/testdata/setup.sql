CREATE TABLE maillist (
    id bigserial PRIMARY KEY,
    email citext UNIQUE NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

CREATE TABLE publishers (
    id bigserial PRIMARY KEY,
    name text UNIQUE NOT NULL,
    link text NOT NULL,
    version integer NOT NULL DEFAULT 1
);

INSERT INTO publishers (name,link) VALUES (
    'sample publisher',
    'http://tester.com'
);

INSERT INTO maillist (email,created_at) VALUES (
    'selami@sahin.com',
    '2022-12-12 12:12:12'
);