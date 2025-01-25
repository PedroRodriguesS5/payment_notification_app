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

func NewService(r *db.Queries) *Service {
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

	userIdConvert, err := tools.ConvertStringToUUID(userID)
	if err != nil {
		return "", fmt.Errorf("error to convert stirng into uuid: %v", err)
	}

	notificationTypeConvert := tools.ConvertStringToPgtypeText(recurringDTO.NotificationType)

	convertedAmount, err := tools.ConvertToNumeric(recurringDTO.Amount)
	if err != nil {
		return "", fmt.Errorf("e rror to convert float in pgtype.Numeric: %v", err)
	}

	dayOfMonthConvert, err := tools.ConvertToInt2(recurringDTO.DayOfMont)
	if err != nil {
		return "", fmt.Errorf("error to conevrt go integer in int2, %v", err)
	}

	if err != nil {
		return "", fmt.Errorf("error to find user id: %v", err.Error())
	}

	getReceiverID, err := s.r.GetReceiverIdByEmail(ctx, recurringDTO.ReceiverEmail)

	if err != nil {
		return "", fmt.Errorf("error to get receiver id: %v", err.Error())
	}

	params := db.CreateRecurringPaymentParams{
		PayerID:          userIdConvert,
		ReceiverID:       getReceiverID,
		Amount:           convertedAmount,
		NotificationType: notificationTypeConvert,
		StartDate:        startDate,
		EndDate:          endDate,
		DayOfMonth:       dayOfMonthConvert,
	}

	createdPayment, err := s.r.CreateRecurringPayment(ctx, params)

	if err != nil {
		return "", fmt.Errorf("error to create payment: %v", err.Error())
	}

	return fmt.Sprintln("Payment created: ", createdPayment), nil
}

func (s *Service) CreateSelfRecurringPayment(ctx context.Context, selfPaymentDTO SelfPaymentDTO, token string) (string, error) {
	startDate, err := tools.ConvertStringToDate(selfPaymentDTO.StartDate)
	if err != nil {
		return "", fmt.Errorf("error to parse data: %v", err)
	}

	endDate, err := tools.ConvertStringToDate(selfPaymentDTO.EndDate)

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

	userIdConvert, err := tools.ConvertStringToUUID(userID)
	if err != nil {
		return "", fmt.Errorf("error to convert stirng into uuid: %v", err)
	}

	notificationTypeConvert := tools.ConvertStringToPgtypeText(selfPaymentDTO.NotificationType)

	convertedAmount, err := tools.ConvertToNumeric(selfPaymentDTO.Amount)
	if err != nil {
		return "", fmt.Errorf("e rror to convert float in pgtype.Numeric: %v", err)
	}

	dayOfMonthConvert, err := tools.ConvertToInt2(selfPaymentDTO.DayOfMont)
	if err != nil {
		return "", fmt.Errorf("error to conevrt go integer in int2, %v", err)
	}

	params := db.CreateSelfRecurringPaymentParams{
		PayerID:          userIdConvert,
		ReceiverName:     selfPaymentDTO.ReceiverName,
		Amount:           convertedAmount,
		NotificationType: notificationTypeConvert,
		StartDate:        startDate,
		EndDate:          endDate,
		DayOfMonth:       dayOfMonthConvert,
	}

	createdPayment, err := s.r.CreateSelfRecurringPayment(ctx, params)

	if err != nil {
		return "", fmt.Errorf("error to create payment: %v", err.Error())
	}

	return fmt.Sprintln("Payment created: ", createdPayment), nil

}

func (s *Service) GetRecurringPayement(ctx context.Context, token string) (*db.RecurringPayment, error) {
	claims, err := infra.VerifyToken(token)
	if err != nil {
		return &db.RecurringPayment{}, fmt.Errorf("error to validate token: %v", err)
	}
	userID, ok := claims["user_id"].(string)
	if !ok || userID == "" {
		return &db.RecurringPayment{}, fmt.Errorf("user id not found in token claims")
	}
	convertedUserId, err := tools.ConvertStringToUUID(userID)

	if err != nil {
		return &db.RecurringPayment{}, fmt.Errorf("error to conevvrt string into UUID, %v", err)
	}

	payment, err := s.r.GetRecurringPaymentInfo(ctx, convertedUserId)

	if err != nil {
		return &db.RecurringPayment{}, fmt.Errorf("error to get recurring payment: %v", err)
	}

	return &db.RecurringPayment{
		RecurringPaymentID: payment.RecurringPaymentID,
		ReceiverName:       payment.ReceiverName,
		PayerName:          payment.PayerName,
		Amount:             payment.Amount,
		PaymentStatus:      payment.PaymentStatus,
		StartDate:          payment.StartDate,
		EndDate:            payment.EndDate,
		DayOfMonth:         payment.DayOfMonth,
	}, nil
}
