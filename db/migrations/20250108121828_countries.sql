-- +goose Up
DROP TABLE IF EXISTS countries;
CREATE TABLE countries (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL
);

TRUNCATE TABLE countries;
INSERT INTO countries (name) VALUES
('Indonesia')
;

-- +goose Down
DROP TABLE IF EXISTS countries;
