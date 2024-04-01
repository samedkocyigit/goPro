package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

// Email struct represents an email object
type Email struct {
	To        string
	FirstName string
	URL       string
	From      string
}

// NewEmail creates a new instance of Email
func NewEmail(user User, url string) *Email {
	return &Email{
		To:        user.Email,
		FirstName: user.Name,
		URL:       url,
		From:      fmt.Sprintf("Jonas Schmedtmann <%s>", os.Getenv("EMAIL_FROM")),
	}
}

// SendWelcome sends a welcome email
func (e *Email) SendWelcome() error {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		return err
	}

	// Create new email message
	m := gomail.NewMessage()
	m.SetHeader("From", e.From)
	m.SetHeader("To", e.To)
	m.SetHeader("Subject", "Welcome to the Natours Family!")

	// HTML body
	html := fmt.Sprintf("<p>Welcome %s to the Natours Family!</p>", e.FirstName)
	m.SetBody("text/html", html)

	// Send email
	d := gomail.NewDialer(os.Getenv("EMAIL_HOST"), 587, os.Getenv("EMAIL_USERNAME"), os.Getenv("EMAIL_PASSWORD"))
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

// SendPasswordReset sends a password reset email
func (e *Email) SendPasswordReset() error {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		return err
	}

	// Create new email message
	m := gomail.NewMessage()
	m.SetHeader("From", e.From)
	m.SetHeader("To", e.To)
	m.SetHeader("Subject", "Your password reset token (valid for only 10 minutes)")

	// HTML body
	html := "<p>Your password reset token (valid for only 10 minutes)</p>"
	m.SetBody("text/html", html)

	// Send email
	d := gomail.NewDialer(os.Getenv("EMAIL_HOST"), 587, os.Getenv("EMAIL_USERNAME"), os.Getenv("EMAIL_PASSWORD"))
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
