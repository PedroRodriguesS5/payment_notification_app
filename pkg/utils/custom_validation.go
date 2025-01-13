package tools

import (
	"regexp"
	"strconv"
	"strings"

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

func ValidateDoc(fl validator.FieldLevel) bool {
	doc := fl.Field().String()
	// Remove caracteres não numéricos
	re := regexp.MustCompile(`[^0-9]`)
	cleanDoc := re.ReplaceAllString(doc, "")

	// Verifica se é CPF ou CNPJ
	if len(cleanDoc) == 11 {
		return validateCPF(cleanDoc)
	} else if len(cleanDoc) == 14 {
		return ValidateCNPJ(cleanDoc)
	}
	return false
}
func validateCPF(cpf string) bool {
	// Verifica se tem 11 dígitos ou se todos os dígitos são iguais
	if len(cpf) != 11 || strings.Repeat(string(cpf[0]), len(cpf)) == cpf {
		return false
	}
	// Calcula o primeiro dígito verificador
	soma := 0
	for i := 0; i < 9; i++ {
		num, _ := strconv.Atoi(string(cpf[i]))
		soma += num * (10 - i)
	}
	resto := soma % 11
	dv1 := 0
	if resto >= 2 {
		dv1 = 11 - resto
	}

	// Calcula o segundo dígito verificador
	soma = 0
	for i := 0; i < 10; i++ {
		num, _ := strconv.Atoi(string(cpf[i]))
		soma += num * (11 - i)
	}
	resto = soma % 11
	dv2 := 0
	if resto >= 2 {
		dv2 = 11 - resto
	}

	// Verifica se os dígitos verificadores estão corretos
	return strconv.Itoa(dv1)+strconv.Itoa(dv2) == cpf[9:]
}

func ValidateCNPJ(cnpj string) bool {
	// Verifica se tem 14 dígitos ou se todos os dígitos são iguais
	if len(cnpj) != 14 || strings.Repeat(string(cnpj[0]), len(cnpj)) == cnpj {
		return false
	}

	// Peso para os cálculos
	pesos1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	pesos2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	// Calcula o primeiro dígito verificador
	soma := 0
	for i := 0; i < 12; i++ {
		num, _ := strconv.Atoi(string(cnpj[i]))
		soma += num * pesos1[i]
	}
	resto := soma % 11
	dv1 := 0
	if resto >= 2 {
		dv1 = 11 - resto
	}

	// Calcula o segundo dígito verificador
	soma = 0
	for i := 0; i < 13; i++ {
		num, _ := strconv.Atoi(string(cnpj[i]))
		soma += num * pesos2[i]
	}
	resto = soma % 11
	dv2 := 0
	if resto >= 2 {
		dv2 = 11 - resto
	}

	// Verifica se os dígitos verificadores estão corretos
	return strconv.Itoa(dv1)+strconv.Itoa(dv2) == cnpj[12:]
}
