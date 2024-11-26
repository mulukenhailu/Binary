package repository

import (
	"context"
	"database/sql"

	"github.com/mulukenhailu/Binary/domain"
	"github.com/mulukenhailu/Binary/internal/database"
)

type rolePermissionRepository struct {
	pg *database.Queries
}

func NewRolePermissionRepository(db *sql.DB) domain.RolePermissionRepository {
	return &rolePermissionRepository{
		pg: database.New(db),
	}
}

// Create implements domain.RolePermissionRepository.
func (rpr *rolePermissionRepository) Create(c context.Context, rolePermission *domain.CreatePermissionDto) error {
	// roleId := rolePermission.RoleId
	// permissionIdList := rolePermission.PermissionIdList

	// for _, permissionId := range permissionIdList{

	// 	createRolePermissionParams := database.CreateRolePermissionParams{
	// 		Roleid: roleId,
	// 		Permissinoid: permissionId,
	// 	}
	// 	rpr.pg.CreateRolePermission(c, createRolePermissionParams)
	// }
	panic("unimplemented")
}

// Delete implements domain.RolePermissionRepository.
func (rpr *rolePermissionRepository) Delete(c context.Context, deletePermission *domain.DeletePermissionDto) error {
	panic("unimplemented")
}

// Update implements domain.RolePermissionRepository.
func (rpr *rolePermissionRepository) Update(c context.Context, updatePermission *domain.UpdatePermissionDto) error {
	panic("unimplemented")
}

// FetchByPermissionId implements domain.RolePermissionRepository.
func (rpr *rolePermissionRepository) FetchByPermissionId(c context.Context, permissionId int32) ([]domain.RolePermission, error) {
	panic("unimplemented")
}

// FetchByRoleId implements domain.RolePermissionRepository.
func (rpr *rolePermissionRepository) FetchByRoleId(c context.Context, roleId int32) ([]domain.RolePermission, error) {
	panic("unimplemented")
}




