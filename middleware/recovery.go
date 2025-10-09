package middleware

import (
	"runtime/debug"

	"github.com/gin-gonic/gin"
	appErrors "github.com/ginanjar-template-golang/shared-pkg/errors"
	"github.com/ginanjar-template-golang/shared-pkg/logger"
	"github.com/ginanjar-template-golang/shared-pkg/response"
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

				var internalErr appErrors.InternalError

				switch e := rec.(type) {
				case string:
					internalErr = appErrors.UnknownError("panic-string", e)
				case error:
					internalErr = appErrors.UnknownError("panic-error", e.Error())
				default:
					internalErr = appErrors.UnknownError("panic-unknown", e)
				}

				logger.Error("Recovered panic", map[string]any{
					"request_id": requestID,
					"error":      internalErr.Message,
					"code":       internalErr.Code,
					"path":       c.FullPath(),
					"method":     c.Request.Method,
					"stack":      string(debug.Stack()),
				})

				response.FromInternalError(c, internalErr)
			}
		}()

		c.Next()
	}
}
