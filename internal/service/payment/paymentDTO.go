package payment

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	tools "github.com/pedroRodriguesS5/payment_notification/pkg/utils"
)

type RecurringPaymentRequestDTO struct {
	ReceiverEmail    string  `json:"receiver_email" validate:"required,email"`
	Amount           float64 `json:"amount" validate:"required,numeric"`
	NotificationType string  `json:"notification_type" validate:"required, alpha"`
	StartDate        string  `json:"start_date" validate:"required,datetime"`
	EndDate          string  `json:"end_date" validate:"required,datetime, verify_date"`
	DayOfMont        int32   `json:"day_of_month" validate:"min=1,max=31, required,numeric"`
}

type SelfPaymentDTO struct {
	ReceiverName     string  `json:"receiver_name" validate:"required, alpha"`
	Amount           float64 `json:"amount" validate:"required,numeric"`
	NotificationType string  `json:"notification_type" validate:"required,alpha"`
	StartDate        string  `json:"start_date" validate:"required,datetime"`
	EndDate          string  `json:"end_date" validate:"required,datetime"`
	DayOfMont        int32   `json:"day_of_month" validate:"required,numeric"`
}

type RecurringPaymentResponseDTO struct {
	RecurringPaymentID int32
	PayerID            uuid.UUID
	ReceiverID         uuid.UUID
	PayerName          string
	ReceiverName       string
	Amount             float64
	NotificationType   string
	StartDate          *time.Time
	EndDate            *time.Time
	DayOfMonth         int16
	PaymentStatus      string
}

func (p *RecurringPaymentRequestDTO) Validate(validate *validator.Validate) error {
	return tools.ValidateFunc[RecurringPaymentRequestDTO](*p, validate)
}
