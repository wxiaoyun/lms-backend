-- +migrate up
CREATE TABLE users (
    id                     BIGSERIAL PRIMARY KEY,
    email                  VARCHAR UNIQUE NOT NULL,
    encrypted_password     VARCHAR        NOT NULL,
    sign_in_count          integer        NOT NULL DEFAULT 0,
    current_sign_in_at     timestamptz,
    last_sign_in_at        timestamptz,
    created_at             created_at,
    updated_at             updated_at,
    deleted_at             deleted_at
);

-- +migrate down
DROP TABLE users;
