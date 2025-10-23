package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ginanjar-template-golang/shared-pkg/errors"
	"github.com/ginanjar-template-golang/shared-pkg/logger"
	httpResponse "github.com/ginanjar-template-golang/shared-pkg/response/http_response"
	"github.com/golang-jwt/jwt/v5"
)

type JWTConfig struct {
	SecretKey string
}

func AuthJWT(cfg JWTConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			logger.Warn("Missing Authorization header", map[string]any{
				"path": c.Request.URL.Path,
			})

			httpResponse.FromAppError(c, errors.ValueMissMatch("Missing Authorization header", nil))
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			return []byte(cfg.SecretKey), nil
		})

		if err != nil || !token.Valid {
			logger.Warn("Invalid JWT token", map[string]any{
				"error": err,
				"path":  c.Request.URL.Path,
			})

			httpResponse.FromAppError(c, errors.InvalidTokenError("Token", err))
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			c.Set("user_claims", claims)
		}

		c.Next()
	}
}
