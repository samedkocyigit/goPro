package models

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID              string
	Username        string
	Email           string
	Password        string
	PasswordConfirm string
}

var users []User

// AddUser adds a new user to the database
func CreateUser(newUser User) error {
	// ID oluşturma
	newUser.ID = uuid.New().String()

	// Kullanıcı adı, e-posta ve şifre alanlarının dolu olması gerekiyor
	if newUser.Username == "" || newUser.Email == "" || newUser.Password == "" {
		return errors.New("username, email, and password are required")
	}

	// Kullanıcıyı listeye ekle
	users = append(users, newUser)

	return nil
}

// GetUserByID retrieves a user from the database by ID
func GetUserByID(id string) (*User, error) {
	for _, user := range users {
		if user.ID == "id" {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

// GetAllUsers retrieves all users from the database
func GetAllUsers() ([]User, error) {
	return users, nil
}

// Get user by Id
func GetUserByEmail(email string) (*User, error) {
	for _, user := range users {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

// UpdateUser updates a user in the database
func UpdateUser(userID string, updatedUser User) error {
	// Kullanıcıyı bul
	var foundUser *User
	for i, user := range users {
		if user.ID == userID {
			foundUser = &users[i]
			break
		}
	}
	if foundUser == nil {
		return errors.New("user not found")
	}

	// Güncelleme işlemi
	if updatedUser.Username != "" {
		foundUser.Username = updatedUser.Username
	}
	if updatedUser.Email != "" {
		foundUser.Email = updatedUser.Email
	}
	if updatedUser.Password != "" {
		foundUser.Password = updatedUser.Password
	}
	if updatedUser.PasswordConfirm != "" {
		foundUser.PasswordConfirm = updatedUser.PasswordConfirm
	}

	return nil
}

// DeleteUserByID deletes a user from the database by ID
func DeleteUserByID(userID string) error {
	// Kullanıcıyı bul ve kaldır
	foundIndex := -1
	for i, user := range users {
		if user.ID == userID {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		return errors.New("user not found")
	}

	// Kullanıcıyı listeden kaldır
	users = append(users[:foundIndex], users[foundIndex+1:]...)

	return nil
}

// UserExistsWithEmail, verilen e-posta adresine sahip bir kullanıcının var olup olmadığını kontrol eder
func UserExistsWithEmail(email string) bool {
	for _, user := range users {
		if user.Email == email {
			return true
		}
	}
	return false
}

// VerifyUserCredentials, kullanıcı giriş bilgilerini doğrular (e-posta ve parola eşleşmesi)
func VerifyUserCredentials(email, password string) (*User, error) {
	// Veritabanından e-posta adresiyle kullanıcıyı getir
	user, err := GetUserByEmail(email)
	if err != nil {
		return nil, err // Kullanıcı bulunamadı hatası
	}

	// Kullanıcının parolasını kontrol et
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("incorrect password") // Parola eşleşmedi hatası
	}

	// Kullanıcı doğrulandı, kullanıcıyı döndür
	return user, nil
}
