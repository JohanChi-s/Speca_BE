-- name: AddTeamToWorkspace :exec
INSERT INTO "Team_Workspace" (
  "team_id",
  "workspace_id"
) VALUES (
  $1, $2
);

-- name: RemoveTeamFromWorkspace :exec
DELETE FROM "Team_Workspace"
WHERE "team_id" = $1 AND "workspace_id" = $2;

-- name: GetWorkspacesByTeamID :many
SELECT "workspace_id" FROM "Team_Workspace"
WHERE "team_id" = $1;
