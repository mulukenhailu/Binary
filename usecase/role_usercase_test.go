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
		RoleName:     "test",
		RegisteredBy: "test",
	}
	t.Run("success", func(t *testing.T) {

		mockRoleRepository.On("Create", mock.Anything, mockAddRole).Return(nil).Once()

		uc := usecase.NewRoleUsecase(mockRoleRepository, time.Second*2)
		err := uc.Create(context.Background(), mockAddRole)

		assert.NoError(t, err)
		mockRoleRepository.AssertExpectations(t)
	})

	t.Run("erro", func(t *testing.T) {
		mockRoleRepository.On("Create", mock.Anything, mockAddRole).Return(errors.New("unexpected")).Once()
		uc := usecase.NewRoleUsecase(mockRoleRepository, time.Second*2)

		err := uc.Create(context.Background(), mockAddRole)

		assert.Error(t, err)
		mockRoleRepository.AssertExpectations(t)

	})
}

func TestUpdate(t *testing.T) {
	mockRoleRepository := mocks.NewRoleRespository(t)
	mockUpdateRoleDto := &domain.UpdateRoleDto{
		PreviousRoleName: "test",
		NewRoleName:      "test",
		NewRegisterName:  "test",
	}

	t.Run("success", func(t *testing.T) {
		mockRoleRepository.On("Update", mock.Anything, mockUpdateRoleDto).Return(nil).Once()

		uc := usecase.NewRoleUsecase(mockRoleRepository, time.Second*2)
		err := uc.Update(context.Background(), mockUpdateRoleDto)

		assert.NoError(t, err)
		mockRoleRepository.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T) {
		mockRoleRepository.On("Update", mock.Anything, mockUpdateRoleDto).Return(errors.New("unexpected")).Once()

		uc := usecase.NewRoleUsecase(mockRoleRepository, time.Second*2)
		err := uc.Update(context.Background(), mockUpdateRoleDto)

		assert.Error(t, err)
		mockRoleRepository.AssertExpectations(t)

	})
}

func TestDelete(t *testing.T) {
	mockRoleRepository := mocks.NewRoleRespository(t)
	mockDeleteRoleDto := domain.DeleteRoleDto{
		RoleId: 1,
	}

	t.Run("success", func(t *testing.T) {
		mockRoleRepository.On("Delete", mock.Anything, mockDeleteRoleDto.RoleId).Return(nil).Once()

		uc := usecase.NewRoleUsecase(mockRoleRepository, time.Second*2)
		err := uc.Delete(context.Background(), mockDeleteRoleDto.RoleId)
		assert.NoError(t, err)
		mockRoleRepository.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T) {
		mockRoleRepository.On("Delete", mock.Anything, mockDeleteRoleDto.RoleId).Return(errors.New("unexpected")).Once()

		uc := usecase.NewRoleUsecase(mockRoleRepository, time.Second*2)
		err := uc.Delete(context.Background(), mockDeleteRoleDto.RoleId)
		assert.Error(t, err)
		mockRoleRepository.AssertExpectations(t)
	})
}

func TestFetchByName(t *testing.T) {
	mockRoleRepository := mocks.NewRoleRespository(t)
	mockFetchByNameDto := domain.FetchByNameDto{
		RoleName: "test",
	}

	mockRole := domain.Role{
		RoleId : 1,
		RoleName : "test",
		RegisteredBy : "test",
		RegisteredDate : "test",
	}

	t.Run("success", func(t *testing.T){
		mockRoleRepository.On("FetchByName", mock.Anything, mockFetchByNameDto.RoleName).Return(mockRole, nil).Once()

		uc := usecase.NewRoleUsecase(mockRoleRepository, time.Second * 2)
		role, err := uc.FetchByName(context.Background(), mockFetchByNameDto.RoleName)
		assert.NoError(t, err)
		assert.Equal(t, mockRole, role)
		mockRoleRepository.AssertExpectations(t)
		
	})

	t.Run("fail", func(t *testing.T) {
		mockRoleRepository.On("FetchByName", mock.Anything, mockFetchByNameDto.RoleName).Return(domain.Role{}, errors.New("unexpected")).Once()

		uc := usecase.NewRoleUsecase(mockRoleRepository, time.Second * 2)
		role, err := uc.FetchByName(context.Background(), mockFetchByNameDto.RoleName)
		assert.Error(t, err)
		assert.Equal(t, domain.Role{}, role)

		mockRoleRepository.AssertExpectations(t)
	})

}

func TestFetchRoles(t *testing.T){
			mockRoleRespository := mocks.NewRoleRespository(t)

			mockRoles := []domain.Role{
					{
						RoleId : 1,
						RoleName : "test1",
						RegisteredBy : "test",
						RegisteredDate : "test",
					},
				{
					RoleId : 2,
					RoleName : "test2",
					RegisteredBy : "test",
					RegisteredDate : "test",
				},
			}

			t.Run("success", func(t *testing.T){
				mockRoleRespository.On("FetchRoles", mock.Anything).Return(mockRoles, nil).Once()

				uc := usecase.NewRoleUsecase(mockRoleRespository, time.Second * 2)
				roles, err := uc.FetchRoles(context.Background())

				assert.NoError(t, err)
				assert.Equal(t, mockRoles, roles)
				mockRoleRespository.AssertExpectations(t)
			})
			
			t.Run("fail", func(t *testing.T){
				mockRoleRespository.On("FetchRoles", mock.Anything).Return([]domain.Role{}, errors.New("unexpected")).Once()

				uc := usecase.NewRoleUsecase(mockRoleRespository, time.Second * 2)
				roles, err := uc.FetchRoles(context.Background())

				assert.Error(t, err)
				assert.Len(t, roles, 0)
				assert.Equal(t, []domain.Role{}, roles)
				mockRoleRespository.AssertExpectations(t)
			})

	}
		
