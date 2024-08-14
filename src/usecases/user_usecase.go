package usecases

import (
	"encoding/json"
	"github.com/ThailanTec/challenger/pousada/domain"
	"github.com/ThailanTec/challenger/pousada/infra/repositories"
	"github.com/ThailanTec/challenger/pousada/src/dto"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

type UserUsecase interface {
	CreateUser(userDTO *dto.UserDTO) (*domain.User, error)
	GetUsers() ([]*domain.User, error)
	GetUserByDocument(document string) (*domain.User, error)
	DeleteUser(id uuid.UUID) error
	UpdateUser(id uuid.UUID, user *dto.UserDTO) (*domain.User, error)
}

type userUsecase struct {
	userRepo  repositories.UserRepository
	redisRepo repositories.RedisRepository
	cacheTTL  time.Duration
	validate  *validator.Validate
}

func NewUserUsecase(ur repositories.UserRepository, redisRepo repositories.RedisRepository, cacheTTL time.Duration) UserUsecase {
	return &userUsecase{userRepo: ur,
		validate:  validator.New(),
		redisRepo: redisRepo,
		cacheTTL:  cacheTTL,
	}
}

func (uc *userUsecase) CreateUser(userDTO *dto.UserDTO) (*domain.User, error) {
	usr, err := domain.NewUser(userDTO)
	if err != nil {
		return nil, err
	}

	err = uc.userRepo.CreateUser(usr)
	if err != nil {
		return nil, err
	}

	userJSON, _ := json.Marshal(usr)
	err = uc.redisRepo.Set(usr.Document, userJSON, uc.cacheTTL)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (uc *userUsecase) GetUsers() ([]*domain.User, error) {
	cacheKey := "all_users"

	cachedUsers, err := uc.redisRepo.Get(cacheKey)
	if err == nil && cachedUsers != "" {
		var users []*domain.User
		err := json.Unmarshal([]byte(cachedUsers), &users)
		if err == nil {
			return users, nil
		}
	}

	users, err := uc.userRepo.GetUsers()
	if err != nil {
		return nil, err
	}

	usersJSON, err := json.Marshal(users)
	if err == nil {
		_ = uc.redisRepo.Set(cacheKey, usersJSON, uc.cacheTTL*time.Minute)
	}

	return users, nil
}

func (uc *userUsecase) GetUserByDocument(document string) (*domain.User, error) {
	cachedUsers, err := uc.redisRepo.Get(document)
	if err == nil && cachedUsers != "" {
		user := &domain.User{}
		err := json.Unmarshal([]byte(cachedUsers), &user)
		if err == nil {
			return user, nil
		}
	}

	users, err := uc.userRepo.GetUserByData(document)
	if err != nil {
		return nil, err
	}

	usersJSON, err := json.Marshal(users)
	if err == nil {
		_ = uc.redisRepo.Set(document, usersJSON, uc.cacheTTL*time.Minute)
	}
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
