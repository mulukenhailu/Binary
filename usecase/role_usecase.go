package usecase

import (
	"context"
	"time"

	"github.com/mulukenhailu/Binary/domain"
)

type roleUsecase struct {
	roleRepository domain.RoleRespository
	contextTimeout time.Duration
}

func NewRoleUsecase(roleRepository domain.RoleRespository, timeout time.Duration) domain.RoleUsecase {
	return &roleUsecase{
		roleRepository: roleRepository,	
		contextTimeout: timeout,
	}
}

func (ru *roleUsecase) Create(c context.Context, role *domain.CreateRoleDto) error {
	ctx, cancel := context.WithTimeout(c, ru.contextTimeout)
	defer cancel()
	return ru.roleRepository.Create(ctx, role)
}


func (ru *roleUsecase) Delete(c context.Context, roleId int32) error {
	ctx, cancel := context.WithTimeout(c, ru.contextTimeout)
	defer cancel()
	return ru.roleRepository.Delete(ctx, roleId)
}



func (ru *roleUsecase) Update(c context.Context, updateRole *domain.UpdateRoleDto) error {
	ctx, cancel := context.WithTimeout(c, ru.contextTimeout)
	defer cancel()
	return ru.roleRepository.Update(ctx, updateRole)
}

func (ru *roleUsecase) FetchRoles(c context.Context) ([]domain.Role, error) {
	ctx, cancel := context.WithTimeout(c, ru.contextTimeout)
	defer cancel()
	return ru.roleRepository.FetchRoles(ctx)
}

func (ru *roleUsecase) FetchByName(c context.Context, roleName string) (domain.Role, error) {
	ctx, cancel := context.WithTimeout(c, ru.contextTimeout)
	defer cancel()
	return ru.roleRepository.FetchByName(ctx, roleName)
}







