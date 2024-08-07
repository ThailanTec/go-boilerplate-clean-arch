package database

import (
	"fmt"
	"github.com/ThailanTec/challenger/pousada/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initialize(cfg config.Config) (*gorm.DB, error) {
	fmt.Println(cfg)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUsername, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
