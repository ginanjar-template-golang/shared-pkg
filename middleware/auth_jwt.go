package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ginanjar-template-golang/shared-pkg/errors"
	"github.com/ginanjar-template-golang/shared-pkg/logger"
	"github.com/ginanjar-template-golang/shared-pkg/response"
	"github.com/ginanjar-template-golang/shared-pkg/translator"
	"github.com/golang-jwt/jwt/v5"
)

type JWTConfig struct {
	SecretKey string
}

func AuthJWT(cfg JWTConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		var t translator.Translator
		if tr, exists := c.Get("translator"); exists {
			t = tr.(translator.Translator)
		} else {
			t = translator.NewTranslator("./translator/messages/en.json")
		}

		if authHeader == "" {
			logger.Warn("Missing Authorization header", map[string]any{
				"path": c.Request.URL.Path,
			})
			response.FromInternalError(c, errors.ValueMissMatch(t, "Invalid Token", nil))
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
			response.FromInternalError(c, errors.ValueMissMatch(t, "Invalid Token", nil))
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
