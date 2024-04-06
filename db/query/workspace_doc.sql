-- name: AddWorkspaceToDocument :exec
INSERT INTO "Workspace_Document" (
  "workspace_id",
  "document_id"
) VALUES (
  $1, $2
);

-- name: RemoveWorkspaceFromDocument :exec
DELETE FROM "Workspace_Document"
WHERE "workspace_id" = $1 AND "document_id" = $2;

-- name: GetDocumentsByWorkspaceID :many
SELECT "document_id" FROM "Workspace_Document"
WHERE "workspace_id" = $1;
