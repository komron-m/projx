-- +migrate Up
CREATE
    EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE
    EXTENSION IF NOT EXISTS "pg_trgm";


CREATE TABLE accounts
(
    id         bigserial primary key,
    owner      varchar     not null,
    balance    bigint      not null,
    currency   varchar     not null,
    created_at timestamptz not null default now()
);

CREATE TABLE transfers
(
    id              bigserial primary key,
    from_account_id bigint      not null references accounts (id),
    to_account_id   bigint      not null references accounts (id),
    amount          bigint      not null,
    created_at      timestamptz not null default now()
);

CREATE TABLE entries
(
    id         bigserial primary key,
    account_id bigint      not null references accounts (id),
    amount     bigint      not null,
    created_at timestamptz not null default now()
);

-- +migrate Down
DROP EXTENSION IF EXISTS "uuid-ossp";
DROP EXTENSION IF EXISTS "pg_trgm";

DROP TABLE IF EXISTS entries;
DROP TABLE IF EXISTS transfers;
DROP TABLE IF EXISTS accounts;
