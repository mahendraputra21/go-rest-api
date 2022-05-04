-- +goose Up
-- +goose StatementBegin
create table public.product
(
    id         serial not null primary key,
    product_name   varchar(255),
    brand_id integer not null,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table public.product;
-- +goose StatementEnd