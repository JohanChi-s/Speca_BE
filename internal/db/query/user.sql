-- name: CreateUser :one
INSERT INTO "User" (
  "email",
  "username",
  "password",
  "firstName",
  "lastName",
  "isAdmin",
  "isActive",
  "isViewer",
  "language",
  "createdAt",
  "updatedAt"
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM "User"
WHERE "id" = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM "User"
WHERE "email" = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM "User"
WHERE "username" = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM "User"
ORDER BY "createdAt"
LIMIT $1
OFFSET $2;

-- name: UpdateUser :exec
UPDATE "User"
SET
  "email" = $2,
  "username" = $3,
  "password" = $4,
  "firstName" = $5,
  "lastName" = $6,
  "isAdmin" = $7,
  "isActive" = $8,
  "isViewer" = $9,
  "language" = $10,
  "updatedAt" = $11
WHERE "id" = $1;

-- name: DeleteUser :exec
DELETE FROM "User"
WHERE "id" = $1;
