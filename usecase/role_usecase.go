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
	return r.roleRepository.Delete(c, roleId)
}



func (r *roleUsercase) Update(c context.Context, role *domain.RoleDto) error {
	return r.roleRepository.Update(c, role)
}

func (r *roleUsercase) FetchRoles(c context.Context) ([]domain.RoleDto, error) {
	return r.roleRepository.FetchRoles(c)
}

func (r *roleUsercase) FetchByName(c context.Context, roleName string) (domain.RoleDto, error) {
	return r.roleRepository.FetchByName(c, roleName)
}







