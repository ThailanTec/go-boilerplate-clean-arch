package domain

import (
	"time"

	"github.com/ThailanTec/challenger/pousada/src/dto"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	Phone     string    `db:"phone"`
	Document  string    `db:"document"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`
}

func NewUser(user *dto.UserDTO) (*User, error) {
	return &User{
		ID:        uuid.New(),
		Name:      user.Name,
		Phone:     user.Phone,
		Document:  user.Document,
		CreatedAt: time.Now(),
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
