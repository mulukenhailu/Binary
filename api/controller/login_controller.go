package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mulukenhailu/Binary/bootstrap"
	"github.com/mulukenhailu/Binary/domain"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Env *bootstrap.Env
}

func (lc *LoginController)Login(c *gin.Context){
	var loginRequestDto domain.LoginRequestDto

	err := c.ShouldBindJSON(&loginRequestDto)
	if err != nil{
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message:err.Error()})
		return 
	}

	DbUser, err := lc.LoginUsecase.FetchByUserName(c, loginRequestDto.UserName)
	if err != nil {
		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: err.Error()})
		return 
	}

	if bcrypt.CompareHashAndPassword([]byte(DbUser.Password), []byte(loginRequestDto.Password)) != nil{
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid Credentials"})
		return 
	}

	user := domain.CreateUserDto{
		RoleId          :DbUser.RoleId,
		UserName        :DbUser.UserName,
		FirstName       :DbUser.FirstName,
		FatherName      :DbUser.FatherName,
		GrandFatherName :DbUser.GrandFatherName,
		Password        :DbUser.Password,
		PhoneNumber     :DbUser.PhoneNumber,
		Address         :DbUser.Address,
		Email           :DbUser.Email,
		RegisteredBy 	:DbUser.RegisteredBy,
	}

	accessToken, err := lc.LoginUsecase.CreateAccessToken(&user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message:err.Error()})
		return 
	}

	signupResponse := domain.LoginResponseDto{
		AccessToken: accessToken,
	}

	c.JSON(http.StatusOK, signupResponse)

}