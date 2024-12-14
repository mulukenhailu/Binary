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

	group.GET("/rolepermission/all", rpc.FetchRolePermissions)
	group.GET("/rolepermission/roleid/:roleId", rpc.FetchByRoleId)
	group.GET("/rolepermission/permissionid/:permissionId", rpc.FetchByPermissionId)
	group.POST("/rolepermission/new", rpc.Create)
	group.POST("/rolepermission/update", rpc.Update)
	
}