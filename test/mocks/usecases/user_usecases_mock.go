package mocks

import (
	"github.com/ThailanTec/challenger/pousada/domain"
	"github.com/ThailanTec/challenger/pousada/src/dto"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type UserUsecaseMock struct {
	mock.Mock
}

func (m *UserUsecaseMock) CreateUser(userDTO *dto.UserDTO) (*domain.User, error) {
	args := m.Called(userDTO)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *UserUsecaseMock) GetUsers() ([]*domain.User, error) {
	args := m.Called()
	return args.Get(0).([]*domain.User), args.Error(1)
}

func (m *UserUsecaseMock) GetUserByDocument(document string) (*domain.User, error) {
	args := m.Called(document)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *UserUsecaseMock) DeleteUser(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *UserUsecaseMock) UpdateUser(id uuid.UUID, userDTO *dto.UserDTO) (*domain.User, error) {
	args := m.Called(id, userDTO)
	return args.Get(0).(*domain.User), args.Error(1)
}
