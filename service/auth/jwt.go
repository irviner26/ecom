package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/irviner26/ecom/config"
)

func CreateJWT(secret []byte, userID int) (string, error) {
	expiration := time.Second * time.Duration(config.Global.JWTExpiration)

	tkn :=  jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userID": strconv.Itoa(userID),
			"expiredAt": time.Now().Add(expiration).Unix(),
		},
	)

	tokenString, err := tkn.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, err
}