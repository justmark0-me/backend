DROP TABLE request_info; -- I've migrated mannually the data from the old database to the new one.
CREATE TABLE if not exists request_info (
    id SERIAL PRIMARY KEY,
    ip VARCHAR(15) NOT NULL,
    os TEXT,
    country TEXT,
    region TEXT,
    city TEXT,
    latitude DOUBLE PRECISION NOT NULL,
    longitude DOUBLE PRECISION NOT NULL,
    user_agent TEXT,
    headers TEXT,
    endpoint TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL
);

CREATE INDEX IF NOT EXISTS request_info_created_at_idx ON request_info(created_at);