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

func TestFetchByUserName(t *testing.T) {
	mockUserRepository := mocks.NewUserRepository(t)
    mockFetchByUserNameDto := domain.FetchByUserNameDto{
        UserName: "test",
    }

    mockFetchByUserNameResponse := domain.User{
        UserId          :1,
        RoleId          :2,
        UserName        :"test",
        FirstName       :"test",
        FatherName      :"test",
        GrandFatherName :"test",
        Password        :"test",
        PhoneNumber     :"test",
        Address         :"test",
        Email           :"test",
        RegisteredBy 	:"test",
        RegisteredDate 	:"test",
    }



    t.Run("success", func(t *testing.T){
        mockUserRepository.On("FetchByUserName", mock.Anything, mockFetchByUserNameDto.UserName).Return(mockFetchByUserNameResponse, nil).Once()
        uc := usecase.NewLoginUsecase(mockUserRepository, time.Second * 2)
        user, err := uc.FetchByUserName(context.Background(), mockFetchByUserNameDto.UserName)

        assert.NoError(t, err)
        assert.Equal(t, user, mockFetchByUserNameResponse)
        mockUserRepository.AssertExpectations(t)
    })
    

    t.Run("fail", func(t *testing.T){
        mockUserRepository.On("FetchByUserName", mock.Anything, mockFetchByUserNameDto.UserName).Return(domain.User{}, errors.New("unexpected")).Once()
        uc := usecase.NewLoginUsecase(mockUserRepository, time.Second * 2)
        user, err := uc.FetchByUserName(context.Background(), mockFetchByUserNameDto.UserName)

        assert.Error(t, err)
        assert.Equal(t, user, domain.User{})
        mockUserRepository.AssertExpectations(t)
    })

    t.Run("fail", func(t *testing.T){
        mockUserRepository.On("FetchByUserName", mock.Anything, mockFetchByUserNameDto.UserName).Return(domain.User{}, errors.New("unexpected")).Once()
        uc := usecase.NewLoginUsecase(mockUserRepository, time.Second * 2)
        user, err := uc.FetchByUserName(context.Background(), mockFetchByUserNameDto.UserName)

        assert.Error(t, err)
        assert.Equal(t, user, domain.User{})
        mockUserRepository.AssertExpectations(t)
    })
    
}