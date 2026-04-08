# Evermos Backend API

Mini Project REST API untuk platform social commerce Evermos, dibangun menggunakan Golang sebagai tugas Project Based Internship (PBI) Rakamin Academy x Evermos Indonesia.

## Tech Stack

- **Go** (Golang) — Bahasa pemrograman utama
- **Fiber v2** — Web framework
- **GORM** — ORM untuk database
- **MySQL** — Database
- **JWT** — Autentikasi
- **bcrypt** — Enkripsi password

## Arsitektur

Project ini menerapkan **Clean Architecture** dengan pemisahan layer:

```
controllers/   → Menerima HTTP request, validasi input, return response
services/      → Business logic
repository/    → Query database (GORM)
models/        → Struct representasi tabel database
middleware/    → JWT authentication, admin authorization
utils/         → Helper (response format, pagination, upload file)
routes/        → Definisi endpoint API
config/        → Konfigurasi database
```

## Cara Menjalankan

### Prasyarat
- Go 1.21+
- MySQL/MariaDB
- Postman (untuk testing)

### Langkah-langkah

1. Clone repository
```bash
git clone <repository-url>
cd evermos
```

2. Buat database MySQL
```sql
CREATE DATABASE rakamin;
```

3. Sesuaikan konfigurasi database di `config/database.go`
```go
host := "127.0.0.1"
port := "3306"
user := "root"
password := ""
dbname := "rakamin"
```

4. Jalankan server
```bash
go run main.go
```

5. Server berjalan di `http://localhost:8080`

6. Untuk mengaktifkan user admin, ubah langsung di database:
```sql
UPDATE users SET is_admin = true WHERE id = <user_id>;
```
Kemudian login ulang agar token baru mengandung role admin.

## Daftar Endpoint API

### Auth
| Method | Endpoint | Deskripsi | Auth |
|--------|----------|-----------|------|
| POST | `/api/v1/auth/register` | Register user baru (toko otomatis terbuat) | - |
| POST | `/api/v1/auth/login` | Login, mendapat JWT token | - |

### User
| Method | Endpoint | Deskripsi | Auth |
|--------|----------|-----------|------|
| GET | `/api/v1/user` | Get profil sendiri | JWT |
| PUT | `/api/v1/user` | Update profil sendiri | JWT |

### Alamat (sub-route User)
| Method | Endpoint | Deskripsi | Auth |
|--------|----------|-----------|------|
| GET | `/api/v1/user/alamat` | Get semua alamat sendiri | JWT |
| GET | `/api/v1/user/alamat/:id` | Get alamat by ID | JWT |
| POST | `/api/v1/user/alamat` | Tambah alamat baru | JWT |
| PUT | `/api/v1/user/alamat/:id` | Update alamat | JWT |
| DELETE | `/api/v1/user/alamat/:id` | Hapus alamat | JWT |

### Toko
| Method | Endpoint | Deskripsi | Auth |
|--------|----------|-----------|------|
| GET | `/api/v1/toko` | Get semua toko (pagination + filter) | - |
| GET | `/api/v1/toko/my` | Get toko sendiri | JWT |
| GET | `/api/v1/toko/:id_toko` | Get toko by ID | - |
| PUT | `/api/v1/toko/:id_toko` | Update toko (pemilik only) | JWT |

### Kategori
| Method | Endpoint | Deskripsi | Auth |
|--------|----------|-----------|------|
| GET | `/api/v1/category` | Get semua kategori (pagination + filter) | - |
| GET | `/api/v1/category/:id` | Get kategori by ID | - |
| POST | `/api/v1/category` | Tambah kategori | Admin |
| PUT | `/api/v1/category/:id` | Update kategori | Admin |
| DELETE | `/api/v1/category/:id` | Hapus kategori | Admin |

### Produk
| Method | Endpoint | Deskripsi | Auth |
|--------|----------|-----------|------|
| GET | `/api/v1/product` | Get semua produk (pagination + filter) | - |
| GET | `/api/v1/product/:id` | Get produk by ID | - |
| POST | `/api/v1/product` | Tambah produk + upload foto | JWT |
| PUT | `/api/v1/product/:id` | Update produk (pemilik only) | JWT |
| DELETE | `/api/v1/product/:id` | Hapus produk (pemilik only) | JWT |

