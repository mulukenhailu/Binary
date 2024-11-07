package repository

import (
	"context"
	"database/sql"
	"github.com/mulukenhailu/Binary/domain"
	"github.com/mulukenhailu/Binary/internal/database"
)

type roleRepository struct {
	pg *database.Queries
}

func NewRoleRepository(db *sql.DB) domain.RoleRespository {
	return &roleRepository{
		pg:database.New(db),
	}
}




// Create implements domain.RoleRespository.
func (r *roleRepository) Create(c context.Context, role *domain.RoleDto) error {

	roleParam := database.CreateRoleParams{
		Rolename : role.RoleName,
		Registeredby : role.RegisteredBy,
	}
	_, err := r.pg.CreateRole(c, roleParam)
	return err
}

// Delete implements domain.RoleRespository.
func (r *roleRepository) Delete(c context.Context, roleId int32) error {
	panic("unimplemented")
}

// Update implements domain.RoleRespository.
func (r *roleRepository) Update(c context.Context, role *domain.RoleDto) error {
	panic("unimplemented")
}

// FetchRoles implements domain.RoleRespository.
func (r *roleRepository) FetchRoles(c context.Context) ([]domain.RoleDto, error) {
	panic("unimplemented")
}

// FetchById implements domain.RoleRespository.
func (r *roleRepository) FetchByName(c context.Context, roleName string) (domain.RoleDto, error) {
	panic("unimplemented")
}





