-- +migrate Up
CREATE TABLE people (
    id                     BIGSERIAL PRIMARY KEY,
    first_name             VARCHAR        NOT NULL,
    last_name              VARCHAR        NOT NULL,
    created_at             created_at,
    updated_at             updated_at,
    deleted_at             deleted_at
);

-- +migrate Down 
DROP TABLE people;