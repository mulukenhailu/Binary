package repository

import (
	"context"
	"database/sql"
	"github.com/mulukenhailu/Binary/domain"
	"github.com/mulukenhailu/Binary/internal/database"
	"github.com/jackc/pgx/v5/pgxpool"
)

type rolePermissionRepository struct {
	pg *database.Queries
	db *sql.DB
}

func NewRolePermissionRepository(db *pgxpool.Pool) domain.RolePermissionRepository {
	return &rolePermissionRepository{
		pg: database.New(db),
	}
}

// Create implements domain.RolePermissionRepository.
func (rpr *rolePermissionRepository) Create(c context.Context, rolePermission *domain.CreatePermissionDto) error {
	roleId := rolePermission.RoleId
	permissionIdList := rolePermission.PermissionIdList
	CreateRolePermissionParamsList := []database.CreateRolePermissionParams{}

	
	tx, err := rpr.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	qtx := rpr.pg.WithTx(tx)

	for _, permissionId := range permissionIdList{
		
		CreateRolePermissionParamsList = append(CreateRolePermissionParamsList, database.CreateRolePermissionParams{
			Roleid: roleId,
			Permissinoid: permissionId,
		} )
		rpr.pg.CreateRolePermission(c, CreateRolePermissionParamsList)
	}
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




