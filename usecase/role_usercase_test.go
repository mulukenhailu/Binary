package usecase_test

import (
	"context"
	"testing"

	"github.com/mulukenhailu/Binary/domain"
	"github.com/mulukenhailu/Binary/domain/mocks"
	"github.com/mulukenhailu/Binary/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreat(t *testing.T) {

	mockRoleRepository := mocks.NewRoleRespository(t)

	t.Run("success", func(t *testing.T){
		// mockRole := domain.Role{
		// 	RoleId:1,
		// 	RoleName: "admin",
		// 	RegisteredBy: "admin",
		// 	RegisteredDate:"2024-11-07 18:43:37.301586",
		// }

		mockAddRole := &domain.RoleDto{
			RoleName: "admin",
			RegisteredBy : "admin",
		}

		mockRoleRepository.On("Create", mock.Anything, mockAddRole).Return(nil).Once()

		u := usecase.NewRoleUsecase(mockRoleRepository)
		err := u.Create(context.Background(), mockAddRole)

		assert.NoError(t, err)


		mockRoleRepository.AssertExpectations(t)


	})
}