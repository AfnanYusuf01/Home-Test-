package helpers

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// APIResponse adalah format standar untuk semua response API yang profesional
type APIResponse struct {
	Success   bool        `json:"success"`
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Timestamp string      `json:"timestamp"`
	Data      interface{} `json:"data,omitempty"`
	Errors    interface{} `json:"errors,omitempty"`
}

// SuccessResponse mengirim response sukses dengan metadata lengkap
func SuccessResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return c.Status(statusCode).JSON(APIResponse{
		Success:   true,
		Code:      statusCode,
		Message:   message,
		Timestamp: time.Now().Format(time.RFC3339),
		Data:      data,
	})
}

// ErrorResponse mengirim response error standar
func ErrorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(APIResponse{
		Success:   false,
		Code:      statusCode,
		Message:   message,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// ValidationErrorResponse menangani error validasi dari validator v10
func ValidationErrorResponse(c *fiber.Ctx, err error) error {
	errors := make(map[string]string)

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			// Mengambil nama field dari json tag jika memungkinkan
			field := strings.ToLower(err.Field())
			
			var msg string
			switch err.Tag() {
			case "required":
				msg = "Field ini wajib diisi"
			case "min":
				msg = fmt.Sprintf("Minimal %s karakter/nilai", err.Param())
			case "max":
				msg = fmt.Sprintf("Maksimal %s karakter/nilai", err.Param())
			case "gt":
				msg = fmt.Sprintf("Nilai harus lebih besar dari %s", err.Param())
			case "numeric":
				msg = "Harus berupa angka"
			default:
				msg = fmt.Sprintf("Tag validasi '%s' gagal", err.Tag())
			}
			errors[field] = msg
		}
	}

	return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
		Success:   false,
		Code:      fiber.StatusBadRequest,
		Message:   "Validasi input gagal",
		Timestamp: time.Now().Format(time.RFC3339),
		Errors:    errors,
	})
}


