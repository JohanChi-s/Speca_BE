-- name: AddUserToTeam :exec
INSERT INTO "User_Team" (
  "user_id",
  "team_id"
) VALUES (
  $1, $2
);

-- name: RemoveUserFromTeam :exec
DELETE FROM "User_Team"
WHERE "user_id" = $1 AND "team_id" = $2;

-- name: GetTeamsByUserID :many
SELECT "team_id" FROM "User_Team"
WHERE "user_id" = $1;

-- name: GetUsersByTeamID :many
SELECT "user_id" FROM "User_Team"
WHERE "team_id" = $1;

-- name: RemoveAllUsersFromTeam :exec
DELETE FROM "User_Team"
WHERE "team_id" = $1;