-- +goose Up
ALTER TABLE dictionary
ADD CONSTRAINT unique_word UNIQUE (word);

-- +goose Down
ALTER TABLE dictionary
DROP CONSTRAINT unique_word;

