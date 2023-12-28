-- +migrate Up
CREATE TABLE
  loans (
    id bigserial PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users (id),
    book_id BIGINT NOT NULL REFERENCES books (id),
    status VARCHAR NOT NULL,
    borrow_date DATE NOT NULL,
    due_date DATE NOT NULL,
    return_date DATE,
    created_at created_at,
    updated_at updated_at,
    deleted_at deleted_at
  );

CREATE INDEX idx_loans_deleted_at ON loans (deleted_at);

CREATE TABLE
  loan_histories (
    id bigserial PRIMARY KEY,
    loan_id BIGINT NOT NULL REFERENCES loans (id),
    ACTION VARCHAR NOT NULL,
    created_at created_at,
    updated_at updated_at,
    deleted_at deleted_at
  );

CREATE INDEX idx_loan_histories_deleted_at ON loan_histories (deleted_at);

CREATE TABLE
  reservations (
    id bigserial PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users (id),
    book_id BIGINT NOT NULL REFERENCES books (id),
    status VARCHAR NOT NULL,
    reservation_date DATE NOT NULL,
    created_at created_at,
    updated_at updated_at,
    deleted_at deleted_at
  );

CREATE INDEX idx_reservations_deleted_at ON reservations (deleted_at);

CREATE TABLE
  fines (
    id bigserial PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users (id),
    loan_id BIGINT NOT NULL REFERENCES loans (id),
    status VARCHAR NOT NULL,
    amount DECIMAL NOT NULL,
    created_at created_at,
    updated_at updated_at,
    deleted_at deleted_at
  );

CREATE INDEX idx_fines_deleted_at ON fines (deleted_at);

-- +migrate Down
DROP TABLE fines;

DROP TABLE reservations;

DROP TABLE loan_histories;

DROP TABLE loans;