package repository

import (
	"context"
	"core_service/internal/model"
)

type CommentRepository interface {
	GetComment(ctx context.Context, id string) (*model.Comment, error)
	CreateComment(ctx context.Context, comment *model.Comment) error
	UpdateComment(ctx context.Context, comment *model.Comment) error
	DeleteComment(ctx context.Context, id string) error
	GetComments(ctx context.Context) ([]model.Comment, error)
	GetCommentByDocumentID(ctx context.Context, documentID string) ([]model.Comment, error)
	GetCommentByUserID(ctx context.Context, userID string) ([]model.Comment, error)
	GetCommentByParentCommentID(ctx context.Context, parentCommentID string) ([]model.Comment, error)
	GetCommentByDocumentIDAndParentCommentID(ctx context.Context, documentID, parentCommentID string) ([]model.Comment, error)
}

func NewCommentRepository(
	repository *Repository,
) CommentRepository {
	return &commentRepository{
		Repository: repository,
	}
}

type commentRepository struct {
	*Repository
}

func (r *commentRepository) GetComment(ctx context.Context, id string) (*model.Comment, error) {
	var comment model.Comment

	return &comment, nil
}

// Create New Comment
func (r *commentRepository) CreateComment(ctx context.Context, comment *model.Comment) error {
	return r.db.Create(comment).Error
}

// Update Comment
func (r *commentRepository) UpdateComment(ctx context.Context, comment *model.Comment) error {
	return r.db.Save(comment).Error
}

// Delete Comment
func (r *commentRepository) DeleteComment(ctx context.Context, id string) error {
	return r.db.Delete(&model.Comment{}, id).Error
}

// Get All Comments return all comments using gorm
func (r *commentRepository) GetComments(ctx context.Context) ([]model.Comment, error) {
	var comments []model.Comment
	return comments, r.db.Find(&comments).Error
}

// Get Comment By Document ID
func (r *commentRepository) GetCommentByDocumentID(ctx context.Context, documentID string) ([]model.Comment, error) {
	var comments []model.Comment
	return comments, r.db.Where("document_id = ?", documentID).Find(&comments).Error
}

// Get Comment By User ID
func (r *commentRepository) GetCommentByUserID(ctx context.Context, userID string) ([]model.Comment, error) {
	var comments []model.Comment
	return comments, r.db.Where("user_id = ?", userID).Find(&comments).Error
}

// Get Comment By Parent Comment ID
func (r *commentRepository) GetCommentByParentCommentID(ctx context.Context, parentCommentID string) ([]model.Comment, error) {
	var comments []model.Comment
	return comments, r.db.Where("parent_comment_id = ?", parentCommentID).Find(&comments).Error
}

// Get Comment By Document ID and Parent Comment ID
func (r *commentRepository) GetCommentByDocumentIDAndParentCommentID(ctx context.Context, documentID, parentCommentID string) ([]model.Comment, error) {
	var comments []model.Comment
	return comments, r.db.Where("document_id = ? AND parent_comment_id = ?", documentID, parentCommentID).Find(&comments).Error
}
