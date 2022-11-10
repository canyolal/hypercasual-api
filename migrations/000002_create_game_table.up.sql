CREATE TABLE IF NOT EXISTS games (
	id bigserial PRIMARY KEY, 
	name text NOT NULL,
	genre text NOT NULL,
	publisher_id bigint NOT NULL REFERENCES publishers ON DELETE CASCADE,
	version integer NOT NULL DEFAULT 1
);