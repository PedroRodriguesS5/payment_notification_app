package user

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/pedroRodriguesS5/payment_notification/pkg/utils"
	"github.com/pedroRodriguesS5/payment_notification/pkg/utils/tools"
	db "github.com/pedroRodriguesS5/payment_notification/project"
)

type Service struct {
	r *db.Queries
}

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

func NewService(r *db.Queries) *Service {
	return &Service{
		r: r,
	}
}

func (s *Service) CreateUser(ctx context.Context, userDTO UserRegisterDTO) (string, error) {
	var bornDate pgtype.Date
	parseDate, err := time.Parse("2006-01-02", userDTO.BornDate)

	if err != nil {
		return "", fmt.Errorf("invalid bornd date: %w", err)
	}

	bornDate.Time = parseDate
	bornDate.Valid = true

	phoneNumber := pgtype.Text{String: userDTO.PhoneNumber, Valid: true}

	encryptedPass, _ := utils.HashPassword(userDTO.Password)
	params := db.CreateUserParams{
		Name:         userDTO.Name,
		SecondName:   userDTO.SecondName,
		Email:        userDTO.Email,
		UserDocument: userDTO.UserDocument,
		Password:     encryptedPass,
		PhoneNumber:  phoneNumber,
		BornDate:     bornDate,
	}

	user, err := s.r.CreateUser(ctx, params)

	if err != nil {
		return "", fmt.Errorf("error to create user: %w", err)
	}

	return fmt.Sprintln("Usuário criado com sucesso: ", user), nil
}

func (s *Service) GetUser(ctx context.Context, userID string) (*db.User, error) {
	parseUUID, err := uuid.Parse(userID)

	if err != nil {
		return nil, fmt.Errorf("error to convert text to uuid: %w", err)
	}

	var pgUUID pgtype.UUID
	pgUUID.Bytes = parseUUID
	pgUUID.Valid = true

	user, err := s.r.GetUser(ctx, pgUUID)

	if err != nil {
		return nil, fmt.Errorf("error to get data from data base: %w", err)
	}

	return &db.User{
		UserID:       user.UserID,
		Name:         user.Name,
		SecondName:   user.SecondName,
		Password:     user.Password,
		Email:        user.Email,
		UserDocument: user.UserDocument,
		BornDate:     user.BornDate,
		PhoneNumber:  user.PhoneNumber,
		CreatedAt:    user.CreatedAt,
	}, nil
}

func (s *Service) GetAllUsers(ctx context.Context) ([]UserResponseDTO, error) {

	u, err := s.r.GetAllUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("error to get data from data base: %v", err)
	}
	var users []UserResponseDTO

	for _, j := range u {
		convertID, _ := tools.ConvertUUIDToString(j.UserID)
		users = append(users, UserResponseDTO{
			UserID: convertID,
			Name:   j.Name,
			Email:  j.Email,
		})
	}
	return users, nil
}

func (s *Service) GetUserByEmail(ctx context.Context, userEmail string) (*db.User, error) {
	user, err := s.r.GetUserByEmail(ctx, userEmail)

	if err != nil {
		return nil, fmt.Errorf("user not found: %v", err.Error())
	}

	return &db.User{
		UserID:   user.UserID,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
