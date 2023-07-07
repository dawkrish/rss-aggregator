-- +goose Up
ALTER TABLE USERS
ADD COLUMN apikey VARCHAR(64) UNIQUE NOT NULL DEFAULT encode(sha256(random()::text::bytea), 'hex');

-- +goose Down
ALTER TABLE USERS
DROP COLUMN apikey;
