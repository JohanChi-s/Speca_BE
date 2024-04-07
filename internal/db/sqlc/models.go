// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type ActionEvent struct {
	ID           string           `json:"id"`
	Action       pgtype.Text      `json:"action"`
	CreatedAt    pgtype.Timestamp `json:"createdAt"`
	Actor        pgtype.Text      `json:"actor"`
	Assignee     pgtype.Text      `json:"assignee"`
	Assigner     pgtype.Text      `json:"assigner"`
	DocumentId   pgtype.Text      `json:"documentId"`
	CollectionId pgtype.Text      `json:"collectionId"`
}

type Collection struct {
	ID                 string           `json:"id"`
	Name               pgtype.Text      `json:"name"`
	Description        pgtype.Text      `json:"description"`
	Icon               pgtype.Text      `json:"icon"`
	IsSaving           pgtype.Bool      `json:"isSaving"`
	CanShare           pgtype.Bool      `json:"canShare"`
	ChildCollectionIds pgtype.Text      `json:"childCollectionIds"`
	DownloadPermission []string         `json:"downloadPermission"`
	CreatedAt          pgtype.Timestamp `json:"createdAt"`
	// Automatically updated by the database
	UpdatedAt          pgtype.Timestamp `json:"updatedAt"`
	WorkspaceId        pgtype.Text      `json:"workspaceId"`
	ParentCollectionId pgtype.Text      `json:"parentCollectionId"`
	OwnerUserId        pgtype.Text      `json:"ownerUserId"`
}

type CollectionMember struct {
	CollectionID pgtype.Text `json:"collection_id"`
	MemberID     pgtype.Text `json:"member_id"`
}

type CollectionTag struct {
	FolderID pgtype.Text `json:"folder_id"`
	TagID    pgtype.Text `json:"tag_id"`
}

type Comment struct {
	ID        string           `json:"id"`
	Content   pgtype.Text      `json:"content"`
	CreatedAt pgtype.Timestamp `json:"createdAt"`
	// Automatically updated by the database
	UpdatedAt       pgtype.Timestamp `json:"updatedAt"`
	UserId          pgtype.Text      `json:"userId"`
	DocumentId      pgtype.Text      `json:"documentId"`
	ParentCommentId pgtype.Text      `json:"parentCommentId"`
}

type Document struct {
	ID          string           `json:"id"`
	Title       pgtype.Text      `json:"title"`
	Text        pgtype.Text      `json:"text"`
	Emoji       pgtype.Text      `json:"emoji"`
	IsPublic    pgtype.Bool      `json:"isPublic"`
	IsFullWidth pgtype.Bool      `json:"isFullWidth"`
	CreatedAt   pgtype.Timestamp `json:"createdAt"`
	// Automatically updated by the database
	UpdatedAt    pgtype.Timestamp `json:"updatedAt"`
	AuthorId     pgtype.Text      `json:"authorId"`
	TeamId       pgtype.Text      `json:"teamId"`
	WorkspaceId  pgtype.Text      `json:"workspaceId"`
	CollectionId pgtype.Text      `json:"collectionId"`
}

type DocumentCollection struct {
	DocumentID   pgtype.Text `json:"document_id"`
	CollectionID pgtype.Text `json:"collection_id"`
}

type DocumentTag struct {
	DocumentID pgtype.Text `json:"document_id"`
	TagID      pgtype.Text `json:"tag_id"`
}

type Member struct {
	ID           string           `json:"id"`
	UserId       pgtype.Text      `json:"userId"`
	CollectionId pgtype.Text      `json:"collectionId"`
	Role         []string         `json:"role"`
	CreatedAt    pgtype.Timestamp `json:"createdAt"`
	// Automatically updated by the database
	UpdatedAt pgtype.Timestamp `json:"updatedAt"`
}

