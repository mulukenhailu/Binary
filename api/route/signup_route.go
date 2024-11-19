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

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	sr := repository.NewUserRepository(db)
	sc := &controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(sr, timeout),
	}

	group.POST("/signup", sc.Signup)
}