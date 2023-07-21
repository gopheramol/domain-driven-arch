// Package handler provides the HTTP handlers for user-related operations.
package handler

import (
	"net/http"
	"strconv"

	"github.com/gopheramol/domain-driven-arch/internal/user/service"
	"github.com/gopheramol/domain-driven-arch/pkg/db/orm/sqlc"

	"github.com/gin-gonic/gin"
)

// UserHandler represents the HTTP handler for user-related operations.
type UserHandler struct {
	userService service.UserService
}

// NewUserHandler creates a new instance of UserHandler with the provided user service.
func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// CreateUser handles the HTTP request to create a new user.
func (h *UserHandler) CreateUser(c *gin.Context) {
	// Bind the incoming JSON request body to a User struct
	var user sqlc.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the UserService's CreateUser method to create the user
	createdUser, err := h.userService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the created user as JSON
	c.JSON(http.StatusCreated, createdUser)
}

// GetAllUsers handles the HTTP request to retrieve all users.
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	// Call the UserService's GetAllUsers method to get all users
	users, err := h.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the list of users as JSON
	c.JSON(http.StatusOK, users)
}

// GetUserByID handles the HTTP request to retrieve a user by ID.
func (h *UserHandler) GetUserByID(c *gin.Context) {
	// Parse the user ID from the URL parameter
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Call the UserService's GetUserByID method to get the user by ID
	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Respond with the user as JSON
	c.JSON(http.StatusOK, user)
}

// UpdateUser handles the HTTP request to update a user.
func (h *UserHandler) UpdateUser(c *gin.Context) {
	// Parse the user ID from the URL parameter
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Bind the incoming JSON request body to an UpdateUserParams struct
	var updatedUser sqlc.UpdateUserParams
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the UserService's UpdateUser method to update the user
	upadateErr := h.userService.UpdateUser(userID, updatedUser)
	if upadateErr != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": upadateErr.Error()})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteUser handles the HTTP request to delete a user.
func (h *UserHandler) DeleteUser(c *gin.Context) {
	// Parse the user ID from the URL parameter
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Call the UserService's DeleteUser method to delete the user
	err = h.userService.DeleteUser(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
