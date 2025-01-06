package domain

import (
	"context"
	"time"

	"github.com/hutamy/golang-clean-architecture/internal/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	GetByID(ctx context.Context, id int) (*entity.User, error)
	FindUserByEmail(ctx context.Context, email string) (*entity.User, error)
}

type LogsRepository interface {
	InsertLog(ctx context.Context, log *entity.Log) error
	FindLogs(ctx context.Context, filter map[string]interface{}) ([]*entity.Log, error)
}

type CacheRepository interface {
	Set(ctx context.Context, key string, value string, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}
