package domain

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaim struct {
	UserName string `json:"user_name"`
	UserId   int32  `json:"user_id"`
	RoleId int32 `json:"role_id"`
	jwt.RegisteredClaims
}