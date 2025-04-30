package commenthandler

import (
	commentresponse "comclas/internal/delivery/http/comment/response"
	customresponse "comclas/pkg/custom_response"
	middlewarepkg "comclas/pkg/middleware"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) FetchCommentByID(c *gin.Context) {
	var (
		_, roleExists = c.Get(middlewarepkg.UserRoleKey)
		id, err       = strconv.Atoi(c.Param("id"))
	)

	if err != nil {
		customresponse.ParseResponseError(c, http.StatusBadRequest, err, "error get id")
		return
	}

	if !roleExists {
		customresponse.ParseResponseError(c, http.StatusBadRequest, nil, InvalidRoleMsg)
		return
	}

	comment, err := h.uc.FetchCommentByID(c, id)
	if err != nil {
		customresponse.ParseResponseError(c, http.StatusInternalServerError, err, "error get all comment")
		return
	}

	commentResponse := commentresponse.ToCommentResponse(comment)

	customresponse.ParseResponseSuccess(c, commentResponse)
}
