-- +goose Up
DROP TABLE IF EXISTS businesses;
CREATE TABLE businesses (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT unsigned NOT NULL,
    name VARCHAR(255) NOT NULL,
    address TEXT,
    business_category_id INT unsigned NOT NULL,
    country_id INT unsigned NOT NULL,
    province_id INT unsigned NOT NULL,
    city_id INT unsigned NOT NULL,
    district_id INT unsigned NOT NULL,
    village_id INT unsigned NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS businesses;
