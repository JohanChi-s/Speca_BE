-- name: CreateActionEvent :one
INSERT INTO "ActionEvent" (
  "action",
  "createdAt",
  "actor",
  "assignee",
  "assigner",
  "documentId",
  "collectionId"
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetActionEventByID :one
SELECT * FROM "ActionEvent"
WHERE "id" = $1 LIMIT 1;

-- name: ListActionEvents :many
SELECT * FROM "ActionEvent"
ORDER BY "createdAt"
LIMIT $1
OFFSET $2;

-- name: UpdateActionEvent :exec
UPDATE "ActionEvent"
SET
  "action" = $2,
  "createdAt" = $3,
  "actor" = $4,
  "assignee" = $5,
  "assigner" = $6,
  "documentId" = $7,
  "collectionId" = $8
WHERE "id" = $1;

-- name: DeleteActionEvent :exec
DELETE FROM "ActionEvent"
WHERE "id" = $1;
