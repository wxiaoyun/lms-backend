-- +migrate Up
CREATE TABLE
  users (
    id BIGSERIAL PRIMARY KEY,
    person_id BIGINT NOT NULL REFERENCES people (id),
    username VARCHAR UNIQUE NOT NULL,
    -- email VARCHAR UNIQUE,
    encrypted_password VARCHAR NOT NULL,
    sign_in_count INTEGER NOT NULL DEFAULT 0,
    current_sign_in_at timestamptz,
    last_sign_in_at timestamptz,
    created_at created_at,
    updated_at updated_at,
    deleted_at deleted_at
  );

CREATE INDEX idx_users_deleted_at ON users (deleted_at);

-- +migrate Down
DROP TABLE users;