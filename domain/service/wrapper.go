package service

type Wrapper struct {
	UserService           UserService
	RoleService           RoleService
	CommentService        CommentService
	ClassificationService ClassificationService
}
