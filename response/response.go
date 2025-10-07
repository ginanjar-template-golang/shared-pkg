package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ginanjar-template-golang/shared-pkg/errors"
	"github.com/google/uuid"
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
	Page      int    `json:"page"`
	Size      int    `json:"size"`
	Limit     int    `json:"limit"`
	TotalRow  int    `json:"total_row"`
	Results   any    `json:"results"`
	RequestID string `json:"request_id"`
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
	c.JSON(http.StatusOK, Response{
		Meta: MetaData{
			RequestID: getRequestID(c),
			Code:      http.StatusOK,
			Message:   message,
		},
		Data: data,
	})
}

func Created(c *gin.Context, message string, data any) {
	c.JSON(http.StatusCreated, Response{
		Meta: MetaData{
			RequestID: getRequestID(c),
			Code:      http.StatusCreated,
			Message:   message,
		},
		Data: data,
	})
}

func Updated(c *gin.Context, message string, data any) {
	c.JSON(http.StatusOK, Response{
		Meta: MetaData{
			RequestID: getRequestID(c),
			Code:      http.StatusOK,
			Message:   message,
		},
		Data: data,
	})
}

func Deleted(c *gin.Context, message string) {
	c.JSON(http.StatusOK, Response{
		Meta: MetaData{
			RequestID: getRequestID(c),
			Code:      http.StatusOK,
			Message:   message,
		},
	})
}

// ========================
// PAGINATION RESPONSE
// ========================
func PaginationResponse(c *gin.Context, page, size, totalRow int, results any) {
	c.JSON(http.StatusOK, Pagination{
		Page:      page,
		Size:      size,
		Limit:     size,
		TotalRow:  totalRow,
		Results:   results,
		RequestID: getRequestID(c),
	})
}

// ========================
// ERROR RESPONSE
// ========================
func FromInternalError(c *gin.Context, err errors.InternalError) {
	c.JSON(err.Code, ResponseError{
		Meta: MetaData{
			RequestID: getRequestID(c),
			Code:      err.Code,
			Message:   err.Message,
		},
		Error: err.Data,
	})
}
