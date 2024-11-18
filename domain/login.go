package domain

import "context"

type LoginRequestDto struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type LoginResponseDto struct {
	AccessToken string `json:"access_token"`
}

type LoginUsecase interface {
	FetchByUserName(c context.Context, userName string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int)(accessToken string, err error)
}