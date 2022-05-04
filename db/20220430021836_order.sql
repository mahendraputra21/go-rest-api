-- +goose Up
-- +goose StatementBegin
create table public.order
(
    id         serial not null primary key,
    price   integer,
    qty integer not null,
    grand_total integer null,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table public.order;
-- +goose StatementEnd