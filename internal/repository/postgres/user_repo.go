package postgres

import (
	"context"

	"github.com/hutamy/golang-clean-architecture/internal/domain"
	"github.com/hutamy/golang-clean-architecture/internal/domain/entity"
)

type UserRepository struct {
	dbManager *DBManager
}

func NewUserRepository(dbManager *DBManager) domain.UserRepository {
	return &UserRepository{dbManager: dbManager}
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) error {
	return r.dbManager.Master.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (*entity.User, error) {
	db := r.dbManager.GetReplica()
	user := &entity.User{}
	err := db.WithContext(ctx).First(user, id).Error
	return user, err
}

func (r *UserRepository) FindUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	db := r.dbManager.GetReplica()
	user := &entity.User{}
	err := db.WithContext(ctx).Where("email = ?", email).First(user).Error
	return user, err
}
