-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp" SCHEMA public;

CREATE TABLE IF NOT EXISTS device (
    id uuid DEFAULT public.uuid_generate_v4(),
    phone_number text NOT NULL,
    pin_code text,
    user_id text,
    is_verified bool DEFAULT false,
    CONSTRAINT device_pk PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS activation_code (
    id uuid DEFAULT public.uuid_generate_v4(),
    device_id uuid NOT NULL,
    code text NOT NULL,
    CONSTRAINT activation_code_pk PRIMARY KEY (id),
    CONSTRAINT fk_device_id FOREIGN KEY (device_id) REFERENCES device (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
