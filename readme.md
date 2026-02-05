# 🅿️ Parking Management System v1

Sistem manajemen parkir berbasis API yang dibangun dengan **Go (Fiber)**, dilengkapi dengan infrastruktur monitoring
modern menggunakan **Loki, Promtail, dan Grafana**.

## 🚀 Fitur Utama

- **High Performance API** menggunakan Go Fiber.
- **Automated Documentation** dengan Swagger.
- **Centralized Logging** menggunakan Grafana Loki.
- **Database Refresh & Seeding** otomatis.
- **Containerized Infrastructure** dengan Docker Compose.

---

## 🛠️ Prasyarat

Sebelum memulai, pastikan Anda sudah menginstal:

- [Go 1.25+](https://go.dev/dl/)
- [Docker & Docker Compose](https://www.docker.com/products/docker-desktop/)
- [Swag CLI](https://github.com/swaggo/swag) (`go install github.com/swaggo/swag/cmd/swag@latest`)

---

## 🏁 Cara Menjalankan

Project ini mendukung dua mode pengembangan menggunakan `Makefile` agar lebih efisien.

### 1. Mode Hybrid (Aplikasi Lokal + Infra Docker)

Mode ini direkomendasikan untuk pengembangan harian karena proses kompilasi lebih cepat.
> **Note:** Pastikan di file `.env` Anda, `DB_HOST` diatur ke `localhost`.

```bash
# Update Swagger, Nyalakan DB & Logging, lalu Run App
make dev
```

### 2. Mode Full Docker (Simulasi Production)

Mode ini menjalankan seluruh sistem (termasuk aplikasi) di dalam kontainer Docker.

> **Note:** Konfigurasi `DB_HOST` akan otomatis diarahkan ke container `parking-db` melalui docker-compose environment
> override.

```bash
# Update Swagger dan Jalankan seluruh kontainer
make docker-dev
```

---

## 📊 Akses Layanan & Monitoring

Setelah menjalankan perintah di atas, Anda dapat mengakses layanan berikut:

| Layanan        | URL                                        | Kredensial        |
|----------------|--------------------------------------------|-------------------|
| **Swagger UI** | `http://localhost:8080/swagger/index.html` | -                 |
| **Grafana**    | `http://localhost:3000`                    | `admin` / `admin` |
| **Metrics**    | `http://localhost:8080/metrics`            | -                 |
| **Loki API**   | `http://localhost:3100`                    | -                 |

---

## 🔍 Panduan Setup Monitoring (Loki)

Agar log aplikasi muncul di Grafana, ikuti langkah berikut (instruksi ini juga muncul di terminal saat Anda menjalankan
`make dev`):

1. Buka **Grafana** di browser.
2. Masuk ke **Connections > Data Sources > Add Data Source**.
3. Pilih **Loki**.
4. Masukkan URL: `http://parking-loki:3100`.
5. Klik **Save & Test**.
6. Buka menu **Explore**, pilih source **Loki**, dan masukkan query: `{job="parking-app"}`.

---

## 📜 Perintah Makefile Lainnya

| Perintah       | Deskripsi                                       |
|----------------|-------------------------------------------------|
| `make build`   | Membangun ulang (rebuild) seluruh image Docker. |
| `make down`    | Menghentikan semua layanan Docker.              |
| `make docs`    | Men-generate ulang dokumentasi Swagger.         |
| `make refresh` | Reset dan migrasi ulang database (Lokal).       |
| `make seed`    | Mengisi database dengan data dummy (Lokal).     |

---

## 📂 Struktur Folder Utama

* `/cmd`: Entry point aplikasi.
* `/internal`: Business logic (Handlers, Services, Repositories).
* `/pkg`: Library/Helper yang bisa digunakan kembali (Config, Logger).
* `/scripts`: Script untuk maintenance database.
* `/docs`: File dokumentasi Swagger (Auto-generated).