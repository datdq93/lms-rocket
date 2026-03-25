package service

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"

	"github.com/lms-rocket/lms-backend/internal/domain"
	"github.com/lms-rocket/lms-backend/internal/repository"
)

// AuthService defines authentication operations
type AuthService interface {
	Register(email, password, name, role string) (*domain.User, error)
	Login(email, password string) (*domain.User, string, string, error)
	RefreshToken(refreshToken string) (string, string, error)
	Logout(userID string) error
	GenerateTokens(user *domain.User) (accessToken, refreshToken string, err error)
}

// authService implements AuthService
type authService struct {
	userRepo    repository.UserRepository
	redisClient *redis.Client
}

// NewAuthService creates a new auth service
func NewAuthService(userRepo repository.UserRepository, redisClient *redis.Client) AuthService {
	return &authService{
		userRepo:    userRepo,
		redisClient: redisClient,
	}
}

func (s *authService) Register(email, password, name, role string) (*domain.User, error) {
	// Check if email already exists
	existingUser, err := s.userRepo.FindByEmail(email)
	if err == nil && existingUser != nil {
		return nil, fmt.Errorf("email already registered")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Create user
	user := &domain.User{
		Email:        email,
		PasswordHash: string(hashedPassword),
		Name:         name,
		Role:         role,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

func (s *authService) Login(email, password string) (*domain.User, string, string, error) {
	// Find user
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, "", "", fmt.Errorf("invalid credentials")
	}

	// Check if user is active
	if !user.IsActive {
		return nil, "", "", fmt.Errorf("account is disabled")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, "", "", fmt.Errorf("invalid credentials")
	}

	// Generate tokens
	accessToken, refreshToken, err := s.GenerateTokens(user)
	if err != nil {
		return nil, "", "", fmt.Errorf("failed to generate tokens: %w", err)
	}

	// Update last login
	now := time.Now()
	user.LastLoginAt = &now
	s.userRepo.Update(user)

	return user, accessToken, refreshToken, nil
}

func (s *authService) RefreshToken(refreshToken string) (string, string, error) {
	// TODO: Implement refresh token validation and rotation
	return "", "", fmt.Errorf("not implemented")
}

func (s *authService) Logout(userID string) error {
	// TODO: Invalidate refresh tokens for user
	return nil
}

func (s *authService) GenerateTokens(user *domain.User) (string, string, error) {
	// Access token - 15 minutes
	accessClaims := jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"role":  user.Role,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(15 * time.Minute).Unix(),
		"type":  "access",
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessString, err := accessToken.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", "", err
	}

	// Refresh token - 7 days
	refreshClaims := jwt.MapClaims{
		"sub":  user.ID,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(7 * 24 * time.Hour).Unix(),
		"type": "refresh",
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshString, err := refreshToken.SignedString([]byte(os.Getenv("JWT_REFRESH_SECRET")))
	if err != nil {
		return "", "", err
	}

	// Store refresh token in Redis if available
	if s.redisClient != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		s.redisClient.Set(ctx, "refresh:"+user.ID, refreshString, 7*24*time.Hour)
	}

	return accessString, refreshString, nil
}
