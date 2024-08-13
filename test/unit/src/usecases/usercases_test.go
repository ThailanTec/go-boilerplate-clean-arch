package usecases

import (
	"errors"
	"testing"

	"github.com/ThailanTec/challenger/pousada/domain"
	"github.com/ThailanTec/challenger/pousada/src/dto"
	"github.com/ThailanTec/challenger/pousada/src/usecases"
	mocks "github.com/ThailanTec/challenger/pousada/test/mocks/repositories"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CreateUser_Success(t *testing.T) {
	// Arrange
	userRepoMock := new(mocks.UserRepositoryMock)
	usecase := usecases.NewUserUsecase(userRepoMock)

	userDTO := &dto.UserDTO{
		Name:     "Belo",
		Phone:    "31994416221",
		Document: "12345678900",
	}

	_ = &domain.User{
		ID:       uuid.New(),
		Name:     userDTO.Name,
		Phone:    userDTO.Phone,
		Document: userDTO.Document,
	}

	userRepoMock.On("CreateUser", mock.AnythingOfType("*domain.User")).Return(nil)

	// Act
	createdUser, err := usecase.CreateUser(userDTO)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, createdUser)
	assert.Equal(t, userDTO.Name, createdUser.Name)
	assert.Equal(t, userDTO.Phone, createdUser.Phone)
	assert.Equal(t, userDTO.Document, createdUser.Document)
	userRepoMock.AssertExpectations(t)
}

func TestCreateUser_Failure(t *testing.T) {
	// Arrange
	userRepoMock := new(mocks.UserRepositoryMock)
	usecase := usecases.NewUserUsecase(userRepoMock)

	userDTO := &dto.UserDTO{
		Name:  "Test User",
		Phone: "1234567890",
	}

	expectedUser := &domain.User{
		Name:  userDTO.Name,
		Phone: userDTO.Phone,
	}

	// Configuração da expectativa do mock para criar o usuário
	userRepoMock.On("CreateUser", expectedUser).Return(errors.New("failed to create user"))

	// Act
	createdUser, err := usecase.CreateUser(userDTO)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, createdUser)
}

func TestCreateUser_MissingField(t *testing.T) {
	// Arrange
	userRepoMock := new(mocks.UserRepositoryMock)
	usecase := usecases.NewUserUsecase(userRepoMock)

	userDTO := &dto.UserDTO{
		Phone:    "1234567890",
		Document: "123456789",
	}
	e := "Key: 'UserDTO.Name' Error:Field validation for 'Name' failed on the 'required' tag"

	userRepoMock.On("CreateUser", mock.AnythingOfType("*domain.User")).Return(nil).Maybe()

	// Act
	createdUser, err := usecase.CreateUser(userDTO)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, createdUser)
	assert.Equal(t, e, err.Error())
	userRepoMock.AssertNotCalled(t, "CreateUser", mock.AnythingOfType("*domain.User"))
}

func TestGetUsers_Success(t *testing.T) {
	// Arrange
	userRepoMock := new(mocks.UserRepositoryMock)
	usecase := usecases.NewUserUsecase(userRepoMock)

	users := []*domain.User{
		{
			ID:       uuid.MustParse("af430404-e5ea-4752-9d89-0c371ec0d9fc"),
			Name:     "user1",
			Document: "doc1",
			Phone:    "phone1",
		},
		{
			ID:       uuid.MustParse("07cfb203-30f7-4cec-b75b-f523822795fb"),
			Name:     "user2",
			Document: "doc2",
			Phone:    "phone2",
		},
	}

	// Act
	userRepoMock.On("GetUsers").Return(users, nil)
	getusers, err := usecase.GetUsers()

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, getusers)
	assert.Equal(t, users, getusers)
	userRepoMock.AssertExpectations(t)
}

func TestGetUserByDocument_Success(t *testing.T) {
	// Arrange
	userRepoMock := new(mocks.UserRepositoryMock)
	usecase := usecases.NewUserUsecase(userRepoMock)
	user := &domain.User{
		ID:       uuid.MustParse("af430404-e5ea-4752-9d89-0c371ec0d9fc"),
		Name:     "user1",
		Document: "doc1",
		Phone:    "phone1",
	}

	// Act
	userRepoMock.On("GetUserByData", "doc1").Return(user, nil)
	getuser, err := usecase.GetUserByDocument("doc1")

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, getuser)
	assert.Equal(t, user, getuser)
	userRepoMock.AssertExpectations(t)
}

func TestGetUserByDocument_Failure(t *testing.T) {
	// Arrange
	userRepoMock := new(mocks.UserRepositoryMock)
	usecase := usecases.NewUserUsecase(userRepoMock)

	userRepoMock.On("GetUserByData", "doc2").Return(nil, domain.ErrGetUserByData)

	// Act
	user, err := usecase.GetUserByDocument("doc2")

	// Assert
	assert.Error(t, err)
	assert.EqualError(t, err, domain.ErrGetUserByData.Error())
	assert.Nil(t, user)
	userRepoMock.AssertExpectations(t)
}

func TestDeleteUser_Success(t *testing.T) {
	// Arrange
	userRepoMock := new(mocks.UserRepositoryMock)
	usecase := usecases.NewUserUsecase(userRepoMock)

	userID := uuid.New()
	userRepoMock.On("DeleteUser", userID).Return(nil)

	// Act
	err := usecase.DeleteUser(userID)

	// Assert
	assert.NoError(t, err)
	userRepoMock.AssertExpectations(t)
}

func TestDeleteUser_Failure(t *testing.T) {
	// Arrange
	userRepoMock := new(mocks.UserRepositoryMock)
	usecase := usecases.NewUserUsecase(userRepoMock)

	userID := uuid.New()
	userRepoMock.On("DeleteUser", userID).Return(errors.New("Erro ao deletar usuário"))

	// Act
	err := usecase.DeleteUser(userID)

	// Assert
	assert.Error(t, err)
	userRepoMock.AssertExpectations(t)
}

func TestUpdateUser_Success(t *testing.T) {
	// Arrange
	userRepoMock := new(mocks.UserRepositoryMock)
	usecase := usecases.NewUserUsecase(userRepoMock)

	userID := uuid.New()
	userDTO := &dto.UserDTO{
		Name:     "John Doe",
		Phone:    "123456789",
		Document: "doc1",
	}

	updatedUser := &domain.User{
		ID:       userID,
		Name:     "John Doe",
		Phone:    "123456789",
		Document: "doc1",
	}

	userRepoMock.On("UpdateUser", userID, mock.AnythingOfType("*domain.User")).Return(updatedUser, nil)

	// Act
	result, err := usecase.UpdateUser(userID, userDTO)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, updatedUser, result)
	userRepoMock.AssertExpectations(t)
}

func TestUpdateUser_ErrorUpdatingUserInRepository(t *testing.T) {
	// Arrange
	userRepoMock := new(mocks.UserRepositoryMock)
	usecase := usecases.NewUserUsecase(userRepoMock)

	userID := uuid.New()
	userDTO := &dto.UserDTO{
		Name:     "John Doe",
		Phone:    "123456789",
		Document: "doc1",
	}

	// Simula um erro na atualização do usuário no repositório
	mockError := errors.New("repository error")
	userRepoMock.On("UpdateUser", userID, mock.AnythingOfType("*domain.User")).Return(nil, mockError)

	// Act
	result, err := usecase.UpdateUser(userID, userDTO)

	// Assert
	assert.Nil(t, result)
	assert.EqualError(t, err, mockError.Error())
	userRepoMock.AssertExpectations(t)
}
