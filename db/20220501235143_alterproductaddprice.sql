-- +goose Up
-- +goose StatementBegin
alter table public.product add price integer null;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
