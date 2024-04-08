package utils

import (
	"encoding/json"
	"net/http"
	"os"
)

// AppError is a custom error type
type AppError struct {
	Message       string
	StatusCode    int
	Status        string
	IsOperational bool
}

// NewAppError creates a new instance of AppError
func NewAppError(message string, statusCode int) *AppError {
	status := "error"
	if statusCode >= 400 && statusCode < 500 {
		status = "fail"
	}

	return &AppError{
		Message:       message,
		StatusCode:    statusCode,
		Status:        status,
		IsOperational: true,
	}
}

func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	response := map[string]interface{}{"error": message}
	json.NewEncoder(w).Encode(response)
}

func RespondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func IsProductionEnvironment() bool {
	environment := os.Getenv("APP_ENV")
	return environment == "production"
}
