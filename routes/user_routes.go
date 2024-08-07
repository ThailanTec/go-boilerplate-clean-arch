package routes

import (
	"github.com/ThailanTec/challenger/pousada/handlers"
	"github.com/ThailanTec/challenger/pousada/repositories"
	"github.com/ThailanTec/challenger/pousada/usecases"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	userRepo := repositories.NewUserRepository(db)
	userUsecase := usecases.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("", userHandler.CreateUser)
	}
}
