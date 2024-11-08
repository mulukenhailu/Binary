package usecase

import (
	"context"
	"time"

	"github.com/mulukenhailu/Binary/domain"
)

type roleUsercase struct {
	roleRepository domain.RoleRespository
	contextTimeout time.Duration
}

func NewRoleUsecase(roleRepository domain.RoleRespository, timeout time.Duration) domain.RoleUsercase {
	return &roleUsercase{
		roleRepository: roleRepository,	
		contextTimeout: timeout,
	}
}

func (r *roleUsercase) Create(c context.Context, role *domain.RoleDto) error {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()
	return r.roleRepository.Create(ctx, role)
}


func (r *roleUsercase) Delete(c context.Context, roleId int32) error {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()
	return r.roleRepository.Delete(ctx, roleId)
}



func (r *roleUsercase) Update(c context.Context, updateRole *domain.UpdateRoleDto) error {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()
	return r.roleRepository.Update(ctx, updateRole)
}

func (r *roleUsercase) FetchRoles(c context.Context) ([]domain.Role, error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()
	return r.roleRepository.FetchRoles(ctx)
}

func (r *roleUsercase) FetchByName(c context.Context, roleName string) (domain.Role, error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()
	return r.roleRepository.FetchByName(ctx, roleName)
}







