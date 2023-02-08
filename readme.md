<h1 align="center">
  Golang Restfull Api With Gorilla/mux
</h1>

## 🛠️ Installation Steps

1. Clone the repository

```bash
git clone https://github.com/ardhisaif/Golang-backend.git
# or
git clone git@github.com:ardhisaif/Golang-backend.git
```

2. Install dependencies

```bash
go get -u ./...
```

3. Run the app

```bash
go run main.go
```

🌟 You are all set!

## 💻 Built with

-   [Golang](https://go.dev/): programming language
-   [gorilla/mux](https://github.com/gorilla/mux): for handle http request
-   [Postgres](https://www.postgresql.org/): for DBMS

<br />
<br />
<br />
<br />


# Gorlilla/mux
- Untuk implementasi Router
- Menjalankan Http requset dan response
- Menjalankan Http handler

# Mux alternatives
- Gin
- Fiber
- Echo

# Microservices
- Terdiri dari banyak service yang saling terhubung
- Setiap service memiliki fungsi atau fiture yang berbeda
- contoh: App gojek memiliki beberapa service atau applikasi serperti GoRide, GoPay, GoFood


<br />
<br />
<br />
<br />

# Dependency Injection
- teknik untuk mengatur dua objek yang saling berhubungan
- proses pemisahan berdasarkan fungsi masing-masing

# Sessions, Cookies, Tokens
### Session
- semua interaksi yang dilakukan user dalam jangka waktu tertentu
- di record melalui session id
- untuk menyimpan data disisi server

### Cookies
- file kecil yang menyimpan kumpulan session

### Token
- menyimpan dan mentransfer data sesnsitif dengan lebih aman(dicompile dalam bentuk signature JWT)
- memiliki header, jenis algoritma, payload

# DB Pooling
metode yang digunakan untuk menjaga koneksi database tetap terbuka sehingga dapat digunakan kembali oleh orang lain.
- Mempercepat response time
- Penggunaan kembali resource database
- Penggunaan memory lebih hemat