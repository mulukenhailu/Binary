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

func NewRoleRouter(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {

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

