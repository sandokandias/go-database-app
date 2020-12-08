CREATE TABLE IF NOT EXISTS customers
(
    customer_id VARCHAR(100) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    address VARCHAR(255) NOT NULL
);

ALTER TABLE orders 
ADD COLUMN customer_id VARCHAR CONSTRAINT fk_customer REFERENCES customers(customer_id);