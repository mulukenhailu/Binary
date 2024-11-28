package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mulukenhailu/Binary/api/controller"
	"github.com/mulukenhailu/Binary/bootstrap"
	"github.com/mulukenhailu/Binary/repository"
	"github.com/mulukenhailu/Binary/usecase"
)

func NewRolePermissionRouter(env *bootstrap.Env, timeout time.Duration, db *pgxpool.Pool, group *gin.RouterGroup) {
	rpr := repository.NewRolePermissionRepository(db)

	rpc := &controller.RolePermissionController{
		RolePermissionUsecase: usecase.NewRolePermissionUsecase(rpr, timeout),
	}

	group.POST("/rolepermission/create", rpc.Create)
}