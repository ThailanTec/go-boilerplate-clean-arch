package repositories

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisRepository interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (string, error)
}

type redisRepository struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisRepository(client *redis.Client) RedisRepository {
	return &redisRepository{
		client: client,
		ctx:    context.Background(),
	}
}

func (r *redisRepository) Set(key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(r.ctx, key, value, expiration).Err()
}

func (r *redisRepository) Get(key string) (string, error) {
	return r.client.Get(r.ctx, key).Result()
}
