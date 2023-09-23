-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "product" (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "created_at" TIMESTAMPTZ NOT NULL  DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ NOT NULL  DEFAULT CURRENT_TIMESTAMP,
    "title" VARCHAR(200) NOT NULL,
    "slug" VARCHAR(200) NOT NULL UNIQUE,
    "description" TEXT,
    "photo" VARCHAR(255),
    "calories" INT NOT NULL,
    "squirrels" DOUBLE PRECISION NOT NULL,
    "fats" DOUBLE PRECISION NOT NULL,
    "carbohydrates" DOUBLE PRECISION NOT NULL,
    "created_by_id" INT NOT NULL REFERENCES "user" ("id") ON DELETE CASCADE,
    "updated_by_id" INT REFERENCES "user" ("id") ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "product";
-- +goose StatementEnd
