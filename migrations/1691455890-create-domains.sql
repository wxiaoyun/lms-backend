-- +migrate up
CREATE DOMAIN created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP;
CREATE DOMAIN creator_id bigint;
CREATE DOMAIN updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP;
CREATE DOMAIN updater_id bigint;
CREATE DOMAIN deleted_at timestamptz;
CREATE DOMAIN deleter_id bigint;

-- +migrate down
DROP DOMAIN created_at;
DROP DOMAIN creator_id;
DROP DOMAIN updated_at;
DROP DOMAIN updater_id;
DROP DOMAIN deleted_at;
DROP DOMAIN deleter_id;
