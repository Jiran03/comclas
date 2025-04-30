package mysql

import (
	"comclas/domain/repository"
	classificationrepository "comclas/internal/repository/mysql/classification"
	commentrepository "comclas/internal/repository/mysql/comment"
	rolerepository "comclas/internal/repository/mysql/role"
	userrepository "comclas/internal/repository/mysql/user"

	"gorm.io/gorm"
)

func Init(db *gorm.DB) *repository.Wrapper {
	return &repository.Wrapper{
		UserRepository:           userrepository.NewUserRepository(db),
		RoleRepository:           rolerepository.NewRoleRepository(db),
		CommentRepository:        commentrepository.NewCommentRepository(db),
		ClassificationRepository: classificationrepository.NewClassificationRepository(db),
	}
}
