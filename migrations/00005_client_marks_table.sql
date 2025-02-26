-- +goose Up
CREATE TABLE marks (id INTEGER PRIMARY KEY, name TEXT);

INSERT INTO
  marks (name)
VALUES
  ('-5%'),
  ('-10%'),
  ('-20%'),
  ('-30%'),
  ('blacklist'),
  ('discount'),
  ('regular'),
  ('VIP'),
  ('Животное'),
  ('Мудак');

-- +goose Down
DROP TABLE marks;
