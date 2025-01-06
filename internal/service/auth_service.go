package service

import (
	"context"
	"errors"
	"time"

	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hutamy/golang-clean-architecture/config"
	"github.com/hutamy/golang-clean-architecture/internal/domain"
	"github.com/hutamy/golang-clean-architecture/pkg/logger"
)

type AuthService struct {
	userRepo domain.UserRepository
	config   *config.Config
	log      *logger.Logger
}

func NewAuthService(userRepo domain.UserRepository, config *config.Config, log *logger.Logger) *AuthService {
	return &AuthService{
		userRepo,
		config,
		log,
	}
}

func (s *AuthService) GenerateJWT(userID string) (string, error) {
	// Set claims
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expiration time (24 hours)
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.config.Jwt.SecretKey))
	if err != nil {
		s.log.Error("Error generating JWT", "error", err.Error())
		return "", err
	}

	return tokenString, nil
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.userRepo.FindUserByEmail(ctx, email)
	if err != nil {
		s.log.Error("Error finding user", "email", email, "error", err.Error())
		return "", errors.New("invalid credentials")
	}

	if user.Password != password {
		s.log.Warn("Incorrect password", "email", email)
		return "", errors.New("invalid credentials")
	}

	token, err := s.GenerateJWT(strconv.Itoa(user.ID))
	if err != nil {
		return "", errors.New("error generating token")
	}

	s.log.Info("User logged in successfully", "email", email, "user_id", user.ID)
	return token, nil
}
