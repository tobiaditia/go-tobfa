Catatan DEV:
1. Penggunaiaan pointer *, digunakan apabila type nya adalah Struct kalau Interface tidak perlu. Referensi [Title](https://dasarpemrogramangolang.novalagung.com/A-pointer.html)

CLI
1. Goose
buat file
goose -dir ./db/migrations create users sql
cek status
goose -dir ./db/migrations mysql "root:@/tobfa_go?parseTime=true" status
naikkan migrasi dan seeder
goose -dir ./db/migrations mysql "root:@/tobfa_go?parseTime=true" up

Error yang pernah ditemukan
goose
goose run: failed to ensure DB version: no next version found
gara2 hapus manual isi table goose_db_version
isi manual dengan version_id 0 dan is_applied 1