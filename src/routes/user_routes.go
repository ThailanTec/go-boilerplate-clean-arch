package routes

import (
	"github.com/ThailanTec/challenger/pousada/infra/repositories"
	"github.com/ThailanTec/challenger/pousada/src/config"
	handler "github.com/ThailanTec/challenger/pousada/src/handlers"
	"github.com/ThailanTec/challenger/pousada/src/middleware"
	"github.com/ThailanTec/challenger/pousada/src/usecases"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func RegisterRoutes(r *gin.Engine, db *sqlx.DB, clientRedis *redis.Client, cfg config.Config, logger *zap.Logger) {
	userRepo := repositories.NewUserRepository(db)
	redisRepo := repositories.NewRedisRepository(clientRedis)
	userUsecase := usecases.NewUserUsecase(userRepo, redisRepo, cfg.RedisTLL)
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
