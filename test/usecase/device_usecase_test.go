package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/mulukenhailu/Binary/domain"
	"github.com/mulukenhailu/Binary/domain/mocks"
	"github.com/mulukenhailu/Binary/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateDevice(t *testing.T) {

	mockDeviceReposoitory := mocks.NewDeviceRepository(t)
	mockCreateDeviceDto := &domain.CreateDeviceDto{
		SerialNumber : "test",
		Port : "test",
		IpAddress : "test",
		Name : "test",
		Campus : "test",
		BlockNumber : "test",
		RegisteredBy : "test",
	}

	

	t.Run("success", func(t *testing.T){
		mockDeviceReposoitory.On("Create", mock.Anything, mockCreateDeviceDto).Return(nil).Once()

		uc := usecase.NewDeviceUsecase(mockDeviceReposoitory, time.Second * 2)
		err := uc.Create(context.Background(), mockCreateDeviceDto)

		assert.NoError(t, err)
		mockDeviceReposoitory.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T){
		mockDeviceReposoitory.On("Create", mock.Anything, mockCreateDeviceDto).Return(errors.New("unexpected")).Once()

		uc := usecase.NewDeviceUsecase(mockDeviceReposoitory, time.Second * 2)
		err := uc.Create(context.Background(), mockCreateDeviceDto)

		assert.Error(t, err)
		mockDeviceReposoitory.AssertExpectations(t)
	})

}


func TestUpdateDevice(t *testing.T){
	mockDeviceReposoitory := mocks.NewDeviceRepository(t)
	mockUpdateDeviceDto := &domain.UpdateDeviceDto{
		DeviceId 		: 1,
		SerialNumber 	: "test",
		Port 			: "test",
		IpAddress 		: "test",
		Name 			: "test",
		Campus 			: "test",
		BlockNumber 	: "test",
		RegisteredBy 	: "test",
	}

	t.Run("success", func(t *testing.T){
		mockDeviceReposoitory.On("Update", mock.Anything, mockUpdateDeviceDto).Return(nil).Once()

		uc := usecase.NewDeviceUsecase(mockDeviceReposoitory, time.Second * 2)
		err := uc.Update(context.Background(), mockUpdateDeviceDto)

		assert.NoError(t, err)
		mockDeviceReposoitory.AssertExpectations((t))
	})

	t.Run("fail", func(t *testing.T){
		mockDeviceReposoitory.On("Update", mock.Anything, mockUpdateDeviceDto).Return(errors.New("enexpected")).Once()

		uc := usecase.NewDeviceUsecase(mockDeviceReposoitory, time.Second * 2)
		err := uc.Update(context.Background(), mockUpdateDeviceDto)

		assert.Error(t, err)
		mockDeviceReposoitory.AssertExpectations((t))
	})
}


func TestDeleteDevice(t *testing.T){
	mockDeviceReposoitory := mocks.NewDeviceRepository(t)
	mockDeleteDeviceDto := domain.DeleteDeviceDto{
		DeviceId: 1,
	}

	t.Run("success", func(t *testing.T){
		mockDeviceReposoitory.On("Delete", mock.Anything, mockDeleteDeviceDto.DeviceId).Return(nil).Once()

		uc := usecase.NewDeviceUsecase(mockDeviceReposoitory, time.Second * 2)
		err := uc.Delete(context.Background(), mockDeleteDeviceDto.DeviceId)

		assert.NoError(t, err)
		mockDeviceReposoitory.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T){
		mockDeviceReposoitory.On("Delete", mock.Anything, mockDeleteDeviceDto.DeviceId).Return(errors.New("enexpected")).Once()

		uc := usecase.NewDeviceUsecase(mockDeviceReposoitory, time.Second * 2)
		err := uc.Delete(context.Background(), mockDeleteDeviceDto.DeviceId)

		assert.Error(t, err)
		mockDeviceReposoitory.AssertExpectations(t)
	})
}


