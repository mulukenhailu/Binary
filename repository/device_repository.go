package repository

import (
	"context"

	"github.com/mulukenhailu/Binary/domain"
	"github.com/mulukenhailu/Binary/internal/database"
	"github.com/mulukenhailu/Binary/internal/utils"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/pgtype"
)

type deviceRepository struct {
	pg *database.Queries
}


func NewdeviceRepository(db *pgxpool.Pool) domain.DeviceRepository {
	return &deviceRepository{
		pg: database.New(db),
	}
}

// Create implements domain.DeviceRepository.
func (dr *deviceRepository) Create(c context.Context, device *domain.CreateDeviceDto) error {
	deviceParam := database.CreateDeviceParams{
		Serialnumber : pgtype.Text{String: device.SerialNumber, Valid: device.SerialNumber != ""},
		Port         : device.Port,
		Ipaddress    : device.IpAddress,
		Name         : pgtype.Text{String: device.Name, Valid: device.Name != ""},
		Campus       : pgtype.Text{String: device.Campus, Valid: device.Campus != ""}, 
		Blocknumber  : pgtype.Text{String: device.BlockNumber, Valid: device.BlockNumber != ""},
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
	dbDevices, err := dr.pg.FetchByCampus(c, pgtype.Text{String: campusName, Valid: campusName != ""})
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
		Serialnumber : pgtype.Text{String: updateDeviceDto.SerialNumber, Valid: updateDeviceDto.SerialNumber != ""},
		Port         : updateDeviceDto.Port,
		Ipaddress    : updateDeviceDto.IpAddress,
		Name         : pgtype.Text{String: updateDeviceDto.Name, Valid: updateDeviceDto.Name != ""},
		Campus       : pgtype.Text{String: updateDeviceDto.Campus, Valid: updateDeviceDto.Campus != ""}, 
		Blocknumber  : pgtype.Text{String: updateDeviceDto.BlockNumber, Valid: updateDeviceDto.BlockNumber != ""},
		Registeredby : updateDeviceDto.RegisteredBy,
	}

	_, err := dr.pg.UpdateDevice(c, updateParam)
	return err
}


