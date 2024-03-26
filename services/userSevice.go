package services

import (
	"errors"
	"goProject/models"
)

// CreateUser validates and adds a new user
func CreateUser(user models.User) error {
	// Validate user fields (e.g., username, email)
	if user.Username == "" || user.Email == "" {
		return errors.New("username and email are required")
	}

	// Add user to the database
	err := models.AddUser(user)
	if err != nil {
		return err
	}

	return nil
}

// GetUserByID retrieves a user by ID
func GetUserByID(id int) (*models.User, error) {
	user, err := models.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetAllUsers retrieves all users
func GetAllUsers() ([]models.User, error) {
	users, err := models.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// UpdateUser updates a user
func UpdateUser(updatedUser models.User) error {
	err := models.UpdateUser(updatedUser)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUserByID deletes a user by ID
func DeleteUserByID(id int) error {
	err := models.DeleteUserByID(id)
	if err != nil {
		return err
	}
	return nil
}