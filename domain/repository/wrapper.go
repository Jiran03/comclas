package repository

type Wrapper struct {
	UserRepository           UserRepository
	RoleRepository           RoleRepository
	CommentRepository        CommentRepository
	ClassificationRepository ClassificationRepository
}
