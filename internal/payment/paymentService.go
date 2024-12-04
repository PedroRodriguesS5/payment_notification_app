package payment

import (
	"context"
	"fmt"

	"github.com/pedroRodriguesS5/payment_notification/pkg/infra"
	tools "github.com/pedroRodriguesS5/payment_notification/pkg/utils"
	db "github.com/pedroRodriguesS5/payment_notification/project"
)

type Service struct {
	r *db.Queries
}

func NewServvice(r *db.Queries) *Service {
	return &Service{
		r: r,
	}
}

// Payments service
func (s *Service) CreateRecurringPayments(ctx context.Context, recurringDTO RecurringPaymentRequestDTO, token string) (string, error) {
	startDate, err := tools.ConvertStringToDate(recurringDTO.StartDate)
	if err != nil {
		return "", fmt.Errorf("error to parse data: %v", err)
	}

	endDate, err := tools.ConvertStringToDate(recurringDTO.EndDate)

	if err != nil {
		return "", fmt.Errorf("error to parse data: %v", err)
	}
	claims, err := infra.VerifyToken(token)
	if err != nil {
		return "", fmt.Errorf("error to validate token: %v", err)
	}

	userID, ok := claims["user_id"].(string)
	if !ok || userID == "" {
		return "", fmt.Errorf("user id not found in token claims")
	}

	convertedUserId, err := tools.ConvertStringToUUID(userID)
	if err != nil {
		return "", fmt.Errorf("error to convert stirng into uuid: %v", err)
	}

	convertedReceiverId, err := tools.ConvertStringToUUID(recurringDTO.ReceiverID)
	if err != nil {
		return "", fmt.Errorf("error to convert data string to pgtype UUDI: %v", err)
	}

	convertedNotificationType := tools.ConvertStringToPgtypeText(recurringDTO.NotificationType)

	params := db.CreateRecurringPaymentParams{
		PayerID:          convertedUserId,
		ReceiverID:       convertedReceiverId,
		Amount:           recurringDTO.Amount,
		NotificationType: convertedNotificationType,
		StartDate:        startDate,
		EndDate:          endDate,
		DayOfMonth:       recurringDTO.DayOfMont,
	}

	createdPayment, err := s.r.CreateRecurringPayment(ctx, params)

	if err != nil {
		return "", fmt.Errorf("error to create payment: %v", err.Error())
	}

	return fmt.Sprintln("Payment created: ", createdPayment), nil
}
