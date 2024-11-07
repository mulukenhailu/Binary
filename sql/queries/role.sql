-- name: CreateRole :one
INSERT INTO role(rolename, registeredby)
VALUES($1, $2)
RETURNING *;