-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp" SCHEMA public;

CREATE TABLE IF NOT EXISTS _user(
    id uuid DEFAULT public.uuid_generate_v4(),
    phone_number text UNIQUE,
    email text UNIQUE,
    password text,
    is_email_verified bool DEFAULT false,
    CONSTRAINT _user_pk PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS code(
    id uuid DEFAULT public.uuid_generate_v4(),
    email text,
    CONSTRAINT code_pk PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE _user IF EXISTS;
DROP TABLE code IF EXISTS;
-- +goose StatementEnd
