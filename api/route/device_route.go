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

func NewDeviceRouter(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	dr := repository.NewdeviceRepository(db)

	dc := &controller.DeviceController{
		DeviceUsecase: usecase.NewDeviceUsecase(dr, timeout),
	}

	group.GET("/devicebycampus", dc.FetchByCampus)
	group.GET("/devices", dc.FetchDevices)
	group.POST("/newdevice", dc.Create)
	group.PUT("/updatedevice", dc.Update)
	group.DELETE("deletedevice", dc.Delete)

}
