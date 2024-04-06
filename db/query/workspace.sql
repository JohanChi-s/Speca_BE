-- name: CreateWorkspace :one
INSERT INTO "Workspace" (
  "name",
  "domain",
  "isPublic",
  "createdAt",
  "updatedAt",
  "url"
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetWorkspaceByID :one
SELECT * FROM "Workspace"
WHERE "id" = $1 LIMIT 1;

-- name: ListWorkspaces :many
SELECT * FROM "Workspace"
ORDER BY "createdAt"
LIMIT $1
OFFSET $2;

-- name: UpdateWorkspace :exec
UPDATE "Workspace"
SET
  "name" = $2,
  "domain" = $3,
  "isPublic" = $4,
  "updatedAt" = $5,
  "url" = $6
WHERE "id" = $1;

-- name: DeleteWorkspace :exec
DELETE FROM "Workspace"
WHERE "id" = $1;
