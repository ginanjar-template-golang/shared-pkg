package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ginanjar-template-golang/shared-pkg/logger"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logger.Error("🔥 Panic recovered", map[string]any{
					"error": r,
					"path":  c.Request.URL.Path,
				})
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"meta": gin.H{
						"code":    500,
						"message": "Internal Server Error",
					},
					"error": r,
				})
			}
		}()
		c.Next()
	}
}
