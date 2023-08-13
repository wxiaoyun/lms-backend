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
    updated_at updated_at
  );

CREATE TABLE
  loan_histories (
    id bigserial PRIMARY KEY,
    loan_id BIGINT NOT NULL REFERENCES loans (id),
    ACTION VARCHAR NOT NULL,
    created_at created_at
  );

CREATE TABLE
  reservations (
    id bigserial PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users (id),
    book_id BIGINT NOT NULL REFERENCES books (id),
    status VARCHAR NOT NULL,
    reservation_date DATE NOT NULL,
    created_at created_at,
    updated_at updated_at
  );

CREATE TABLE
  fines (
    id bigserial PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users (id),
    loan_id BIGINT NOT NULL REFERENCES loans (id),
    status VARCHAR NOT NULL,
    amount DECIMAL NOT NULL,
    created_at created_at,
    updated_at updated_at
  );

-- +migrate Down
DROP TABLE fines;

DROP TABLE reservations;

DROP TABLE loan_histories;

DROP TABLE loans;