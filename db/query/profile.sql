-- name: CreateProfile :one
INSERT INTO "Profile" (
  "userId",
  "fullName",
  "age",
  "address",
  "avatarUrl",
  "createdAt",
  "updatedAt"
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetProfileByID :one
SELECT * FROM "Profile"
WHERE "id" = $1 LIMIT 1;

-- name: GetProfileByUserID :one
SELECT * FROM "Profile"
WHERE "userId" = $1 LIMIT 1;

-- name: ListProfiles :many
SELECT * FROM "Profile"
ORDER BY "createdAt"
LIMIT $1
OFFSET $2;

-- name: UpdateProfile :exec
UPDATE "Profile"
SET
  "fullName" = $2,
  "age" = $3,
  "address" = $4,
  "avatarUrl" = $5,
  "updatedAt" = $6
WHERE "id" = $1;

-- name: DeleteProfile :exec
DELETE FROM "Profile"
WHERE "id" = $1;
