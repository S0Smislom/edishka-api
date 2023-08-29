-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "user" (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "created_at" TIMESTAMPTZ NOT NULL  DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ NOT NULL  DEFAULT CURRENT_TIMESTAMP,
    "phone" VARCHAR(20) NOT NULL UNIQUE,
    "first_name" VARCHAR(200),
    "last_name" VARCHAR(200),
    "is_superuser" BOOLEAN NOT NULL DEFAULT FALSE,
    "is_staff" BOOLEAN NOT NULL DEFAULT FALSE,
    "code" VARCHAR(255) NOT NULL,
    "birthday" DATE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "user";
-- +goose StatementEnd
