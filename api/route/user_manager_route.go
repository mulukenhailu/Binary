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

	group.GET("/user/all", umc.FetchUsers)
	group.PUT("/user/update", umc.Update)
	group.GET("/user/username/:userName", umc.FetchByUserName)
	group.GET("/user/role/:roleId", umc.FetchByRoleId)
	group.DELETE("/user/delete/:userId", umc.Delete)
}