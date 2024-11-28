package domain

import (
	"context"
)

type Role struct{
	RoleId 			int32  `json:"role_id"`
	RoleName 		string `json:"role_name"`
	RegisteredBy 	string `json:"registered_by"`
	RegisteredDate 	string `json:"registered_date"`
}
type CreateRoleDto struct{
	RoleName 		string `json:"role_name" binding:"required"`
	RegisteredBy 	string `json:"registered_by" binding:"required"`
}

type UpdateRoleDto struct{
	RoleId 			int32  `json:"role_id" binding:"required"`
	RoleName 		string `json:"role_name" binding:"required"`
	RegisteredBy 	string `json:"registered_by" binding:"required"`
}

type FetchByNameDto struct{
	RoleName string `json:"role_name" binding:"required"`
}

type DeleteRoleDto struct{
	RoleId 	int32 `json:"role_id" binding:"required"`
}


type RoleRespository interface{
	Create(c context.Context, createRoleDto *CreateRoleDto) 			error
	Update(c context.Context, updateRoleDto *UpdateRoleDto) 			error
	Delete(c context.Context, roleId int32) 							error
	FetchRoles(c context.Context) 										([]Role, error)
	FetchByName(c context.Context, roleName string)						(Role, error)
}

type RoleUsecase interface{
	Create(c context.Context, role *CreateRoleDto) 						error
	Update(c context.Context, updateRoleDto *UpdateRoleDto) 			error
	Delete(c context.Context, roleId int32) 							error
	FetchRoles(c context.Context) 										([]Role, error)
	FetchByName(c context.Context, roleName string)						(Role, error)
}


