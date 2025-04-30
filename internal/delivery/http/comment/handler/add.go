package commenthandler

import (
	"comclas/domain/entity"
	commentrequest "comclas/internal/delivery/http/comment/request"
	commentresponse "comclas/internal/delivery/http/comment/response"
	customresponse "comclas/pkg/custom_response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCommentClassification(c *gin.Context) {
	var (
		comment commentrequest.Comment
	)
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	commentClass, err := h.uc.CommentClassification(c, entity.CommentDTO{
		Value: comment.Value,
	})
	if err != nil {
		customresponse.ParseResponseError(c, http.StatusInternalServerError, err, "error add comment")
		return
	}

	commentClassResponse := commentresponse.ToClassificationResponse(commentClass)

	customresponse.ParseResponseSuccess(c, commentClassResponse)
}
