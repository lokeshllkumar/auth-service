package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenToken(username string)  (string, error) {
	expTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(secretKey)
}

func ValidateToken(token string) (*Claims, error) {
	claims := &Claims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !parsedToken.Valid {
		return nil, err
	}

	return claims, nil
}