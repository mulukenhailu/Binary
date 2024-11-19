package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mulukenhailu/Binary/domain"
)

func CreateAccessToken(user *domain.CreateUserDto, secret string, expiry int) (accessToken string, err error) {

	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()
	claims := &domain.JwtCustomClaim{
		UserName			:user.UserName,
		RoleId				:user.RoleId,
		RegisteredClaims	:jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(exp, 0)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil{
		return "", err
	}
	return signedToken, nil
}