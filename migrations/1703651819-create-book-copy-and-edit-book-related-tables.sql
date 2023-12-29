-- +migrate Up
CREATE TABLE
  book_copies (
    id BIGSERIAL PRIMARY KEY,
    book_id BIGINT NOT NULL REFERENCES books (id),
    status VARCHAR NOT NULL,
    created_at created_at,
    updated_at updated_at,
    deleted_at deleted_at
  );

CREATE INDEX idx_book_copies_deleted_at ON book_copies (deleted_at);

ALTER TABLE loans
ADD COLUMN book_copy_id BIGINT NOT NULL REFERENCES book_copies (id),
DROP COLUMN book_id;

ALTER TABLE reservations
ADD COLUMN book_copy_id BIGINT NOT NULL REFERENCES book_copies (id),
DROP COLUMN book_id;

-- +migrate Down
ALTER TABLE reservations
ADD COLUMN book_id BIGINT REFERENCES books (id),
DROP COLUMN book_copy_id;

ALTER TABLE loans
ADD COLUMN book_id BIGINT REFERENCES books (id),
DROP COLUMN book_copy_id;

DROP TABLE book_copies;