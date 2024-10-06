package service

import (
	"errors"
	"time"

	"github.com/botsgalaxy/go-rest-api-v2/internal/models"
	"github.com/botsgalaxy/go-rest-api-v2/internal/repository"
	"github.com/botsgalaxy/go-rest-api-v2/internal/utils"
)

// AuthService handles authentication business logic
type AuthService interface {
	Register(req *models.RegisterRequest) (*models.UserResponse, error)
	Login(req *models.LoginRequest) (string, *models.UserResponse, error)
}

type authService struct {
	userRepo   repository.UserRepository
	jwtSecret  string
	jwtExpiry  time.Duration
}

// NewAuthService creates a new auth service
func NewAuthService(userRepo repository.UserRepository, jwtSecret string, jwtExpiry time.Duration) AuthService {
	return &authService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
		jwtExpiry: jwtExpiry,
	}
}

func (s *authService) Register(req *models.RegisterRequest) (*models.UserResponse, error) {
	// Check if user already exists
	existingUser, _ := s.userRepo.FindByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	existingUser, _ = s.userRepo.FindByUsername(req.Username)
	if existingUser != nil {
		return nil, errors.New("username already taken")
	}

	// Create new user
	user := &models.User{
		Email:     req.Email,
		Username:  req.Username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      "user",
		IsActive:  true,
	}

	// Hash password
	if err := user.HashPassword(req.Password); err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Save user
	if err := s.userRepo.Create(user); err != nil {
		return nil, errors.New("failed to create user")
	}

	userResponse := user.ToResponse()
	return &userResponse, nil
}

func (s *authService) Login(req *models.LoginRequest) (string, *models.UserResponse, error) {
	// Find user by email
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	// Check if user is active
	if !user.IsActive {
		return "", nil, errors.New("user account is deactivated")
	}

	// Verify password
	if !user.CheckPassword(req.Password) {
		return "", nil, errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID, user.Email, user.Role, s.jwtSecret, s.jwtExpiry)
	if err != nil {
		return "", nil, errors.New("failed to generate token")
	}

	userResponse := user.ToResponse()
	return token, &userResponse, nil
}
