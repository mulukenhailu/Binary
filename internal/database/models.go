// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"database/sql"
)

type Device struct {
	Deviceid       int32
	Serialnumber   sql.NullString
	Port           string
	Ipaddress      string
	Name           sql.NullString
	Campus         sql.NullString
	Blocknumber    sql.NullString
	Registeredby   string
	Registereddate sql.NullTime
}

type Role struct {
	Roleid         int32
	Rolename       string
	Registeredby   string
	Registereddate sql.NullTime
}
