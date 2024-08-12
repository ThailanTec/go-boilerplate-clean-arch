package handler

import (
	"net/http"

	"github.com/ThailanTec/challenger/pousada/domain"
	"github.com/ThailanTec/challenger/pousada/src/dto"
	"github.com/ThailanTec/challenger/pousada/src/usecases"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	UserUsecase usecases.UserUsecase
}

func NewUserHandler(uc usecases.UserUsecase) *UserHandler {
	return &UserHandler{UserUsecase: uc}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user dto.UserDTO
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := h.UserUsecase.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	output := domain.OutputUser(u)

	c.JSON(http.StatusCreated, output)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	u, err := h.UserUsecase.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, u)

}

func (h *UserHandler) GetUserByDocument(c *gin.Context) {
	document := c.Param("document")

	u, err := h.UserUsecase.GetUserByDocument(document)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	output := domain.OutputUser(u)

	c.JSON(http.StatusOK, output)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	userID := uuid.MustParse(id)

	err := h.UserUsecase.DeleteUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusNoContent)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	userID := uuid.MustParse(id)
	var user dto.UserDTO
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usr, err := h.UserUsecase.UpdateUser(userID, &user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	output := domain.OutputUser(usr)
	println(output)

	c.JSON(http.StatusOK, output)
}
