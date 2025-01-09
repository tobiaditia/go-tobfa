-- +goose Up
DROP TABLE IF EXISTS provinces;
CREATE TABLE provinces (
    id INT PRIMARY KEY AUTO_INCREMENT,
    code VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL
);

TRUNCATE TABLE provinces;
INSERT INTO provinces (code, name) VALUES
('11', 'Aceh (NAD)'),
('51', 'Bali'),
('36', 'Banten'),
('17', 'Bengkulu'),
('34', 'DI Yogyakarta'),
('31', 'DKI Jakarta'),
('75', 'Gorontalo'),
('15', 'Jambi'),
('32', 'Jawa Barat'),
('33', 'Jawa Tengah'),
('35', 'Jawa Timur'),
('61', 'Kalimantan Barat'),
('63', 'Kalimantan Selatan'),
('62', 'Kalimantan Tengah'),
('64', 'Kalimantan Timur'),
('65', 'Kalimantan Utara'),
('19', 'Kepulauan Bangka Belitung'),
('21', 'Kepulauan Riau'),
('18', 'Lampung'),
('81', 'Maluku'),
('82', 'Maluku Utara'),
('52', 'Nusa Tenggara Barat (NTB)'),
('53', 'Nusa Tenggara Timur (NTT)'),
('91', 'Papua'),
('92', 'Papua Barat'),
('92', 'Papua Barat Daya'),
('95', 'Papua Pegunungan'),
('93', 'Papua Selatan'),
('94', 'Papua Tengah'),
('14', 'Riau'),
('76', 'Sulawesi Barat'),
('73', 'Sulawesi Selatan'),
('72', 'Sulawesi Tengah'),
('74', 'Sulawesi Tenggara'),
('71', 'Sulawesi Utara'),
('13', 'Sumatera Barat'),
('16', 'Sumatera Selatan'),
('12', 'Sumatera Utara')
;

-- +goose Down
DROP TABLE IF EXISTS provinces;
