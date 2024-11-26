package user

import "time"

type UserRegisterDTO struct {
	Name         string `json:"name"`
	SecondName   string `json:"second_name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	PhoneNumber  string `json:"phone_number"`
	UserDocument string `json:"user_document"`
	BornDate     string `json:"born_date"`
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
