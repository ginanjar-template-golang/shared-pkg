package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ginanjar-template-golang/shared-pkg/logger"
	"github.com/ginanjar-template-golang/shared-pkg/translator"
	"github.com/ginanjar-template-golang/shared-pkg/utils"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		reqID := c.GetHeader("X-Request-ID")
		if reqID == "" {
			reqID = utils.NewRequestID()
		}
		c.Set("request_id", reqID)

		lang := c.GetHeader("Accept-Language")
		fmt.Println("Accept-Language: ", lang)
		translator.InitGlobalTranslator(lang)

		logger.Info("Request Start", map[string]any{
			"request_id": reqID,
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"client_ip":  c.ClientIP(),
		})

		c.Next()

		duration := time.Since(start)
		logger.Info("Request End", map[string]any{
			"request_id": reqID,
			"status":     c.Writer.Status(),
			"latency_ms": duration.Milliseconds(),
		})
	}
}
