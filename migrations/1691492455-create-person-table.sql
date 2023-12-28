-- +migrate Up
CREATE TABLE
  people (
    id BIGSERIAL PRIMARY KEY,
    full_name VARCHAR NOT NULL,
    preferred_name VARCHAR NOT NULL,
    language_preference VARCHAR NOT NULL,
    created_at created_at,
    updated_at updated_at,
    deleted_at deleted_at
  );

CREATE INDEX idx_people_deleted_at ON people (deleted_at);

-- +migrate Down 
DROP TABLE people;