-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists "recipe" (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "created_at" TIMESTAMPTZ NOT NULL  DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ NOT NULL  DEFAULT CURRENT_TIMESTAMP,
    "title" VARCHAR(200) NOT NULL,
    "slug" VARCHAR(200) NOT NULL UNIQUE,
    "description" TEXT,
    "photo" VARCHAR(255),
    "cooking_time" INT NOT NULL,
    "preparing_time" INT,
    "kitchen" VARCHAR(100) NOT NULL,
    "difficulty_level" VARCHAR(100) NOT NULL,
    "published" BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE if not EXISTS "recipe_step" (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "created_at" TIMESTAMPTZ NOT NULL  DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ NOT NULL  DEFAULT CURRENT_TIMESTAMP,
    "title" VARCHAR(200) NOT NULL,
    "description" TEXT,
    "ordering" INT NOT NULL DEFAULT 0,
    "photo" VARCHAR(255),
    "recipe_id" INT NOT NULL REFERENCES "recipe" ("id") ON DELETE CASCADE
);

CREATE TABLE if not EXISTS "step_product" (
    "recipe_step_id" int not null REFERENCES "recipe_step" ("id") ON DELETE CASCADE,
    "product_id" int not null REFERENCES "product" ("id") on delete CASCADE,
    "amount" DOUBLE PRECISION NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if EXISTS "step_product";
drop table if EXISTS "recipe_step";
drop table if EXISTS "recipe";
-- +goose StatementEnd
