CREATE TABLE IF NOT EXISTS games (
	id bigserial PRIMARY KEY, 
	name text NOT NULL,
	genre text NOT NULL,
	publisher_name text NOT NULL,
	version integer NOT NULL DEFAULT 1
);