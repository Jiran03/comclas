package commenthandler

import (
	"comclas/domain/entity"
	commentrequest "comclas/internal/delivery/http/comment/request"
	customresponse "comclas/pkg/custom_response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateComment(c *gin.Context) {
	var (
		comment commentrequest.Comment
		id, err = strconv.Atoi(c.Param("id"))
	)

	if err != nil {
		customresponse.ParseResponseError(c, http.StatusBadRequest, err, "error get id")
		return
	}

	if err = c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.uc.UpdateComment(c, id, entity.CommentDTO{
		Value: comment.Value,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	customresponse.ParseResponseSuccess(c, nil)
}
