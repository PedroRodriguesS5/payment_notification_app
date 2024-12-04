package payment

type RecurringPaymentRequestDTO struct {
	ReceiverID       string  `json:"receiverID"`
	Amount           float64 `json:"amount"`
	NotificationType string  `json:"notification_type"`
	StartDate        string  `json:"start_date"`
	EndDate          string  `json:"end_date"`
	DayOfMont        int32   `json:"day_of_month"`
}
