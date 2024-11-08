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

func TestCreat(t *testing.T) {

	mockRoleRepository := mocks.NewRoleRespository(t)
	mockAddRole := &domain.RoleDto{
		RoleName: "test",
		RegisteredBy : "test",
	}
	t.Run("success", func(t *testing.T){
		
		mockRoleRepository.On("Create", mock.Anything, mockAddRole).Return(nil).Once()

		uc := usecase.NewRoleUsecase(mockRoleRepository, time.Second * 2)
		err := uc.Create(context.Background(), mockAddRole)

		assert.NoError(t, err)
		mockRoleRepository.AssertExpectations(t)
	})

	t.Run("erro", func(t *testing.T) {
		mockRoleRepository.On("Create", mock.Anything, mockAddRole).Return(errors.New("unexpected")).Once()
		uc := usecase.NewRoleUsecase(mockRoleRepository, time.Second * 2)

		err := uc.Create(context.Background(), mockAddRole)

		assert.Error(t, err)
		mockRoleRepository.AssertExpectations(t)

	})
}

func TestUpdate(t *testing.T){
	mockRoleRepository := mocks.NewRoleRespository(t)
	mockUpdateRoleDto := &domain.UpdateRoleDto{
		PreviousRoleName : "test",
		NewRoleName : "test",
		NewRegisterName : "test",
	}

	t.Run("success", func(t *testing.T) {
		mockRoleRepository.On("Update", mock.Anything, mockUpdateRoleDto).Return(nil).Once()

		uc := usecase.NewRoleUsecase(mockRoleRepository, time.Second * 2)
		err := uc.Update(context.Background(), mockUpdateRoleDto)

		assert.NoError(t, err)
		mockRoleRepository.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T){
		mockRoleRepository.On("Update", mock.Anything, mockUpdateRoleDto).Return(errors.New("unexpected")).Once()

		uc := usecase.NewRoleUsecase(mockRoleRepository, time.Second * 2)
		err := uc.Update(context.Background(), mockUpdateRoleDto)

		assert.Error(t, err)
		mockRoleRepository.AssertExpectations(t)

	})
}
