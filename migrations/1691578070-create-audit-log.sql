-- +migrate Up
CREATE TABLE audit_logs (
    id                     BIGSERIAL PRIMARY KEY,
    user_id                BIGINT         NOT NULL REFERENCES users(id),
    action                 VARCHAR        NOT NULL,
    created_at             created_at
);

-- +migrate Down
DROP TABLE audit_log;