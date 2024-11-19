package route

import (
	"database/sql"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/mulukenhailu/Binary/api/controller"
	"github.com/mulukenhailu/Binary/bootstrap"
	"github.com/mulukenhailu/Binary/repository"
	"github.com/mulukenhailu/Binary/usecase"
)

func NewUserManagerRouter(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup){
	umr := repository.NewUserRepository(db)
	umc := &controller.UserManagementController{
		UserManagementUsecase: usecase.NewUserManagerUsecase(umr, timeout),
	}

	group.PUT("/updateuser", umc.Update)
	group.DELETE("deleteuser", umc.Delete)
	group.GET("/userbyrolename", umc.FetchByRoleId)
	group.GET("/userbyusername", umc.FetchByUserName)
	group.GET("/allusers", umc.FetchUsers)
}