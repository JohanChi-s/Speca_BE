-- name: AddUserToWorkspace :exec
INSERT INTO "User_Workspace" (
  "user_id",
  "workspace_id"
) VALUES (
  $1, $2
);

-- name: RemoveUserFromWorkspace :exec
DELETE FROM "User_Workspace"
WHERE "user_id" = $1 AND "workspace_id" = $2;

-- name: GetWorkspacesByUserID :many
SELECT "workspace_id" FROM "User_Workspace"
WHERE "user_id" = $1;

-- name: GetUsersByWorkspaceID :many
SELECT "user_id" FROM "User_Workspace"
WHERE "workspace_id" = $1;

-- name: RemoveAllUsersFromWorkspace :exec
DELETE FROM "User_Workspace"
WHERE "workspace_id" = $1;
