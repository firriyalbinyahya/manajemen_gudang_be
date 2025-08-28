package auth

import (
	"fmt"
	"manajemen_gudang_be/config"
	"manajemen_gudang_be/entity"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaim struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(user *entity.User) (string, *time.Time, error) {
	var (
		secret string
		expiry int
	)

	secret = config.Config.JWT.AccessSecret
	expiry = config.Config.JWT.AccessExpiryInSec

	expiredAt := time.Now().Add(time.Second * time.Duration(expiry))
	claims := &JWTClaim{
		ID:       fmt.Sprint(user.ID),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "firriyal-bin-yahya",
		},
	}

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	if err != nil {
		return "", nil, err
	}

	return tokenString, &expiredAt, nil
}

func VerifyToken(tokenString string) (*entity.User, error) {
	secret := config.Config.JWT.AccessSecret

	token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaim); ok && token.Valid {
		user := &entity.User{}
		idInt, _ := strconv.ParseUint(claims.ID, 10, 64)
		user.ID = idInt
		user.Username = claims.Username

		return user, nil
	}

	return nil, err
}
