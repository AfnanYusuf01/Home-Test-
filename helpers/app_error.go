package helpers

import "net/http"

// AppError adalah custom error type yang menyimpan HTTP status code
// sehingga handler bisa menentukan response yang tepat tanpa string matching
type AppError struct {
	Code    int
	Message string
}

// Error mengimplementasikan interface error
func (e *AppError) Error() string {
	return e.Message
}

// NewNotFoundError membuat error 404
func NewNotFoundError(message string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

// NewBadRequestError membuat error 400
func NewBadRequestError(message string) *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

// NewInternalError membuat error 500
func NewInternalError(message string) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

// NewConflictError membuat error 409 (misal: duplicate entry)
func NewConflictError(message string) *AppError {
	return &AppError{
		Code:    http.StatusConflict,
		Message: message,
	}
}

// GetStatusCode mengembalikan HTTP status code dari error
// Jika error bukan AppError, default ke 500
func GetStatusCode(err error) int {
	if appErr, ok := err.(*AppError); ok {
		return appErr.Code
	}
	return http.StatusInternalServerError
}
