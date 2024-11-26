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

func NewRoleRouter(env *bootstrap.Env, timeout time.Duration, db *pgxpool.Pool, group *gin.RouterGroup) {

	rr := repository.NewRoleRepository(db)

	rc := &controller.RoleController{
		RoleUsecase:  usecase.NewRoleUsecase(rr, timeout),
	}

	group.GET("/role", rc.FetchByName)
	group.GET("/roles", rc.FetchRoles)
	group.POST("/newrole", rc.Create)
	group.PUT("/updaterole", rc.Update)
	group.DELETE("/deleterole", rc.Delete)

}

