package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID              int
	Username        string
	Email           string
	Password        string
	PasswordConfirm string
}

var users []User

// AddUser adds a new user to the database
func AddUser(user User) error {
	// Check if user already exists
	for _, u := range users {
		if u.ID == user.ID {
			return errors.New("user already exists")
		}
	}

	// Generate ID for the new user
	user.ID = len(users) + 1

	// Append user to the users slice
	users = append(users, user)

	fmt.Println("User added:", user)
	return nil
}

// GetUserByID retrieves a user from the database by ID
func GetUserByID(id int) (*User, error) {
	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

// GetAllUsers retrieves all users from the database
func GetAllUsers() ([]User, error) {
	return users, nil
}

// UpdateUser updates a user in the database
func UpdateUser(updatedUser User) error {
	for i, user := range users {
		if user.ID == updatedUser.ID {
			users[i] = updatedUser
			fmt.Println("User updated:", updatedUser)
			return nil
		}
	}
	return errors.New("user not found")
}

// DeleteUserByID deletes a user from the database by ID
func DeleteUserByID(id int) error {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			fmt.Println("User deleted with ID:", id)
			return nil
		}
	}
	return errors.New("user not found")
}
