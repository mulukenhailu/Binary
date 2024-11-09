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

func (r *roleUsecase) Create(c context.Context, role *domain.RoleDto) error {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()
	return r.roleRepository.Create(ctx, role)
}


func (r *roleUsecase) Delete(c context.Context, roleId int32) error {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()
	return r.roleRepository.Delete(ctx, roleId)
}



func (r *roleUsecase) Update(c context.Context, updateRole *domain.UpdateRoleDto) error {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()
	return r.roleRepository.Update(ctx, updateRole)
}

func (r *roleUsecase) FetchRoles(c context.Context) ([]domain.Role, error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()
	return r.roleRepository.FetchRoles(ctx)
}

func (r *roleUsecase) FetchByName(c context.Context, roleName string) (domain.Role, error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()
	return r.roleRepository.FetchByName(ctx, roleName)
}







