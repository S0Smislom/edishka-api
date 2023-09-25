-- +goose Up
-- +goose StatementBegin
alter table "product" ADD "suggested_by_user" BOOLEAN DEFAULT TRUE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table "product" drop column "suggested_by_user";
-- +goose StatementEnd
