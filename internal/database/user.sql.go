// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: user.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO AppUser(
    RoleId, 
    UserName, 
    FirstName, 
    FatherName, 
    GrandFatherName, 
    Password, 
    PhoneNumber, 
    Address, 
    Email, 
    RegisteredBy
    )
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING userid, roleid, username, firstname, fathername, grandfathername, password, phonenumber, address, email, registeredby, registereddate
`

type CreateUserParams struct {
	Roleid          int32
	Username        string
	Firstname       string
	Fathername      string
	Grandfathername string
	Password        string
	Phonenumber     string
	Address         string
	Email           pgtype.Text
	Registeredby    string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (Appuser, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Roleid,
		arg.Username,
		arg.Firstname,
		arg.Fathername,
		arg.Grandfathername,
		arg.Password,
		arg.Phonenumber,
		arg.Address,
		arg.Email,
		arg.Registeredby,
	)
	var i Appuser
	err := row.Scan(
		&i.Userid,
		&i.Roleid,
		&i.Username,
		&i.Firstname,
		&i.Fathername,
		&i.Grandfathername,
		&i.Password,
		&i.Phonenumber,
		&i.Address,
		&i.Email,
		&i.Registeredby,
		&i.Registereddate,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :one
DELETE FROM AppUser
where UserId = $1
RETURNING userid, roleid, username, firstname, fathername, grandfathername, password, phonenumber, address, email, registeredby, registereddate
`

func (q *Queries) DeleteUser(ctx context.Context, userid int32) (Appuser, error) {
	row := q.db.QueryRow(ctx, deleteUser, userid)
	var i Appuser
	err := row.Scan(
		&i.Userid,
		&i.Roleid,
		&i.Username,
		&i.Firstname,
		&i.Fathername,
		&i.Grandfathername,
		&i.Password,
		&i.Phonenumber,
		&i.Address,
		&i.Email,
		&i.Registeredby,
		&i.Registereddate,
	)
	return i, err
}

const fetchByRoleName = `-- name: FetchByRoleName :many
SELECT 
AppUser.UserId, 
AppUser.RoleId, 
AppUser.UserName, 
AppUser.FirstName, 
AppUser.FatherName, 
AppUser.GrandFatherName, 
AppUser.Password,
AppUser.PhoneNumber, 
AppUser.Address, 
AppUser.Email, 
AppUser.RegisteredBy, 
AppUser.RegisteredDate 
FROM AppUser
INNER JOIN Role ON AppUser.RoleId = Role.RoleId
WHERE Role.RoleId = $1
`

func (q *Queries) FetchByRoleName(ctx context.Context, roleid int32) ([]Appuser, error) {
	rows, err := q.db.Query(ctx, fetchByRoleName, roleid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Appuser
	for rows.Next() {
		var i Appuser
		if err := rows.Scan(
			&i.Userid,
			&i.Roleid,
			&i.Username,
			&i.Firstname,
			&i.Fathername,
			&i.Grandfathername,
			&i.Password,
			&i.Phonenumber,
			&i.Address,
			&i.Email,
			&i.Registeredby,
			&i.Registereddate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const fetchByUserName = `-- name: FetchByUserName :one
SELECT userid, roleid, username, firstname, fathername, grandfathername, password, phonenumber, address, email, registeredby, registereddate FROM AppUser
WHERE UserName = $1
`

func (q *Queries) FetchByUserName(ctx context.Context, username string) (Appuser, error) {
	row := q.db.QueryRow(ctx, fetchByUserName, username)
	var i Appuser
	err := row.Scan(
		&i.Userid,
		&i.Roleid,
		&i.Username,
		&i.Firstname,
		&i.Fathername,
		&i.Grandfathername,
		&i.Password,
		&i.Phonenumber,
		&i.Address,
		&i.Email,
		&i.Registeredby,
		&i.Registereddate,
	)
	return i, err
}

const fetchUsers = `-- name: FetchUsers :many
SELECT userid, roleid, username, firstname, fathername, grandfathername, password, phonenumber, address, email, registeredby, registereddate FROM AppUser
ORDER BY UserId
`

func (q *Queries) FetchUsers(ctx context.Context) ([]Appuser, error) {
	rows, err := q.db.Query(ctx, fetchUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Appuser
	for rows.Next() {
		var i Appuser
		if err := rows.Scan(
			&i.Userid,
			&i.Roleid,
			&i.Username,
			&i.Firstname,
			&i.Fathername,
			&i.Grandfathername,
			&i.Password,
			&i.Phonenumber,
			&i.Address,
			&i.Email,
			&i.Registeredby,
			&i.Registereddate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :one
UPDATE AppUser
SET 
    RoleId = $2, 
    UserName= $3, 
    FirstName = $4, 
    FatherName = $5, 
    GrandFatherName = $6, 
    Password= $7, 
    PhoneNumber = $8, 
    Address = $9, 
    Email = $10, 
    RegisteredBy = $11
WHERE UserId = $1
RETURNING userid, roleid, username, firstname, fathername, grandfathername, password, phonenumber, address, email, registeredby, registereddate
`

type UpdateUserParams struct {
	Userid          int32
	Roleid          int32
	Username        string
	Firstname       string
	Fathername      string
	Grandfathername string
	Password        string
	Phonenumber     string
	Address         string
	Email           pgtype.Text
	Registeredby    string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (Appuser, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.Userid,
		arg.Roleid,
		arg.Username,
		arg.Firstname,
		arg.Fathername,
		arg.Grandfathername,
		arg.Password,
		arg.Phonenumber,
		arg.Address,
		arg.Email,
		arg.Registeredby,
	)
	var i Appuser
	err := row.Scan(
		&i.Userid,
		&i.Roleid,
		&i.Username,
		&i.Firstname,
		&i.Fathername,
		&i.Grandfathername,
		&i.Password,
		&i.Phonenumber,
		&i.Address,
		&i.Email,
		&i.Registeredby,
		&i.Registereddate,
	)
	return i, err
}
