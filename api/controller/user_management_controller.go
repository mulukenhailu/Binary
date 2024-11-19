package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mulukenhailu/Binary/domain"
)

type UserManagementController struct {
	UserManagementUsecase domain.UserManagementUsecase
}


func(umc *UserManagementController)Update(c *gin.Context){
	var updateUserDto domain.UpdateUserDto

	err := c.ShouldBindJSON(&updateUserDto)
	if err != nil{
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
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
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
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
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
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
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
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