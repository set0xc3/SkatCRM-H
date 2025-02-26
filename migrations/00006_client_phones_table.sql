-- +goose Up
CREATE TABLE client_phones (
  client_id INTEGER,
  phone TEXT,
  FOREIGN KEY (client_id) REFERENCES clients (id)
);

-- +goose Down
DROP TABLE client_phones;
