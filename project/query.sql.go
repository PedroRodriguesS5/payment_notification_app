// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package sqlc_db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createRecurringPayment = `-- name: CreateRecurringPayment :one
INSERT INTO recurring_payment (payer_id, receiver_id, payer_name, receiver_name, amount, notification_type, start_date, end_date, day_of_month)
SELECT 
    $1, 
    $2, 
    u1.name || ' ' || u1.second_name AS payer_name,   
    u2.name || ' ' || u2.second_name AS receiver_name, 
    $3, 
    $4, 
    $5, 
    $6, 
    $7
FROM users u1, users u2
WHERE u1.user_id = $1 AND u2.user_id = $2
RETURNING recurring_payment_id
`

type CreateRecurringPaymentParams struct {
	PayerID          pgtype.UUID
	ReceiverID       pgtype.UUID
	Amount           pgtype.Numeric
	NotificationType pgtype.Text
	StartDate        pgtype.Date
	EndDate          pgtype.Date
	DayOfMonth       pgtype.Int2
}

func (q *Queries) CreateRecurringPayment(ctx context.Context, arg CreateRecurringPaymentParams) (int32, error) {
	row := q.db.QueryRow(ctx, createRecurringPayment,
		arg.PayerID,
		arg.ReceiverID,
		arg.Amount,
		arg.NotificationType,
		arg.StartDate,
		arg.EndDate,
		arg.DayOfMonth,
	)
	var recurring_payment_id int32
	err := row.Scan(&recurring_payment_id)
	return recurring_payment_id, err
}

const createSelfRecurringPayment = `-- name: CreateSelfRecurringPayment :one
INSERT INTO recurring_payment(payer_id,receiver_name,payer_name, amount, notification_type,  start_date, end_date, day_of_month)
SELECT
    $1,
    $2,
    u.name || ' ' || u.second_name AS payer_name, 
    $3, 
    $4, 
    $5, 
    $6,
    $7
FROM users u WHERE u.user_id = $1
RETURNING recurring_payment_id
`

type CreateSelfRecurringPaymentParams struct {
	PayerID          pgtype.UUID
	ReceiverName     string
	Amount           pgtype.Numeric
	NotificationType pgtype.Text
	StartDate        pgtype.Date
	EndDate          pgtype.Date
	DayOfMonth       pgtype.Int2
}

func (q *Queries) CreateSelfRecurringPayment(ctx context.Context, arg CreateSelfRecurringPaymentParams) (int32, error) {
	row := q.db.QueryRow(ctx, createSelfRecurringPayment,
		arg.PayerID,
		arg.ReceiverName,
		arg.Amount,
		arg.NotificationType,
		arg.StartDate,
		arg.EndDate,
		arg.DayOfMonth,
	)
	var recurring_payment_id int32
	err := row.Scan(&recurring_payment_id)
	return recurring_payment_id, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users(name,second_name, email,password ,phone_number, user_document, born_date)
VALUES($1, $2, $3, $4, $5, $6, $7)
RETURNING user_id
`

type CreateUserParams struct {
	Name         string
	SecondName   string
	Email        string
	Password     string
	PhoneNumber  pgtype.Text
	UserDocument string
	BornDate     pgtype.Date
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (pgtype.UUID, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Name,
		arg.SecondName,
		arg.Email,
		arg.Password,
		arg.PhoneNumber,
		arg.UserDocument,
		arg.BornDate,
	)
	var user_id pgtype.UUID
	err := row.Scan(&user_id)
	return user_id, err
}

const getReceiverIdByEmail = `-- name: GetReceiverIdByEmail :one
SELECT user_id
FROM users
WHERE email = $1
`

func (q *Queries) GetReceiverIdByEmail(ctx context.Context, email string) (pgtype.UUID, error) {
	row := q.db.QueryRow(ctx, getReceiverIdByEmail, email)
	var user_id pgtype.UUID
	err := row.Scan(&user_id)
	return user_id, err
}

const getRecurringPaymentInfo = `-- name: GetRecurringPaymentInfo :one
SELECT 
    rp.recurring_payment_id,
    rp.amount,
    rp.start_date,
    rp.end_date,
    rp.payer_name,  
    rp.receiver_name,
    rp.day_of_month,
    rp.payment_status,
    u.email AS receiver_email
FROM recurring_payment rp
JOIN users u ON u.user_id = rp.receiver_id
WHERE u.user_id = $1
`

