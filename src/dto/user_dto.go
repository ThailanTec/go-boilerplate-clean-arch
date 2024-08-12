package dto

import (
	"github.com/google/uuid"
)

type UserDTO struct {
	Name     string `json:"name" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Document string `json:"document" validate:"required"`
}

type UserResponseDTO struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Document string    `json:"email"`
	Phone    string    `json:"phone"`
}
