package routes

import (
	"github.com/ThailanTec/challenger/pousada/infra/repositories"
	"github.com/ThailanTec/challenger/pousada/src/handlers"
	"github.com/ThailanTec/challenger/pousada/src/usecases"
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
		userRoutes.GET("", userHandler.GetUser)
		userRoutes.GET(":document", userHandler.GetUserByDocument)
		userRoutes.DELETE(":id", userHandler.DeleteUser)
		userRoutes.PUT(":id", userHandler.UpdateUser)
	}
}
