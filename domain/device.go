package domain

import "context"

type CreateDeviceDto struct{
	SerialNumber string  `json:"serial_number"`
    Port string  `json:"port"`
    IpAddress string  `json:"ip_address"`
    Name string  `json:"name"`
    Campus string  `json:"campus"`
    BlockNumber string  `json:"block_number"`
    RegisteredBy string  `json:"registered_by"`
}

type UpdateDeviceDto struct{
	DeviceId int32 `json:"device_id"`
	SerialNumber string  `json:"serial_number"`
    Port string  `json:"port"`
    IpAddress string  `json:"ip_address"`
    Name string  `json:"name"`
    Campus string  `json:"campus"`
    BlockNumber string  `json:"block_number"`
    RegisteredBy string  `json:"registered_by"`
}

type Device struct{
	DeviceId int32 `json:"device_id"`
	SerialNumber string  `json:"serial_number"`
    Port string  `json:"port"`
    IpAddress string  `json:"ip_address"`
    Name string  `json:"name"`
    Campus string  `json:"campus"`
    BlockNumber string  `json:"block_number"`
    RegisteredBy string  `json:"registered_by"`
	RegisteredDate string `json:"registered_date"`
}


type DeleteDeviceDto struct{
    DeviceId int32  `json:"device_id"`
}

type FetchByCampusDto struct{
    CampusName string `json:"campus_name"`
}


type DeviceUsecase interface {
	Create(c context.Context, createDeviceDto *CreateDeviceDto) error
	Update(c context.Context, updateDeviceDto *UpdateDeviceDto) error
	Delete(c context.Context, deviceId int32)                   error
	FetchDevices(c context.Context)                             ([]Device, error)
	FetchByCampus(c context.Context, campusName string)         ([]Device, error)
}


type DeviceRepository interface{
	Create(c context.Context, device *CreateDeviceDto)           error
	Update(c context.Context, updateDeviceDto *UpdateDeviceDto)  error
	Delete(c context.Context, deviceId int32)                    error
	FetchDevices(c context.Context)                              ([]Device, error)
	FetchByCampus(c context.Context, campusName string)          ([]Device, error)
}




