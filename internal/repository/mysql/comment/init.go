package commentrepository

import (
	"comclas/domain/entity"
	"comclas/domain/repository"
	"fmt"

	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

// Create implements repository.CommentRepository.
func (r *commentRepository) Create(comment *entity.Comment) (int, error) {
	commentModel := DomainToModel(comment)

	// Eksekusi create dan tangkap error
	result := r.db.Create(commentModel)
	if result.Error != nil {
		return 0, result.Error
	}

	// Pastikan ID terisi
	if commentModel.ID == 0 {
		return 0, fmt.Errorf("failed to get inserted ID")
	}

	return commentModel.ID, nil
}

// Delete implements repository.CommentRepository.
func (r *commentRepository) Delete(id int) error {
	return r.db.Delete(&Comment{}, "id = ?", id).Error
}

// GetAll implements repository.CommentRepository.
func (r *commentRepository) GetAll() ([]*entity.Comment, error) {
	var (
		commentModels []Comment
		comments      []*entity.Comment
	)
	err := r.db.Find(&commentModels).Error
	if err != nil {
		return nil, err
	}

	for _, v := range commentModels {
		comment, err := ModelToDomain(v)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

// GetByID implements repository.CommentRepository.
func (r *commentRepository) GetByID(id int) (*entity.Comment, error) {
	var comment Comment
	err := r.db.First(&comment, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return ModelToDomain(comment)
}

// Update implements repository.CommentRepository.
func (r *commentRepository) Update(comment *entity.Comment) error {
	commentModel := DomainToModel(comment)
	return r.db.Save(commentModel).Error
}

func NewCommentRepository(db *gorm.DB) repository.CommentRepository {
	return &commentRepository{db}
}
