package commentrepository

import "time"

type Comment struct {
	ID        int        `gorm:"primaryKey;autoIncrement"`
	Value     string     `gorm:"type:text"`
	CreatedBy string     `gorm:"type:varchar(100)"`
	CreatedAt *time.Time `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
	DeletedAt *time.Time `gorm:"index"`
}

func (m Comment) GetTableName() string {
	return `comments`
}
