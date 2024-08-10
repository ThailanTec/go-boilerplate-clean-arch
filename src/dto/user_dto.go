package dto

import "github.com/google/uuid"

type UserDTO struct {
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Document string `json:"document" binding:"required"`
}

type UserResponseDTO struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Document string    `json:"email"`
	Phone    string    `json:"phone"`
}
