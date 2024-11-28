package repository

import (
	"context"

	"github.com/mulukenhailu/Binary/domain"
	"github.com/mulukenhailu/Binary/internal/database"
	"github.com/mulukenhailu/Binary/internal/utils"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/pgtype"
)

type userRespository struct {
	pg *database.Queries
}

func NewUserRepository(db *pgxpool.Pool) domain.UserRepository {
	return &userRespository{
		pg: database.New(db),
	}
}


// Create implements domain.UserRepository.
func (ur *userRespository) Create(c context.Context, createUserDto *domain.CreateUserDto) error {
	userParams :=  database.CreateUserParams{
		Roleid          :createUserDto.RoleId,
		Username        :createUserDto.UserName,
		Firstname       :createUserDto.FirstName,
		Fathername      :createUserDto.FatherName,
		Grandfathername :createUserDto.GrandFatherName,
		Password        :createUserDto.Password,
		Phonenumber     :createUserDto.PhoneNumber,
		Address         :createUserDto.Address,
		Email           :pgtype.Text{String: createUserDto.Email, Valid: createUserDto.Email != ""},
		Registeredby    :createUserDto.RegisteredBy,
	}

	_, err := ur.pg.CreateUser(c, userParams)
	return err
}

// Delete implements domain.UserRepository.
func (ur *userRespository) Delete(c context.Context, userId int32) error {
	_, err := ur.pg.DeleteUser(c, userId)
	return err
}

// FetchByRoleName implements domain.UserRepository.
func (ur *userRespository) FetchByRoleId(c context.Context, roleId int32) ([]domain.User, error) {
	dbUsers, err := ur.pg.FetchByRoleName(c, roleId)
	domainUsers := utils.ConvertDbUserToDomainUser(dbUsers)
	return domainUsers, err
}

// FetchByUserName implements domain.UserRepository.
func (ur *userRespository) FetchByUserName(c context.Context, userName string) (domain.User, error) {
	dbUser, err := ur.pg.FetchByUserName(c, userName)
	domainUsers := utils.ConvertDbUserToDomainUser([]database.Appuser{dbUser})
	return domainUsers[0], err
}

// FetchUsers implements domain.UserRepository.
func (ur *userRespository) FetchUsers(c context.Context) ([]domain.User, error) {
	dbUsers, err := ur.pg.FetchUsers(c)
	domainUsers := utils.ConvertDbUserToDomainUser(dbUsers)
	return domainUsers, err
}

// Update implements domain.UserRepository.
func (ur *userRespository) Update(c context.Context, updateUserDto *domain.UpdateUserDto) error {
	userUpdateParam := database.UpdateUserParams{
		Userid           :updateUserDto.UserId,
		Roleid           :updateUserDto.RoleId,
		Username         :updateUserDto.UserName,
		Firstname        :updateUserDto.FirstName,
		Fathername       :updateUserDto.FatherName,
		Grandfathername  :updateUserDto.GrandFatherName,
		Password         :updateUserDto.Password,
		Phonenumber      :updateUserDto.PhoneNumber,
		Address          :updateUserDto.Address,
		Email            :pgtype.Text{String: updateUserDto.Email, Valid: updateUserDto.Email != ""},
		Registeredby     :updateUserDto.RegisteredBy,
	}

	_, err := ur.pg.UpdateUser(c, userUpdateParam)
	return err
}

