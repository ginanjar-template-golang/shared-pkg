package middleware

import (
	"github.com/gin-gonic/gin"
	errHandler "github.com/ginanjar-template-golang/shared-pkg/errors"
	"github.com/ginanjar-template-golang/shared-pkg/response"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rec := recover(); rec != nil {
				var internalErr errHandler.AppError

				switch e := rec.(type) {
				case string:
					internalErr = errHandler.UnknownError("panic-string", e)
				case error:
					internalErr = errHandler.UnknownError("panic-error", e.Error())
				default:
					internalErr = errHandler.UnknownError("panic-unknown", e)
				}

				response.FromInternalError(c, internalErr)
				c.Abort()
			}
		}()

		c.Next()
	}
}
