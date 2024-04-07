// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: comment.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createComment = `-- name: CreateComment :one
INSERT INTO "Comment" (
  "content",
  "createdAt",
  "updatedAt",
  "userId",
  "documentId",
  "parentCommentId"
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING id, content, "createdAt", "updatedAt", "userId", "documentId", "parentCommentId"
`

type CreateCommentParams struct {
	Content         pgtype.Text      `json:"content"`
	CreatedAt       pgtype.Timestamp `json:"createdAt"`
	UpdatedAt       pgtype.Timestamp `json:"updatedAt"`
	UserId          pgtype.Text      `json:"userId"`
	DocumentId      pgtype.Text      `json:"documentId"`
	ParentCommentId pgtype.Text      `json:"parentCommentId"`
}

func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error) {
	row := q.db.QueryRow(ctx, createComment,
		arg.Content,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserId,
		arg.DocumentId,
		arg.ParentCommentId,
	)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserId,
		&i.DocumentId,
		&i.ParentCommentId,
	)
	return i, err
}

const deleteComment = `-- name: DeleteComment :exec
DELETE FROM "Comment"
WHERE "id" = $1
`

func (q *Queries) DeleteComment(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteComment, id)
	return err
}

const getCommentByID = `-- name: GetCommentByID :one
SELECT id, content, "createdAt", "updatedAt", "userId", "documentId", "parentCommentId" FROM "Comment"
WHERE "id" = $1 LIMIT 1
`

func (q *Queries) GetCommentByID(ctx context.Context, id string) (Comment, error) {
	row := q.db.QueryRow(ctx, getCommentByID, id)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserId,
		&i.DocumentId,
		&i.ParentCommentId,
	)
	return i, err
}

const listComments = `-- name: ListComments :many
SELECT id, content, "createdAt", "updatedAt", "userId", "documentId", "parentCommentId" FROM "Comment"
ORDER BY "createdAt"
LIMIT $1
OFFSET $2
`

type ListCommentsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListComments(ctx context.Context, arg ListCommentsParams) ([]Comment, error) {
	rows, err := q.db.Query(ctx, listComments, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Comment{}
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserId,
			&i.DocumentId,
			&i.ParentCommentId,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateComment = `-- name: UpdateComment :exec
UPDATE "Comment"
SET
  "content" = $2,
  "updatedAt" = $3,
  "userId" = $4,
  "documentId" = $5,
  "parentCommentId" = $6
WHERE "id" = $1
`

type UpdateCommentParams struct {
	ID              string           `json:"id"`
	Content         pgtype.Text      `json:"content"`
	UpdatedAt       pgtype.Timestamp `json:"updatedAt"`
	UserId          pgtype.Text      `json:"userId"`
	DocumentId      pgtype.Text      `json:"documentId"`
	ParentCommentId pgtype.Text      `json:"parentCommentId"`
}

func (q *Queries) UpdateComment(ctx context.Context, arg UpdateCommentParams) error {
	_, err := q.db.Exec(ctx, updateComment,
		arg.ID,
		arg.Content,
		arg.UpdatedAt,
		arg.UserId,
		arg.DocumentId,
		arg.ParentCommentId,
	)
	return err
}
