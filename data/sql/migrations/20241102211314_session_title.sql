-- +goose Up
-- +goose StatementBegin
alter table session
add column title text not null default '';
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
alter table session drop column title;
-- +goose StatementEnd