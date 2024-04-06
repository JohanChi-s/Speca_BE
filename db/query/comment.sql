-- name: CreateComment :one
INSERT INTO "Comment" (
  "content",
  "createdAt",
  "updatedAt",
  "userId",
  "documentId",
  "parentCommentId"
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetCommentByID :one
SELECT * FROM "Comment"
WHERE "id" = $1 LIMIT 1;

-- name: ListComments :many
SELECT * FROM "Comment"
ORDER BY "createdAt"
LIMIT $1
OFFSET $2;

-- name: UpdateComment :exec
UPDATE "Comment"
SET
  "content" = $2,
  "updatedAt" = $3,
  "userId" = $4,
  "documentId" = $5,
  "parentCommentId" = $6
WHERE "id" = $1;

-- name: DeleteComment :exec
DELETE FROM "Comment"
WHERE "id" = $1;
