-- name: GetUser :one
SELECT * FROM users WHERE user_id = $1;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: ListPayers :many
SELECT DISTINCT u.user_id, u.email, u.name, u.second_name
FROM users u
JOIN recurring_payment rp ON u.user_id = rp.payer_id
WHERE rp.receiver_id = $1;

-- name: ListReceivers :many
SELECT DISTINCT u.user_id, u.email, u.name, u.second_name
FROM users u
JOIN recurring_payment rp ON u.user_id = rp.receiver_id
WHERE rp.payer_id = $1;

-- name: CreateUser :one
INSERT INTO users(name,second_name, email,password ,phone_number, user_document, born_date)
VALUES($1, $2, $3, $4, $5, $6, $7)
RETURNING user_id;

-- name: UpdateUser :exec
UPDATE users
SET name = $2,
    second_name =$3,
    email = $4,
    born_date = $5
WHERE user_id = $1;

-- name: CreateCharge :one
INSERT INTO recurring_payment(payer_id, receiver_id, amount, start_date, end_date, day_of_month)
VALUES($1, $2, $3, $4, $5, $6)
RETURNING recurring_payment_id;

-- name: GetUserByEmail :one
SELECT user_id, email, password
FROM users
WHERE email = $1;

-- name: UpdateCharge :exec
UPDATE recurring_payment
SET payment_status = $2,
    end_date = $3,
    day_of_month = $4
WHERE recurring_payment_id = $1;

-- name: CreateNotification :one
INSERT INTO notification(recurring_payment_id, notification_type)
VALUES($1, $2)
RETURNING notification_id;  -- Corrigi para retornar o ID correto da notificação

-- name: GetPaymentInfo :one
SELECT 
    rp.recurring_payment_id,
    rp.amount,
    rp.start_date,
    rp.end_date,
    rp.day_of_month,
    rp.payment_status,
    u.email AS receiver_email
FROM recurring_payment rp
JOIN users u ON u.user_id = rp.receiver_id
WHERE rp.payment_status = 'active' 
AND CURRENT_DATE BETWEEN rp.start_date AND rp.end_date
AND EXTRACT(DAY FROM CURRENT_DATE) = rp.day_of_month;
