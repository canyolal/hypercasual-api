CREATE TABLE IF NOT EXISTS games (
	id bigserial PRIMARY KEY, 
	name text NOT NULL,
	genre text NOT NULL,
	publisher_name text NOT NULL,
	created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
	version integer NOT NULL DEFAULT 1
);