-- name: CreateDocument :one
INSERT INTO "Document" (
  "title",
  "text",
  "emoji",
  "isPublic",
  "isFullWidth",
  "createdAt",
  "updatedAt",
  "authorId",
  "teamId",
  "workspaceId",
  "collectionId"
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING *;

-- name: GetDocumentByID :one
SELECT * FROM "Document"
WHERE "id" = $1 LIMIT 1;

-- name: ListDocuments :many
SELECT * FROM "Document"
ORDER BY "createdAt"
LIMIT $1
OFFSET $2;

-- name: UpdateDocument :exec
UPDATE "Document"
SET
  "title" = $2,
  "text" = $3,
  "emoji" = $4,
  "isPublic" = $5,
  "isFullWidth" = $6,
  "updatedAt" = $7,
  "authorId" = $8,
  "teamId" = $9,
  "workspaceId" = $10,
  "collectionId" = $11
WHERE "id" = $1;

-- name: DeleteDocument :exec
DELETE FROM "Document"
WHERE "id" = $1;
