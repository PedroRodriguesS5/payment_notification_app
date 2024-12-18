package user

import (
	"time"

	"github.com/go-playground/validator/v10"
	tools "github.com/pedroRodriguesS5/payment_notification/pkg/utils"
)

type UserRegisterDTO struct {
	Name         string `json:"name" validate:"required,min=3,max=20,alpha" errormgs:"O nome é obrigatório e deve conter entre 3 e 20 caracteres usando apenas letras e espaços"`
	SecondName   string `json:"second_name" validate:"required,min=2,max=40,alpha" errormgs:"O sobrenome é obrigatório e deve conter entre 2 e 40 caracteres usando apenas letras e espaços"`
	Email        string `json:"email" validate:"required,email" errormgs:"O email é obrigatório e deve ser válido ou já existe no banco"`
	Password     string `json:"password" validate:"required,min=8,password_strength" errormgs:"A senha é obrigatória e deve ter ao menos 8 caracteres, incluindo uma letra maiúscula, uma minúscula, um número e um caractere especial"`
	PhoneNumber  string `json:"phone_number" validate:"required,min=11,max=15" errormgs:"O número de telefone é obrigatório e deve conter 11 dígitos"`
	UserDocument string `json:"user_document" validate:"required,min=11,max=14,numeric" errormgs:"O CPF ou CNPJ é obrigatório e deve conter 11 ou 14 números"`
	BornDate     string `json:"born_date" validate:"required,datetime=2006-01-02" errormgs:"A data de nascimento é obrigatória e deve estar no formato YYYY-MM-DD"`
}

type LoginUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponseDTO struct {
	UserID       string     `json:"user_id"`
	Name         string     `json:"name"`
	SecondName   string     `json:"second_name,omitempty"`
	Email        string     `json:"email"`
	Password     string     `json:"-"` // Hides the password field entirely
	UserDocument string     `json:"user_document,omitempty"`
	PhoneNumber  *string    `json:"phone_number,omitempty"`
	BornDate     *time.Time `json:"born_date,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
}

func (u *UserRegisterDTO) Validate(validate *validator.Validate) error {
	return tools.ValidateFunc[UserRegisterDTO](*u, validate)
}
