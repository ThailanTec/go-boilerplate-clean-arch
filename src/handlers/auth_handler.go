package handler

import (
	"net/http"

	"github.com/ThailanTec/challenger/pousada/src/dto"
	"github.com/ThailanTec/challenger/pousada/src/usecases"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUsecase *usecases.AuthUsecase
}

func NewAuthHandler(authUsecase *usecases.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase: authUsecase,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var document dto.LoginDTO
	if err := c.ShouldBindJSON(&document); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authUsecase.Login(document.Document)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *AuthHandler) Validate(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
		return
	}

	user, err := h.authUsecase.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
