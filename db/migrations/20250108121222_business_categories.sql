-- +goose Up
DROP TABLE IF EXISTS business_categories;
CREATE TABLE business_categories (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

TRUNCATE TABLE business_categories;
INSERT INTO business_categories (name, created_at, updated_at) VALUES
('Ikan Gurame', NOW(), NOW());

-- +goose Down
DROP TABLE IF EXISTS business_categories;
