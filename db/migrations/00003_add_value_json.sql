-- +goose Up
ALTER TABLE conditions ADD COLUMN value JSONB NOT NULL DEFAULT '{}'::JSONB;

-- +goose Down
ALTER TABLE conditions DROP COLUMN value;
