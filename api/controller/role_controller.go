package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mulukenhailu/Binary/domain"
)

type RoleController struct {
	RoleUsecase domain.RoleUsercase
}



func (rc *RoleController) Create(c *gin.Context){
	var role domain.RoleDto

	err := c.ShouldBindJSON(&role)
	if err != nil{
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return 
	}

	err = rc.RoleUsecase.Create(c, &role)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message:err.Error()})
		return 
	}

	c.JSON(http.StatusOK, domain.SucessResponse{Message: "role created."})
}


func (rc *RoleController) Delete(c *gin.Context){
	
}


func (rc *RoleController) Update(c *gin.Context){
	
}


func (rc *RoleController) FetchRoles(c *gin.Context){
	
}


func (rc *RoleController) FetchByName(c *gin.Context){
	
}