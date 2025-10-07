package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    "SUCCESS",
		Message: message,
		Data:    data,
	})
}

func Error(c *gin.Context, code string, message string, statusCode int) {
	c.JSON(statusCode, Response{
		Code:    code,
		Message: message,
	})
}
