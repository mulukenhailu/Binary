-- name: CreateRole :one
INSERT INTO role(rolename, registeredby)
VALUES($1, $2)
RETURNING *;

-- name: DeleteRole :one
DELETE FROM Role
WHERE RoleId = $1
RETURNING *;

-- name: UpdateRole :one
UPDATE Role
SET RoleName = $2, RegisteredBy = $3
WHERE RoleName = $1
RETURNING *;

-- name: FetchRoles :many
SELECT * FROM Role
ORDER BY RoleId;

-- name: FetchByName :one
SELECT * FROM Role
WHERE RoleName = $1;

