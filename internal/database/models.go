// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"database/sql"
)

type Role struct {
	Roleid         int32
	Rolename       string
	Registeredby   string
	Registereddate sql.NullTime
}
