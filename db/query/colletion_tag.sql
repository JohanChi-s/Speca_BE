-- name: AddTagToCollection :exec
INSERT INTO "Collection_Tag" (
  "folder_id",
  "tag_id"
) VALUES (
  $1, $2
);

-- name: RemoveTagFromCollection :exec
DELETE FROM "Collection_Tag"
WHERE "folder_id" = $1 AND "tag_id" = $2;

-- name: GetTagsByCollectionID :many
SELECT "tag_id" FROM "Collection_Tag"
WHERE "folder_id" = $1;
