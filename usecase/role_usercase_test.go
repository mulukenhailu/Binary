package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/mulukenhailu/Binary/domain"
	"github.com/mulukenhailu/Binary/domain/mocks"
	"github.com/mulukenhailu/Binary/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreat(t *testing.T) {

	mockRoleRepository := mocks.NewRoleRespository(t)
	mockAddRole := &domain.RoleDto{
		RoleName: "admin",
		RegisteredBy : "admin",
	}


	t.Run("success", func(t *testing.T){
		
		mockRoleRepository.On("Create", mock.Anything, mockAddRole).Return(nil).Once()

		u := usecase.NewRoleUsecase(mockRoleRepository)
		err := u.Create(context.Background(), mockAddRole)

		assert.NoError(t, err)
		mockRoleRepository.AssertExpectations(t)
	})

	t.Run("erro", func(t *testing.T) {
		mockRoleRepository.On("Create", mock.Anything, mockAddRole).Return(errors.New("Unexpected")).Once()
		u := usecase.NewRoleUsecase(mockRoleRepository)

		err := u.Create(context.Background(), mockAddRole)

		assert.Error(t, err)
		mockRoleRepository.AssertExpectations(t)

	})
}