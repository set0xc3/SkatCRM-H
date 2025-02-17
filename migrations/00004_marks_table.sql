-- +goose Up
CREATE TABLE marks (id INTEGER PRIMARY KEY, name TEXT);

-- +goose Down
DROP TABLE marks;
