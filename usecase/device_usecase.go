package usecase

import (
	"context"
	"github.com/mulukenhailu/Binary/domain"
	"time"
)

type deviceUsecase struct {
	deviceRepository domain.DeviceRepository
	contextTimeout   time.Duration
}

func NewDeviceUsecase(deviceRepository domain.DeviceRepository, timeout time.Duration) domain.DeviceUsecase {
	return &deviceUsecase{
		deviceRepository: deviceRepository,
		contextTimeout:   timeout,
	}
}



// Create implements domain.DeviceUsecase.
func (du *deviceUsecase) Create(c context.Context, device *domain.CreateDeviceDto) error {
	ctx, cancel := context.WithTimeout(c, du.contextTimeout)
	defer cancel()
	return du.deviceRepository.Create(ctx, device)
}

// Delete implements domain.DeviceUsecase.
func (du *deviceUsecase) Delete(c context.Context, deviceId int32) error {
	ctx, cancel := context.WithTimeout(c, du.contextTimeout)
	defer cancel()
	return du.deviceRepository.Delete(ctx, deviceId)
}

// FetchByCampus implements domain.DeviceUsecase.
func (du *deviceUsecase) FetchByCampus(c context.Context, campusName string) ([]domain.Device, error) {
	ctx, cancel := context.WithTimeout(c, du.contextTimeout)
	defer cancel()
	return du.deviceRepository.FetchByCampus(ctx, campusName)
}

// FetchDevices implements domain.DeviceUsecase.
func (du *deviceUsecase) FetchDevices(c context.Context) ([]domain.Device, error) {
	ctx, cancel := context.WithTimeout(c, du.contextTimeout)
	defer cancel()
	return du.deviceRepository.FetchDevices(ctx)
}

// Update implements domain.DeviceUsecase.
func (du *deviceUsecase) Update(c context.Context, updateDeviceDto *domain.UpdateDeviceDto) error {
	ctx, cancel := context.WithTimeout(c, du.contextTimeout)
	defer cancel()
	return du.deviceRepository.Update(ctx, updateDeviceDto)
}


