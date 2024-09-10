package main

import (
	"log"
	"os"

	"github.com/ThailanTec/challenger/pousada/infra/database"
	"github.com/ThailanTec/challenger/pousada/infra/database/migrations"
	"github.com/ThailanTec/challenger/pousada/src/config"
	"github.com/ThailanTec/challenger/pousada/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	logger, err := config.InitLogger()
	if err != nil {
		log.Fatalf("Failed To init a logger", err)
	}

	db, err := database.PostgresClient(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	redis := database.RedisClient(cfg)

	err = migrations.Migrations(db)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	r := gin.Default()
	routes.RegisterRoutes(r, db, redis, cfg, logger)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err = r.Run(":" + port)
	logger.Info("Aplicação iniciada na porta: " + port)
	if err != nil {
		return
	}
}
