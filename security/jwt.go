package security

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type customClaims struct {
	UserId int64 `json:"user_id"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte("this is the secret for signature")

func ValidateToken(tokenString string) (bool, int64) {
	token, err := jwt.ParseWithClaims(tokenString, &customClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

	if err != nil {
		return false, -1
	}
	if !token.Valid {
		return false, -1
	}
	claims, ok := token.Claims.(*customClaims)
	if !ok {
		return false, -1
	}

	return true, claims.UserId
}

func GenToken(userId int64) (string, error) {
	claims := customClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour * time.Duration(1))),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := tokenClaims.SignedString(jwtSecret)
	return tokenString, err
}
