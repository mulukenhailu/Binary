package domain

import "context"

type SignupResponseDto struct{
	AccessToken string `json:"access_token"`
}

type SignupUsecase interface {
	Create(c context.Context, createUserDto *CreateUserDto) error
	FetchByUserName(c context.Context, userName string) (User, error)
}