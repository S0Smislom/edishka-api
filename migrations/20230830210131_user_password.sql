-- +goose Up
-- +goose StatementBegin
ALTER TABLE "user" ADD "password" VARCHAR(200);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "user" DROP COLUMN "password";
-- +goose StatementEnd