**Filter produk:** `?nama_produk=xxx&category_id=1&toko_id=1&min_harga=1000&max_harga=50000`

### Transaksi
| Method | Endpoint | Deskripsi | Auth |
|--------|----------|-----------|------|
| GET | `/api/v1/trx` | Get semua transaksi sendiri (pagination) | JWT |
| GET | `/api/v1/trx/:id` | Get detail transaksi | JWT |
| POST | `/api/v1/trx` | Buat transaksi baru | JWT |

### Provinsi & Kota (API Wilayah)
| Method | Endpoint | Deskripsi | Auth |
|--------|----------|-----------|------|
| GET | `/api/v1/provcity/listprovincies` | Get semua provinsi | - |
| GET | `/api/v1/provcity/detailprovince/:prov_id` | Get detail provinsi | - |
| GET | `/api/v1/provcity/listcities/:prov_id` | Get kota by provinsi | - |
| GET | `/api/v1/provcity/detailcity/:city_id` | Get detail kota | - |

## Contoh Request & Response

### Register
```json
// POST /api/v1/auth/register
// Request:
{
    "nama": "John Doe",
    "email": "john@mail.com",
    "no_telp": "081234567890",
    "kata_sandi": "password123",
    "tanggal_lahir": "2000-01-15",
    "pekerjaan": "Developer",
    "id_provinsi": "32",
    "id_kota": "3204"
}

// Response (201):
{
    "status": true,
    "message": "Register berhasil",
    "errors": null,
    "data": {
        "token": "eyJhbGciOiJ...",
        "user": {
            "id": 1,
            "nama": "John Doe",
            "email": "john@mail.com",
            "toko": {
                "id": 1,
                "nama_toko": "Toko John Doe"
            }
        }
    }
}
```

### Login
```json
// POST /api/v1/auth/login
// Request:
{
    "no_telp": "081234567890",
    "kata_sandi": "password123"
}

// Response (200):
{
    "status": true,
    "message": "Login berhasil",
    "errors": null,
    "data": {
        "token": "eyJhbGciOiJ...",
        "user": { ... }
    }
}
```

### Create Transaksi
```json
// POST /api/v1/trx
// Header: Authorization: Bearer <token>
// Request:
{
    "alamat_kirim": 1,
    "method_bayar": "COD",
    "detail_trx": [
        { "product_id": 2, "kuantitas": 3 },
        { "product_id": 3, "kuantitas": 1 }
    ]
}

// Response (201):
{
    "status": true,
    "message": "Transaksi berhasil dibuat",
    "errors": null,
    "data": {
        "trx": {
            "id": 1,
            "harga_total": 580000,
            "kode_invoice": "INV-20260408-001",
            "method_bayar": "COD"
        },
        "detail_trx": [ ... ]
    }
}
```

## Fitur Utama

- **Autentikasi JWT** — Token berlaku 24 jam
- **Auto-create Toko** — Toko otomatis dibuat saat register
- **Upload File** — Foto produk melalui multipart/form-data
- **Role Admin** — Hanya admin yang bisa mengelola kategori
- **Validasi Kepemilikan** — User hanya bisa mengelola data miliknya sendiri
- **Pagination** — Semua list endpoint mendukung `?page=1&limit=10`
- **Filtering** — Filter produk, toko, dan kategori berdasarkan parameter
- **Log Produk** — Snapshot data produk saat transaksi dibuat
- **DB Transaction** — Rollback otomatis jika terjadi error saat transaksi
- **API Wilayah** — Integrasi dengan API wilayah Indonesia (emsifa.com)

## Struktur Database

Tabel yang digunakan:
- `users` — Data pengguna
- `tokos` — Data toko (1 user = 1 toko)
- `alamats` — Alamat pengiriman
- `kategoris` — Kategori produk
- `produks` — Data produk
- `foto_produks` — Foto produk
- `trxes` — Transaksi
- `detail_trxes` — Detail transaksi
- `log_produks` — Snapshot produk saat transaksi
