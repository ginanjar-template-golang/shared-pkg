package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ginanjar-template-golang/shared-pkg/logger"
	"github.com/ginanjar-template-golang/shared-pkg/translator"
	"github.com/ginanjar-template-golang/shared-pkg/utils"
)

func sanitizeHeaders(c *gin.Context) map[string]string {
	headers := map[string]string{
		"Content-Type":  c.GetHeader("Content-Type"),
		"User-Agent":    c.GetHeader("User-Agent"),
		"Authorization": "[REDACTED]",
	}

	for _, h := range []string{"X-Forwarded-For", "X-Request-ID", "Accept-Language"} {
		if val := c.GetHeader(h); val != "" {
			headers[h] = val
		}
	}

	return headers
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		reqID := c.GetHeader("X-Request-ID")
		if reqID == "" {
			reqID = utils.NewRequestID()
		}
		c.Set("request_id", reqID)
		utils.SetRequestID(reqID)

		lang := c.GetHeader("Accept-Language")
		if lang == "" {
			lang = "en"
		}
		translator.InitGlobalTranslator(lang)

		var bodyData map[string]any
		if c.Request.Body != nil {
			bodyBytes, _ := io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // reset body
			if len(bodyBytes) > 0 {
				_ = json.Unmarshal(bodyBytes, &bodyData)
				bodyData = utils.SanitizeMap(bodyData)
			}
		}

		queryParams := map[string]any{}
		for k, v := range c.Request.URL.Query() {
			queryParams[k] = v
		}

		logger.Trace("Request Start", map[string]any{
			"request_id": reqID,
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"client_ip":  c.ClientIP(),
			"lang":       lang,
			"query":      queryParams,
			"body":       bodyData,
			"headers":    sanitizeHeaders(c),
		})

		c.Next()

		duration := time.Since(start)

		logger.Trace("Request End", map[string]any{
			"request_id": reqID,
			"status":     c.Writer.Status(),
			"latency_ms": duration.Milliseconds(),
		})
	}
}
