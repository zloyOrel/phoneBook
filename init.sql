-- Инициализация базы данных phonebook
-- Запуск: psql -U postgres -d phonebook -f init.sql

CREATE TABLE IF NOT EXISTS organizations (
    id         SERIAL PRIMARY KEY,
    name       TEXT    NOT NULL,
    is_default BOOLEAN NOT NULL DEFAULT false
);

CREATE TABLE IF NOT EXISTS departments (
    id              SERIAL PRIMARY KEY,
    parent_id       INT  REFERENCES departments(id) ON DELETE CASCADE,
    organization_id INT  REFERENCES organizations(id) ON DELETE SET NULL,
    name            TEXT NOT NULL,
    sort_order      INT  NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS contacts (
    id             SERIAL PRIMARY KEY,
    department_id  INT  NOT NULL REFERENCES departments(id) ON DELETE CASCADE,
    room           TEXT NOT NULL DEFAULT '',
    position       TEXT NOT NULL DEFAULT '',
    full_name      TEXT NOT NULL DEFAULT '',
    phone_city     TEXT NOT NULL DEFAULT '',
    phone_mobile   TEXT NOT NULL DEFAULT '',
    phone_internal TEXT NOT NULL DEFAULT '',
    email          TEXT NOT NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS admins (
    id            SERIAL PRIMARY KEY,
    username      TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL
);
