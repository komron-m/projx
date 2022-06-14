-- +migrate Up
CREATE TABLE users
(
    id         uuid primary key        default uuid_generate_v1(),
    email      varchar unique not null,
    fullname   varchar        not null,
    enabled    bool           not null default false,
    attributes json           null,
    created_at timestamptz    not null default now(),
    updated_at timestamptz    not null default now()
);

-- +migrate Down
DROP TABLE users;
