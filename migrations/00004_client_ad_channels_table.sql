-- +goose Up
CREATE TABLE ad_channels (id INTEGER PRIMARY KEY, name TEXT);

INSERT INTO
  ad_channels (name)
VALUES
  ('Интернет'),
  ('Партнер'),
  ('По рекомендации'),
  ('Постоянные клиенты'),
  ('Проходящий поток'),
  ('СЦ ТРУД');

-- +goose Down
DROP TABLE ad_channels;
