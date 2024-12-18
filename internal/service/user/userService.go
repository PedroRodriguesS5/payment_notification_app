package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/pedroRodriguesS5/payment_notification/pkg/infra"
	tools "github.com/pedroRodriguesS5/payment_notification/pkg/utils"
	db "github.com/pedroRodriguesS5/payment_notification/project"
)

type Service struct {
	r *db.Queries
}

func NewService(r *db.Queries) *Service {
	return &Service{
		r: r,
	}
}

func (s *Service) CreateUser(ctx context.Context, userDTO UserRegisterDTO) (string, error) {
	parseDate, err := tools.ConvertStringToDate(userDTO.BornDate)
	if err != nil {
		return "", fmt.Errorf("invalid bornd date: %w", err)
	}

	phoneNumber := pgtype.Text{String: userDTO.PhoneNumber, Valid: true}

	encryptedPass, _ := infra.HashPassword(userDTO.Password)
	params := db.CreateUserParams{
		Name:         userDTO.Name,
		SecondName:   userDTO.SecondName,
		Email:        userDTO.Email,
		UserDocument: userDTO.UserDocument,
		Password:     encryptedPass,
		PhoneNumber:  phoneNumber,
		BornDate:     parseDate,
	}

	user, err := s.r.CreateUser(ctx, params)

	if err != nil {
		return "", fmt.Errorf("erro ao criar usuário: %w", err)
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
