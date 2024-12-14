package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mulukenhailu/Binary/domain"
	"github.com/mulukenhailu/Binary/internal/utils"
)

type RolePermissionController struct {
	RolePermissionUsecase domain.RolePermissionUsecase
}

func (rpc *RolePermissionController) Create(c *gin.Context) {
	var CreatePermissionDto domain.CreatePermissionDto

	err := c.ShouldBindJSON(&CreatePermissionDto)
	if err != nil{

		var validationErrors validator.ValidationErrors
		if errors, ok := err.(validator.ValidationErrors); ok{
			validationErrors = errors
		}

		errorMessage := make(map[string]string)
		for _, e := range validationErrors{

			field := e.Field()
			switch field{
			case "RoleId":
				errorMessage["RoleId"] = "role_id is required"
			case "PermissionIdList":
				errorMessage["PermissionIdList"] = "PermissionIdList is required"
			default:
				errorMessage["UnknownField"] = "Invalid field provided"
			}
		}
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Validation failed", Errors:  errorMessage})
		return 
	}

	err = rpc.RolePermissionUsecase.Create(c, &CreatePermissionDto)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return 
	}

	c.JSON(http.StatusCreated, domain.SucessResponse{Message: "created successfully."})
}

func (rpc *RolePermissionController) FetchRolePermissions(c *gin.Context) {
	rolePermission, err := rpc.RolePermissionUsecase.FetchRolePermissions(c)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message:err.Error()})
		return 
	}
	c.JSON(http.StatusOK, rolePermission)
}

func (rpc *RolePermissionController) Update(c *gin.Context) {

	var updatePermissionDto domain.UpdatePermissionDto

	err := c.ShouldBindJSON(&updatePermissionDto)
	if err != nil{

		var validationErrors validator.ValidationErrors
		if errors, ok := err.(validator.ValidationErrors); ok{
			validationErrors = errors
		}

		errorMessage := make(map[string]string)
		for _, e := range validationErrors{

			field := e.Field()
			switch field{
			// case "RolePermissionId":
			// 	errorMessage["RolePermissionId"] = "rolePermissionId is required"
			case "RoleId":
				errorMessage["RoleId"] = "role_id is required"
			case "PermissionIdList":
				errorMessage["PermissionIdList"] = "PermissionIdList is required"
			default:
				errorMessage["UnknownField"] = "Invalid field provided"
			}
		}
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Validation failed", Errors:  errorMessage})
		return 
	}

	err = rpc.RolePermissionUsecase.Update(c, &updatePermissionDto)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return 
	}

	c.JSON(http.StatusCreated, domain.SucessResponse{Message: "updated successfully."})

}


func (rpc *RolePermissionController) FetchByRoleId(c *gin.Context) {
	roleId := c.Param("roleId")

	bit32, err := utils.ConverParamID(roleId)
	if err != nil{
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	rolePermission, err := rpc.RolePermissionUsecase.FetchByRoleId(c, *bit32)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message:err.Error()})
		return 
	}

	c.JSON(http.StatusOK, rolePermission)
}


func (rpc *RolePermissionController) FetchByPermissionId(c *gin.Context) {
	roleId := c.Param("permissionId")

	bit32, err := utils.ConverParamID(roleId)
	if err != nil{
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	rolePermission, err := rpc.RolePermissionUsecase.FetchByPermissionId(c, *bit32)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message:err.Error()})
		return 
	}

	c.JSON(http.StatusOK, rolePermission)
}

