package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mulukenhailu/Binary/domain"
	"github.com/mulukenhailu/Binary/internal/utils"
)

type RoleController struct {
	RoleUsecase domain.RoleUsecase
}



func (rc *RoleController) Create(c *gin.Context){
	var RoleDto domain.CreateRoleDto

	err := c.ShouldBindJSON(&RoleDto)
	if err != nil{

		var validationErrors validator.ValidationErrors
		if errors, ok := err.(validator.ValidationErrors); ok{
			validationErrors = errors
		}

		errorMessage := make(map[string]string)
		for _, e := range validationErrors{

			field := e.Field()
			switch field {
			case "RoleName":
				errorMessage["RoleName"] = "RoleName is required"
			case "RegisteredBy":
				errorMessage["RegisteredBy"] = "RegisteredBy is required"
			default:
				errorMessage["UnknownField"] = "Invalid field provided"
			}
			
		}

		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Validation failed", Errors:  errorMessage})
		return 
	}

	RoleDto.RoleName = utils.ConvertToSmallLetter(RoleDto.RoleName)
	
	err = rc.RoleUsecase.Create(c, &RoleDto)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message:err.Error()})
		return 
	}

	c.JSON(http.StatusCreated, domain.SucessResponse{Message: "role created."})
}

func (rc *RoleController) Update(c *gin.Context){

	var UpdateRoleDto domain.UpdateRoleDto
	err := c.ShouldBindJSON(&UpdateRoleDto)
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
				errorMessage["RoleId"] = "RoleId is required"
			case "RoleName":
				errorMessage["RoleName"] = "RoleName is required"
			case "RegisteredBy":
				errorMessage["RegisteredBy"] = "RegisteredBy is required"
			default:
				errorMessage["UnknownField"] = "Invalid field provided"
			}
			
		}

		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Validation failed", Errors:  errorMessage})
		return 
	}

	UpdateRoleDto.RoleName = utils.ConvertToSmallLetter(UpdateRoleDto.RoleName)
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


func (rc *RoleController) Delete(c *gin.Context){
	roleId := c.Param("roleId")

	bit32, err := utils.ConverParamID(roleId)
	if err != nil{
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = rc.RoleUsecase.Delete(c, *bit32)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message:err.Error()})
		return 
	}

	c.JSON(http.StatusOK, domain.SucessResponse{Message: "role Deleted."})
}


func (rc *RoleController) FetchByName(c *gin.Context){
	roleName := c.Param("roleName")
	roleName = utils.ConvertToSmallLetter(roleName)
	
	role, err :=rc.RoleUsecase.FetchByName(c, roleName)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return 
	}

	c.JSON(http.StatusOK, role)
}