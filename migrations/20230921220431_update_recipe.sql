-- +goose Up
-- +goose StatementBegin
alter TABLE "recipe" DROP COLUMN "photo";
CREATE TABLE if not EXISTS "recipe_gallery" (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "created_at" TIMESTAMPTZ NOT NULL  DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ NOT NULL  DEFAULT CURRENT_TIMESTAMP,
    "ordering" INT NOT NULL DEFAULT 0,
    "published" BOOLEAN NOT NULL DEFAULT TRUE,
    "photo" VARCHAR(255),
    "recipe_id" INT NOT NULL REFERENCES "recipe" ("id") ON DELETE CASCADE,
    "created_by_id" INT NOT NULL REFERENCES "user" ("id") ON DELETE CASCADE,
    "updated_by_id" INT REFERENCES "user" ("id") ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "recipe" ADD "photo" VARCHAR(255);
drop table if EXISTS "recipe_gallery";
-- +goose StatementEnd
