-- name: CreateRolePermission :copyfrom
INSERT INTO RolePermission(RoleId, PermissinoId)
VALUES($1, $2);

-- update operation use the functionality of 
-- 1.delete and then only then
-- 2.create
-- this way avoid the need to loop


-- name: DeleteRolePermission :many
DELETE FROM RolePermission
WHERE  RoleId = $1
RETURNING *;


-- name: FetchRolePermissions :many
SELECT * FROM RolePermission
ORDER BY RoleId;

-- name: FetchRolePermissionByRoleId :many
SELECT * FROM RolePermission
WHERE RoleId = $1;

-- name: FetchRolePermissionByPermissionId :many
SELECT * FROM RolePermission
WHERE PermissinoId = $1;


