package route

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/mulukenhailu/Binary/api/controller"
	"github.com/mulukenhailu/Binary/bootstrap"
	"github.com/mulukenhailu/Binary/repository"
	"github.com/mulukenhailu/Binary/usecase"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewUserManagerRouter(env *bootstrap.Env, timeout time.Duration, db *pgxpool.Pool, group *gin.RouterGroup){
	umr := repository.NewUserRepository(db)
	umc := &controller.UserManagementController{
		UserManagementUsecase: usecase.NewUserManagerUsecase(umr, timeout),
	}

	group.PUT("/user/updateuser", umc.Update)
	group.DELETE("/user/deleteuser", umc.Delete)
	group.GET("/user/userbyroleid", umc.FetchByRoleId)
	group.GET("/user/userbyusername", umc.FetchByUserName)
	group.GET("/user/allusers", umc.FetchUsers)
}