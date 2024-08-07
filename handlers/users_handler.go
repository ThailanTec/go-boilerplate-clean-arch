package handler

import (
	"github.com/ThailanTec/challenger/pousada/domain"
	"github.com/ThailanTec/challenger/pousada/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	UserUsecase usecases.UserUsecase
}

func NewUserHandler(uc usecases.UserUsecase) *UserHandler {
	return &UserHandler{UserUsecase: uc}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.UserUsecase.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
