// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: role_permission.sql

package database

import (
	"context"
)

const createRolePermission = `-- name: CreateRolePermission :one
INSERT INTO RolePermission(RoleId, PermissinoId)
VALUES($1, $2)
RETURNING rolepermissionid, roleid, permissinoid
`

type CreateRolePermissionParams struct {
	Roleid       int32
	Permissinoid int32
}

func (q *Queries) CreateRolePermission(ctx context.Context, arg CreateRolePermissionParams) (Rolepermission, error) {
	row := q.db.QueryRowContext(ctx, createRolePermission, arg.Roleid, arg.Permissinoid)
	var i Rolepermission
	err := row.Scan(&i.Rolepermissionid, &i.Roleid, &i.Permissinoid)
	return i, err
}

const deleteRolePermission = `-- name: DeleteRolePermission :many

DELETE FROM RolePermission
WHERE  RoleId = $1
RETURNING rolepermissionid, roleid, permissinoid
`

// update operation use the functionality of
// 1.delete and then
// 2.update
// this way avoid the need to loop
func (q *Queries) DeleteRolePermission(ctx context.Context, roleid int32) ([]Rolepermission, error) {
	rows, err := q.db.QueryContext(ctx, deleteRolePermission, roleid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Rolepermission
	for rows.Next() {
		var i Rolepermission
		if err := rows.Scan(&i.Rolepermissionid, &i.Roleid, &i.Permissinoid); err != nil {
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

const fetchByPermissionId = `-- name: FetchByPermissionId :many
SELECT rolepermissionid, roleid, permissinoid FROM RolePermission
WHERE PermissinoId = $1
`

func (q *Queries) FetchByPermissionId(ctx context.Context, permissinoid int32) ([]Rolepermission, error) {
	rows, err := q.db.QueryContext(ctx, fetchByPermissionId, permissinoid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Rolepermission
	for rows.Next() {
		var i Rolepermission
		if err := rows.Scan(&i.Rolepermissionid, &i.Roleid, &i.Permissinoid); err != nil {
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

const fetchByRoleId = `-- name: FetchByRoleId :many
SELECT rolepermissionid, roleid, permissinoid FROM RolePermission
WHERE RoleId = $1
`

func (q *Queries) FetchByRoleId(ctx context.Context, roleid int32) ([]Rolepermission, error) {
	rows, err := q.db.QueryContext(ctx, fetchByRoleId, roleid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Rolepermission
	for rows.Next() {
		var i Rolepermission
		if err := rows.Scan(&i.Rolepermissionid, &i.Roleid, &i.Permissinoid); err != nil {
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