func TestFetchDevices(t *testing.T){
	mockDeviceReposoitory := mocks.NewDeviceRepository(t)
	mockDomainDevices := []domain.Device{
		{
			DeviceId 		: 1,
			SerialNumber 	:"test1",
			Port 			:"test1",
			IpAddress 		:"test",
			Name 			:"test",
			Campus			:"test",
			BlockNumber 	:"test",
			RegisteredBy 	:"test",
			RegisteredDate 	:"test",
		},
		{
			DeviceId 		: 2,
			SerialNumber 	:"test2",
			Port 			:"test2",
			IpAddress 		:"test",
			Name 			:"test",
			Campus			:"test",
			BlockNumber 	:"test",
			RegisteredBy 	:"test",
			RegisteredDate 	:"test",
		},
	}

	t.Run("success", func(t *testing.T){
		mockDeviceReposoitory.On("FetchDevices", mock.Anything).Return(mockDomainDevices, nil).Once()

		uc := usecase.NewDeviceUsecase(mockDeviceReposoitory, time.Second * 2)
		devices, err := uc.FetchDevices(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, devices, mockDomainDevices)
		assert.Len(t, devices, len(mockDomainDevices))
		mockDeviceReposoitory.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T){
		mockDeviceReposoitory.On("FetchDevices", mock.Anything).Return([]domain.Device{}, errors.New("unexpected")).Once()

		uc := usecase.NewDeviceUsecase(mockDeviceReposoitory, time.Second * 2)
		devices, err := uc.FetchDevices(context.Background())

		assert.Error(t, err)
		assert.Equal(t, devices, []domain.Device{})
		assert.Len(t, devices, 0)
		mockDeviceReposoitory.AssertExpectations(t)
	})
}


func TestFetchByCampus(t *testing.T){
	mockDeviceReposoitory := mocks.NewDeviceRepository(t)

	mockFetchDeviceByCampusDto := domain.FetchByCampusDto{
		CampusName: "test",
	}

	mockDomainDevices := []domain.Device{
		{
			DeviceId 		: 1,
			SerialNumber 	:"test1",
			Port 			:"test1",
			IpAddress 		:"test",
			Name 			:"test",
			Campus			:"test",
			BlockNumber 	:"test",
			RegisteredBy 	:"test",
			RegisteredDate 	:"test",
		},
		{
			DeviceId 		: 2,
			SerialNumber 	:"test2",
			Port 			:"test2",
			IpAddress 		:"test",
			Name 			:"test",
			Campus			:"test",
			BlockNumber 	:"test",
			RegisteredBy 	:"test",
			RegisteredDate 	:"test",
		},
	}
	t.Run("success", func(t *testing.T){
		mockDeviceReposoitory.On("FetchByCampus", mock.Anything, mockFetchDeviceByCampusDto.CampusName).Return(mockDomainDevices, nil).Once()

		uc := usecase.NewDeviceUsecase(mockDeviceReposoitory, time.Second * 2)
		devices, err := uc.FetchByCampus(context.Background(), mockFetchDeviceByCampusDto.CampusName)

		assert.NoError(t, err)
		assert.Equal(t, devices, mockDomainDevices)
		assert.Len(t, devices, len(mockDomainDevices))
		mockDeviceReposoitory.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T){
		mockDeviceReposoitory.On("FetchByCampus", mock.Anything, mockFetchDeviceByCampusDto.CampusName).Return([]domain.Device{}, errors.New("unexpected")).Once()

		uc := usecase.NewDeviceUsecase(mockDeviceReposoitory, time.Second * 2)
		devices, err := uc.FetchByCampus(context.Background(), mockFetchDeviceByCampusDto.CampusName)

		assert.Error(t, err)
		assert.Equal(t, devices, []domain.Device{})
		assert.Len(t, devices, 0)
		mockDeviceReposoitory.AssertExpectations(t)
	})

}

