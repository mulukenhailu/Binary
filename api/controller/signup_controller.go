package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mulukenhailu/Binary/domain"
	"github.com/mulukenhailu/Binary/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
}

func (sc *SignupController)Signup(c *gin.Context){
	var signupDto domain.CreateUserDto

	err := c.ShouldBindJSON(&signupDto)
	if err != nil{

		var validationErrors validator.ValidationErrors
		if errors, ok := err.(validator.ValidationErrors); ok{
			validationErrors = errors
		}

		errorMessage := make(map[string]string)
		for _, e := range validationErrors{

			field := e.Field()
			switch field {
			case "RoleId":
				errorMessage["RoleId"] = "Role ID is required"
			case "UserName":
				errorMessage["UserName"] = "User Name is required"
			case "FirstName":
				errorMessage["FirstName"] = "First Name is required"
			case "FatherName":
				errorMessage["FatherName"] = "Father Name is required"
			case "GrandFatherName":
				errorMessage["GrandFatherName"] = "Grandfather Name is required"
			case "Password":
				errorMessage["Password"] = "Password is required"
			case "PhoneNumber":
				errorMessage["PhoneNumber"] = "Phone Number is required"
			case "Address":
				errorMessage["Address"] = "Address is required"
			case "Email":
				errorMessage["Email"] = "Email is required"
			case "RegisteredBy":
				errorMessage["RegisteredBy"] = "Registered By is required"
			default:
				errorMessage["UnknownField"] = "Invalid field provided"
			}
			
		}

		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Validation failed", Errors:  errorMessage})
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
	signupDto.UserName = utils.ConvertToSmallLetter(signupDto.UserName)

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



