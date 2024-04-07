-- name: CreateTag :one
INSERT INTO "Tag" (
  "name",
  "description"
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetTagByID :one
SELECT * FROM "Tag"
WHERE "id" = $1 LIMIT 1;

-- name: GetTagByName :one
SELECT * FROM "Tag"
WHERE "name" = $1 LIMIT 1;

-- name: ListTags :many
SELECT * FROM "Tag"
ORDER BY "id"
LIMIT $1
OFFSET $2;

-- name: UpdateTag :exec
UPDATE "Tag"
SET
  "name" = $2,
  "description" = $3
WHERE "id" = $1;

-- name: DeleteTag :exec
DELETE FROM "Tag"
WHERE "id" = $1;
