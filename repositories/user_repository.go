package repositories

import (
	"github.com/ThailanTec/challenger/pousada/domain"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
)

type UserRepository interface {
	Save(user *domain.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) Save(user *domain.User) error {
	id := rand.Int()
	user.ID = strconv.Itoa(id)
	return repo.db.Create(user).Error
}
