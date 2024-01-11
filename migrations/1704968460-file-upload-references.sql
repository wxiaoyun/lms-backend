-- +migrate Up
CREATE TABLE
  file_upload_references (
    id BIGSERIAL PRIMARY KEY,
    file_upload_id BIGINT NOT NULL REFERENCES file_uploads (id),
    attachable_id BIGINT NOT NULL,
    attachable_type VARCHAR(255) NOT NULL,
    UNIQUE (file_upload_id, attachable_id, attachable_type),
    created_at created_at,
    updated_at updated_at,
    deleted_at deleted_at
  );

CREATE INDEX idx_file_upload_references_deleted_at ON file_upload_references (deleted_at);

-- +migrate Down
DROP TABLE file_upload_references;