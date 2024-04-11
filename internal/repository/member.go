package repository

import (
	"context"
	"core_service/internal/model"
)

type MemberRepository interface {
	GetMember(ctx context.Context, id int64) (*model.Member, error)
	CreateMember(ctx context.Context, member *model.Member) error
	UpdateMember(ctx context.Context, member *model.Member) error
	DeleteMember(ctx context.Context, id int64) error
	GetMembers(ctx context.Context) ([]model.Member, error)
	GetMemberByUserID(ctx context.Context, userID uint) ([]model.Member, error)
	GetMembersByCollectionID(ctx context.Context, collectionID uint) ([]model.Member, error)
	GetMembersByWorkspaceID(ctx context.Context, workspaceID uint) ([]model.Member, error)
	GetMembersByRole(ctx context.Context, role string) ([]model.Member, error)
	RemoveMemberFromCollection(ctx context.Context, memberID, collectionID uint) error
	AddMemberToWorkspace(ctx context.Context, memberID, workspaceID uint) error
	RemoveMemberFromWorkspace(ctx context.Context, memberID, workspaceID uint) error
	AddMemberToDocument(ctx context.Context, memberID, documentID uint) error
	RemoveMemberFromDocument(ctx context.Context, memberID, documentID uint) error
}

func NewMemberRepository(
	repository *Repository,
) MemberRepository {
	return &memberRepository{
		Repository: repository,
	}
}

type memberRepository struct {
	*Repository
}

func (r *memberRepository) GetMember(ctx context.Context, id int64) (*model.Member, error) {
	var member model.Member

	return &member, nil
}

// Create New Member
func (r *memberRepository) CreateMember(ctx context.Context, member *model.Member) error {
	return r.db.Create(member).Error
}

// Update Member
func (r *memberRepository) UpdateMember(ctx context.Context, member *model.Member) error {
	return r.db.Save(member).Error
}

// Delete Member
func (r *memberRepository) DeleteMember(ctx context.Context, id int64) error {
	return r.db.Delete(&model.Member{}, id).Error
}

// Get All Members
func (r *memberRepository) GetMembers(ctx context.Context) ([]model.Member, error) {
	var members []model.Member
	return members, r.db.Find(&members).Error
}

// Get Member By User ID
func (r *memberRepository) GetMemberByUserID(ctx context.Context, userID uint) ([]model.Member, error) {
	var members []model.Member
	return members, r.db.Where("user_id = ?", userID).Find(&members).Error
}

// Get Members By Collection ID
func (r *memberRepository) GetMembersByCollectionID(ctx context.Context, collectionID uint) ([]model.Member, error) {
	var members []model.Member
	return members, r.db.Where("collection_id = ?", collectionID).Find(&members).Error
}

// Get Members By Workspace ID
func (r *memberRepository) GetMembersByWorkspaceID(ctx context.Context, workspaceID uint) ([]model.Member, error) {
	var members []model.Member
	return members, r.db.Where("workspace_id = ?", workspaceID).Find(&members).Error
}

// Get Members By Role
func (r *memberRepository) GetMembersByRole(ctx context.Context, role string) ([]model.Member, error) {
	var members []model.Member
	return members, r.db.Where("role = ?", role).Find(&members).Error
}

// RemoveMemberFromCollection
func (r *memberRepository) RemoveMemberFromCollection(ctx context.Context, memberID, collectionID uint) error {
	return r.db.Where("user_id = ? AND collection_id = ?", memberID, collectionID).Delete(&model.Member{}).Error
}

// AddMemberToWorkspace
func (r *memberRepository) AddMemberToWorkspace(ctx context.Context, memberID, workspaceID uint) error {
	return r.db.Create(&model.Member{
		UserID:      memberID,
		WorkspaceID: &workspaceID,
	}).Error
}

// RemoveMemberFromWorkspace
func (r *memberRepository) RemoveMemberFromWorkspace(ctx context.Context, memberID, workspaceID uint) error {
	return r.db.Where("user_id = ? AND workspace_id = ?", memberID, workspaceID).Delete(&model.Member{}).Error
}

// Add Member To Document
func (r *memberRepository) AddMemberToDocument(ctx context.Context, memberID, documentID uint) error {
	return r.db.Create(&model.Member{
		UserID:     memberID,
		DocumentID: &documentID,
	}).Error
}

// Remove Member From Document
func (r *memberRepository) RemoveMemberFromDocument(ctx context.Context, memberID, documentID uint) error {
	return r.db.Where("user_id = ? AND document_id = ?", memberID, documentID).Delete(&model.Member{}).Error
}
