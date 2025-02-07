-- +goose Up
CREATE TABLE IF NOT EXISTS calls (
  id TEXT PRIMARY KEY, -- Уникальный идентификатор звонка
  id2 TEXT -- Дополнительный идентификатор
);

-- +goose Down
DROP TABLE calls;
