package usecase

import (
	"context"
	"time"

	"github.com/mulukenhailu/Binary/domain"
)

type rolePermission struct {
	rolePermissionRepository domain.RolePermissionRepository
	contextTimeout time.Duration
}


func NewRolePermissionUsecase(rolePermissionRepository domain.RolePermissionRepository, timeout time.Duration) domain.RolePermissionUsecase {
	return &rolePermission{
		rolePermissionRepository: rolePermissionRepository,
		contextTimeout: timeout,
	}
}


// Create implements domain.RolePermissionUsecase.
func (rpu *rolePermission) Create(c context.Context, CreatePermission *domain.CreatePermissionDto) error {
	ctx, cancel := context.WithTimeout(c, rpu.contextTimeout)
	defer cancel()
	return rpu.rolePermissionRepository.Create(ctx, CreatePermission)
}

// Update implements domain.RolePermissionUsecase.
func (rpu *rolePermission) Update(c context.Context, updatePermission *domain.UpdatePermissionDto) error {
	ctx, cancel := context.WithTimeout(c, rpu.contextTimeout)
	defer cancel()
	return rpu.rolePermissionRepository.Update(ctx, updatePermission)
}

// Delete implements domain.RolePermissionUsecase.
func (rpu *rolePermission) Delete(c context.Context, roleId int32) error {
	ctx, cancel := context.WithTimeout(c, rpu.contextTimeout)
	defer cancel()
	return rpu.rolePermissionRepository.Delete(ctx, roleId)
}

// FetchByPermissionId implements domain.RolePermissionUsecase.
func (rpu *rolePermission) FetchByPermissionId(c context.Context, permissionId int32) ([]domain.RolePermission, error) {
	ctx, cancel := context.WithTimeout(c, rpu.contextTimeout)
	defer cancel()
	return rpu.rolePermissionRepository.FetchByPermissionId(ctx, permissionId)
}

// FetchByRoleId implements domain.RolePermissionUsecase.
func (rpu *rolePermission) FetchByRoleId(c context.Context, roleId int32) ([]domain.RolePermission, error) {
	ctx, cancel := context.WithTimeout(c, rpu.contextTimeout)
	defer cancel()
	return rpu.rolePermissionRepository.FetchByRoleId(ctx, roleId)
}


// FetchRolePermissions implements domain.RolePermissionUsecase.
func (rpu *rolePermission) FetchRolePermissions(c context.Context) ([]domain.RolePermission, error) {
	ctx, cancel := context.WithTimeout(c, rpu.contextTimeout)
	defer cancel()
	return rpu.rolePermissionRepository.FetchRolePermissions(ctx)
}



