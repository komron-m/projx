-- +migrate Up
INSERT INTO accounts (owner, balance, currency)
VALUES ('owner1', 100, 'USD'),
       ('owner2', 100, 'USD'),
       ('owner3', 100, 'USD');

-- +migrate Down
DELETE
FROM accounts;
