package utils

import (
	"github.com/mulukenhailu/Binary/domain"
	"github.com/mulukenhailu/Binary/internal/database"
)

func ConvertDbRolesToDomainRoles(dbRoles []database.Role) []domain.Role{
	domainRoles := make([]domain.Role, len(dbRoles))
	for i, dbRole := range dbRoles{
		domainRoles[i] = domain.Role{
			RoleId 			: dbRole.Roleid,
			RoleName 		: dbRole.Rolename,
			RegisteredBy 	: dbRole.Registeredby,
			RegisteredDate 	: dbRole.Registereddate.Time.String(),
		}
	}
	return domainRoles
}

func ConvertDbDevicesToDomainDevices(dbDevices []database.Device)[]domain.Device{
	domainDevices := make([]domain.Device, len(dbDevices))
	for i, dbDevice := range dbDevices{
		domainDevices[i] = domain.Device{
			DeviceId 			: dbDevice.Deviceid,
			SerialNumber  		: dbDevice.Serialnumber.String,
			Port 				: dbDevice.Port,
			IpAddress  			: dbDevice.Ipaddress,
			Name  				: dbDevice.Name.String,
			Campus  			: dbDevice.Campus.String,
			BlockNumber  		: dbDevice.Blocknumber.String,
			RegisteredBy  		: dbDevice.Registeredby,
			RegisteredDate  	: dbDevice.Registereddate.Time.String(),
		}
	}
	return domainDevices
}