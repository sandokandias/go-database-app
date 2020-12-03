CREATE TABLE IF NOT EXISTS orders
(
    order_id VARCHAR(100) PRIMARY KEY,
    amount INT NOT NULL,
    created_at TIMESTAMP
);