package repositories

import (
	"errors"
	"testing"
	"time"

	"github.com/ThailanTec/challenger/pousada/domain"
	"github.com/ThailanTec/challenger/pousada/src/dto"
	"github.com/ThailanTec/challenger/pousada/src/usecases"
	mocks "github.com/ThailanTec/challenger/pousada/test/mocks/repositories"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mocks.NewUserRepositoryMock(ctrl)
	usecase := usecases.NewUserUsecase(userRepoMock)

	user := &domain.User{
		ID:        uuid.New(),
		Name:      "John Doe",
		Phone:     "123456789",
		Document:  "doc1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	userRepoMock.EXPECT().CreateUser(gomock.Any()).Return(nil) // Accepts any user object

	result, err := usecase.CreateUser(&dto.UserDTO{
		Name:     "John Doe",
		Phone:    "123456789",
		Document: "doc1",
	})

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user.Name, result.Name)
}

func TestCreateUser_Failure(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mocks.NewUserRepositoryMock(ctrl)
	usecase := usecases.NewUserUsecase(userRepoMock)

	userDTO := &dto.UserDTO{
		Name:     "John Doe",
		Phone:    "123456789",
		Document: "doc1",
	}
	// action
	expectedErr := errors.New("erro ao criar o usuário")
	userRepoMock.EXPECT().CreateUser(gomock.Any()).Return(expectedErr)

	result, err := usecase.CreateUser(userDTO)
	// assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, expectedErr, err)
}

func TestGetUsers_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mocks.NewUserRepositoryMock(ctrl)
	usecase := usecases.NewUserUsecase(userRepoMock)

	expectedErr := errors.New("erro ao obter usuários")
	userRepoMock.EXPECT().GetUsers().Return(nil, expectedErr)

	result, err := usecase.GetUsers()

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, expectedErr, err)
}

func TestGetUsers_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mocks.NewUserRepositoryMock(ctrl)
	usecase := usecases.NewUserUsecase(userRepoMock)

	mockUsers := []*domain.User{
		{
			ID:        uuid.New(),
			Name:      "John Doe",
			Phone:     "123456789",
			Document:  "doc1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        uuid.New(),
			Name:      "Jane Doe",
			Phone:     "987654321",
			Document:  "doc2",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	userRepoMock.EXPECT().GetUsers().Return(mockUsers, nil)

	result, err := usecase.GetUsers()

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, len(mockUsers), len(result))
	assert.Equal(t, mockUsers, result)
}

func TestGetUserByDocument_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mocks.NewUserRepositoryMock(ctrl)
	usecase := usecases.NewUserUsecase(userRepoMock)

	mockUser := &domain.User{
		ID:        uuid.New(),
		Name:      "John Doe",
		Phone:     "123456789",
		Document:  "doc1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	userRepoMock.EXPECT().GetUserByData("doc1").Return(mockUser, nil)

	result, err := usecase.GetUserByDocument("doc1")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, mockUser, result)
}

func TestGetUserByDocument_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mocks.NewUserRepositoryMock(ctrl)
	usecase := usecases.NewUserUsecase(userRepoMock)

	userRepoMock.EXPECT().GetUserByData("doc1").Return(nil, errors.New("user not found"))

	result, err := usecase.GetUserByDocument("doc1")

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestDeleteUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mocks.NewUserRepositoryMock(ctrl)
	usecase := usecases.NewUserUsecase(userRepoMock)

	userID := uuid.New()

	userRepoMock.EXPECT().DeleteUser(userID).Return(nil)

	err := usecase.DeleteUser(userID)

	assert.NoError(t, err)
}

func TestDeleteUser_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mocks.NewUserRepositoryMock(ctrl)
	usecase := usecases.NewUserUsecase(userRepoMock)

	userID := uuid.New()

	userRepoMock.EXPECT().DeleteUser(userID).Return(errors.New("error deleting user"))

	err := usecase.DeleteUser(userID)

	assert.Error(t, err)
}

func TestUpdateUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mocks.NewUserRepositoryMock(ctrl)
	usecase := usecases.NewUserUsecase(userRepoMock)

	userID := uuid.New()
	userDTO := &dto.UserDTO{
		Name:     "John Doe",
		Phone:    "123456789",
		Document: "doc1",
	}

	updatedUser := &domain.User{
		ID:        userID,
		Name:      "John Doe",
		Phone:     "123456789",
		Document:  "doc1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	userRepoMock.EXPECT().UpdateUser(userID, gomock.Any()).Return(updatedUser, nil)

	result, err := usecase.UpdateUser(userID, userDTO)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, updatedUser, result)
}

func TestUpdateUser_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mocks.NewUserRepositoryMock(ctrl)
	usecase := usecases.NewUserUsecase(userRepoMock)

	userID := uuid.New()
	userDTO := &dto.UserDTO{
		Name:     "John Doe",
		Phone:    "123456789",
		Document: "doc1",
	}

	userRepoMock.EXPECT().UpdateUser(userID, gomock.Any()).Return(nil, errors.New("error updating user"))

	result, err := usecase.UpdateUser(userID, userDTO)

	assert.Error(t, err)
	assert.Nil(t, result)
}
