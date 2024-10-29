package middleware

import (
	"AI-Recruitment-backend/internal/global"
	"AI-Recruitment-backend/internal/global/response"
	j "AI-Recruitment-backend/internal/pkg/jwt"
	"AI-Recruitment-backend/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get token string from header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			response.Error(c, http.StatusUnauthorized, response.CodeInvalidToken, "invalid token", "")
			c.Abort()
			return
		}
		tokenString = tokenString[7:]

		// parse token
		tokenClaims, err := jwt.ParseJwtToken(tokenString, []byte(global.Config.Jwt.Key))
		if err != nil {
			response.Error(c, http.StatusUnauthorized, response.CodeInvalidToken, "parse token failed", err.Error())
			c.Abort()
			return
		}

		// verify token
		if ok := j.VerifyJwtToken(tokenClaims); !ok {
			response.Error(c, http.StatusUnauthorized, response.CodeUnauthorized, "invalid token")
			c.Abort()
			return
		}

		c.Set("user", map[string]string{
			"id":   tokenClaims.Subject,
			"role": tokenClaims.Role.String(),
		})

		c.Next()
	}
}
