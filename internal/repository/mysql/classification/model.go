package classificationrepository

import "time"

type Classification struct {
	ID        int
	CommentID int
	Type      string
	CreatedBy string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func (m Classification) GetTableName() string {
	return `classifications`
}