type Position struct {
	ID        string           `json:"id"`
	Line      pgtype.Int4      `json:"line"`
	Col       pgtype.Int4      `json:"col"`
	ToLine    pgtype.Text      `json:"toLine"`
	ToCol     pgtype.Text      `json:"toCol"`
	Content   pgtype.Text      `json:"content"`
	CreatedAt pgtype.Timestamp `json:"createdAt"`
	// Automatically updated by the database
	UpdatedAt  pgtype.Timestamp `json:"updatedAt"`
	CommentsId pgtype.Text      `json:"commentsId"`
}

type Profile struct {
	ID        string           `json:"id"`
	UserId    pgtype.Text      `json:"userId"`
	FullName  pgtype.Text      `json:"fullName"`
	Age       pgtype.Int4      `json:"age"`
	Address   pgtype.Text      `json:"address"`
	AvatarUrl pgtype.Text      `json:"avatarUrl"`
	CreatedAt pgtype.Timestamp `json:"createdAt"`
	// Automatically updated by the database
	UpdatedAt pgtype.Timestamp `json:"updatedAt"`
}

type Tag struct {
	ID          string      `json:"id"`
	Name        pgtype.Text `json:"name"`
	Description pgtype.Text `json:"description"`
}

type Team struct {
	ID              string           `json:"id"`
	Name            pgtype.Text      `json:"name"`
	AvatarUrl       pgtype.Text      `json:"avatarUrl"`
	SubDomain       pgtype.Text      `json:"subDomain"`
	Theme           pgtype.Text      `json:"theme"`
	CanComment      pgtype.Bool      `json:"canComment"`
	CanShare        pgtype.Bool      `json:"canShare"`
	InviteRequired  pgtype.Bool      `json:"inviteRequired"`
	DefaultUserRole pgtype.Text      `json:"defaultUserRole"`
	CreatedAt       pgtype.Timestamp `json:"createdAt"`
	// Automatically updated by the database
	UpdatedAt pgtype.Timestamp `json:"updatedAt"`
	Url       pgtype.Text      `json:"url"`
}

type TeamDocument struct {
	TeamID     pgtype.Text `json:"team_id"`
	DocumentID pgtype.Text `json:"document_id"`
}

type TeamWorkspace struct {
	TeamID      pgtype.Text `json:"team_id"`
	WorkspaceID pgtype.Text `json:"workspace_id"`
}

type User struct {
	ID        string      `json:"id"`
	Email     pgtype.Text `json:"email"`
	Username  pgtype.Text `json:"username"`
	Password  pgtype.Text `json:"password"`
	FirstName pgtype.Text `json:"firstName"`
	LastName  pgtype.Text `json:"lastName"`
	IsAdmin   pgtype.Bool `json:"isAdmin"`
	IsActive  pgtype.Bool `json:"isActive"`
	IsViewer  pgtype.Bool `json:"isViewer"`
	Language  pgtype.Text `json:"language"`
	// Automatically updated by the database
	LastActiveAt pgtype.Timestamp `json:"lastActiveAt"`
	Roles        []string         `json:"roles"`
	CreatedAt    pgtype.Timestamp `json:"createdAt"`
	// Automatically updated by the database
	UpdatedAt pgtype.Timestamp `json:"updatedAt"`
}

type UserDocument struct {
	UserID     pgtype.Text `json:"user_id"`
	DocumentID pgtype.Text `json:"document_id"`
}

type UserTeam struct {
	UserID pgtype.Text `json:"user_id"`
	TeamID pgtype.Text `json:"team_id"`
}

type UserWorkspace struct {
	UserID      pgtype.Text `json:"user_id"`
	WorkspaceID pgtype.Text `json:"workspace_id"`
}

type Workspace struct {
	ID        string           `json:"id"`
	Name      pgtype.Text      `json:"name"`
	Domain    pgtype.Text      `json:"domain"`
	IsPublic  pgtype.Bool      `json:"isPublic"`
	CreatedAt pgtype.Timestamp `json:"createdAt"`
	// Automatically updated by the database
	UpdatedAt pgtype.Timestamp `json:"updatedAt"`
	Url       pgtype.Text      `json:"url"`
}

type WorkspaceDocument struct {
	WorkspaceID pgtype.Text `json:"workspace_id"`
	DocumentID  pgtype.Text `json:"document_id"`
}
