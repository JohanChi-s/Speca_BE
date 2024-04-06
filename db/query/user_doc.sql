-- name: AddUserToDocument :exec
INSERT INTO "User_Document" (
  "user_id",
  "document_id"
) VALUES (
  $1, $2
);

-- name: RemoveUserFromDocument :exec
DELETE FROM "User_Document"
WHERE "user_id" = $1 AND "document_id" = $2;

-- name: GetDocumentsByUserID :many
SELECT "document_id" FROM "User_Document"
WHERE "user_id" = $1;
