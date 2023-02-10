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

func IsValidToken(tokenString string, userId int64) bool {
	token, err := jwt.ParseWithClaims(tokenString, &customClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil})

	if err != nil {
		return false
	}
	if !token.Valid {
		return false
	}
	claims, ok := token.Claims.(*customClaims)
	if !ok {
		return false
	}
	if claims.UserId != userId {
		return false
	}

	return true
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