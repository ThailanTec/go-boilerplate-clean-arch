package domain

import (
	"time"

	"github.com/ThailanTec/challenger/pousada/src/dto"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID
	Name      string
	Phone     string
	Document  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func NewUser(user *dto.UserDTO) (*User, error) {
	return &User{
		ID:        uuid.New(),
		Name:      user.Name,
		Phone:     user.Phone,
		Document:  user.Document,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func OutputUser(user *User) *dto.UserResponseDTO {
	return &dto.UserResponseDTO{
		ID:       user.ID,
		Name:     user.Name,
		Document: user.Document,
		Phone:    user.Phone,
	}
}
