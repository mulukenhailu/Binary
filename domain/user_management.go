/* 
This domain is responsible for user management, including the following functionalities:  

- Updating users  
- Deleting users  
- Fetching all available users  
- filltering user based on roleName

The functionality for creating new users is handled by the **Sign-up** domain.  
*/


package domain

import (
	"context"
)


type UserManagementUsecase interface{
	Update(c context.Context, updateUserDto *UpdateUserDto) error 
	Delete(c context.Context, userId int32) 				error
	FetchByRoleId(c context.Context, roleId int32)			([]User, error)
	FetchByUserName(c context.Context, userName string)		(User, error)
	FetchUsers(c context.Context)							([]User, error)
}

