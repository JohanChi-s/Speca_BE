-- name: AddTeamToDocument :exec
INSERT INTO "Team_Document" (
  "team_id",
  "document_id"
) VALUES (
  $1, $2
);

-- name: RemoveTeamFromDocument :exec
DELETE FROM "Team_Document"
WHERE "team_id" = $1 AND "document_id" = $2;

-- name: GetDocumentsByTeamID :many
SELECT "document_id" FROM "Team_Document"
WHERE "team_id" = $1;
