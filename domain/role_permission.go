package domain

import "context"

type RolePermission struct {
	RolePermissionId 	int32 `json:"role_permission_id"`
	RoleId       		int32  `json:"role_id"`
	PermissionId  		int32 `json:"permission_id"`
}

type CreatePermissionDto struct{
	RoleId 				int32 `json:"role_id" binding:"required"`
	PermissionIdList 	[]int32 `json:"permission_id_list" binding:"required"`
}

type UpdatePermissionDto struct{
	RoleId 				int32 `json:"role_id" binding:"required"`
	PermissionIdList 	[]int32 `json:"permission_id_list" binding:"required"`
}


type DeletePermissionDto struct{
	RoleId int32 `json:"role_id" binding:"required"`
}

type RolePermissionUsecase interface {
	Create(c context.Context, createPermission *CreatePermissionDto) 	error
	Update(c context.Context, updatePermission *UpdatePermissionDto) 	error 
	Delete(c context.Context, deletePermission *DeletePermissionDto) 	error
	FetchByRoleId(c context.Context, roleId int32) 						([]RolePermission, error)
	FetchByPermissionId(c context.Context, permissionId int32) 			([]RolePermission, error)
} 

type RolePermissionRepository interface {
	Create(c context.Context, createPermission *CreatePermissionDto) 	error
	Update(c context.Context, updatePermission *UpdatePermissionDto) 	error 
	Delete(c context.Context, deletePermission *DeletePermissionDto) 	error
	FetchByRoleId(c context.Context, roleId int32) 						([]RolePermission, error)
	FetchByPermissionId(c context.Context, permissionId int32) 			([]RolePermission, error)
} 

