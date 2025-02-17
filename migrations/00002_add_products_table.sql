-- +goose Up
CREATE TABLE products (
  id TEXT PRIMARY KEY,
  id2 TEXT,
  name TEXT,
  serial_number TEXT,
  article TEXT,
  date TEXT,
  quantity TEXT, -- Кол-во
  retail_price TEXT, -- Розничная цена
  purchase_price TEXT, -- Закупочная цена
  exchange_rate_pc TEXT, -- Курс ПК
  exchange_rate_pr TEXT, -- Курс ПР
  warehouse TEXT, -- Склад
  location TEXT,
  customer_order TEXT, -- Заказ клиента
  supplier_order TEXT, -- Заказ поставщику
  supplier TEXT -- Поставщик
);

-- +goose Down
DROP TABLE products;
