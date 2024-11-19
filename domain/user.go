package domain

import "context"

type User struct {
	UserId          int32  `json:"user_id"`
	RoleId          int32  `json:"role_id"`
	UserName        string `json:"user_name"`
	FirstName       string `json:"first_name"`
	FatherName      string `json:"father_name"`
	GrandFatherName string `json:"grand_father_name"`
	Password        string `json:"password"`
	PhoneNumber     string `json:"phone_number"`
	Address         string `json:"address"`
	Email           string `json:"email"`
	RegisteredBy 	string  `json:"registered_by"`
	RegisteredDate 	string `json:"registered_date"`
}

type CreateUserDto struct {
	RoleId          int32  `json:"role_id"`
	UserName        string `json:"user_name"`
	FirstName       string `json:"first_name"`
	FatherName      string `json:"father_name"`
	GrandFatherName string `json:"grand_father_name"`
	Password        string `json:"password"`
	PhoneNumber     string `json:"phone_number"`
	Address         string `json:"address"`
	Email           string `json:"email"`
	RegisteredBy 	string  `json:"registered_by"`
}

type UpdateUserDto struct {
	RoleId          int32  `json:"role_id"`
	UserId          int32  `json:"user_id"`
	UserName        string `json:"user_name"`
	FirstName       string `json:"first_name"`
	FatherName      string `json:"father_name"`
	GrandFatherName string `json:"grand_father_name"`
	Password        string `json:"password"`
	PhoneNumber     string `json:"phone_number"`
	Address         string `json:"address"`
	Email           string `json:"email"`
	RegisteredBy 	string  `json:"registered_by"`
}

type DeleteUserDto struct{
	UserId int32 `json:"user_id"`
}

type FetchByRoleNameDto struct{
	RoleId int32 `json:"role_id"`
}

type FetchByUserNameDto struct{
	UserName string  `json:"user_name"`
}

type UserRepository interface {
	Create(c context.Context, createUserDto *CreateUserDto) error
	Update(c context.Context, updateUserDto *UpdateUserDto) error 
	Delete(c context.Context, userId int32) 				error
	FetchByRoleId(c context.Context, roleId int32)		([]User, error)
	FetchByUserName(c context.Context, userName string)		(User, error)
	FetchUsers(c context.Context)							([]User, error)
}

