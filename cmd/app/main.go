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

	db, err := database.Initialize(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = migrations.Migrate(db)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	r := gin.Default()
	routes.RegisterRoutes(r, db, cfg)

	port := os.Getenv("PORT")
	println(port)
	if port == "" {
		port = "8080"
	}

	err = r.Run(":" + port)
	if err != nil {
		return
	}
}
