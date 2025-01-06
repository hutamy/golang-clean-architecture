package domain

import (
	"context"

	"github.com/hutamy/golang-clean-architecture/internal/domain/entity"
)

type UserService interface {
	Register(ctx context.Context, user *entity.User) error
	GetUser(ctx context.Context, id int) (*entity.User, error)
}

type AuthService interface {
	Login(ctx context.Context, email, password string) (string, error)
}
