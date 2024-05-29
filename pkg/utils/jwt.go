package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte("MyVeryVerySecretKey")

type JWTClaims struct {
	UserId   int    `json:"user_id"`
	UserType string `json:"user_type"`
	jwt.StandardClaims
}

func GenerateJWT(userId int, userType string) (string, error) {
	claims := &JWTClaims{
		UserId:   userId,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 24).Unix(), // 24 Hours
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateJWT(jwtToken string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(
		jwtToken,
		&JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)

	if !ok {
		return nil, errors.New("invalid JWT claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("token expired")
	}
	return claims, err
}
