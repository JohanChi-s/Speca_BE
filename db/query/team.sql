-- name: CreateTeam :one
INSERT INTO "Team" (
  "name",
  "avatarUrl",
  "subDomain",
  "theme",
  "canComment",
  "canShare",
  "inviteRequired",
  "defaultUserRole",
  "createdAt",
  "updatedAt",
  "url"
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING *;

-- name: GetTeamByID :one
SELECT * FROM "Team"
WHERE "id" = $1 LIMIT 1;

-- name: ListTeams :many
SELECT * FROM "Team"
ORDER BY "createdAt"
LIMIT $1
OFFSET $2;

-- name: UpdateTeam :exec
UPDATE "Team"
SET
  "name" = $2,
  "avatarUrl" = $3,
  "subDomain" = $4,
  "theme" = $5,
  "canComment" = $6,
  "canShare" = $7,
  "inviteRequired" = $8,
  "defaultUserRole" = $9,
  "updatedAt" = $10,
  "url" = $11
WHERE "id" = $1;

-- name: DeleteTeam :exec
DELETE FROM "Team"
WHERE "id" = $1;
