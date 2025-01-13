package payment

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	tools "github.com/pedroRodriguesS5/payment_notification/pkg/utils"
)

type RecurringPaymentRequestDTO struct {
	ReceiverID       string  `json:"receiverID" valid:"IsUUID, required"`
	Amount           float64 `json:"amount" valid:"IsFLoat, required"`
	NotificationType string  `json:"notification_type" valid:"IsString, required"`
	StartDate        string  `json:"start_date" valid:"datetime(2006-01-02),required"`
	EndDate          string  `json:"end_date" valid:"datetime(2006-01-02),required"`
	DayOfMont        int32   `json:"day_of_month" valid:"IsInt, range(1|31), required"`
}

type SelfPaymentDTO struct {
	ReceiverName     string  `json:"receiver_name" valid:"IsString, required"`
	Amount           float64 `json:"amount" valid:"IsFloat, required"`
	NotificationType string  `json:"notification_type" valid:"IsString,required"`
	StartDate        string  `json:"start_date" valid:"datetime(2006-01-02),required"`
	EndDate          string  `json:"end_date" valid:"datetime(2006-01-02),required"`
	DayOfMont        int32   `json:"day_of_month" valid:"IsInt, required"`
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
