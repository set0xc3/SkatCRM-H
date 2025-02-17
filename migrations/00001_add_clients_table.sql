-- +goose Up
CREATE TABLE clients (
  id TEXT PRIMARY KEY,
  id2 TEXT,
  mark TEXT,
  contractor TEXT,
  full_name TEXT,
  type TEXT,
  phones TEXT,
  email TEXT,
  legal_address TEXT,
  physical_address TEXT,
  registration_date TEXT,
  ad_channel TEXT,
  reg_data_1 TEXT,
  reg_data_2 TEXT,
  note TEXT,
  request_count TEXT,
  birthday TEXT,
  income TEXT -- Доход
);

-- +goose Down
DROP TABLE clients;
