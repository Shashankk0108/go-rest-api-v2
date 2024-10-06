package service

import (
	"errors"

	"github.com/botsgalaxy/go-rest-api-v2/internal/models"
	"github.com/botsgalaxy/go-rest-api-v2/internal/repository"
)

// UserService handles user business logic
type UserService interface {
	GetUserByID(id uint) (*models.UserResponse, error)
	GetAllUsers(limit, offset int) ([]models.UserResponse, error)
	UpdateUser(id uint, req *models.UpdateUserRequest) (*models.UserResponse, error)
	DeleteUser(id uint) error
}

type userService struct {
	userRepo repository.UserRepository
}

// NewUserService creates a new user service
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetUserByID(id uint) (*models.UserResponse, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	userResponse := user.ToResponse()
	return &userResponse, nil
}

func (s *userService) GetAllUsers(limit, offset int) ([]models.UserResponse, error) {
	users, err := s.userRepo.FindAll(limit, offset)
	if err != nil {
		return nil, err
	}

	var userResponses []models.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, user.ToResponse())
	}

	return userResponses, nil
}

func (s *userService) UpdateUser(id uint, req *models.UpdateUserRequest) (*models.UserResponse, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if req.Username != "" {
		// Check if username is already taken
		existingUser, _ := s.userRepo.FindByUsername(req.Username)
		if existingUser != nil && existingUser.ID != id {
			return nil, errors.New("username already taken")
		}
		user.Username = req.Username
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, errors.New("failed to update user")
	}

	userResponse := user.ToResponse()
	return &userResponse, nil
}

func (s *userService) DeleteUser(id uint) error {
	_, err := s.userRepo.FindByID(id)
	if err != nil {
		return err
	}

	return s.userRepo.Delete(id)
}
