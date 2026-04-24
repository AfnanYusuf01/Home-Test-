# Paket Data API

REST API sederhana untuk pengelolaan pembelian Paket Data (Golang + Fiber + PostgreSQL + GORM).

## Tech Stack

- **Golang** - Bahasa pemrograman
- **Fiber v2** - Framework HTTP
- **PostgreSQL** - Database
- **GORM** - ORM
- **Docker** - Containerization

## Struktur Project

```
├── config/          # Konfigurasi database
├── models/          # Model & DTO
├── repository/      # Repository layer (akses database)
├── service/         # Service layer (business logic & validasi)
├── handler/         # Handler layer (HTTP handler)
├── routes/          # Routing
├── postman/         # Postman Collection
├── main.go          # Entry point
├── docker-compose.yml
└── Dockerfile
```

## Cara Menjalankan

### Prasyarat

- Go v1.21+
- PostgreSQL v15+
- Git

### 1. Clone Repository

```bash
git clone https://github.com/AfnanYusuf01/-Take-Home-Test-.git
cd -Take-Home-Test-
```

### 2. Setup Database

**Opsi A: Docker (Rekomendasi)**
```bash
docker-compose up -d db
```

**Opsi B: Manual**
```sql
CREATE DATABASE paket_data_db;
```

### 3. Konfigurasi Environment

```bash
cp .env.example .env
```

Edit `.env` sesuai konfigurasi database:
```
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

### Docker (Full Stack)

```bash
docker-compose up --build
```

## API Endpoints

### User (`/api/users`)

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/api/users` | Semua user |
| GET | `/api/users/:id` | User by ID |
| POST | `/api/users` | Buat user |
| PUT | `/api/users/:id` | Update user |
| DELETE | `/api/users/:id` | Hapus user |

### Paket Data (`/api/paket-data`)

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/api/paket-data` | Semua paket |
| GET | `/api/paket-data/:id` | Paket by ID |
| POST | `/api/paket-data` | Buat paket |
| PUT | `/api/paket-data/:id` | Update paket |
| DELETE | `/api/paket-data/:id` | Hapus paket |

### Transaksi (`/api/transaksi`)

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/api/transaksi` | Semua transaksi |
| GET | `/api/transaksi/:id` | Transaksi by ID |
| POST | `/api/transaksi` | Buat transaksi |

## Contoh Request

**Create User:**
```json
POST /api/users
{ "name": "Afnan Yusuf", "phone_number": "081234567890" }
```

**Create Paket Data:**
```json
POST /api/paket-data
{ "name": "Paket Internet 10GB", "price": 50000, "quota": 10, "active_period": 30 }
```

**Create Transaksi:**
```json
POST /api/transaksi
{ "user_id": 1, "paket_data_id": 1 }
```

## Postman Collection

Import file `postman/Paket_Data_API.postman_collection.json` ke Postman.

## Catatan

- Tidak ada authentication (JWT, dll)
- Harga transaksi otomatis dari harga paket data saat transaksi
- Database auto-migrate saat aplikasi pertama kali jalan
