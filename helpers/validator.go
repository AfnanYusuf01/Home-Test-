package helpers

import "github.com/go-playground/validator/v10"

var Validate = validator.New()

// ValidateStruct memvalidasi struct menggunakan tag validate
func ValidateStruct(s interface{}) error {
	return Validate.Struct(s)
}
