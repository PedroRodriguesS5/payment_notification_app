package tools

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func Passwordvalidation(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	minLength := regexp.MustCompile(`.{8,}`) // Pelo menos 8 caracteres
	hasUpper := regexp.MustCompile(`[A-Z]`)  // Pelo menos uma letra maiúscula
	hasLower := regexp.MustCompile(`[a-z]`)  // Pelo menos uma letra minúscula
	hasNumber := regexp.MustCompile(`\d`)    // Pelo menos um número
	hasSpecial := regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\{\}\[\]:;"'<>,\.\?/\\|~]`)

	return minLength.MatchString(password) &&
		hasUpper.MatchString(password) &&
		hasLower.MatchString(password) &&
		hasNumber.MatchString(password) &&
		hasSpecial.MatchString(password)
}
