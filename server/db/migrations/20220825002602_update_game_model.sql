-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

ALTER TABLE Model ADD COLUMN NumPlayers INT NOT NULL AFTER Name;
UPDATE Model SET NumPlayers = 12 WHERE Name = "normal";

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

ALTER TABLE Model DROP NumPlayers;