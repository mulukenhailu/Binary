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

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db *pgxpool.Pool, group *gin.RouterGroup) {
	sr := repository.NewUserRepository(db)
	sc := &controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(sr, timeout),
	}

	group.POST("/user/signup", sc.Signup)
}