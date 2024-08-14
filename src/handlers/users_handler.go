package handler

import (
	"go.uber.org/zap"
	"net/http"

	"github.com/ThailanTec/challenger/pousada/domain"
	"github.com/ThailanTec/challenger/pousada/src/dto"
	"github.com/ThailanTec/challenger/pousada/src/usecases"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	UserUsecase usecases.UserUsecase
	Logger      *zap.Logger
}

func NewUserHandler(uc usecases.UserUsecase, logger *zap.Logger) *UserHandler {
	return &UserHandler{UserUsecase: uc,
		Logger: logger}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user dto.UserDTO
	h.Logger.Info("CreateUser called")
	if err := c.ShouldBindJSON(&user); err != nil {
		h.Logger.Error("Error binding JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := h.UserUsecase.CreateUser(&user)
	if err != nil {
		h.Logger.Error("Error creating user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	output := domain.OutputUser(u)
	h.Logger.Info("User created successfully", zap.String("user_id", u.ID.String()))
	c.JSON(http.StatusCreated, output)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	u, err := h.UserUsecase.GetUsers()
	h.Logger.Info("GetUser called", zap.String("user_id", c.Param("id")))
	if err != nil {
		h.Logger.Error("Error getting users", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, u)

}

func (h *UserHandler) GetUserByDocument(c *gin.Context) {
	document := c.Param("document")
	h.Logger.Info("GetUserByDocument called", zap.String("document", document))
	u, err := h.UserUsecase.GetUserByDocument(document)
	if err != nil {
		h.Logger.Error("Error getting user by document", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	output := domain.OutputUser(u)

	c.JSON(http.StatusOK, output)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	userID := uuid.MustParse(id)
	h.Logger.Info("DeleteUser called", zap.String("user_id", userID.String()))

	err := h.UserUsecase.DeleteUser(userID)
	if err != nil {
		h.Logger.Error("Error deleting user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	h.Logger.Info("User deleted successfully", zap.String("user_id", userID.String()))
	c.Status(http.StatusNoContent)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	h.Logger.Info("UpdateUser called", zap.String("user_id", id))
	userID := uuid.MustParse(id)
	var user dto.UserDTO
	if err := c.ShouldBindJSON(&user); err != nil {
		h.Logger.Error("Error binding JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usr, err := h.UserUsecase.UpdateUser(userID, &user)
	if err != nil {
		h.Logger.Error("Error updating user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	output := domain.OutputUser(usr)
	h.Logger.Info("User updated successfully", zap.String("user_id", userID.String()))
	c.JSON(http.StatusOK, output)
}
