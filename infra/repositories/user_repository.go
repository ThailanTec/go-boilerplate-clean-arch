package repositories

import (
	"errors"

	"github.com/ThailanTec/challenger/pousada/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *domain.User) error
	GetUsers() ([]*domain.User, error)
	GetUserByData(document string) (*domain.User, error)
	DeleteUser(id uuid.UUID) error
	UpdateUser(id uuid.UUID, user *domain.User) (*domain.User, error)
	GetUserByID(id uuid.UUID) (*domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) CreateUser(user *domain.User) error {
	user.ID = uuid.New()
	return repo.db.Create(user).Error
}

func (repo *userRepository) GetUsers() ([]*domain.User, error) {
	var users []*domain.User
	result := repo.db.Find(&users)
	return users, result.Error
}

func (repo *userRepository) GetUserByData(document string) (*domain.User, error) {
	var user domain.User
	result := repo.db.Where("document = ?", document).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, domain.ErrGetUserByData
	}

	return &user, result.Error
}

func (repo *userRepository) GetUserByID(id uuid.UUID) (*domain.User, error) {
	var user domain.User
	result := repo.db.Where("id = ?", id).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, domain.ErrIDNotFound
	}

	return &user, result.Error
}

func (repo *userRepository) DeleteUser(id uuid.UUID) error {
	var user domain.User
	result := repo.db.Where("id = ?", id).Delete(&user)
	return result.Error
}

func (repo *userRepository) UpdateUser(id uuid.UUID, user *domain.User) (*domain.User, error) {
	tx := repo.db.Begin()

	if tx.Error != nil {
		return nil, tx.Error
	}

	req := tx.Model(&domain.User{}).Where("id = ?", id).Omit("ID").Updates(user)

	if req.Error != nil {
		tx.Rollback()
		return nil, req.Error
	}

	if req.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("err to update a user")
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return user, nil
}
