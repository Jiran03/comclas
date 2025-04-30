package commentresponse

import (
	"comclas/domain/entity"
)

type Comment struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

func ToCommentResponse(data *entity.Comment) *Comment {
	commentResponse := &Comment{
		ID:    data.GetID(),
		Value: data.GetValue(),
	}

	return commentResponse
}

func ToListCommentResponse(data []*entity.Comment) []*Comment {
	var commentResponse []*Comment
	for _, val := range data {
		comment := ToCommentResponse(val)

		commentResponse = append(commentResponse, comment)
	}

	return commentResponse
}

type Classification struct {
	Comment        string `json:"comment"`
	Classification string `json:"classification"`
}

func ToClassificationResponse(data *entity.Classification) *Classification {
	var (
		commentValue string
	)

	if data.GetComment() != nil {
		commentValue = data.GetComment().GetValue()
	}

	classificationResponse := &Classification{
		Comment:        commentValue,
		Classification: data.GetType(),
	}

	return classificationResponse
}
