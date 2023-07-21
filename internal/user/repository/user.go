// Package repository provides the implementation of the UserRepository interface for user data operations.
package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/gopheramol/domain-driven-arch/pkg/db/orm/sqlc"
)

// UserRepository is an interface for user data operations.
type UserRepository interface {
	Create(user sqlc.CreateUserParams) (sqlc.User, error)
	GetByID(id int) (sqlc.User, error)
	GetAll() ([]sqlc.User, error)
	Update(user sqlc.UpdateUserParams) error
	Delete(id int) error
}

// UserRepositoryImpl is the implementation of the UserRepository interface.
type UserRepositoryImpl struct {
	db *sql.DB
}

// NewUserRepository creates a new instance of UserRepositoryImpl with the provided SQL database.
func NewUserRepository(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

// Create creates a new user in the database.
func (r *UserRepositoryImpl) Create(user sqlc.CreateUserParams) (sqlc.User, error) {
	ctx := context.Background()
	queries := sqlc.New(r.db)

	createdUser, err := queries.CreateUser(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	return createdUser, nil
}

// GetByID retrieves a user from the database based on the provided ID.
func (r *UserRepositoryImpl) GetByID(id int) (sqlc.User, error) {
	ctx := context.Background()
	queries := sqlc.New(r.db)

	user, err := queries.GetUser(ctx, int64(id))
	if err != nil {
		return sqlc.User{}, fmt.Errorf("user with ID %d not found", id)
	}
	return user, nil
}

// GetAll retrieves all users from the database.
func (r *UserRepositoryImpl) GetAll() ([]sqlc.User, error) {
	ctx := context.Background()
	queries := sqlc.New(r.db)

	users, err := queries.ListUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("users not found")
	}
	return users, nil
}

// Update updates an existing user in the database.
func (r *UserRepositoryImpl) Update(user sqlc.UpdateUserParams) error {
	ctx := context.Background()
	queries := sqlc.New(r.db)

	err := queries.UpdateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("user update failed")
	}
	return nil
}

// Delete deletes a user from the database based on the provided ID.
func (r *UserRepositoryImpl) Delete(id int) error {
	ctx := context.Background()
	queries := sqlc.New(r.db)

	err := queries.DeleteUser(ctx, int64(id))
	if err != nil {
		return fmt.Errorf("user delete failed")
	}
	return nil
}
