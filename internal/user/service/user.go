// Package service provides the implementation of the UserService interface for user-related operations.
package service

import (
	"time"

	"github.com/gopheramol/domain-driven-arch/internal/user/repository"
	"github.com/gopheramol/domain-driven-arch/pkg/db/orm/sqlc"
)

// UserService is an interface for user-related operations.
type UserService interface {
	CreateUser(user sqlc.User) (sqlc.User, error)
	GetUserByID(id int) (sqlc.User, error)
	GetAllUsers() ([]sqlc.User, error)
	UpdateUser(id int, updatedUser sqlc.UpdateUserParams) error
	DeleteUser(id int) error
}

// UserServiceImpl is the implementation of the UserService interface.
type UserServiceImpl struct {
	userRepo repository.UserRepository
}

// NewUserService creates a new instance of UserServiceImpl with the provided UserRepository.
func NewUserService(userRepo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

// CreateUser creates a new user with the provided data.
func (s *UserServiceImpl) CreateUser(user sqlc.User) (sqlc.User, error) {
	// Implement logic to create a user using the user repository
	userData := sqlc.CreateUserParams{
		Name:      user.Name,
		Email:     user.Email,
		Mobile:    user.Mobile,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdUser, err := s.userRepo.Create(userData)
	if err != nil {
		return sqlc.User{}, nil
	}
	return createdUser, nil
}

// GetUserByID retrieves a user from the database based on the provided ID.
func (s *UserServiceImpl) GetUserByID(id int) (sqlc.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return sqlc.User{}, err
	}
	return user, nil
}

// GetAllUsers retrieves all users from the database.
func (s *UserServiceImpl) GetAllUsers() ([]sqlc.User, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		return []sqlc.User{}, err
	}
	return users, nil
}

// UpdateUser updates an existing user in the database.
func (s *UserServiceImpl) UpdateUser(id int, user sqlc.UpdateUserParams) error {
	user.ID = int64(id)
	err := s.userRepo.Update(user)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes a user from the database based on the provided ID.
func (s *UserServiceImpl) DeleteUser(id int) error {
	err := s.userRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
