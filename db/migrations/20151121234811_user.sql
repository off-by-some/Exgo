
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "user" (
    id uuid NOT NULL,

    -- FIXME: Unique constraint on one of these?
    username text,
    email text,

    -- Our hash must be salty
    -- so we avoid lookup/rainbow table attacks
    -- password_salt bytea,
    -- We're gonna use PBKDF2
    -- that way hashing is so computatianally intensive
    -- that dictionary attacks / brute-forcing is nearly
    -- impossible
    -- Varying the number of iterations by some random
    -- amount helps somehow
    -- password_iterations integer,
    -- Don't store the password, just the hash
    -- password_hash bytea,

    -- dafaq is body?
    body text,
    PRIMARY KEY(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE "user";
DROP EXTENSION "uuid-ossp";
