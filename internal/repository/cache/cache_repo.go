package cache

import (
	"context"
	"time"

	"github.com/hutamy/golang-clean-architecture/internal/domain"
	"github.com/redis/go-redis/v9"
)

type CacheRepository struct {
	client *redis.Client
}

func NewCacheRepository(dbManager *DBManager) domain.CacheRepository {
	return &CacheRepository{
		client: dbManager.GetClient(),
	}
}

func (r *CacheRepository) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *CacheRepository) Get(ctx context.Context, key string) (string, error) {
	result, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Key not found
	}
	return result, err
}

func (r *CacheRepository) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}
