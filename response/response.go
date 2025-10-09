package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/ginanjar-template-golang/shared-pkg/constants"
	"github.com/ginanjar-template-golang/shared-pkg/errors"
	"github.com/ginanjar-template-golang/shared-pkg/logger"
)

// MetaData standard response meta
type MetaData struct {
	RequestID string `json:"request_id"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
}

// Standard response format
type Response struct {
	Meta MetaData `json:"meta"`
	Data any      `json:"data,omitempty"`
}

type ResponseError struct {
	Meta  MetaData `json:"meta"`
	Error any      `json:"error,omitempty"`
}

// Pagination response format
type Pagination struct {
	Page     int `json:"page"`
	Size     int `json:"size"`
	Limit    int `json:"limit"`
	TotalRow int `json:"total_row"`
	Results  any `json:"results"`
}

// helper untuk auto-generate request_id jika belum ada
func getRequestID(c *gin.Context) string {
	reqID := c.GetString("request_id")
	if reqID == "" {
		reqID = uuid.NewString()
		c.Set("request_id", reqID)
	}
	return reqID
}

// ========================
// SUCCESS RESPONSES
// ========================
func Success(c *gin.Context, message string, data any) {
	reqID := getRequestID(c)

	// log otomatis
	logger.Info(message, map[string]any{
		"request_id": reqID,
		"status":     http.StatusOK,
		"method":     c.Request.Method,
		"path":       c.FullPath(),
	})

	c.JSON(http.StatusOK, Response{
		Meta: MetaData{
			RequestID: reqID,
			Code:      constants.SuccessOK,
			Message:   message,
		},
		Data: data,
	})
}

func Created(c *gin.Context, message string, data any) {
	reqID := getRequestID(c)

	logger.Info(message, map[string]any{
		"request_id": reqID,
		"status":     http.StatusCreated,
		"method":     c.Request.Method,
		"path":       c.FullPath(),
	})

	c.JSON(http.StatusCreated, Response{
		Meta: MetaData{
			RequestID: reqID,
			Code:      constants.SuccessCreated,
			Message:   message,
		},
		Data: data,
	})
}

func Updated(c *gin.Context, message string, data any) {
	reqID := getRequestID(c)

	logger.Info(message, map[string]any{
		"request_id": reqID,
		"status":     http.StatusOK,
		"method":     c.Request.Method,
		"path":       c.FullPath(),
	})

	c.JSON(http.StatusOK, Response{
		Meta: MetaData{
			RequestID: reqID,
			Code:      constants.SuccessOK,
			Message:   message,
		},
		Data: data,
	})
}

func Deleted(c *gin.Context, message string) {
	reqID := getRequestID(c)

	logger.Info(message, map[string]any{
		"request_id": reqID,
		"status":     http.StatusOK,
		"method":     c.Request.Method,
		"path":       c.FullPath(),
	})

	c.JSON(http.StatusOK, Response{
		Meta: MetaData{
			RequestID: reqID,
			Code:      constants.SuccessOK,
			Message:   message,
		},
	})
}

// ========================
// PAGINATION RESPONSE
// ========================
func PaginationResponse(c *gin.Context, message string, data Pagination) {
	reqID := getRequestID(c)

	logger.Info(message, map[string]any{
		"request_id": reqID,
		"page":       data.Page,
		"limit":      data.Limit,
		"total":      data.TotalRow,
		"status":     http.StatusOK,
		"method":     c.Request.Method,
		"path":       c.FullPath(),
	})

	c.JSON(http.StatusOK, Response{
		Meta: MetaData{
			RequestID: reqID,
			Code:      constants.SuccessOK,
			Message:   message,
		},
		Data: data,
	})
}

// ========================
// ERROR RESPONSE
// ========================
func FromInternalError(c *gin.Context, err errors.InternalError) {
	reqID := getRequestID(c)

	c.JSON(err.Code, ResponseError{
		Meta: MetaData{
			RequestID: reqID,
			Code:      err.Code,
			Message:   err.Message,
		},
		Error: err.Data,
	})
}
