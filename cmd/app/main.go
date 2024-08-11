package main

import (
	"github.com/ThailanTec/challenger/pousada/infra/database"
	"github.com/ThailanTec/challenger/pousada/infra/database/migrations"
	"github.com/ThailanTec/challenger/pousada/src/config"
	"github.com/ThailanTec/challenger/pousada/src/routes"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	// Carregar configurações
	cfg := config.LoadConfig()

	// Inicializar banco de dados
	db, err := database.Initialize(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrar o banco de dados
	err = migrations.Migrate(db)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Configurar Gin
	r := gin.Default()
	routes.RegisterRoutes(r, db, cfg)

	// Iniciar servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err = r.Run(":" + port)
	if err != nil {
		return
	}
}
