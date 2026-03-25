package service

import (
	"github.com/lms-rocket/lms-backend/internal/domain"
	"github.com/lms-rocket/lms-backend/internal/repository"
)

// UserService defines user operations
type UserService interface {
	GetProfile(userID string) (*domain.User, error)
	UpdateProfile(userID string, updates map[string]interface{}) (*domain.User, error)
	ChangePassword(userID, currentPassword, newPassword string) error
	ListUsers(page, limit int) ([]domain.User, int64, error)
	GetUserByID(id string) (*domain.User, error)
	UpdateUser(id string, updates map[string]interface{}) (*domain.User, error)
	DeleteUser(id string) error
}

// userService implements UserService
type userService struct {
	userRepo repository.UserRepository
}

// NewUserService creates a new user service
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetProfile(userID string) (*domain.User, error) {
	return s.userRepo.FindByID(userID)
}

func (s *userService) UpdateProfile(userID string, updates map[string]interface{}) (*domain.User, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	// Apply updates
	if name, ok := updates["name"].(string); ok {
		user.Name = name
	}
	if avatarURL, ok := updates["avatar_url"].(string); ok {
		user.AvatarURL = &avatarURL
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) ChangePassword(userID, currentPassword, newPassword string) error {
	// TODO: Implement
	return nil
}

func (s *userService) ListUsers(page, limit int) ([]domain.User, int64, error) {
	return s.userRepo.List(page, limit)
}

func (s *userService) GetUserByID(id string) (*domain.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *userService) UpdateUser(id string, updates map[string]interface{}) (*domain.User, error) {
	// TODO: Implement
	return nil, nil
}

func (s *userService) DeleteUser(id string) error {
	return s.userRepo.Delete(id)
}
