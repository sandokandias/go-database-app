CREATE TABLE IF NOT EXISTS items
(
    item_id VARCHAR(100) PRIMARY KEY,
    order_id VARCHAR(100),
    description VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    quantity INT NOT NULL,
    CONSTRAINT fk_order FOREIGN KEY(order_id) REFERENCES orders(order_id)
);