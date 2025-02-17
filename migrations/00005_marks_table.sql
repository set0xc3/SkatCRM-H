-- +goose Up
CREATE TABLE ad_channels (id INTEGER PRIMARY KEY, name TEXT);

-- +goose Down
DROP TABLE ad_channels;
