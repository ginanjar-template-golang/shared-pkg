package response

import (
	"github.com/gin-gonic/gin"
	"github.com/ginanjar-template-golang/shared-pkg/errors"
	"github.com/ginanjar-template-golang/shared-pkg/translator"
)

type SuccessResponse struct {
	Meta MetaData `json:"meta_data"`
	Data any      `json:"data"`
}

type ErrorResponse struct {
	Meta   MetaData `json:"meta_data"`
	Errors any      `json:"errors"`
}

// Format sukses
func NewSuccessResponse(requestID, code, lang string, results any) SuccessResponse {
	message := translator.T(lang, code)
	return SuccessResponse{
		Meta: MetaData{
			Status:    "success",
			RequestID: requestID,
			Code:      code,
			Message:   message,
		},
		Data: results,
	}
}

// Format error
func NewErrorResponse(requestID string, err error, lang string) ErrorResponse {
	appErr := errors.ParseError(err)
	message := translator.T(lang, appErr.Code)
	return ErrorResponse{
		Meta: MetaData{
			Status:    "error",
			RequestID: requestID,
			Code:      appErr.Code,
			Message:   message,
		},
		Errors: appErr.Detail,
	}
}

// Kirim JSON pakai gin.Context
func WriteJSON(c *gin.Context, status int, payload any) {
	c.JSON(status, payload)
}
