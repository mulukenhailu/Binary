// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: role.sql

package database

import (
	"context"
)

const createRole = `-- name: CreateRole :one
INSERT INTO role(rolename, registeredby)
VALUES($1, $2)
RETURNING roleid, rolename, registeredby, registereddate
`

type CreateRoleParams struct {
	Rolename     string
	Registeredby string
}

func (q *Queries) CreateRole(ctx context.Context, arg CreateRoleParams) (Role, error) {
	row := q.db.QueryRowContext(ctx, createRole, arg.Rolename, arg.Registeredby)
	var i Role
	err := row.Scan(
		&i.Roleid,
		&i.Rolename,
		&i.Registeredby,
		&i.Registereddate,
	)
	return i, err
}

const deleteRole = `-- name: DeleteRole :one
DELETE FROM Role
WHERE RoleId = $1
RETURNING roleid, rolename, registeredby, registereddate
`

func (q *Queries) DeleteRole(ctx context.Context, roleid int32) (Role, error) {
	row := q.db.QueryRowContext(ctx, deleteRole, roleid)
	var i Role
	err := row.Scan(
		&i.Roleid,
		&i.Rolename,
		&i.Registeredby,
		&i.Registereddate,
	)
	return i, err
}

const fetchByName = `-- name: FetchByName :one
SELECT roleid, rolename, registeredby, registereddate FROM Role
WHERE RoleName = $1
`

func (q *Queries) FetchByName(ctx context.Context, rolename string) (Role, error) {
	row := q.db.QueryRowContext(ctx, fetchByName, rolename)
	var i Role
	err := row.Scan(
		&i.Roleid,
		&i.Rolename,
		&i.Registeredby,
		&i.Registereddate,
	)
	return i, err
}

const fetchRoles = `-- name: FetchRoles :many
SELECT roleid, rolename, registeredby, registereddate FROM Role
ORDER BY RoleId
`

func (q *Queries) FetchRoles(ctx context.Context) ([]Role, error) {
	rows, err := q.db.QueryContext(ctx, fetchRoles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Role
	for rows.Next() {
		var i Role
		if err := rows.Scan(
			&i.Roleid,
			&i.Rolename,
			&i.Registeredby,
			&i.Registereddate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const resetRoleId = `-- name: ResetRoleId :one
SELECT setval('role_roleid_seq', (SELECT MAX(RoleId) FROM Role))
`

func (q *Queries) ResetRoleId(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, resetRoleId)
	var setval int64
	err := row.Scan(&setval)
	return setval, err
}

const updateRole = `-- name: UpdateRole :one
UPDATE Role
SET RoleName = $2, RegisteredBy = $3
WHERE RoleName = $1
RETURNING roleid, rolename, registeredby, registereddate
`

type UpdateRoleParams struct {
	Rolename     string
	Rolename_2   string
	Registeredby string
}

func (q *Queries) UpdateRole(ctx context.Context, arg UpdateRoleParams) (Role, error) {
	row := q.db.QueryRowContext(ctx, updateRole, arg.Rolename, arg.Rolename_2, arg.Registeredby)
	var i Role
	err := row.Scan(
		&i.Roleid,
		&i.Rolename,
		&i.Registeredby,
		&i.Registereddate,
	)
	return i, err
}
