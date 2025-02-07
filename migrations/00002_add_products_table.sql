-- +goose Up
CREATE TABLE products (
  id TEXT PRIMARY KEY, -- Уникальный идентификатор продукта
  id2 TEXT, -- № изделия
  name TEXT, -- Наименование
  serial_number TEXT, -- Серийный номер
  article TEXT, -- Артикул
  date TEXT, -- Дата
  quantity TEXT, -- Кол-во
  retail_price TEXT, -- Розничная цена
  purchase_price TEXT, -- Закупочная цена
  exchange_rate_pc TEXT, -- Курс ПК
  exchange_rate_pr TEXT, -- Курс ПР
  warehouse TEXT, -- Склад
  location TEXT, -- Локация
  customer_order TEXT, -- Заказ клиента
  supplier_order TEXT, -- Заказ поставщику
  supplier TEXT -- Поставщик
);

-- +goose Down
DROP TABLE products;
