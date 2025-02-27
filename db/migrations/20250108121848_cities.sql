-- +goose Up
DROP TABLE IF EXISTS cities;
CREATE TABLE cities (
    id INT PRIMARY KEY AUTO_INCREMENT,
    code VARCHAR(255) NOT NULL,
    full_code VARCHAR(255) NOT NULL,
    province_id INT NOT NULL,
    type INT NOT NULL,
    name VARCHAR(255) NOT NULL
);

TRUNCATE TABLE cities;
INSERT INTO cities (code, full_code, province_id, type, name) VALUES
('05', '1105', 1, 2, 'Aceh Barat'),
('12', '1112', 1, 2, 'Aceh Barat Daya'),
('20', '5320', 23, 2, 'Sabu Raijua'),
('73', '3373', 10, 1, 'Salatiga'),
('06', '1106', 1, 2, 'Aceh Besar'),
('14', '1114', 1, 2, 'Aceh Jaya'),
('01', '1101', 1, 2, 'Aceh Selatan'),
('10', '1110', 1, 2, 'Aceh Singkil'),
('16', '1116', 1, 2, 'Aceh Tamiang'),
('04', '1104', 1, 2, 'Aceh Tengah'),
('02', '1102', 1, 2, 'Aceh Tenggara'),
('03', '1103', 1, 2, 'Aceh Timur'),
('05', '7505', 7, 2, 'Gorontalo Utara'),
('08', '1108', 1, 2, 'Aceh Utara'),
('06', '1306', 36, 2, 'Agam'),
('05', '5305', 23, 2, 'Alor'),
('71', '8171', 20, 1, 'Ambon'),
('06', '7306', 32, 2, 'Gowa'),
('72', '6472', 15, 1, 'Samarinda'),
('09', '1209', 38, 2, 'Asahan'),
('01', '6101', 12, 2, 'Sambas'),
('04', '9304', 28, 2, 'Asmat'),
('17', '1217', 38, 2, 'Samosir'),
('03', '5103', 2, 2, 'Badung'),
('27', '3527', 11, 2, 'Sampang'),
('11', '6311', 13, 2, 'Balangan'),
('03', '6103', 12, 2, 'Sanggau'),
('10', '9110', 24, 2, 'Sarmi'),
('25', '3525', 11, 2, 'Gresik'),
('15', '3315', 10, 2, 'Grobogan'),
('10', '6210', 14, 2, 'Gunung Mas'),
('03', '3403', 5, 2, 'Gunungkidul'),
('03', '1503', 8, 2, 'Sarolangun'),
('78', '1278', 38, 1, 'Gunungsitoli'),
('73', '1373', 36, 1, 'Sawahlunto'),
('01', '8201', 21, 2, 'Halmahera Barat'),
('09', '6109', 12, 2, 'Sekadau'),
('04', '8204', 21, 2, 'Halmahera Selatan'),
('05', '1705', 4, 2, 'Seluma'),
('02', '8202', 21, 2, 'Halmahera Tengah'),
('22', '3322', 10, 2, 'Semarang'),
('74', '3374', 10, 1, 'Semarang'),
('06', '8206', 21, 2, 'Halmahera Timur'),
('06', '8106', 20, 2, 'Seram Bagian Barat'),
('03', '8203', 21, 2, 'Halmahera Utara'),
('05', '8105', 20, 2, 'Seram Bagian Timur'),
('06', '6306', 13, 2, 'Hulu Sungai Selatan'),
('07', '6307', 13, 2, 'Hulu Sungai Tengah'),
('08', '6308', 13, 2, 'Hulu Sungai Utara'),
('73', '3673', 3, 1, 'Serang'),
('04', '3604', 3, 2, 'Serang'),
('16', '1216', 38, 2, 'Humbang Hasundutan'),
('71', '6471', 15, 1, 'Balikpapan'),
('04', '1404', 30, 2, 'Indragiri Hilir'),
('02', '1402', 30, 2, 'Indragiri Hulu'),
('71', '1171', 1, 1, 'Banda Aceh'),
('71', '1871', 19, 1, 'Bandar Lampung'),
('73', '3273', 9, 1, 'Bandung'),
('04', '3204', 9, 2, 'Bandung'),
('17', '3217', 9, 2, 'Bandung Barat'),
('01', '7201', 33, 2, 'Banggai'),
('07', '7207', 33, 2, 'Banggai Kepulauan'),
('11', '7211', 33, 2, 'Banggai Laut'),
('01', '1901', 17, 2, 'Bangka'),
('05', '1905', 17, 2, 'Bangka Barat'),
('03', '1903', 17, 2, 'Bangka Selatan'),
('04', '1904', 17, 2, 'Bangka Tengah'),
('26', '3526', 11, 2, 'Bangkalan'),
('06', '5106', 2, 2, 'Bangli'),
('79', '3279', 9, 1, 'Banjar'),
('03', '6303', 13, 2, 'Banjar'),
('72', '6372', 13, 1, 'Banjarbaru'),
('71', '6371', 13, 1, 'Banjarmasin'),
('04', '3304', 10, 2, 'Banjarnegara'),
('03', '7303', 32, 2, 'Bantaeng'),
('02', '3402', 5, 2, 'Bantul'),
('07', '1607', 37, 2, 'Banyuasin'),
('02', '3302', 10, 2, 'Banyumas'),
('10', '3510', 11, 2, 'Banyuwangi'),
('04', '6304', 13, 2, 'Barito Kuala'),
('04', '6204', 14, 2, 'Barito Selatan'),
('13', '6213', 14, 2, 'Barito Timur'),
('05', '6205', 14, 2, 'Barito Utara'),
('11', '7311', 32, 2, 'Barru'),
('71', '2171', 18, 1, 'Batam'),
('25', '3325', 10, 2, 'Batang'),
('04', '1504', 8, 2, 'Batanghari'),
('79', '3579', 11, 1, 'Batu'),
('19', '1219', 38, 2, 'Batu Bara'),
('72', '7472', 34, 1, 'Bau Bau'),
('75', '3275', 9, 1, 'Bekasi'),
('16', '3216', 9, 2, 'Bekasi'),
('02', '1902', 17, 2, 'Belitung'),
('06', '1906', 17, 2, 'Belitung Timur'),
('04', '5304', 23, 2, 'Belu'),
('17', '1117', 1, 2, 'Bener Meriah'),
('03', '1403', 30, 2, 'Bengkalis'),
('07', '6107', 12, 2, 'Bengkayang'),
('18', '1218', 38, 2, 'Serdang Bedagai'),
('71', '1771', 4, 1, 'Bengkulu'),
('01', '1701', 4, 2, 'Bengkulu Selatan'),
('07', '6207', 14, 2, 'Seruyan'),
('12', '3212', 9, 2, 'Indramayu'),
('08', '1408', 30, 2, 'Siak'),
('07', '9407', 29, 2, 'Intan Jaya'),
('73', '3173', 6, 1, 'Jakarta Barat'),
('73', '1273', 38, 1, 'Sibolga'),
('09', '1709', 4, 2, 'Bengkulu Tengah'),
('71', '3171', 6, 1, 'Jakarta Pusat'),
('14', '7314', 32, 2, 'Sidenreng Rappang'),
('03', '1703', 4, 2, 'Bengkulu Utara'),
('74', '3174', 6, 1, 'Jakarta Selatan'),
('15', '3515', 11, 2, 'Sidoarjo'),
('75', '3175', 6, 1, 'Jakarta Timur'),
('10', '7210', 33, 2, 'Sigi'),
('03', '6403', 15, 2, 'Berau'),
('06', '9106', 24, 2, 'Biak Numfor'),
('06', '5206', 22, 2, 'Bima'),
('72', '3172', 6, 1, 'Jakarta Utara'),
('72', '5272', 22, 1, 'Bima'),
('71', '1571', 8, 1, 'Jambi'),
('03', '1303', 36, 2, 'Sijunjung'),
('75', '1275', 38, 1, 'Binjai'),
('03', '9103', 24, 2, 'Jayapura'),
('01', '2101', 18, 2, 'Bintan'),
('07', '5307', 23, 2, 'Sikka'),
('11', '1111', 1, 2, 'Bireuen'),
('08', '1208', 38, 2, 'Simalungun'),
('72', '7172', 35, 1, 'Bitung'),
('71', '9171', 24, 1, 'Jayapura'),
('09', '1109', 1, 2, 'Simeulue'),
('01', '9501', 27, 2, 'Jayawijaya'),
('09', '3509', 11, 2, 'Jember'),
('72', '6172', 12, 1, 'Singkawang'),
('05', '3505', 11, 2, 'Blitar'),
('01', '5101', 2, 2, 'Jembrana'),
('07', '7307', 32, 2, 'Sinjai'),
('72', '3572', 11, 1, 'Blitar'),
('04', '7304', 32, 2, 'Jeneponto'),
('16', '3316', 10, 2, 'Blora'),
('20', '3320', 10, 2, 'Jepara'),
('05', '6105', 12, 2, 'Sintang'),
('02', '7502', 7, 2, 'Boalemo'),
('17', '3517', 11, 2, 'Jombang'),
('01', '3201', 9, 2, 'Bogor'),
('12', '3512', 11, 2, 'Situbondo'),
('08', '9208', 25, 2, 'Kaimana'),
('71', '3271', 9, 1, 'Bogor'),
('04', '3404', 5, 2, 'Sleman'),
('01', '1401', 30, 2, 'Kampar'),
('22', '3522', 11, 2, 'Bojonegoro'),
('02', '1302', 36, 2, 'Solok'),
('01', '7101', 35, 2, 'Bolaang Mongondow'),
('11', '7111', 35, 2, 'Bolaang Mongondow Selatan'),
('72', '1372', 36, 1, 'Solok'),
('03', '6203', 14, 2, 'Kapuas'),
('11', '1311', 36, 2, 'Solok Selatan'),
('10', '7110', 35, 2, 'Bolaang Mongondow Timur'),
('06', '6106', 12, 2, 'Kapuas Hulu'),
('13', '3313', 10, 2, 'Karanganyar'),
('07', '5107', 2, 2, 'Karangasem'),
('08', '7108', 35, 2, 'Bolaang Mongondow Utara'),
('15', '3215', 9, 2, 'Karawang'),
('06', '7406', 34, 2, 'Bombana'),
('02', '2102', 18, 2, 'Karimun'),
('11', '3511', 11, 2, 'Bondowoso'),
('06', '1206', 38, 2, 'Karo'),
('12', '7312', 32, 2, 'Soppeng'),
('06', '6206', 14, 2, 'Katingan'),
('04', '1704', 4, 2, 'Kaur'),
('01', '9201', 26, 2, 'Sorong'),
('11', '6111', 12, 2, 'Kayong Utara'),
('71', '9271', 26, 1, 'Sorong'),
('04', '9204', 26, 2, 'Sorong Selatan'),
('05', '3305', 10, 2, 'Kebumen'),
('14', '3314', 10, 2, 'Sragen'),
('06', '3506', 11, 2, 'Kediri'),
('13', '3213', 9, 2, 'Subang'),
('75', '1175', 1, 1, 'Subulussalam'),
('72', '3272', 9, 1, 'Sukabumi'),
('71', '3571', 11, 1, 'Kediri'),
('08', '7308', 32, 2, 'Bone'),
('11', '9111', 24, 2, 'Keerom'),
('02', '3202', 9, 2, 'Sukabumi'),
('03', '7503', 7, 2, 'Bone Bolango'),
('24', '3324', 10, 2, 'Kendal'),
('74', '6474', 15, 1, 'Bontang'),
('71', '7471', 34, 1, 'Kendari'),
('02', '9302', 28, 2, 'Boven Digoel'),
('08', '1708', 4, 2, 'Kepahiang'),
('09', '3309', 10, 2, 'Boyolali'),
('08', '6208', 14, 2, 'Sukamara'),
('05', '2105', 18, 2, 'Kepulauan Anambas'),
('11', '3311', 10, 2, 'Sukoharjo'),
('12', '5312', 23, 2, 'Sumba Barat'),
('29', '3329', 10, 2, 'Brebes'),
('75', '1375', 36, 1, 'Bukittinggi'),
('08', '5108', 2, 2, 'Buleleng'),
('02', '7302', 32, 2, 'Bulukumba'),
('18', '5318', 23, 2, 'Sumba Barat Daya')
;

-- +goose Down
DROP TABLE IF EXISTS cities;
