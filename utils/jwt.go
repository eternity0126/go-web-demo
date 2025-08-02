package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"time"
)

type JwtCustomClaims struct {
	ID   uint
	Name string
	jwt.RegisteredClaims
}

var stSigningKey = []byte(viper.GetString("jwt.signing_key"))

func GenerateToken(id uint, name string) (string, error) {
	iJwtCustomClaims := JwtCustomClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(viper.GetDuration("jwt.token_expire") * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "Token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, iJwtCustomClaims)
	return token.SignedString(stSigningKey)
}

func ParseToken(tokenStr string) (JwtCustomClaims, error) {
	iJwtCustomClaims := JwtCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, &iJwtCustomClaims, func(token *jwt.Token) (interface{}, error) {
		return stSigningKey, nil
	})

	if err == nil && !token.Valid {
		err = errors.New("Invalid token")
	}

	return iJwtCustomClaims, err
}

//func IsTokenValid(tokenStr string) bool {
//	_, err := ParseToken(tokenStr)
//	if err != nil {
//		return false
//	}
//	return true
//}
