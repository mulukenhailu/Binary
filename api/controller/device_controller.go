package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mulukenhailu/Binary/domain"

)

type DeviceController struct {
	DeviceUsecase domain.DeviceUsecase
}


func (dc *DeviceController)Create(c *gin.Context){
	var DeviceDto domain.CreateDeviceDto

	err := c.ShouldBindJSON(&DeviceDto)
	if err != nil{
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return 
	}

	err = dc.DeviceUsecase.Create(c, &DeviceDto)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return 
	}

	c.JSON(http.StatusCreated, domain.SucessResponse{Message: "device created."})
}

func (dc *DeviceController)Update(c *gin.Context){
	var updateDeviceDto domain.UpdateDeviceDto

	err := c.ShouldBindJSON(&updateDeviceDto)
	if err != nil{
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
	}

	err = dc.DeviceUsecase.Update(c, &updateDeviceDto)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return 
	}

	c.JSON(http.StatusOK, domain.SucessResponse{Message: "device updated."})
	
}

func (dc *DeviceController)Delete(c *gin.Context){
	var DeleteDeviceDto  domain.DeleteDeviceDto

	err := c.ShouldBindJSON(&DeleteDeviceDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return 
	}

	err = dc.DeviceUsecase.Delete(c, DeleteDeviceDto.DeviceId)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return 
	}

	c.JSON(http.StatusOK, domain.SucessResponse{Message: "device deleted."})

}

func (dc *DeviceController)FetchDevices(c *gin.Context){
	devices, err := dc.DeviceUsecase.FetchDevices(c)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return 
	}
	c.JSON(http.StatusOK, devices)
}

func (dc *DeviceController)FetchByCampus(c *gin.Context){
	var FetchByCampusDto domain.FetchByCampusDto

	err := c.ShouldBindJSON(&FetchByCampusDto)
	if err != nil{
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return 
	}

	devices, err := dc.DeviceUsecase.FetchByCampus(c, FetchByCampusDto.CampusName)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return 
	}

	c.JSON(http.StatusOK, devices)
}