package usecases

import (
	"github.com/ThailanTec/challenger/pousada/domain"
	"github.com/ThailanTec/challenger/pousada/infra/repositories"
	"github.com/ThailanTec/challenger/pousada/src/dto"
	"github.com/google/uuid"
)

type UserUsecase interface {
	CreateUser(userDTO *dto.UserDTO) (*domain.User, error)
	GetUsers() ([]*domain.User, error)
	GetUserByDocument(document string) (*domain.User, error)
	DeleteUser(id uuid.UUID) error
	UpdateUser(id uuid.UUID, user *dto.UserDTO) (*domain.User, error)
}

type userUsecase struct {
	userRepo repositories.UserRepository
}

func NewUserUsecase(ur repositories.UserRepository) UserUsecase {
	return &userUsecase{userRepo: ur}
}

func (uc *userUsecase) CreateUser(userDTO *dto.UserDTO) (*domain.User, error) {
	user, err := domain.NewUser(userDTO)
	if err != nil {
		return nil, err
	}

	err = uc.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *userUsecase) GetUsers() ([]*domain.User, error) {
	users, err := uc.userRepo.GetUsers()

	return users, err
}

func (uc *userUsecase) GetUserByDocument(document string) (*domain.User, error) {
	users, err := uc.userRepo.GetUserByData(document)

	return users, err
}

func (uc *userUsecase) DeleteUser(id uuid.UUID) error {
	err := uc.userRepo.DeleteUser(id)

	return err
}

func (uc *userUsecase) UpdateUser(id uuid.UUID, user *dto.UserDTO) (*domain.User, error) {
	usr, err := domain.NewUser(user)
	if err != nil {
		return nil, err
	}

	users, err := uc.userRepo.UpdateUser(id, usr)

	return users, err
}
