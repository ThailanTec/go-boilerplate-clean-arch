package config

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

type Config struct {
	JWTSecret            string
	JWTExpirationMinutes int
	DBUsername           string
	DBPassword           string
	DBName               string
	DBHost               string
	DBPort               string
	RedisADR             string
	RedisPassword        string
	RedisDB              int
	RedisTLL             time.Duration
}

func LoadConfig() Config {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	config := Config{
		JWTSecret:            viper.GetString("JWTSecret"),
		JWTExpirationMinutes: viper.GetInt("JWTExpirationMinutes"),
		DBUsername:           viper.GetString("DB_USERNAME"),
		DBPassword:           viper.GetString("DB_PASSWORD"),
		DBName:               viper.GetString("DB_NAME"),
		DBHost:               viper.GetString("DB_HOST"),
		DBPort:               viper.GetString("DB_PORT"),
		RedisADR:             viper.GetString("REDIS_ADR"),
		RedisPassword:        viper.GetString("REDIS_PASSWORD"),
		RedisDB:              viper.GetInt("REDIS_DB"),
		RedisTLL:             viper.GetDuration("REDIS_TLL"),
	}

	return config
}
