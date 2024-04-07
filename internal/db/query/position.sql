-- name: CreatePosition :one
INSERT INTO "Position" (
  "line",
  "col",
  "toLine",
  "toCol",
  "content",
  "createdAt",
  "updatedAt",
  "commentsId"
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetPositionByID :one
SELECT * FROM "Position"
WHERE "id" = $1 LIMIT 1;

-- name: ListPositions :many
SELECT * FROM "Position"
ORDER BY "createdAt"
LIMIT $1
OFFSET $2;

-- name: UpdatePosition :exec
UPDATE "Position"
SET
  "line" = $2,
  "col" = $3,
  "toLine" = $4,
  "toCol" = $5,
  "content" = $6,
  "updatedAt" = $7,
  "commentsId" = $8
WHERE "id" = $1;

-- name: DeletePosition :exec
DELETE FROM "Position"
WHERE "id" = $1;
