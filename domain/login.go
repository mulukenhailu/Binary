package domain

import "context"

type LoginRequestDto struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponseDto struct {
    AccessToken string `json:"access_token"`
}

type LoginUsecase interface {
	FetchByUserName(c context.Context, userName string) (User, error)
	CreateAccessToken(user *CreateUserDto, secret string, expiry int)(accessToken string, err error)
}