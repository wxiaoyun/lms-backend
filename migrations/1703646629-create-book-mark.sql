-- +migrate Up
CREATE TABLE
  bookmarks (
    id bigserial PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users (id),
    book_id BIGINT NOT NULL REFERENCES books (id),
    created_at created_at,
    updated_at updated_at,
    deleted_at deleted_at
  );

CREATE INDEX idx_bookmarks_deleted_at ON bookmarks (deleted_at);

-- +migrate Down
DROP TABLE bookmarks;