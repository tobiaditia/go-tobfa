-- +goose Up
DROP TABLE IF EXISTS business_transaction_items;
CREATE TABLE business_transaction_items (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

TRUNCATE TABLE business_transaction_items;
INSERT INTO business_transaction_items (name, created_at, updated_at) VALUES
('Pakan', NOW(), NOW())
;

-- +goose Down
DROP TABLE IF EXISTS business_transaction_items;
