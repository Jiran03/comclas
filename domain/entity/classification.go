package entity

type ClassificationType string

const (
	Neutral  ClassificationType = "neutral"
	Positive ClassificationType = "positive"
	Negative ClassificationType = "negative"
)

type Classification struct {
	id                 int
	commentID          int
	comment            *Comment
	classificationType ClassificationType
	createdby          string
}

type ClassificationDTO struct {
	ID                 int
	CommentID          int
	ClassificationType ClassificationType
	CreatedBy          string
}

func NewClassification(dto ClassificationDTO) (*Classification, error) {
	Classification := &Classification{
		id:                 dto.ID,
		commentID:          dto.CommentID,
		classificationType: dto.ClassificationType,
		createdby:          dto.CreatedBy,
	}

	return Classification, nil
}

func (i *Classification) GetID() int {
	return i.id
}

func (i *Classification) GetCommentID() int {
	return i.commentID
}

func (i *Classification) GetComment() *Comment {
	return i.comment
}

func (i *Classification) SetCommment(comment *Comment) {
	i.comment = comment
}

func (i *Classification) GetType() string {
	return string(i.classificationType)
}

func (i *Classification) GetCreatedBy() string {
	return i.createdby
}
