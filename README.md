# GolangCRUD_Project
 Ini adalah Project CRUD menggunakan API Golang ber framework Gin, Database menggunakan PostgreSQL dan frontend menggunakan Tailwind CSS.
 
Saat ini, Project masih dalam proses pembuatan front-end UI dan pembuatan Javascript untuk pengaplikasian API kedalam front-end. Untuk CRUD API dapat dijalankan menggunakan Postman.
 
 ## Perencanaan Project
 
 GOLANG GIN API = ✔️
 CRUD = ✔️
 Desain Frontend = ✔️
 Javascript = ❌
 
 ## Sebelum Penjalanan
 
 Harap diperhatikan! Sebelum menjalankannya, lakukan cara berikut:
 
 1. Silahkan download master atau clone repository ini.
```
git clone https://github.com/adam-ghafara/GolangCRUD_Project/
```
 2. Silahkan Buat folder baru bernama "Autobros"
 3. Copy seluruh data pada repo ini kedalam folder tersebut.
 4. Jangan lupa untuk membuat database yang disediakan repo didalam Folder "BUAT DATABASE" silahkan masukkan sql kedalam PostgreSQL **(Wajib menggunakan PostgreSQL, MySQL atau DB lainnya tidak dapat jalan).**

## Penjalanan API

Jika sudah dilakukan cara tersebut, silahkan buka folder menggunakan VScode. Buka Terminal dan Jalankan perintah berikut:
```
go run app.go
```

## Test API

Tes API menggunakan Postman dengan cara berikut:
1. Silahkan Buka Postman
2. Buat Request Baru, dan silahkan lakukan praktek berikut:
### POST
1. Pilih Post dan Masukkan Link berikut kedalam URL
```
localhost:8000/autobros/tambahservis
```
2. Buka Body -> Pilih RAW, Masukkan Script berikut:
```
{
  "id_jenis": 0, // Pilih 0 = Mobil, atau 1 = Sepeda Motor
  "nama_pemilik": "Isi Sendiri",
  "nama_kendaraan": "Isi Sendiri",
  "nomor_kendaraan": Isi Sendiri dengan Angka,
  "detail_servis": "Isi Sendiri",
  "status_servis": "Di isi Pending/Dalam Perbaikan/Selesai."
}
```
3. Jika Sudah, Send Request dan lihat Hasilnya.
### GET
1. Pilih Get dan Masukkan Link berikut kedalam URL
```
localhost:8000/autobros/tabelservis
```
2. Langsung Jalankan dan lihat Hasilnya.
### PUT
1. Pilih Put Masukkan Link berikut kedalam URL
```
localhost:8000/autobros/(ANGKA ID DALAM DATABASE)
```
2. Isi seperti sebelumnya, silahkan ubah salah satu data yang ingin di ubah.
```
{
  "id_jenis": 0, // Pilih 0 = Mobil, atau 1 = Sepeda Motor
  "nama_pemilik": "Isi Sendiri",
  "nama_kendaraan": "Isi Sendiri",
  "nomor_kendaraan": Isi Sendiri dengan Angka,
  "detail_servis": "Isi Sendiri",
  "status_servis": "Di isi Pending/Dalam Perbaikan/Selesai."
}
```
3. Jalankan dan lihat Hasilnya.
### DELETE
1. Pilih Delete dan Masukkan Link berikut kedalam URL
```
localhost:8000/autobros/(ANGKA ID DALAM DATABASE)
```
2. Silahkan Request langsung dan lihat Hasilnya.

<h2> SELAMAT MENCOBA ▶▶▶ </h2> 
