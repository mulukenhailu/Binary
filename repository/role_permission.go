package repository

import (
	"context"

	// "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mulukenhailu/Binary/domain"
	"github.com/mulukenhailu/Binary/internal/database"
	"github.com/mulukenhailu/Binary/internal/utils"
)

type rolePermissionRepository struct {
	pg *database.Queries
	db *pgxpool.Pool
}

func NewRolePermissionRepository(db *pgxpool.Pool) domain.RolePermissionRepository {
	return &rolePermissionRepository{
		pg: database.New(db),
		db:db,
	}
}

// Create implements domain.RolePermissionRepository.
func (rpr *rolePermissionRepository) Create(c context.Context, createPermission *domain.CreatePermissionDto) error {
		roleId := createPermission.RoleId
		permissionIdList := createPermission.PermissionIdList

		CreateRolePermissionParamsList := make([]database.CreateRolePermissionParams, len(permissionIdList))
		for i, permissionId := range permissionIdList {
			CreateRolePermissionParamsList[i] = database.CreateRolePermissionParams{
				Roleid:      roleId,
				Permissinoid: permissionId,
			}
		}
		
		tx, err := rpr.db.Begin(c)
		if err != nil {
			return err
		}
				defer tx.Rollback(c)

				tq := rpr.pg.WithTx(tx)
				_, err = tq.CreateRolePermission(c, CreateRolePermissionParamsList)
				if err != nil {
					return err
				}
		
		err = tx.Commit(c)
		if err != nil {
			return err
		}

    	return nil

		// ctx := context.Background()
		// conn, err := rpr.db.Acquire(ctx)
		// if err != nil{
		// 	return err
		// }

		// tx, err :=conn.Begin(ctx)
		// if err != nil{
		// 	return err
		// }

		// _, err = rpr.pg.WithTx(tx).CreateRolePermission(ctx, CreateRolePermissionParamsList)
		// if err != nil{
		// 	return err
		// }

		// defer conn.Release()
		// return tx.Commit(ctx)





	// tx, err := rpr.database.(*pgxpool.Pool).Begin()
	
    // tx, err := rpr.db.Begin(c)
    // if err != nil {
    //     return err
    // }
    // defer tx.Rollback(c) 

    // qtx := rpr.pg.WithTx(tx)
	// _, err = qtx.CreateRolePermission(c, CreateRolePermissionParamsList)
	// if err != nil {
	// 	return err 
	// }

    // return tx.Commit(c)
}


// Delete implements domain.RolePermissionRepository.
func (rpr *rolePermissionRepository) Delete(c context.Context, deletePermission *domain.DeletePermissionDto) error {
	_, err := rpr.pg.DeleteRolePermission(c, deletePermission.RoleId)
	return err
}

// Update implements domain.RolePermissionRepository.
func (rpr *rolePermissionRepository) Update(c context.Context, updatePermission *domain.UpdatePermissionDto) error {
	roleId := updatePermission.RoleId
	permissionListId := updatePermission.PermissionIdList

	deletePermissionDto := domain.DeletePermissionDto{
		RoleId: roleId,
	}
	err := rpr.Delete(c, &deletePermissionDto)
	if err != nil{
		return err
	}

	createPermissionDto := domain.CreatePermissionDto{
		RoleId: roleId,
		PermissionIdList: permissionListId,
	}
	err = rpr.Create(c, &createPermissionDto)

	return err
}

// FetchByPermissionId implements domain.RolePermissionRepository.
func (rpr *rolePermissionRepository) FetchByPermissionId(c context.Context, permissionId int32) ([]domain.RolePermission, error) {
	dbRolePermissionId, err := rpr.pg.FetchRolePermissionByPermissionId(c, permissionId)
	domainRolePermissionId := utils.ConvertDbRolePermToDomainRolePerm(dbRolePermissionId)
	return domainRolePermissionId, err
}

// FetchByRoleId implements domain.RolePermissionRepository.
func (rpr *rolePermissionRepository) FetchByRoleId(c context.Context, roleId int32) ([]domain.RolePermission, error) {
	roleRoleId, err := rpr.pg.FetchRolePermissionByRoleId(c, roleId)
	domainroleRoleId := utils.ConvertDbRolePermToDomainRolePerm(roleRoleId)
	return domainroleRoleId, err
}




