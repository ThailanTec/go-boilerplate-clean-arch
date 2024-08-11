package routes

import (
	"github.com/ThailanTec/challenger/pousada/infra/repositories"
	"github.com/ThailanTec/challenger/pousada/src/config"
	"github.com/ThailanTec/challenger/pousada/src/handlers"
	"github.com/ThailanTec/challenger/pousada/src/middleware"
	"github.com/ThailanTec/challenger/pousada/src/usecases"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB, cfg config.Config) {
	userRepo := repositories.NewUserRepository(db)
	userUsecase := usecases.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)
	authUsecase := usecases.NewAuthUsecase(userRepo, cfg)
	authHandler := handler.NewAuthHandler(authUsecase)

	r.POST("", userHandler.CreateUser)
	r.POST("/login", authHandler.Login)

	userRoutes := r.Group("/users")
	userRoutes.Use(middleware.JWTAuthMiddleware())
	{

		userRoutes.GET("", userHandler.GetUser)
		userRoutes.GET(":document", userHandler.GetUserByDocument)
		userRoutes.DELETE(":id", userHandler.DeleteUser)
		userRoutes.PUT(":id", userHandler.UpdateUser)
	}
}
