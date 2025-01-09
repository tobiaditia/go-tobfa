-- +goose Up
DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    handphone VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

TRUNCATE TABLE users;
INSERT INTO users (name, email, password, handphone, created_at, updated_at) VALUES
('tobfa', 'me@tobfa.id', '12345678', '085111111111', NOW(), NOW());

-- +goose Down
DROP TABLE IF EXISTS users;
