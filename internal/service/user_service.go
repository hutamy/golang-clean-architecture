package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/hutamy/golang-clean-architecture/internal/domain"
	"github.com/hutamy/golang-clean-architecture/internal/domain/entity"
	"github.com/hutamy/golang-clean-architecture/pkg/logger"
)

type UserService struct {
	userRepo  domain.UserRepository
	logsRepo  domain.LogsRepository
	cacheRepo domain.CacheRepository
	log       *logger.Logger
}

func NewUserService(
	userRepo domain.UserRepository,
	logsRepo domain.LogsRepository,
	cacheRepo domain.CacheRepository,
	log *logger.Logger,
) domain.UserService {
	return &UserService{
		userRepo,
		logsRepo,
		cacheRepo,
		log,
	}
}

func (s *UserService) Register(ctx context.Context, user *entity.User) error {
	err := s.userRepo.Create(ctx, user)
	if err != nil {
		s.log.Error("Failed to create user: %v", err)
		s.logsRepo.InsertLog(ctx, &entity.Log{Message: err.Error(), Level: "error"})
		return err
	}

	return nil
}

func (s *UserService) GetUser(ctx context.Context, id int) (*entity.User, error) {
	// Check cache first
	cacheKey := "user:" + string(rune(id))
	cachedUser, err := s.cacheRepo.Get(ctx, cacheKey)
	if err == nil {
		var user entity.User
		json.Unmarshal([]byte(cachedUser), &user)
		return &user, nil
	}

	res, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		s.log.Error("Failed to get user: %v", err)
		s.logsRepo.InsertLog(ctx, &entity.Log{Message: err.Error(), Level: "error"})
		return nil, err
	}

	// Cache the result
	userJSON, _ := json.Marshal(res)
	duration := 24 * time.Hour
	s.cacheRepo.Set(ctx, cacheKey, string(userJSON), duration)
	return res, nil
}
