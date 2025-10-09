package middleware

import (
	"fmt"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	appErrors "github.com/ginanjar-template-golang/shared-pkg/errors"
	"github.com/ginanjar-template-golang/shared-pkg/logger"
	"github.com/ginanjar-template-golang/shared-pkg/response"
	"github.com/ginanjar-template-golang/shared-pkg/translator"
	"github.com/google/uuid"
)

// Recovery menangkap panic dan kembalikan response standar
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rec := recover(); rec != nil {
				requestID := c.GetString("request_id")
				if requestID == "" {
					requestID = uuid.NewString()
					c.Set("request_id", requestID)
				}

				var t translator.Translator
				if tr, exists := c.Get("translator"); exists {
					t = tr.(translator.Translator)
				} else {
					t = translator.NewTranslator("./translator/messages/en.json")
				}

				switch err := rec.(type) {

				// ==========================
				// CASE 1: panic with InternalError
				// ==========================
				case appErrors.InternalError:
					logger.Error("Recovered InternalError", map[string]any{
						"request_id": requestID,
						"code":       err.Code,
						"message":    err.Message,
						"data":       err.Data,
						"path":       c.FullPath(),
						"method":     c.Request.Method,
					})
					response.FromInternalError(c, err)
					return

				// ==========================
				// CASE 2: panic with string
				// ==========================
				case string:
					msg := fmt.Sprintf("Panic: %s", err)
					logger.Error(msg, map[string]any{
						"request_id": requestID,
						"path":       c.FullPath(),
						"method":     c.Request.Method,
						"stack":      string(debug.Stack()),
					})
					internalErr := appErrors.UnknownError(t, "panic-string", err)
					response.FromInternalError(c, internalErr)
					return

				// ==========================
				// CASE 3: panic with general error
				// ==========================
				case error:
					logger.Error(err.Error(), map[string]any{
						"request_id": requestID,
						"path":       c.FullPath(),
						"method":     c.Request.Method,
						"stack":      string(debug.Stack()),
					})
					internalErr := appErrors.UnknownError(t, "panic-error", err.Error())
					response.FromInternalError(c, internalErr)
					return

				// ==========================
				// CASE 4: unknown panic type
				// ==========================
				default:
					logger.Error("Unknown panic", map[string]any{
						"request_id": requestID,
						"path":       c.FullPath(),
						"method":     c.Request.Method,
						"stack":      string(debug.Stack()),
					})
					internalErr := appErrors.UnknownError(t, "panic-unknown", rec)
					response.FromInternalError(c, internalErr)
					return
				}
			}
		}()

		c.Next()
	}
}
