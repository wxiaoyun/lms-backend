-- +migrate Up
CREATE TABLE
  books (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR NOT NULL,
    author VARCHAR NOT NULL,
    ISBN VARCHAR NOT NULL,
    publisher VARCHAR NOT NULL,
    publication_date DATE NOT NULL,
    genre VARCHAR NOT NULL,
    LANGUAGE VARCHAR NOT NULL,
    created_at created_at,
    updated_at updated_at,
    deleted_at deleted_at
  );

-- +migrate Down
DROP TABLE books;