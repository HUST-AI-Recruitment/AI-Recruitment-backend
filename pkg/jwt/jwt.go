package jwt

import (
	"AI-Recruitment-backend/pkg/common"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Role common.Role `json:"role"`
	jwt.StandardClaims
}

func GenerateJwtToken(id string, role common.Role, expire int64, issuer string) *jwt.Token {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(expire) * time.Second)

	claims := &Claims{
		Role: role,
		StandardClaims: jwt.StandardClaims{
			Subject:   id,
			IssuedAt:  nowTime.Unix(),
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token
}

func GenerateJwtTokenString(token *jwt.Token, jwtKey []byte) (string, error) {
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}

func ParseJwtToken(tokenString string, jwtKey []byte) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if token == nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
