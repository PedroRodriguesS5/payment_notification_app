-- Tabela Users
CREATE TABLE users (
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),  -- UUID usado como tipo de chave primária
    name VARCHAR(60) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    user_document VARCHAR(15) NOT NULL UNIQUE,
    phone_number VARCHAR(20),
    born_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabela Recurring_Payments
CREATE TABLE recurring_payment (
    recurring_payment_id SERIAL PRIMARY KEY,
    payer_id UUID,
    receiver_id UUID,
    amount DECIMAL(10, 2) NOT NULL,
    start_date DATE,
    end_date DATE,
    day_of_month SMALLINT,
    payment_status VARCHAR(20) DEFAULT 'active',
    FOREIGN KEY (payer_id) REFERENCES Users(user_id),
    FOREIGN KEY (receiver_id) REFERENCES Users(user_id)
);

-- Tabela Payments
CREATE TABLE notification (
    notification_id SERIAL PRIMARY KEY,
    user_id UUID,  -- Usar SERIAL para auto-incremento
    recurring_payment_id INT,
    amount DECIMAL(10, 2) NOT NULL,
    notification_type VARCHAR(20),  
    notification_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR(20) DEFAULT 'panding',
   FOREIGN KEY (recurring_payment_id) REFERENCES Recurring_Payments(recurring_payment_id)
);
