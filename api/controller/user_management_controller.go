package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mulukenhailu/Binary/domain"
	"github.com/go-playground/validator/v10"
)

type UserManagementController struct {
	UserManagementUsecase domain.UserManagementUsecase
}


func(umc *UserManagementController)Update(c *gin.Context){
	var updateUserDto domain.UpdateUserDto

	err := c.ShouldBindJSON(&updateUserDto)
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
			case "UserId":
				errorMessage["UserId"] = "User ID is required"
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

		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Validation Failed", Errors:  errorMessage})
		return 
	}

	err = umc.UserManagementUsecase.Update(c, &updateUserDto)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return 
	}

	c.JSON(http.StatusOK, domain.SucessResponse{Message: "update successfully"})
}

func(umc *UserManagementController)Delete(c *gin.Context){
	var deleteUserDto domain.DeleteUserDto

	err := c.ShouldBindJSON(&deleteUserDto)
	if err != nil{
		var validationErrors validator.ValidationErrors
		if errors, ok := err.(validator.ValidationErrors); ok{
			validationErrors = errors
		}

		errorMessage := make(map[string]string)
		for _, e := range validationErrors{

			field := e.Field()
			switch field {
			case "UserId":
				errorMessage["UserId"] = "User ID is required"
			default:
				errorMessage["UnknownField"] = "Invalid field provided"
			}	
		}
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Validation failed", Errors:  errorMessage})
		return 
	}

	err = umc.UserManagementUsecase.Delete(c, deleteUserDto.UserId)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return 
	}

	c.JSON(http.StatusOK, domain.SucessResponse{Message: "delelted sucessfully"})

}

func(umc *UserManagementController)FetchByRoleId(c *gin.Context){

	var FetchByRoleNameDto domain.FetchByRoleNameDto

	err := c.ShouldBindJSON(&FetchByRoleNameDto)
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
			default:
				errorMessage["UnknownField"] = "Invalid field provided"
			}	
		}
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Validation failed", Errors:  errorMessage})
		return 
	}

	users, err := umc.UserManagementUsecase.FetchByRoleId(c,  FetchByRoleNameDto.RoleId)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return 
	}
	c.JSON(http.StatusOK, users)
}

func(umc *UserManagementController)FetchByUserName(c *gin.Context){
	var FetchByRoleNameDto domain.FetchByUserNameDto

	err := c.ShouldBindJSON(&FetchByRoleNameDto)
	if err != nil{
		var validationErrors validator.ValidationErrors
		if errors, ok := err.(validator.ValidationErrors); ok{
			validationErrors = errors
		}

		errorMessage := make(map[string]string)
		for _, e := range validationErrors{

			field := e.Field()
			switch field {
			case "UserName":
				errorMessage["UserName"] = "UserName is required"
			default:
				errorMessage["UnknownField"] = "Invalid field provided"
			}	
		}
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Validation failed", Errors:  errorMessage})
		return 
	}
	 
	user, err := umc.UserManagementUsecase.FetchByUserName(c, FetchByRoleNameDto.UserName)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return 
	}
	
	c.JSON(http.StatusOK, user)
}

func(umc *UserManagementController)FetchUsers(c *gin.Context){
	
	users, err := umc.UserManagementUsecase.FetchUsers(c)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return 
	}

	c.JSON(http.StatusOK, users)
}