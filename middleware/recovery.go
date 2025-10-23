package middleware

import (
	"github.com/gin-gonic/gin"
	appError "github.com/ginanjar-template-golang/shared-pkg/errors"
	httpResponse "github.com/ginanjar-template-golang/shared-pkg/response/http_response"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rec := recover(); rec != nil {
				var internalErr appError.AppError

				switch e := rec.(type) {
				case string:
					internalErr = appError.UnknownError("panic-string", e)
				case error:
					internalErr = appError.UnknownError("panic-error", e.Error())
				default:
					internalErr = appError.UnknownError("panic-unknown", e)
				}

				httpResponse.FromAppError(c, internalErr)
				c.Abort()
			}
		}()

		c.Next()
	}
}
