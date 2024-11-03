-- +goose Up
-- +goose StatementBegin
create table chart (
    id bigserial primary key,
    charts jsonb not null,
    multipliers jsonb not null,
    description text not null,
    created_at timestamp default current_timestamp not null
);

create table response (
    id bigserial primary key,
    session_id bigint not null,
    query_id bigint not null,
    sources text [] not null,
    filenames text [] not null,
    description text not null,
    created_at timestamp default current_timestamp not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table chart;
drop table response;
-- +goose StatementEnd
