-- +goose Up
-- +goose StatementBegin
create table public.brand
(
    id           serial not null primary key,
    brand_name   varchar(255),
    created_at   timestamp with time zone default now(),
    updated_at   timestamp with time zone default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table public.brand;
-- +goose StatementEnd