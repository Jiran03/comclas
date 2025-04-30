package entity

type Comment struct {
	id        int
	value     string
	createdby string
}

type CommentDTO struct {
	ID        int
	Value     string
	CreatedBy string
}

func NewComment(dto CommentDTO) (*Comment, error) {
	comment := &Comment{
		id:        dto.ID,
		value:     dto.Value,
		createdby: dto.CreatedBy,
	}

	return comment, nil
}

func (i *Comment) GetID() int {
	return i.id
}

func (i *Comment) GetValue() string {
	return i.value
}

func (i *Comment) GetCreatedBy() string {
	return i.createdby
}
