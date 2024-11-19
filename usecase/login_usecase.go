package usecase

import (
	"context"
	"time"

	"github.com/mulukenhailu/Binary/domain"
	"github.com/mulukenhailu/Binary/internal/utils"
)

type loginUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewLoginUsecase(userRepository domain.UserRepository, contextTimeout time.Duration) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: contextTimeout,
	}
}

// FetchByUserName implements domain.LoginUsecase.
func (lu *loginUsecase) FetchByUserName(c context.Context, userName string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.FetchByUserName(ctx, userName)
}

// CreateAccessToken implements domain.LoginUsecase.
func (lu *loginUsecase) CreateAccessToken(user *domain.CreateUserDto, secret string, expiry int) (accessToken string, err error) {
	return utils.CreateAccessToken(user, secret, expiry)
}




