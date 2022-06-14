-- name: CreateUser :one
INSERT INTO users (fullname, email, enabled, attributes)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1
LIMIT 1;

-- name: GetUsers :many
SELECT *
FROM users
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: GetUsersByIds :many
SELECT *
FROM users
WHERE id = ANY (sqlc.arg(ids)::uuid[])
ORDER BY id;

-- name: UpdateUser :one
UPDATE users
SET fullname=$2,
    email=$3,
    enabled=$4,
    attributes=$5,
    updated_at=now()
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE
FROM users
WHERE id = $1;
