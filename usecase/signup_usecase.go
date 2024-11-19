package usecase

import (
	"context"
	"time"

	"github.com/mulukenhailu/Binary/domain"
)

type signupUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewSignupUsecase(userRepository domain.UserRepository, contextTimeout time.Duration) domain.SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
		contextTimeout: contextTimeout,
	}
}

// Create implements domain.SignupUsecase.
func (su *signupUsecase) Create(c context.Context, createUserDto *domain.CreateUserDto) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.Create(ctx, createUserDto)
}


// FetchByUserName implements domain.SignupUsecase.
func (su *signupUsecase) FetchByUserName(c context.Context, userName string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.FetchByUserName(ctx, userName)
}




