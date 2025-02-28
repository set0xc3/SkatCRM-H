-- +goose Up
CREATE TABLE clients (
  id TEXT PRIMARY KEY,
  id2 TEXT,
  mark TEXT,
  contractor INTEGER,
  full_name TEXT,
  type TEXT,
  email TEXT,
  legal_address TEXT,
  physical_address TEXT,
  registration_date TEXT,
  ad_channel TEXT,
  reg_data_1 TEXT,
  reg_data_2 TEXT,
  note TEXT,
  request_count INTEGER,
  birthday TEXT,
  income NUMERIC -- NUMERIC (или REAL) для числового значения
);

CREATE TABLE client_phones (
  client_id INTEGER,
  phone TEXT,
  FOREIGN KEY (client_id) REFERENCES clients (id)
);

-- +goose Down
DROP TABLE client_phones;

DROP TABLE clients;
