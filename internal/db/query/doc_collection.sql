-- name: AddDocumentToCollection :exec
INSERT INTO "Document_Collection" (
  "document_id",
  "collection_id"
) VALUES (
  $1, $2
);

-- name: RemoveDocumentFromCollection :exec
DELETE FROM "Document_Collection"
WHERE "document_id" = $1 AND "collection_id" = $2;

-- name: GetCollectionsByDocumentID :many
SELECT "collection_id" FROM "Document_Collection"
WHERE "document_id" = $1;
