-- +migrate Up
DROP DOMAIN IF EXISTS created_at;

CREATE DOMAIN created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP;

DROP DOMAIN IF EXISTS creator_id;

CREATE DOMAIN creator_id BIGINT;

DROP DOMAIN IF EXISTS updated_at;

CREATE DOMAIN updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP;

DROP DOMAIN IF EXISTS updater_id;

CREATE DOMAIN updater_id BIGINT;

DROP DOMAIN IF EXISTS deleted_at;

CREATE DOMAIN deleted_at timestamptz;

DROP DOMAIN IF EXISTS deleter_id;

CREATE DOMAIN deleter_id BIGINT;

-- +migrate Down
DROP DOMAIN creator_id;

DROP DOMAIN updated_at;

DROP DOMAIN updater_id;

DROP DOMAIN deleted_at;

DROP DOMAIN deleter_id;
