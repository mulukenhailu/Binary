-- name: CreateRolePermission :copyfrom
INSERT INTO RolePermission(RoleId, PermissinoId)
VALUES($1, $2);

-- update operation use the functionality of 
-- 1.delete and then 
-- 2.update
-- this way avoid the need to loop

-- name: DeleteRolePermission :many
DELETE FROM RolePermission
WHERE  RoleId = $1
RETURNING *;

-- name: FetchByRoleId :many
SELECT * FROM RolePermission
WHERE RoleId = $1;

-- name: FetchByPermissionId :many
SELECT * FROM RolePermission
WHERE PermissinoId = $1;


