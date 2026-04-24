# Paket Data API

REST API untuk pengelolaan pembelian Paket Data menggunakan Golang, Fiber, PostgreSQL, dan GORM.

## Tech Stack

| Teknologi | Keterangan |
|-----------|------------|
| **Golang** | Bahasa pemrograman utama |
| **Fiber v2** | Framework HTTP yang cepat |
| **PostgreSQL** | Relational Database |
| **GORM** | ORM untuk Golang |
| **Validator v10** | Validasi input request |

## Struktur Project

```
├── config/          # Konfigurasi database
├── helpers/         # Response helper, validator & custom error types
├── models/          # Model database, Request DTO & Response DTO
├── repository/      # Repository layer (akses database)
├── service/         # Service layer (business logic & validasi)
├── handler/         # Handler layer (HTTP request/response)
├── routes/          # Routing endpoint
├── postman/         # Postman Collection
├── main.go          # Entry point aplikasi
└── README.md
```

### Arsitektur Layer

```
Request → Handler → Service → Repository → Database
                      ↓
                  Validasi &
                Business Logic
```

- **Handler**: Menerima HTTP request, parsing input, mengembalikan response (menggunakan Response DTO)
- **Service**: Validasi data, business logic, error wrapping
- **Repository**: Operasi CRUD ke database via GORM
- **Helpers**: Custom error types (`AppError`), unified response format, validator

## Cara Menjalankan

### Prasyarat

- Go v1.21+
- PostgreSQL v15+ (terpasang lokal)
- Git

### 1. Clone Repository

```bash
git clone https://github.com/AfnanYusuf01/-Take-Home-Test-.git
cd -Take-Home-Test-
```

### 2. Setup Database

Buat database PostgreSQL secara manual:
```sql
CREATE DATABASE paket_data_db;
```

### 3. Konfigurasi Environment

```bash
cp .env.example .env
```

Edit `.env` sesuai konfigurasi database lokal Anda:
```env
APP_PORT=3000
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=paket_data_db
```

### 4. Install Dependencies & Jalankan

```bash
go mod download
go run main.go
```

Server berjalan di `http://localhost:3000`

## Fitur Profesional & Clean Code

1. **Layered Architecture** — Pemisahan tanggung jawab yang jelas antara Handler → Service → Repository → Model
2. **Soft Delete** — Data User dan Paket Data tidak benar-benar dihapus dari database, melainkan ditandai dengan `deleted_at`. Memungkinkan audit data dan pemulihan jika diperlukan
3. **Response DTO** — Model database dan response JSON dipisahkan. Field internal seperti `deleted_at` tidak pernah terekspos ke client
4. **Robust Validation** — Menggunakan `go-playground/validator` dengan pesan error yang deskriptif per-field
5. **Unified Response Format** — Semua API response memiliki struktur konsisten (`success`, `code`, `message`, `timestamp`, `data`, `errors`)
6. **Global Error Handling** — Error ditangani secara terpusat di `main.go` menggunakan custom `AppError`, mencegah kebocoran detail teknis ke client
7. **Historical Integrity** — Transaksi yang sudah terjadi tetap tersimpan lengkap dengan data User & Paket Data meskipun keduanya sudah di-soft-delete

## Format Response

Semua endpoint menggunakan format response yang konsisten:

**Success Response:**
```json
{
  "success": true,
  "code": 200,
  "message": "Berhasil mengambil data user",
  "timestamp": "2026-04-25T04:00:00+07:00",
  "data": { ... }
}
```

**Error Response (General):**
```json
{
  "success": false,
  "code": 404,
  "message": "User dengan ID tersebut tidak ditemukan",
  "timestamp": "2026-04-25T04:00:00+07:00"
}
```

**Error Response (Validation):**
```json
{
  "success": false,
  "code": 400,
  "message": "Validasi input gagal",
  "timestamp": "2026-04-25T04:00:00+07:00",
  "errors": {
    "name": "Field ini wajib diisi",
    "phone_number": "Minimal 10 karakter/nilai"
  }
}
```

## API Endpoints

### User (`/api/users`)

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| `GET` | `/api/users` | Mengambil semua user aktif |
| `GET` | `/api/users/:id` | Mengambil user berdasarkan ID |
| `POST` | `/api/users` | Membuat user baru |
| `PUT` | `/api/users/:id` | Mengupdate user |
| `DELETE` | `/api/users/:id` | Soft delete user |

### Paket Data (`/api/paket-data`)

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| `GET` | `/api/paket-data` | Mengambil semua paket data aktif |
| `GET` | `/api/paket-data/:id` | Mengambil paket data berdasarkan ID |
| `POST` | `/api/paket-data` | Membuat paket data baru |
| `PUT` | `/api/paket-data/:id` | Mengupdate paket data |
| `DELETE` | `/api/paket-data/:id` | Soft delete paket data |

### Transaksi (`/api/transaksi`)

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| `GET` | `/api/transaksi` | Mengambil semua transaksi |
| `GET` | `/api/transaksi/:id` | Mengambil transaksi berdasarkan ID |
| `POST` | `/api/transaksi` | Membuat transaksi baru |

## Relasi & Soft Delete Logic

- **User & Paket Data**: Menggunakan **Soft Delete** via GORM `DeletedAt`. Saat dihapus, field `deleted_at` diisi timestamp. Data tidak muncul di query `Find`/`First` biasa, tapi tetap ada di database
- **Transaksi**: Data historis yang tidak pernah dihapus. Relasi ke User dan Paket Data menggunakan `Unscoped Preload` sehingga data yang sudah di-soft-delete tetap tampil pada transaksi (untuk keperluan reporting/audit)
- **Validasi Transaksi Baru**: Hanya User dan Paket Data yang masih aktif (belum di-soft-delete) yang bisa digunakan untuk membuat transaksi baru

## Postman Collection

Import file `postman/Paket_Data_API.postman_collection.json` ke Postman untuk testing semua endpoint.

## Catatan

- Tidak menggunakan authentication (JWT, dll) sesuai ketentuan test
- Harga transaksi otomatis diambil dari harga paket data saat transaksi dibuat
- Database auto-migrate saat aplikasi pertama kali dijalankan
- Soft delete memastikan integritas data historis transaksi
