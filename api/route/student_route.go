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

	 group.GET("/student/all", rc.FetchStudents)
	 group.POST("/student/new", rc.Create)
	 group.PUT("/student/update", rc.Update)
	 group.GET("/student/studentid/:studentId", rc.FetchByStudentId)
	 group.DELETE("/student/delete/:studentInformationId", rc.Delete)
	
}