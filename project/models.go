// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Notification struct {
	NotificationID     int32
	PaymentID          int32
	RecurringPaymentID pgtype.Int4
	Amount             pgtype.Numeric
	NotificationType   pgtype.Text
	NotificationDate   pgtype.Timestamp
	Status             pgtype.Text
}

type RecurringPayment struct {
	RecurringPaymentID int32
	PayerID            pgtype.UUID
	ReceiverID         pgtype.UUID
	Amount             pgtype.Numeric
	StartDate          pgtype.Date
	EndDate            pgtype.Date
	DayOfMonth         pgtype.Int2
	PaymentStatus      pgtype.Text
}

type User struct {
	UserID       pgtype.UUID
	Name         string
	Email        string
	Password     string
	UserDocument string
	PhoneNumber  pgtype.Text
	BornDate     pgtype.Date
	CreatedAt    pgtype.Timestamp
}
