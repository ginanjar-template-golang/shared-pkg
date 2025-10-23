package http_response

import (
	"net/http"

	"github.com/gin-gonic/gin"

	httpCode "github.com/ginanjar-template-golang/shared-pkg/constants/http_code"
	errHandler "github.com/ginanjar-template-golang/shared-pkg/errors"
	"github.com/ginanjar-template-golang/shared-pkg/logger"
	"github.com/ginanjar-template-golang/shared-pkg/translator"
	"github.com/ginanjar-template-golang/shared-pkg/utils"
)

// MetaData standard response meta
type MetaData struct {
	RequestID string `json:"request_id"`
	HttpCode  int    `json:"http_code"`
	Message   string `json:"message"`
}

// Standard response format
type Response struct {
	Meta       MetaData   `json:"meta"`
	Pagination Pagination `json:"pagination,omitempty"`
	Results    any        `json:"data,omitempty"`
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
	Results  any `json:"results,omitempty"`
}

// helper untuk auto-generate request_id jika belum ada
func getRequestID(c *gin.Context) string {
	reqID := c.GetString("request_id")
	if reqID == "" {
		reqID = utils.NewRequestID()
		c.Set("request_id", reqID)
	}
	return reqID
}

// ========================
// SUCCESS RESPONSES
// ========================
func Success(c *gin.Context, messageKey string, data any) {
	reqID := getRequestID(c)

	logger.Info(translator.GetMessageByLang(messageKey), map[string]any{
		"request_id": reqID,
		"status":     httpCode.SuccessOK,
		"method":     c.Request.Method,
		"path":       c.FullPath(),
		"results":    data,
	})

	c.JSON(http.StatusOK, Response{
		Meta: MetaData{
			RequestID: reqID,
			HttpCode:  httpCode.SuccessOK,
			Message:   translator.GetMessageGlobal(messageKey),
		},
		Results: data,
	})
}

func Created(c *gin.Context, messageKey string, data any) {
	reqID := getRequestID(c)

	logger.Info(translator.GetMessageByLang(messageKey), map[string]any{
		"request_id": reqID,
		"status":     httpCode.SuccessCreated,
		"method":     c.Request.Method,
		"path":       c.FullPath(),
		"results":    data,
	})

	c.JSON(http.StatusCreated, Response{
		Meta: MetaData{
			RequestID: reqID,
			HttpCode:  httpCode.SuccessCreated,
			Message:   translator.GetMessageGlobal(messageKey),
		},
		Results: data,
	})
}

func Updated(c *gin.Context, messageKey string, data any) {
	reqID := getRequestID(c)

	logger.Info(translator.GetMessageByLang(messageKey), map[string]any{
		"request_id": reqID,
		"status":     httpCode.SuccessOK,
		"method":     c.Request.Method,
		"path":       c.FullPath(),
		"results":    data,
	})

	c.JSON(http.StatusOK, Response{
		Meta: MetaData{
			RequestID: reqID,
			HttpCode:  httpCode.SuccessOK,
			Message:   translator.GetMessageGlobal(messageKey),
		},
		Results: data,
	})
}

func Deleted(c *gin.Context, messageKey string) {
	reqID := getRequestID(c)

	logger.Info(translator.GetMessageByLang(messageKey), map[string]any{
		"request_id": reqID,
		"status":     httpCode.SuccessOK,
		"method":     c.Request.Method,
		"path":       c.FullPath(),
	})

	c.JSON(http.StatusOK, Response{
		Meta: MetaData{
			RequestID: reqID,
			HttpCode:  httpCode.SuccessOK,
			Message:   translator.GetMessageGlobal(messageKey),
		},
	})
}

// ========================
// PAGINATION RESPONSE
// ========================
func PaginationResponse(c *gin.Context, messageKey string, data Pagination) {
	reqID := getRequestID(c)

	logger.Info(translator.GetMessageByLang(messageKey), map[string]any{
		"request_id": reqID,
		"page":       data.Page,
		"limit":      data.Limit,
		"total":      data.TotalRow,
		"status":     httpCode.SuccessOK,
		"method":     c.Request.Method,
		"path":       c.FullPath(),
		"results":    data.Results,
	})

	c.JSON(http.StatusOK, Response{
		Meta: MetaData{
			RequestID: reqID,
			HttpCode:  httpCode.SuccessOK,
			Message:   translator.GetMessageGlobal(messageKey),
		},
		Pagination: Pagination{
			Page:     data.Page,
			Size:     data.Size,
			Limit:    data.Limit,
			TotalRow: data.TotalRow,
		},
		Results: data.Results,
	})
}

// ========================
// ERROR RESPONSE
// ========================
func FromAppError(c *gin.Context, err error) {
	reqID := getRequestID(c)

	if internalErr, ok := err.(errHandler.AppError); ok {
		c.JSON(internalErr.HttpCode, ResponseError{
			Meta: MetaData{
				RequestID: reqID,
				HttpCode:  internalErr.HttpCode,
				Message:   translator.GetMessageGlobal(internalErr.MessageKey),
			},
			Error: internalErr.Data,
		})
		return
	}

	// fallback
	c.JSON(http.StatusInternalServerError, ResponseError{
		Meta: MetaData{
			RequestID: reqID,
			HttpCode:  httpCode.InternalServerError,
			Message:   "Unexpected error",
		},
		Error: err.Error(),
	})
}
