package server

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/kamakuni/rails-jwt/auth/app/uuid"
)

func CreateAccessToken(userId string, now time.Time, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": "123456789",
		"exp": now.Add(10 * time.Minute).Unix(),
		"iat": now.Unix(),
	})
	return token.SignedString([]byte(secret))
}

func CreateRefreshToken() (string, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return "", nil
	}
	return uuid.String(), nil
}
