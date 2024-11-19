package usecase

import (
	"context"
	"time"

	"github.com/mulukenhailu/Binary/domain"
)

type userManagementUsecase struct {
	userRepository domain.UserRepository
	contexTimeout  time.Duration
}


func NewUserManagerUsecase(userRepository domain.UserRepository, contexTimeout time.Duration) domain.UserManagementUsecase {
	return &userManagementUsecase{
		userRepository: userRepository,
		contexTimeout:  contexTimeout,
	}
}

// Update implements domain.UserManagerUsecase.
func (umu *userManagementUsecase) Update(c context.Context, updateUserDto *domain.UpdateUserDto) error {
	ctx, cancel := context.WithTimeout(c, umu.contexTimeout)
	defer cancel()
	return umu.userRepository.Update(ctx, updateUserDto)
}

// Delete implements domain.UserManagerUsecase.
func (umu *userManagementUsecase) Delete(c context.Context, userId int32) error {
	ctx, cancel := context.WithTimeout(c, umu.contexTimeout)
	defer cancel()
	return umu.userRepository.Delete(ctx, userId)
}

// FetchByRoleName implements domain.UserManagerUsecase.
func (umu *userManagementUsecase) FetchByRoleId(c context.Context, roleId int32) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(c, umu.contexTimeout)
	defer cancel()
	return umu.userRepository.FetchByRoleId(ctx, roleId)
}

// FetchByUserName implements domain.UserManagerUsecase.
func (umu *userManagementUsecase) FetchByUserName(c context.Context, userName string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, umu.contexTimeout)
	defer cancel()
	return  umu.userRepository.FetchByUserName(ctx, userName)
}

// FetchUsers implements domain.UserManagerUsecase.
func (umu *userManagementUsecase) FetchUsers(c context.Context) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(c, umu.contexTimeout)
	defer cancel()
	return umu.userRepository.FetchUsers(ctx)
}



