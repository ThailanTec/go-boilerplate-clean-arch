package database

import (
	"context"
	"github.com/ThailanTec/challenger/pousada/src/config"
	"github.com/redis/go-redis/v9"
)

func RedisClient(cfg config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisADR,      // Endereço do Redis
		Password: cfg.RedisPassword, // Sem senha
		DB:       cfg.RedisDB,       // Usando o banco de dados 0
	})

	return rdb
}

func GetRedisContext() context.Context {
	return context.Background()
}
