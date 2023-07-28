package validators

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

// Kểm tra xem trường title có chứa từ "" hay không
// nếu có trả về true hoặc ngược lại
func ValidateCoolTitle(field validator.FieldLevel) bool {
	return strings.Contains(field.Field().String(), "")
}
