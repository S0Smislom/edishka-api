-- +goose Up
-- +goose StatementBegin
ALTER TABLE "user" ADD "photo" VARCHAR(255);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter TABLE "user" DROP COLUMN "photo";
-- +goose StatementEnd
