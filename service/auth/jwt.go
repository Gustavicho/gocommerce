package auth

import (
	"time"

	"github.com/Gustavicho/gocommerce/config"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(secret []byte, userID uint) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JwtExpInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   userID,
		"expireAt": time.Now().Add(expiration).Unix(),
	})

	tokeString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokeString, nil
}
