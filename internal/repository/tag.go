package repository

import (
	"context"
	"core_service/internal/model"
)

type TagRepository interface {
	GetTag(ctx context.Context, id string) (*model.Tag, error)
	CreateTag(ctx context.Context, tag *model.Tag) error
	UpdateTag(ctx context.Context, tag *model.Tag) error
	DeleteTag(ctx context.Context, id string) error
	GetTags(ctx context.Context) ([]model.Tag, error)
	GetTagByDocumentID(ctx context.Context, documentID string) ([]model.Tag, error)
	GetTagByCollectionID(ctx context.Context, collectionID string) ([]model.Tag, error)

	AssignTagToDocument(ctx context.Context, tagID, documentID string) error
	UnassignTagFromDocument(ctx context.Context, tagID, documentID string) error
	GetTagsByDocumentID(ctx context.Context, documentID string) ([]model.Tag, error)

	AssignTagToCollection(ctx context.Context, tagID, collectionID string) error
	UnassignTagFromCollection(ctx context.Context, tagID, collectionID string) error
	GetTagsByCollectionID(ctx context.Context, collectionID string) ([]model.Tag, error)
}

func NewTagRepository(
	repository *Repository,
) TagRepository {
	return &tagRepository{
		Repository: repository,
	}
}

type tagRepository struct {
	*Repository
}

func (r *tagRepository) GetTag(ctx context.Context, id string) (*model.Tag, error) {
	var tag model.Tag

	return &tag, nil
}

// Create New Tag
func (r *tagRepository) CreateTag(ctx context.Context, tag *model.Tag) error {
	return r.db.Create(tag).Error
}

// Update Tag
func (r *tagRepository) UpdateTag(ctx context.Context, tag *model.Tag) error {
	return r.db.Save(tag).Error
}

// Delete Tag
func (r *tagRepository) DeleteTag(ctx context.Context, id string) error {
	return r.db.Delete(&model.Tag{}, id).Error
}

// Get All Tags
func (r *tagRepository) GetTags(ctx context.Context) ([]model.Tag, error) {
	var tags []model.Tag
	return tags, r.db.Find(&tags).Error
}

// Get Tag By Document ID
func (r *tagRepository) GetTagByDocumentID(ctx context.Context, documentID string) ([]model.Tag, error) {
	var tags []model.Tag
	return tags, r.db.Where("document_id = ?", documentID).Find(&tags).Error
}

// Get Tags By Collection ID
func (r *tagRepository) GetTagByCollectionID(ctx context.Context, collectionID string) ([]model.Tag, error) {
	var tags []model.Tag
	return tags, r.db.Where("collection_id = ?", collectionID).Find(&tags).Error
}

// AssignTagToDocument assign tag to document
func (r *tagRepository) AssignTagToDocument(ctx context.Context, tagID, documentID string) error {
	return r.db.Exec("INSERT INTO document_tag (document_id, tag_id) VALUES (?, ?)", documentID, tagID).Error
}

// UnassignTagFromDocument unassign tag from document
func (r *tagRepository) UnassignTagFromDocument(ctx context.Context, tagID, documentID string) error {
	return r.db.Exec("DELETE FROM document_tag WHERE document_id = ? AND tag_id = ?", documentID, tagID).Error
}

// GetTagsByDocumentID get tags by document id
func (r *tagRepository) GetTagsByDocumentID(ctx context.Context, documentID string) ([]model.Tag, error) {
	var tags []model.Tag
	return tags, r.db.Raw("SELECT t.* FROM tag t JOIN document_tag dt ON t.id = dt.tag_id WHERE dt.document_id = ?", documentID).Scan(&tags).Error
}

// AssignTagToCollection assign tag to collection
func (r *tagRepository) AssignTagToCollection(ctx context.Context, tagID, collectionID string) error {
	return r.db.Exec("INSERT INTO collection_tag (collection_id, tag_id) VALUES (?, ?)", collectionID, tagID).Error
}

// UnassignTagFromCollection unassign tag from collection
func (r *tagRepository) UnassignTagFromCollection(ctx context.Context, tagID, collectionID string) error {
	return r.db.Exec("DELETE FROM collection_tag WHERE collection_id = ? AND tag_id = ?", collectionID, tagID).Error
}

// GetTagsByCollectionID get tags by collection id
func (r *tagRepository) GetTagsByCollectionID(ctx context.Context, collectionID string) ([]model.Tag, error) {
	var tags []model.Tag
	return tags, r.db.Raw("SELECT t.* FROM tag t JOIN collection_tag ct ON t.id = ct.tag_id WHERE ct.collection_id = ?", collectionID).Scan(&tags).Error
}
