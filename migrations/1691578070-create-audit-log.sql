-- +migrate Up
CREATE TABLE
  audit_logs (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users (id),
    ACTION VARCHAR NOT NULL,
    date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at created_at
  );

-- +migrate Down
DROP TABLE audit_log;