type GetRecurringPaymentInfoRow struct {
	RecurringPaymentID int32
	Amount             pgtype.Numeric
	StartDate          pgtype.Date
	EndDate            pgtype.Date
	PayerName          string
	ReceiverName       string
	DayOfMonth         pgtype.Int2
	PaymentStatus      pgtype.Text
	ReceiverEmail      string
}

func (q *Queries) GetRecurringPaymentInfo(ctx context.Context, userID pgtype.UUID) (GetRecurringPaymentInfoRow, error) {
	row := q.db.QueryRow(ctx, getRecurringPaymentInfo, userID)
	var i GetRecurringPaymentInfoRow
	err := row.Scan(
		&i.RecurringPaymentID,
		&i.Amount,
		&i.StartDate,
		&i.EndDate,
		&i.PayerName,
		&i.ReceiverName,
		&i.DayOfMonth,
		&i.PaymentStatus,
		&i.ReceiverEmail,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT user_id, name, second_name, email, password, user_document, phone_number, born_date, created_at FROM users WHERE user_id = $1
`

func (q *Queries) GetUser(ctx context.Context, userID pgtype.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUser, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Name,
		&i.SecondName,
		&i.Email,
		&i.Password,
		&i.UserDocument,
		&i.PhoneNumber,
		&i.BornDate,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT user_id,name, second_name email, password
FROM users
WHERE email = $1
`

type GetUserByEmailRow struct {
	UserID   pgtype.UUID
	Name     string
	Email    string
	Password string
}

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (GetUserByEmailRow, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i GetUserByEmailRow
	err := row.Scan(
		&i.UserID,
		&i.Name,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const listPayers = `-- name: ListPayers :many
SELECT DISTINCT u.user_id, u.email, u.name, u.second_name
FROM users u
JOIN recurring_payment rp ON u.user_id = rp.payer_id
WHERE rp.receiver_id = $1
`

type ListPayersRow struct {
	UserID     pgtype.UUID
	Email      string
	Name       string
	SecondName string
}

func (q *Queries) ListPayers(ctx context.Context, receiverID pgtype.UUID) ([]ListPayersRow, error) {
	rows, err := q.db.Query(ctx, listPayers, receiverID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListPayersRow
	for rows.Next() {
		var i ListPayersRow
		if err := rows.Scan(
			&i.UserID,
			&i.Email,
			&i.Name,
			&i.SecondName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listReceivers = `-- name: ListReceivers :many
SELECT DISTINCT u.user_id, u.email, u.name, u.second_name
FROM users u
JOIN recurring_payment rp ON u.user_id = rp.receiver_id
WHERE rp.payer_id = $1
`

type ListReceiversRow struct {
	UserID     pgtype.UUID
	Email      string
	Name       string
	SecondName string
}

func (q *Queries) ListReceivers(ctx context.Context, payerID pgtype.UUID) ([]ListReceiversRow, error) {
	rows, err := q.db.Query(ctx, listReceivers, payerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListReceiversRow
	for rows.Next() {
		var i ListReceiversRow
		if err := rows.Scan(
			&i.UserID,
			&i.Email,
			&i.Name,
			&i.SecondName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCharge = `-- name: UpdateCharge :exec
UPDATE recurring_payment
SET payment_status = $2,
    end_date = $3,
    day_of_month = $4
WHERE recurring_payment_id = $1
RETURNING recurring_payment_id
`

type UpdateChargeParams struct {
	RecurringPaymentID int32
	PaymentStatus      pgtype.Text
	EndDate            pgtype.Date
	DayOfMonth         pgtype.Int2
}

func (q *Queries) UpdateCharge(ctx context.Context, arg UpdateChargeParams) error {
	_, err := q.db.Exec(ctx, updateCharge,
		arg.RecurringPaymentID,
		arg.PaymentStatus,
		arg.EndDate,
		arg.DayOfMonth,
	)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET name = $2,
    second_name =$3,
    email = $4,
    born_date = $5
WHERE user_id = $1
`

type UpdateUserParams struct {
	UserID     pgtype.UUID
	Name       string
	SecondName string
	Email      string
	BornDate   pgtype.Date
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser,
		arg.UserID,
		arg.Name,
		arg.SecondName,
		arg.Email,
		arg.BornDate,
	)
	return err
}
