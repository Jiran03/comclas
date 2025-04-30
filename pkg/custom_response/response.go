package customresponse

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

type ResponseWrapper struct {
	Status struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"status"`
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta,omitempty"`
}

func ParseResponseSuccess(c *gin.Context, data interface{}) {
	fmt.Println("data:", data)

	// Create the response wrapper
	response := ResponseWrapper{
		Status: struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		}{
			Code:    "0", // 0 = success
			Message: "success",
		},
		Data: data,
	}

	// Parse JSON from request body
	if _, err := json.Marshal(&data); err != nil && data != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": gin.H{
				"code":    "1",
				"message": fmt.Sprintf("invalid JSON: %v", err),
			},
		})
		return
	}

	// Respond with the wrapped JSON
	c.JSON(http.StatusOK, response)
}

func ParseResponseError(c *gin.Context, httpCode int, err error, message string) {
	messageErr := message
	if err != nil {
		messageErr += (", " + err.Error())
	}

	log.Println(messageErr)

	// Create the response wrapper
	response := ResponseWrapper{
		Status: struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		}{
			Code:    "1",
			Message: message,
		},
		Data: nil,
	}

	// Respond with the wrapped JSON
	c.JSON(httpCode, response)
}
