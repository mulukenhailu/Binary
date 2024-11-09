package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mulukenhailu/Binary/domain"
)

type RoleController struct {
	RoleUsecase domain.RoleUsecase
}



func (rc *RoleController) Create(c *gin.Context){
	var RoleDto domain.RoleDto

	err := c.ShouldBindJSON(&RoleDto)
	if err != nil{
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return 
	}

	err = rc.RoleUsecase.Create(c, &RoleDto)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message:err.Error()})
		return 
	}

	c.JSON(http.StatusOK, domain.SucessResponse{Message: "role created."})
}


func (rc *RoleController) Delete(c *gin.Context){
	var DeleteRoleDto domain.DeleteRoleDto

	err := c.ShouldBindJSON(&DeleteRoleDto)
	if err != nil{
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return 
	}

	err = rc.RoleUsecase.Delete(c, DeleteRoleDto.RoleId)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message:err.Error()})
		return 
	}

	c.JSON(http.StatusOK, domain.SucessResponse{Message: "role Deleted."})
}


func (rc *RoleController) Update(c *gin.Context){

	var UpdateRoleDto domain.UpdateRoleDto
	err := c.ShouldBindJSON(&UpdateRoleDto)
	if err != nil{
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return 
	}

	err = rc.RoleUsecase.Update(c, &UpdateRoleDto)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message:err.Error()})
		return 
	}

	c.JSON(http.StatusOK, domain.SucessResponse{Message: "role Updated."})
}


func (rc *RoleController) FetchRoles(c *gin.Context){
	
	roles, err := rc.RoleUsecase.FetchRoles(c)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message:err.Error()})
		return 
	}
	c.JSON(http.StatusOK, roles)
}


func (rc *RoleController) FetchByName(c *gin.Context){
	var FetchByNameDto domain.FetchByNameDto

	err := c.ShouldBindJSON(&FetchByNameDto)
	if err != nil{
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return 
	}

	role, err :=rc.RoleUsecase.FetchByName(c, FetchByNameDto.RoleName)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return 
	}

	c.JSON(http.StatusOK, role)
}