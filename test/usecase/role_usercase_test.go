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

func TestCreateRole(t *testing.T) {

	mockRoleRepository := mocks.NewRoleRespository(t)
	mockAddRoleDto := &domain.CreateRoleDto{
		RoleName:     "test",
		RegisteredBy: "test",
	}
	t.Run("success", func(t *testing.T) {

		mockRoleRepository.On("Create", mock.Anything, mockAddRoleDto).Return(nil).Once()

		uc := usecase.NewRoleUsecase(mockRoleRepository, time.Second*2)
		err := uc.Create(context.Background(), mockAddRoleDto)

		assert.NoError(t, err)
		mockRoleRepository.AssertExpectations(t)
	})

	t.Run("erro", func(t *testing.T) {
		mockRoleRepository.On("Create", mock.Anything, mockAddRoleDto).Return(errors.New("unexpected")).Once()
		uc := usecase.NewRoleUsecase(mockRoleRepository, time.Second*2)

		err := uc.Create(context.Background(), mockAddRoleDto)

		assert.Error(t, err)
		mockRoleRepository.AssertExpectations(t)

	})

}

func TestUpdateRole(t *testing.T) {
	mockRoleRepository := mocks.NewRoleRespository(t)
	mockUpdateRoleDto := &domain.UpdateRoleDto{
		RoleId: 		1,
		RoleName:		"test",
		RegisteredBy:  	"test",
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

func TestDeleteRole(t *testing.T) {
	mockRoleRepository := mocks.NewRoleRespository(t)

	roleId := int32(1)

	t.Run("success", func(t *testing.T) {
		mockRoleRepository.On("Delete", mock.Anything, roleId).Return(nil).Once()

		uc := usecase.NewRoleUsecase(mockRoleRepository, time.Second*2)
		err := uc.Delete(context.Background(), roleId)
		assert.NoError(t, err)
		mockRoleRepository.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T) {
		mockRoleRepository.On("Delete", mock.Anything, roleId).Return(errors.New("unexpected")).Once()

		uc := usecase.NewRoleUsecase(mockRoleRepository, time.Second*2)
		err := uc.Delete(context.Background(), roleId)
		assert.Error(t, err)
		mockRoleRepository.AssertExpectations(t)
	})
}

func TestFetchByNameRole(t *testing.T) {
	mockRoleRepository := mocks.NewRoleRespository(t)
	roleName := "test"

	mockRole := domain.Role{
		RoleId : 1,
		RoleName : "test",
		RegisteredBy : "test",
		RegisteredDate : "test",
	}

	t.Run("success", func(t *testing.T){
		mockRoleRepository.On("FetchByName", mock.Anything, roleName).Return(mockRole, nil).Once()

		uc := usecase.NewRoleUsecase(mockRoleRepository, time.Second * 2)
		role, err := uc.FetchByName(context.Background(), roleName)
		assert.NoError(t, err)
		assert.Equal(t, mockRole, role)
		mockRoleRepository.AssertExpectations(t)
		
	})

	t.Run("fail", func(t *testing.T) {
		mockRoleRepository.On("FetchByName", mock.Anything, roleName).Return(domain.Role{}, errors.New("unexpected")).Once()

		uc := usecase.NewRoleUsecase(mockRoleRepository, time.Second * 2)
		role, err := uc.FetchByName(context.Background(), roleName)
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
		
