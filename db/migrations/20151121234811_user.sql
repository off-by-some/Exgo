
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "user" (
    id uuid NOT NULL,
    username text,
    email text,

    -- TODO: We all know just by looking at it.
    password text,

    body text,
    PRIMARY KEY(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE "user";
DROP EXTENSION "uuid-ossp";
