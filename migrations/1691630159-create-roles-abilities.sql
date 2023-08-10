-- +migrate Up
CREATE TABLE
  abilities (
    id BIGSERIAL PRIMARY KEY,
    NAME VARCHAR UNIQUE NOT NULL,
    description VARCHAR NOT NULL,
    created_at created_at
  );

CREATE TABLE
  roles (
    id BIGSERIAL PRIMARY KEY,
    NAME VARCHAR UNIQUE NOT NULL,
    created_at created_at
  );

CREATE TABLE
  role_abilities (
    id BIGSERIAL PRIMARY KEY,
    role_id BIGINT NOT NULL REFERENCES roles (id),
    ability_id BIGINT NOT NULL REFERENCES abilities (id)
  );

CREATE TABLE
  user_roles (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users (id),
    role_id BIGINT NOT NULL REFERENCES roles (id)
  );

-- +migrate Down
DROP TABLE user_roles;

DROP TABLE role_abilites;

DROP TABLE roles;

DROP TABLE abilities;