package route

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/mulukenhailu/Binary/api/controller"
	"github.com/mulukenhailu/Binary/repository"
	"github.com/mulukenhailu/Binary/bootstrap"
	"github.com/mulukenhailu/Binary/usecase"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db *pgxpool.Pool, group *gin.RouterGroup) {
	lr := repository.NewUserRepository(db)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(lr, timeout),
		Env:env,
	}

	group.POST("/user/login", lc.Login)
}