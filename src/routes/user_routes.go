package routes

import (
	"github.com/ThailanTec/challenger/pousada/infra/repositories"
	"github.com/ThailanTec/challenger/pousada/src/config"
	handler "github.com/ThailanTec/challenger/pousada/src/handlers"
	"github.com/ThailanTec/challenger/pousada/src/middleware"
	"github.com/ThailanTec/challenger/pousada/src/usecases"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB, cfg config.Config, logger *zap.Logger) {
	userRepo := repositories.NewUserRepository(db)
	userUsecase := usecases.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase, logger)
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
