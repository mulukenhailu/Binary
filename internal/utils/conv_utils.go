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

func ConvertDbUserToDomainUser(dbUsers []database.Appuser)[]domain.User{
	domainUsers := make([]domain.User, len(dbUsers))

	for i, dbUser := range dbUsers{
		domainUsers[i] = domain.User{
			UserId          :dbUser.Userid,
			RoleId          :dbUser.Roleid,
			UserName        :dbUser.Username,
			FirstName       :dbUser.Firstname,
			FatherName      :dbUser.Fathername,
			GrandFatherName :dbUser.Grandfathername,
			Password        :dbUser.Password,
			PhoneNumber     :dbUser.Phonenumber,
			Address         :dbUser.Address,
			Email          	:dbUser.Email.String,
			RegisteredBy 	:dbUser.Registeredby,
			RegisteredDate 	:dbUser.Registereddate.Time.String(),
		}
	}
	return domainUsers
}

func ConvertDbRolePermToDomainRolePerm(dbRolepermission []database.Rolepermission)[]domain.RolePermission{
	domainRolePermission := make([]domain.RolePermission, len(dbRolepermission))

	for i, dbRolePermission := range dbRolepermission{
		domainRolePermission[i] = domain.RolePermission{
			RolePermissionId : dbRolePermission.Rolepermissionid,
			RoleId       	 : dbRolePermission.Roleid,
			PermissionId  	 : dbRolePermission.Permissinoid,
		}
	}
	return domainRolePermission
}




