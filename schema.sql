-- Tabela Users
CREATE TABLE users (
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),  -- UUID usado como tipo de chave primária
    name VARCHAR(20) NOT NULL,
    second_name VARCHAR(40) NOT NULL,
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
    payer_id UUID NOT NULL,
    receiver_id UUID NOT NULL,
    payer_name VARCHAR(60) NOT NULL,  
    receiver_name VARCHAR(60) NOT NULL,  
    amount DECIMAL(10, 2) NOT NULL,
    notification_type VARCHAR(20) DEFAULT 'email',
    start_date DATE,
    end_date DATE,
    day_of_month SMALLINT,
    payment_status VARCHAR(20) DEFAULT 'active',
    FOREIGN KEY (payer_id) REFERENCES Users(user_id),
    FOREIGN KEY (receiver_id) REFERENCES Users(user_id)
);

-- Tabela Payments
CREATE TABLE notification (
    notification_id SERIAL PRIMARY KEY, -- Usar SERIAL para auto-incremento
    recurring_payment_id INTEGER,
    notification_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    notification_status VARCHAR(20) DEFAULT 'pending',
   FOREIGN KEY (recurring_payment_id) REFERENCES Recurring_Payments(recurring_payment_id)
);

CREATE OR REPLACE FUNCTION generate_payment_notifications()
RETURNS TRIGGER AS $$
BEGIN
    -- Busca o amount e o notification_type do recurring_payment associado
    SELECT amount, notification_type 
    INTO recurring_amount, recurring_notification_type
    FROM recurring_payment
    WHERE recurring_payment_id = NEW.recurring_payment_id;

    -- Insere notificação apenas se não houver para o mesmo dia
    IF NOT EXISTS (
        SELECT 1 
        FROM notification 
        WHERE recurring_payment_id = NEW.recurring_payment_id 
        AND notification_date::date = CURRENT_DATE
    ) THEN
        INSERT INTO notification (
            recurring_payment_id,
            amount,
            day_of_month,
            notification_type,
            notification_status
        )
        VALUES (
            NEW.recurring_payment_id,
            NEW.amount, -- Valor buscado diretamente
            NEW.day_of_month, -- Mantém a informação específica
            NEW.notificaiton_type, -- Tipo buscado diretamente
            'pending'
        );
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;



CREATE INDEX idx_recurring_payment_id ON recurring_payment(recurring_payment_id);
CREATE INDEX idx_recurring_payment_start_date ON recurring_payment(start_date);
CREATE INDEX idx_recurring_payment_end_date ON recurring_payment(end_date);
CREATE INDEX idx_recurring_payment_day_of_month ON recurring_payment(day_of_month);