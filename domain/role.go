package domain

import (
	"context"
)


type RoleDto struct{
	RoleName string `json:"role_name"`
	RegisteredBy string `json:"registered_by"`
}

type UpdateRoleDto struct{
	PreviousRoleName string `json:"revious_role_name"`
	NewRoleName string `json:"new_role_name"`
	NewRegisterName string `json:"new_register_name"`
}

type Role struct{
	RoleId int32 `json:"role_id"`
	RoleName string `json:"role_name"`
	RegisteredBy string `json:"registered_by"`
	RegisteredDate string `json:"registered_date"`
}

type FetchByNameDto struct{
	RoleName string `json:"role_name"`
}

type DeleteRoleDto struct{
	RoleId int32 `json:"role_id"`
}



type RoleRespository interface{
	Create(c context.Context, role *RoleDto) 					error
	Update(c context.Context, updateRole *UpdateRoleDto) 		error
	Delete(c context.Context, roleId int32) 					error
	FetchRoles(c context.Context) 								([]Role, error)
	FetchByName(c context.Context, roleName string)				(Role, error)
}

type RoleUsercase interface{
	Create(c context.Context, role *RoleDto) 					error
	Update(c context.Context, updateRole *UpdateRoleDto) 		error
	Delete(c context.Context, roleId int32) 					error
	FetchRoles(c context.Context) 								([]Role, error)
	FetchByName(c context.Context, roleName string)				(Role, error)
}
