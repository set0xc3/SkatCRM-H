-- +goose Up
CREATE TABLE clients (
  id TEXT PRIMARY KEY, -- Уникальный идентификатор клиента
  id2 TEXT, -- Дополнительный идентификатор
  mark TEXT, -- Метка или отметка
  contractor TEXT, -- Контрагент
  full_name TEXT, -- Полное имя клиента
  type TEXT, -- Тип клиента
  phones TEXT, -- Номера телефонов (можно хранить как JSON или строку)
  email TEXT UNIQUE, -- Email клиента (уникальное поле)
  legal_address TEXT, -- Юридический адрес
  physical_address TEXT, -- Физический адрес
  registration_date TEXT, -- Дата регистрации
  ad_channel TEXT, -- Канал рекламы
  reg_data_1 TEXT, -- Дополнительные данные регистрации 1
  reg_data_2 TEXT, -- Дополнительные данные регистрации 2
  note TEXT, -- Примечания
  request_count TEXT, -- Количество запросов
  birthday TEXT, -- День рождения
  income TEXT -- Доход
);

-- +goose Down
DROP TABLE clients;
