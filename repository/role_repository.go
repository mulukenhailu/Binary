package repository

import (
	"context"
	"database/sql"

	"github.com/mulukenhailu/Binary/domain"
	"github.com/mulukenhailu/Binary/internal/database"
	"github.com/mulukenhailu/Binary/internal/utils"
)

type roleRepository struct {
	pg *database.Queries
	db *sql.DB
}

func NewRoleRepository(db *sql.DB) domain.RoleRespository {
	return &roleRepository{
		pg:database.New(db),
		db:db,
	}
}

// Create implements domain.RoleRespository.
func (rr *roleRepository) Create(c context.Context, role *domain.RoleDto) error {

	tx, err := rr.db.Begin()
	if err != nil{
		return err
	}

	defer tx.Rollback()

	roleParam := database.CreateRoleParams{
		Rolename : role.RoleName,
		Registeredby : role.RegisteredBy,
	}

	qtx := rr.pg.WithTx(tx)

	_, err = qtx.CreateRole(c, roleParam)
	if err != nil{
		return err
	}
	
	_, err = qtx.ResetRoleId(c)
	if err != nil{
		return err
	}

	return tx.Commit()
}

// Delete implements domain.RoleRespository.
func (rr *roleRepository) Delete(c context.Context, roleId int32) error {
	_, err := rr.pg.DeleteRole(c, roleId)
	return err
}

// Update implements domain.RoleRespository.
func (rr *roleRepository) Update(c context.Context, updatedRole *domain.UpdateRoleDto) error {

	newRole := database.UpdateRoleParams{
		Rolename     :updatedRole.PreviousRoleName,
		Rolename_2   :updatedRole.NewRoleName,
		Registeredby :updatedRole.NewRegisterName,
	}

	_, err := rr.pg.UpdateRole(c, newRole)
	return err
}

// FetchRoles implements domain.RoleRespository.
func (rr *roleRepository) FetchRoles(c context.Context) ([]domain.Role, error) {

	dbRoles, err := rr.pg.FetchRoles(c)

	domainRoles := utils.ConvertDbRolesToDomainRoles(dbRoles)
	return domainRoles, err

}

// FetchById implements domain.RoleRespository.
func (rr *roleRepository) FetchByName(c context.Context, roleName string) (domain.Role, error) {

	dbRole, err := rr.pg.FetchByName(c, roleName)

	domainRoles := utils.ConvertDbRolesToDomainRoles([]database.Role{dbRole})
	return domainRoles[0], err
}





