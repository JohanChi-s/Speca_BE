-- name: AddTagToDocument :exec
INSERT INTO "Document_Tag" (
  "document_id",
  "tag_id"
) VALUES (
  $1, $2
);

-- name: RemoveTagFromDocument :exec
DELETE FROM "Document_Tag"
WHERE "document_id" = $1 AND "tag_id" = $2;

-- name: GetTagsByDocumentID :many
SELECT "tag_id" FROM "Document_Tag"
WHERE "document_id" = $1;
