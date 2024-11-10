package repository

import (
	"context"
	"database/sql"

	"github.com/mulukenhailu/Binary/domain"
	"github.com/mulukenhailu/Binary/internal/database"
	"github.com/mulukenhailu/Binary/internal/utils"
)

type deviceRepository struct {
	pg *database.Queries
}


func NewdeviceRepository(db *sql.DB) domain.DeviceRepository {
	return &deviceRepository{
		pg: database.New(db),
	}
}

// Create implements domain.DeviceRepository.
func (dr *deviceRepository) Create(c context.Context, device *domain.CreateDeviceDto) error {
	deviceParam := database.CreateDeviceParams{
		Serialnumber : sql.NullString{String: device.SerialNumber, Valid: device.SerialNumber != ""},
		Port         : device.Port,
		Ipaddress    : device.IpAddress,
		Name         : sql.NullString{String: device.Name, Valid: device.Name != ""},
		Campus       : sql.NullString{String: device.Campus, Valid: device.Campus != ""}, 
		Blocknumber  : sql.NullString{String: device.BlockNumber, Valid: device.BlockNumber != ""},
		Registeredby : device.RegisteredBy,
	}
	_, err := dr.pg.CreateDevice(c, deviceParam)
	return err
}

// Delete implements domain.DeviceRepository.
func (dr *deviceRepository) Delete(c context.Context, deviceId int32) error {
	_, err := dr.pg.DeleteDevice(c, deviceId)
	return err
}

// FetchByCampus implements domain.DeviceRepository.
func (dr *deviceRepository) FetchByCampus(c context.Context, campusName string) ([]domain.Device, error) {
	dbDevices, err := dr.pg.FetchByCampus(c, sql.NullString{String: campusName, Valid: campusName != ""})
	domainDevices := utils.ConvertDbDevicesToDomainDevices(dbDevices)
	return domainDevices, err
}

// FetchDevices implements domain.DeviceRepository.
func (dr *deviceRepository) FetchDevices(c context.Context) ([]domain.Device, error) {
	DbDevices, err := dr.pg.FetchDevices(c)
	domainDevices := utils.ConvertDbDevicesToDomainDevices(DbDevices)
	return domainDevices, err
}

// Update implements domain.DeviceRepository.
func (dr *deviceRepository) Update(c context.Context, updateDeviceDto *domain.UpdateDeviceDto) error {
	updateParam := database.UpdateDeviceParams{
		Deviceid     : updateDeviceDto.DeviceId,
		Serialnumber : sql.NullString{String: updateDeviceDto.SerialNumber, Valid: updateDeviceDto.SerialNumber != ""},
		Port         : updateDeviceDto.Port,
		Ipaddress    : updateDeviceDto.IpAddress,
		Name         : sql.NullString{String: updateDeviceDto.Name, Valid: updateDeviceDto.Name != ""},
		Campus       : sql.NullString{String: updateDeviceDto.Campus, Valid: updateDeviceDto.Campus != ""}, 
		Blocknumber  : sql.NullString{String: updateDeviceDto.BlockNumber, Valid: updateDeviceDto.BlockNumber != ""},
		Registeredby : updateDeviceDto.RegisteredBy,
	}

	_, err := dr.pg.UpdateDevice(c, updateParam)
	return err
}


