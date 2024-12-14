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

func NewDeviceRouter(env *bootstrap.Env, timeout time.Duration, db *pgxpool.Pool, group *gin.RouterGroup) {
	dr := repository.NewdeviceRepository(db)

	dc := &controller.DeviceController{
		DeviceUsecase: usecase.NewDeviceUsecase(dr, timeout),
	}

	
	group.GET("/device/all", dc.FetchDevices)
	group.POST("/device/new", dc.Create)
	group.PUT("/device/update", dc.Update)
	group.GET("/device/campusname/:campusName", dc.FetchByCampus)
	group.DELETE("/device/delete/:deviceId", dc.Delete)

}
