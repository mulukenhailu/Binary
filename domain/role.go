package domain

import (
	"context"
)


type RoleDto struct{
	RoleName string `json:"role_name"`
	RegisteredBy string `json:"registered_by"`
}

type Role struct{
	RoleId int32 `json:"role_id"`
	RoleName string `json:"role_name"`
	RegisteredBy string `json:"registered_by"`
	RegisteredDate string `json:"registered_date"`
}



type RoleRespository interface{
	Create(c context.Context, role *RoleDto) 			error
	Update(c context.Context, role *RoleDto) 			error
	Delete(c context.Context, roleId int32) 			error
	FetchByName(c context.Context, roleName string)		(RoleDto, error)
	FetchRoles(c context.Context) 						([]RoleDto, error)
	
}

type RoleUsercase interface{
	Create(c context.Context, role *RoleDto) 			error
	Update(c context.Context, role *RoleDto) 			error
	Delete(c context.Context, roleId int32) 			error
	FetchByName(c context.Context, roleName string)		(RoleDto, error)
	FetchRoles(c context.Context) 						([]RoleDto, error)
}
