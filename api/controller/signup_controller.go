package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mulukenhailu/Binary/domain"
	"golang.org/x/crypto/bcrypt"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
}

func (sc *SignupController)Signup(c *gin.Context){
	var signupDto domain.CreateUserDto

	err := c.ShouldBindJSON(&signupDto)
	if err != nil{
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return 
	}

	_, err = sc.SignupUsecase.FetchByUserName(c, signupDto.UserName)
	if err == nil {
		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "user already exit with the given user-name"})
		return 
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(signupDto.Password), bcrypt.DefaultCost)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message:err.Error()})
		return 
	}

	signupDto.Password = string(hashPassword)

	user := domain.CreateUserDto{
		RoleId          :signupDto.RoleId,
		UserName        :signupDto.UserName,
		FirstName       :signupDto.FirstName,
		FatherName      :signupDto.FatherName,
		GrandFatherName :signupDto.GrandFatherName,
		Password        :signupDto.Password,
		PhoneNumber     :signupDto.PhoneNumber,
		Address         :signupDto.Address,
		Email           :signupDto.Email,
		RegisteredBy 	:signupDto.RegisteredBy,
	}

	err = sc.SignupUsecase.Create(c, &user)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return 
	}

	c.JSON(http.StatusOK, domain.SucessResponse{Message: "created successfully"})
}



