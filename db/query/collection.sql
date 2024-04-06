-- name: CreateCollection :one
INSERT INTO "Collection" (
  "name",
  "description",
  "icon",
  "isSaving",
  "canShare",
  "childCollectionIds",
  "downloadPermission",
  "createdAt",
  "updatedAt",
  "workspaceId",
  "parentCollectionId",
  "ownerUserId"
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
) RETURNING *;

-- name: GetCollectionByID :one
SELECT * FROM "Collection"
WHERE "id" = $1 LIMIT 1;

-- name: ListCollections :many
SELECT * FROM "Collection"
ORDER BY "createdAt"
LIMIT $1
OFFSET $2;

-- name: UpdateCollection :exec
UPDATE "Collection"
SET
  "name" = $2,
  "description" = $3,
  "icon" = $4,
  "isSaving" = $5,
  "canShare" = $6,
  "childCollectionIds" = $7,
  "downloadPermission" = $8,
  "updatedAt" = $9,
  "workspaceId" = $10,
  "parentCollectionId" = $11,
  "ownerUserId" = $12
WHERE "id" = $1;

-- name: DeleteCollection :exec
DELETE FROM "Collection"
WHERE "id" = $1;
