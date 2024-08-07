package usecases

import (
	"github.com/ThailanTec/challenger/pousada/domain"
	"github.com/ThailanTec/challenger/pousada/repositories"
)

type UserUsecase interface {
	CreateUser(user *domain.User) error
}

type userUsecase struct {
	userRepo repositories.UserRepository
}

func NewUserUsecase(ur repositories.UserRepository) UserUsecase {
	return &userUsecase{userRepo: ur}
}

func (uc *userUsecase) CreateUser(user *domain.User) error {
	return uc.userRepo.Save(user)
}
