-- +goose Up
DROP TABLE IF EXISTS business_transactions;
CREATE TABLE business_transactions (
    id INT PRIMARY KEY AUTO_INCREMENT,
    business_id INT unsigned NOT NULL,
    business_transaction_type_id INT unsigned NOT NULL,
    business_transaction_item_id INT unsigned NOT NULL,
    total INT unsigned NOT NULL,
    quantity INT unsigned NOT NULL,
    date DATE NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS business_transactions;
