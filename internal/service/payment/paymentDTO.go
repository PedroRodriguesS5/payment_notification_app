package payment

import (
	"time"

	"github.com/google/uuid"
)

type RecurringPaymentRequestDTO struct {
	ReceiverID       string  `json:"receiverID"`
	Amount           float64 `json:"amount"`
	NotificationType string  `json:"notification_type"`
	StartDate        string  `json:"start_date"`
	EndDate          string  `json:"end_date"`
	DayOfMont        int32   `json:"day_of_month"`
}

type SelfPaymentDTO struct {
	ReceiverName     string  `json:"receiver_name"`
	Amount           float64 `json:"amount"`
	NotificationType string  `json:"notification_type"`
	StartDate        string  `json:"start_date"`
	EndDate          string  `json:"end_date"`
	DayOfMont        int32   `json:"day_of_month"`
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
