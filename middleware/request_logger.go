package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ginanjar-template-golang/shared-pkg/logger"
	"github.com/google/uuid"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		reqID := c.GetHeader("X-Request-ID")
		if reqID == "" {
			reqID = uuid.NewString()
		}
		c.Set("request_id", reqID)

		logger.Info("📥 Request Start", map[string]any{
			"request_id": reqID,
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"client_ip":  c.ClientIP(),
		})

		c.Next()

		duration := time.Since(start)
		logger.Info("📤 Request End", map[string]any{
			"request_id": reqID,
			"status":     c.Writer.Status(),
			"latency_ms": duration.Milliseconds(),
		})
	}
}
