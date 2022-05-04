-- +goose Up
-- +goose StatementBegin
alter table public.order add product_id int not null;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
