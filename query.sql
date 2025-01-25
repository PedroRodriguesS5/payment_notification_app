-- name: GetUser :one
SELECT * FROM users WHERE user_id = $1;

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

-- name: CreateRecurringPayment :one
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
RETURNING recurring_payment_id;

-- name: CreateSelfRecurringPayment :one
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
RETURNING recurring_payment_id;

-- name: GetUserByEmail :one
SELECT user_id,name, second_name email, password
FROM users
WHERE email = $1;

-- name: GetReceiverIdByEmail :one
SELECT user_id
FROM users
WHERE email = $1;

-- name: UpdateCharge :exec
UPDATE recurring_payment
SET payment_status = $2,
    end_date = $3,
    day_of_month = $4
WHERE recurring_payment_id = $1
RETURNING recurring_payment_id;

-- name: GetRecurringPaymentInfo :one
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
WHERE u.user_id = $1;



-- -- name: GetNotificationInfo :one
-- SELECT
--     n.notification_id,
--     n.recurring_payment_id,
--     rp.payer_name,
--     rp.receiver_name,
--     rp.amount,
--     rp.notification_type,
--     n.notification_date,
--     n.notification_status
-- FROM notification n
-- JOIN recurring_payment rp ON n.recurring_payment_id = rp.recurring_payment_id;
        
-- -- name: GetNotificationsByStatus :many
-- SELECT 
--     notification_id, 
--     rp.amount, 
--     rp.notification_type, 
--     notification_date 
-- FROM notification
-- JOIN recurring_payment rp ON notification.recurring_payment_id = rp.recurring_payment_id
-- WHERE notification_status = $1;

