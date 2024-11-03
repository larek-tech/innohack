-- +goose Up
-- +goose StatementBegin
create table users(
    id bigserial primary key,
    email text not null unique,
    password text not null,
    created_at timestamp default current_timestamp not null
);
create table session(
    id uuid primary key,
    user_id bigint not null,
    created_at timestamp default current_timestamp not null,
    updated_at timestamp default current_timestamp not null,
    is_deleted boolean default false not null
);
create table query(
    id bigserial primary key,
    session_id bigint not null,
    prompt text not null,
    created_at timestamp default current_timestamp not null
);
create table response(
    id bigserial primary key,
    session_id bigint not null,
    query_id bigint not null,
    sources text [] not null,
    filenames text [] not null,
    charts jsonb not null,
    description text not null,
    multipliers jsonb not null,
    created_at timestamp default current_timestamp not null
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
drop table users;
drop table session;
drop table query;
drop table response;
-- +goose StatementEnd