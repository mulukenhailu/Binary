package utils

import (
	"github.com/mulukenhailu/Binary/domain"
	"github.com/mulukenhailu/Binary/internal/database"
)

func ConvertDbRolesToDomainRoles(dbRoles []database.Role) []domain.Role{
	domainRoles := make([]domain.Role, len(dbRoles))
	for i, dbRole := range dbRoles{
		domainRoles[i] = domain.Role{
			RoleId : dbRole.Roleid,
			RoleName : dbRole.Rolename,
			RegisteredBy : dbRole.Registeredby,
			RegisteredDate : dbRole.Registereddate.Time.String(),
		}
	}
	return domainRoles
}