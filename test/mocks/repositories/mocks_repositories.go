package mocks

import (
	"github.com/ThailanTec/challenger/pousada/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) CreateUser(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepositoryMock) GetUsers() ([]*domain.User, error) {
	args := m.Called()
	return args.Get(0).([]*domain.User), args.Error(1)
}

func (m *UserRepositoryMock) GetUserByData(document string) (*domain.User, error) {
	args := m.Called(document)
	var user *domain.User
	if args.Get(0) != nil {
		user = args.Get(0).(*domain.User)
	}
	return user, args.Error(1)
}

func (m *UserRepositoryMock) DeleteUser(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *UserRepositoryMock) UpdateUser(id uuid.UUID, user *domain.User) (*domain.User, error) {
	args := m.Called(id, user)
	if result := args.Get(0); result != nil {
		return result.(*domain.User), args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *UserRepositoryMock) GetUserByID(id uuid.UUID) (*domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.User), args.Error(1)
}
