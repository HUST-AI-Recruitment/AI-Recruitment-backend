package jwt

import (
	"AI-Recruitment-backend/internal/config"
	"AI-Recruitment-backend/pkg/jwt"
	"time"
)

func VerifyJwtToken(claims *jwt.Claims) bool {
	if time.Now().Unix() > claims.ExpiresAt || claims.Issuer != config.C.User.Jwt.Issuer {
		return false
	}
	// TODO: add more verification
	return true
}
