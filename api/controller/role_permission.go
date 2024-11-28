package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mulukenhailu/Binary/domain"
	"github.com/go-playground/validator/v10"
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

func (rpc *RolePermissionController) Update(c *gin.Context) {

}

func (rpc *RolePermissionController) Delete(c *gin.Context) {

}

func (rpc *RolePermissionController) FetchByPermissionId(c *gin.Context) {

}

func (rpc *RolePermissionController) FetchByRoleId(c *gin.Context) {

}
