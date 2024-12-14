package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mulukenhailu/Binary/domain"
	"github.com/mulukenhailu/Binary/internal/utils"
)

type DeviceController struct {
	DeviceUsecase domain.DeviceUsecase
}


func (dc *DeviceController)Create(c *gin.Context){
	var DeviceDto domain.CreateDeviceDto

	err := c.ShouldBindJSON(&DeviceDto)
	if err != nil{

		var validationErrors validator.ValidationErrors
		if errors, ok := err.(validator.ValidationErrors); ok{
			validationErrors = errors
		}

		errorMessage := make(map[string]string)
		for _, e := range validationErrors{

			field := e.Field()
			switch field {
			case "SerialNumber":
				errorMessage["SerialNumber"] = "SerialNumber is required"
			case "Port":
				errorMessage["Port"] = "Port is required"
			case "IpAddress":
				errorMessage["IpAddress"] = "IpAddress is required"
			case "Name":
				errorMessage["Name"] = "Name is required"
			case "Campus":
				errorMessage["Campus"] = "Campus is required"
			case "BlockNumber":
				errorMessage["BlockNumber"] = "BlockNumber is required"
			case "RegisteredBy":
				errorMessage["RegisteredBy"] = "RegisteredBy is required"
			default:
				errorMessage["UnknownField"] = "Invalid field provided"
			}
			
		}
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Validation failed", Errors:  errorMessage})
		return 
	}

	DeviceDto.Campus = utils.ConvertToSmallLetter(DeviceDto.Campus)

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
		var validationErrors validator.ValidationErrors
		if errors, ok := err.(validator.ValidationErrors); ok{
			validationErrors = errors
		}

		errorMessage := make(map[string]string)
		for _, e := range validationErrors{

			field := e.Field()
			switch field {
			case "DeviceId":
				errorMessage["DeviceId"] = "DeviceId is required"
			case "SerialNumber":
				errorMessage["SerialNumber"] = "SerialNumber is required"
			case "Port":
				errorMessage["Port"] = "Port is required"
			case "IpAddress":
				errorMessage["IpAddress"] = "IpAddress is required"
			case "Name":
				errorMessage["Name"] = "Name is required"
			case "Campus":
				errorMessage["Campus"] = "Campus is required"
			case "BlockNumber":
				errorMessage["BlockNumber"] = "BlockNumber is required"
			case "RegisteredBy":
				errorMessage["RegisteredBy"] = "RegisteredBy is required"
			default:
				errorMessage["UnknownField"] = "Invalid field provided"
			}
			
		}
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Validation failed", Errors:  errorMessage})
		return 
	}

	updateDeviceDto.Campus = utils.ConvertToSmallLetter(updateDeviceDto.Campus)
	err = dc.DeviceUsecase.Update(c, &updateDeviceDto)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return 
	}

	c.JSON(http.StatusOK, domain.SucessResponse{Message: "device updated."})
	
}

func (dc *DeviceController)Delete(c *gin.Context){
	deviceId := c.Param("deviceId")
	
	bit32, err := utils.ConverParamID(deviceId)
	if err != nil{
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = dc.DeviceUsecase.Delete(c, *bit32)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return 
	}

	c.JSON(http.StatusOK, domain.SucessResponse{Message: "device deleted."})

}



func (dc *DeviceController)FetchByCampus(c *gin.Context){
	campusName := c.Param("campusName")
	campusName = utils.ConvertToSmallLetter((campusName))


	devices, err := dc.DeviceUsecase.FetchByCampus(c, campusName)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return 
	}

	c.JSON(http.StatusOK, devices)
}

func (dc *DeviceController)FetchDevices(c *gin.Context){
	devices, err := dc.DeviceUsecase.FetchDevices(c)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return 
	}
	c.JSON(http.StatusOK, devices)
}