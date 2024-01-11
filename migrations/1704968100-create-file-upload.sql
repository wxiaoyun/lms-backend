-- +migrate Up
CREATE TABLE
  file_uploads (
    id BIGSERIAL PRIMARY KEY,
    file_name VARCHAR(255) NOT NULL,
    file_path VARCHAR(255) NOT NULL,
    content_type VARCHAR(255) NOT NULL,
    created_at created_at,
    updated_at updated_at,
    deleted_at deleted_at
  );

CREATE INDEX idx_file_uploads_deleted_at ON file_uploads (deleted_at);

-- +migrate Down
DROP TABLE file_uploads;