package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ThailanTec/challenger/pousada/domain"
	"github.com/ThailanTec/challenger/pousada/src/dto"
	handler "github.com/ThailanTec/challenger/pousada/src/handlers"
	mocks "github.com/ThailanTec/challenger/pousada/test/mocks/usecases"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestCreateUser_Success tests the successful creation of a user
func TestCreateUser_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Arrange
	userUsecaseMock := new(mocks.UserUsecaseMock)
	userHandler := handler.NewUserHandler(userUsecaseMock)

	// Creating a test server with the handler
	router := gin.Default()
	router.POST("/users", userHandler.CreateUser)

	userDTO := dto.UserDTO{
		Name:     "John Doe",
		Phone:    "123456789",
		Document: "doc1",
	}

	user := &domain.User{
		ID:        uuid.New(),
		Name:      "John Doe",
		Phone:     "123456789",
		Document:  "doc1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	userUsecaseMock.On("CreateUser", mock.Anything).Return(user, nil)

	// Act
	body, _ := json.Marshal(userDTO)
	req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusCreated, w.Code)
	userUsecaseMock.AssertExpectations(t)
}

// TestCreateUser_InternalServerError tests the scenario where the usecase returns an error
func TestCreateUser_InternalServerError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Arrange
	userUsecaseMock := new(mocks.UserUsecaseMock)
	userHandler := handler.NewUserHandler(userUsecaseMock)

	router := gin.Default()
	router.POST("/users", userHandler.CreateUser)

	userDTO := dto.UserDTO{
		Name:     "John Doe",
		Phone:    "123456789",
		Document: "doc1",
	}

	userUsecaseMock.On("CreateUser", mock.Anything).Return(nil, errors.New("something went wrong"))

	// Act
	body, _ := json.Marshal(userDTO)
	req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	userUsecaseMock.AssertExpectations(t)
}

func TestGetUser_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Arrange
	userUsecaseMock := new(mocks.UserUsecaseMock)
	users := []*domain.User{
		{
			ID:        uuid.New(),
			Name:      "John Doe",
			Phone:     "123456789",
			Document:  "doc1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	userUsecaseMock.On("GetUsers").Return(users, nil)

	userHandler := handler.NewUserHandler(userUsecaseMock)

	router := gin.Default()
	router.GET("/users", userHandler.GetUser)

	// Act
	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	userUsecaseMock.AssertExpectations(t)

	var responseBody []dto.UserDTO
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	if err != nil {
		t.Errorf("Error unmarshalling response body: %v", err)
	}
	assert.Equal(t, len(users), len(responseBody))
}

func TestGetUser_Failure(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// Arrange
	userUsecaseMock := new(mocks.UserUsecaseMock)
	expectedError := errors.New("failed to fetch users")
	userUsecaseMock.On("GetUsers").Return([]*domain.User{}, expectedError)

	userHandler := handler.NewUserHandler(userUsecaseMock)

	router := gin.Default()
	router.GET("/users", userHandler.GetUser)

	// Act
	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var responseBody gin.H
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	if err != nil {
		t.Errorf("Error unmarshalling response body: %v", err)
	}

	assert.Equal(t, expectedError.Error(), responseBody["error"])
}

func TestGetUserByDocument_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Arrange
	userUsecaseMock := new(mocks.UserUsecaseMock)
	userHandler := handler.NewUserHandler(userUsecaseMock)

	expectedUser := &domain.User{}
	userUsecaseMock.On("GetUserByDocument", mock.Anything).Return(expectedUser, nil)

	router := gin.Default()
	router.GET("/users/:document", userHandler.GetUserByDocument)

	// Act
	document := "some_document"
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", document), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	userUsecaseMock.AssertExpectations(t)

	// Optional: Assert response body
	var responseBody dto.UserDTO
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
}

func TestGetUserByDocument_Failure(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Arrange
	userUsecaseMock := new(mocks.UserUsecaseMock)
	userHandler := handler.NewUserHandler(userUsecaseMock)

	expectedError := errors.New("failed to get user")
	userUsecaseMock.On("GetUserByDocument", mock.Anything).Return(&domain.User{}, expectedError)

	router := gin.Default()
	router.GET("/users/:document", userHandler.GetUserByDocument)

	// Act
	document := "some_document"
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", document), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	userUsecaseMock.AssertExpectations(t)
}

func TestDeleteUser_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Arrange
	userUsecaseMock := new(mocks.UserUsecaseMock)
	userHandler := handler.NewUserHandler(userUsecaseMock)

	userID := uuid.New()
	userUsecaseMock.On("DeleteUser", userID).Return(nil)

	router := gin.Default()
	router.DELETE("/users/:id", userHandler.DeleteUser)

	// Act
	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/users/%s", userID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusNoContent, w.Code)
	userUsecaseMock.AssertExpectations(t)
}

func TestDeleteUser_Failure(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Arrange
	userUsecaseMock := new(mocks.UserUsecaseMock)
	userHandler := handler.NewUserHandler(userUsecaseMock)

	userID := uuid.New()
	expectedError := errors.New("failed to delete user")
	userUsecaseMock.On("DeleteUser", userID).Return(expectedError)

	router := gin.Default()
	router.DELETE("/users/:id", userHandler.DeleteUser)

	// Act
	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/users/%s", userID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	userUsecaseMock.AssertExpectations(t)
}

func TestUpdateUser_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Arrange
	userUsecaseMock := new(mocks.UserUsecaseMock)
	userHandler := handler.NewUserHandler(userUsecaseMock)

	userID := uuid.New()
	userDTO := dto.UserDTO{
		// ... populate with user data
	}
	expectedUser := &domain.User{
		// ... populate with expected user data
	}
	userUsecaseMock.On("UpdateUser", userID, mock.Anything).Return(expectedUser, nil)

	router := gin.Default()
	router.PUT("/users/:id", userHandler.UpdateUser)

	// Act
	body, _ := json.Marshal(userDTO)
	req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/users/%s", userID), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	userUsecaseMock.AssertExpectations(t)
}

func TestUpdateUser_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Arrange
	userUsecaseMock := new(mocks.UserUsecaseMock)
	userHandler := handler.NewUserHandler(userUsecaseMock)

	userID := uuid.New()

	router := gin.Default()
	router.PUT("/users/:id", userHandler.UpdateUser)

	// Act
	invalidJSON := []byte("invalid json")
	req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/users/%s", userID), bytes.NewBuffer(invalidJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	userUsecaseMock.AssertExpectations(t)
}

func TestUpdateUser_UsecaseError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Arrange
	userUsecaseMock := new(mocks.UserUsecaseMock)
	userHandler := handler.NewUserHandler(userUsecaseMock)

	userID := uuid.New()
	userDTO := dto.UserDTO{
		// ... populate with user data
	}
	expectedError := errors.New("failed to update user")
	userUsecaseMock.On("UpdateUser", userID, mock.Anything).Return(&domain.User{}, expectedError)

	router := gin.Default()
	router.PUT("/users/:id", userHandler.UpdateUser)

	// Act
	body, _ := json.Marshal(userDTO)
	req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/users/%s", userID), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	userUsecaseMock.AssertExpectations(t)
}
