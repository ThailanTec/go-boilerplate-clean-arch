package usecases

import (
	"errors"
	"fmt"
	"github.com/ThailanTec/challenger/pousada/domain"
	"github.com/ThailanTec/challenger/pousada/infra/auth"
	"github.com/ThailanTec/challenger/pousada/infra/repositories"
	"github.com/ThailanTec/challenger/pousada/src/config"
)

type AuthUsecase struct {
	userRepo repositories.UserRepository
	cfg      config.Config
}

func NewAuthUsecase(userRepo repositories.UserRepository, cfg config.Config) *AuthUsecase {
	return &AuthUsecase{
		userRepo: userRepo,
		cfg:      cfg,
	}
}

func (u *AuthUsecase) Login(document string) (string, error) {
	user, err := u.userRepo.GetUserByData(document)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := auth.GenerateJWT(user.ID, u.cfg)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *AuthUsecase) ValidateToken(token string) (*domain.User, error) {
	claims, err := auth.ValidateJWT(token, u.cfg)
	if err != nil {
		return nil, err
	}
	// TODO: implementar lógica correta
	user, err := u.userRepo.GetUserByID(claims.UserID)
	fmt.Println(claims.UserID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
