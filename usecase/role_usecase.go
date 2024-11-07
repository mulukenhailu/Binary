package usecase

import (
	"context"
	"github.com/mulukenhailu/Binary/domain"
)

type roleUsercase struct {
	roleRepository domain.RoleRespository
}

func NewRoleUsecase(roleRepository domain.RoleRespository) domain.RoleUsercase {
	return &roleUsercase{
		roleRepository: roleRepository,
	}
}


func (r *roleUsercase) Create(c context.Context, role *domain.RoleDto) error {
	return r.roleRepository.Create(c, role)
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







