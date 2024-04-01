package utils

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
