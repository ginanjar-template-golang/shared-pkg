package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ginanjar-template-golang/shared-pkg/utils"
)

// Middleware untuk inject request_id dan language ke context
func RequestContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = utils.GenerateRequestID()
		}
		lang := c.GetHeader("Accept-Language")
		if lang == "" {
			lang = "id"
		}

		c.Set("request_id", requestID)
		c.Set("lang", lang)

		c.Writer.Header().Set("X-Request-ID", requestID)
		c.Next()
	}
}
