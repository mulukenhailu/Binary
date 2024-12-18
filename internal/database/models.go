// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Appuser struct {
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
	Registereddate  pgtype.Timestamp
}

type Device struct {
	Deviceid       int32
	Serialnumber   pgtype.Text
	Port           string
	Ipaddress      string
	Name           pgtype.Text
	Campus         pgtype.Text
	Blocknumber    pgtype.Text
	Registeredby   string
	Registereddate pgtype.Timestamp
}

type Permission struct {
	Permissionid   int32
	Permissionname string
	Registeredby   string
	Registereddate pgtype.Timestamp
}

type Role struct {
	Roleid         int32
	Rolename       string
	Registeredby   string
	Registereddate pgtype.Timestamp
}

type Rolepermission struct {
	Rolepermissionid int32
	Roleid           int32
	Permissinoid     int32
}

type Student struct {
	Studentinformationid int32
	Studentid            string
	Firstname            string
	Fathername           string
	Grandfathername      string
	Yearofregistration   string
	Phonenumber          string
	Religion             string
	Sex                  string
	Status               string
	Cardnumber           pgtype.Text
	Photo                pgtype.Text
	Registeredby         string
	Registereddate       pgtype.Timestamp
}
