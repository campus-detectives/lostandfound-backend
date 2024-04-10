CREATE TABLE IF NOT EXISTS item (
    id bigserial PRIMARY KEY,
    embedding text,
    image text,
    found_time timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    found_by bigint NOT NULL REFERENCES users ON DELETE CASCADE,
    location text,
    claimed bool NOT NULL default false,
    claimed_by text,
    category text,
    version integer NOT NULL DEFAULT 1
);
