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


func NewStudentRouter(env *bootstrap.Env, timeout time.Duration, db *pgxpool.Pool, group *gin.RouterGroup){
	 sr := repository.NewStudentRepository(db)

	 rc := &controller.StudentController{
		StudentUsecase: usecase.NewStudentUsecase(sr, timeout),
	 }

	 group.POST("/student/create", rc.Create)
	 group.POST("/student/update", rc.Update)
}