// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	AddDocumentToCollection(ctx context.Context, arg AddDocumentToCollectionParams) error
	AddTagToCollection(ctx context.Context, arg AddTagToCollectionParams) error
	AddTagToDocument(ctx context.Context, arg AddTagToDocumentParams) error
	AddTeamToDocument(ctx context.Context, arg AddTeamToDocumentParams) error
	AddTeamToWorkspace(ctx context.Context, arg AddTeamToWorkspaceParams) error
	AddUserToDocument(ctx context.Context, arg AddUserToDocumentParams) error
	AddUserToTeam(ctx context.Context, arg AddUserToTeamParams) error
	AddUserToWorkspace(ctx context.Context, arg AddUserToWorkspaceParams) error
	AddWorkspaceToDocument(ctx context.Context, arg AddWorkspaceToDocumentParams) error
	CreateActionEvent(ctx context.Context, arg CreateActionEventParams) (ActionEvent, error)
	CreateCollection(ctx context.Context, arg CreateCollectionParams) (Collection, error)
	CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error)
	CreateDocument(ctx context.Context, arg CreateDocumentParams) (Document, error)
	CreatePosition(ctx context.Context, arg CreatePositionParams) (Position, error)
	CreateProfile(ctx context.Context, arg CreateProfileParams) (Profile, error)
	CreateTag(ctx context.Context, arg CreateTagParams) (Tag, error)
	CreateTeam(ctx context.Context, arg CreateTeamParams) (Team, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateWorkspace(ctx context.Context, arg CreateWorkspaceParams) (Workspace, error)
	DeleteActionEvent(ctx context.Context, id uuid.UUID) error
	DeleteCollection(ctx context.Context, id uuid.UUID) error
	DeleteComment(ctx context.Context, id uuid.UUID) error
	DeleteDocument(ctx context.Context, id uuid.UUID) error
	DeletePosition(ctx context.Context, id uuid.UUID) error
	DeleteProfile(ctx context.Context, id uuid.UUID) error
	DeleteTag(ctx context.Context, id uuid.UUID) error
	DeleteTeam(ctx context.Context, id uuid.UUID) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
	DeleteWorkspace(ctx context.Context, id uuid.UUID) error
	GetActionEventByID(ctx context.Context, id uuid.UUID) (ActionEvent, error)
	GetCollectionByID(ctx context.Context, id uuid.UUID) (Collection, error)
	GetCollectionsByDocumentID(ctx context.Context, documentID pgtype.UUID) ([]pgtype.UUID, error)
	GetCommentByID(ctx context.Context, id uuid.UUID) (Comment, error)
	GetDocumentByID(ctx context.Context, id uuid.UUID) (Document, error)
	GetDocumentsByTeamID(ctx context.Context, teamID pgtype.UUID) ([]pgtype.UUID, error)
	GetDocumentsByUserID(ctx context.Context, userID pgtype.UUID) ([]pgtype.UUID, error)
	GetDocumentsByWorkspaceID(ctx context.Context, workspaceID pgtype.UUID) ([]pgtype.UUID, error)
	GetPositionByID(ctx context.Context, id uuid.UUID) (Position, error)
	GetProfileByID(ctx context.Context, id uuid.UUID) (Profile, error)
	GetProfileByUserID(ctx context.Context, userid pgtype.UUID) (Profile, error)
	GetTagByID(ctx context.Context, id uuid.UUID) (Tag, error)
	GetTagByName(ctx context.Context, name pgtype.Text) (Tag, error)
	GetTagsByCollectionID(ctx context.Context, folderID pgtype.UUID) ([]pgtype.UUID, error)
	GetTagsByDocumentID(ctx context.Context, documentID pgtype.UUID) ([]pgtype.UUID, error)
	GetTeamByID(ctx context.Context, id uuid.UUID) (Team, error)
	GetTeamsByUserID(ctx context.Context, userID pgtype.UUID) ([]pgtype.UUID, error)
	GetUserByEmail(ctx context.Context, email pgtype.Text) (User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (User, error)
	GetUserByUsername(ctx context.Context, username pgtype.Text) (User, error)
	GetUsersByTeamID(ctx context.Context, teamID pgtype.UUID) ([]pgtype.UUID, error)
	GetUsersByWorkspaceID(ctx context.Context, workspaceID pgtype.UUID) ([]pgtype.UUID, error)
	GetWorkspaceByID(ctx context.Context, id uuid.UUID) (Workspace, error)
	GetWorkspacesByTeamID(ctx context.Context, teamID pgtype.UUID) ([]pgtype.UUID, error)
	GetWorkspacesByUserID(ctx context.Context, userID pgtype.UUID) ([]pgtype.UUID, error)
	ListActionEvents(ctx context.Context, arg ListActionEventsParams) ([]ActionEvent, error)
	ListCollections(ctx context.Context, arg ListCollectionsParams) ([]Collection, error)
	ListComments(ctx context.Context, arg ListCommentsParams) ([]Comment, error)
	ListDocuments(ctx context.Context, arg ListDocumentsParams) ([]Document, error)
	ListPositions(ctx context.Context, arg ListPositionsParams) ([]Position, error)
	ListProfiles(ctx context.Context, arg ListProfilesParams) ([]Profile, error)
	ListTags(ctx context.Context, arg ListTagsParams) ([]Tag, error)
	ListTeams(ctx context.Context, arg ListTeamsParams) ([]Team, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	ListWorkspaces(ctx context.Context, arg ListWorkspacesParams) ([]Workspace, error)
	RemoveAllUsersFromTeam(ctx context.Context, teamID pgtype.UUID) error
	RemoveAllUsersFromWorkspace(ctx context.Context, workspaceID pgtype.UUID) error
	RemoveDocumentFromCollection(ctx context.Context, arg RemoveDocumentFromCollectionParams) error
	RemoveTagFromCollection(ctx context.Context, arg RemoveTagFromCollectionParams) error
	RemoveTagFromDocument(ctx context.Context, arg RemoveTagFromDocumentParams) error
	RemoveTeamFromDocument(ctx context.Context, arg RemoveTeamFromDocumentParams) error
	RemoveTeamFromWorkspace(ctx context.Context, arg RemoveTeamFromWorkspaceParams) error
	RemoveUserFromDocument(ctx context.Context, arg RemoveUserFromDocumentParams) error
	RemoveUserFromTeam(ctx context.Context, arg RemoveUserFromTeamParams) error
	RemoveUserFromWorkspace(ctx context.Context, arg RemoveUserFromWorkspaceParams) error
	RemoveWorkspaceFromDocument(ctx context.Context, arg RemoveWorkspaceFromDocumentParams) error
	UpdateActionEvent(ctx context.Context, arg UpdateActionEventParams) error
	UpdateCollection(ctx context.Context, arg UpdateCollectionParams) error
	UpdateComment(ctx context.Context, arg UpdateCommentParams) error
	UpdateDocument(ctx context.Context, arg UpdateDocumentParams) error
	UpdatePosition(ctx context.Context, arg UpdatePositionParams) error
	UpdateProfile(ctx context.Context, arg UpdateProfileParams) error
	UpdateTag(ctx context.Context, arg UpdateTagParams) error
	UpdateTeam(ctx context.Context, arg UpdateTeamParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
	UpdateWorkspace(ctx context.Context, arg UpdateWorkspaceParams) error
}

var _ Querier = (*Queries)(nil)
