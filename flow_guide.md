```
1. MODELS                          ✅ (udah selesai)
   └── internal/models/*.go

2. DATABASE
   └── pkg/database/postgres.go    → koneksi + AutoMigrate

3. REPOSITORIES
   └── internal/repositories/*.go  → semua query database (CRUD)

4. DTOs
   ├── dto/requests/*.go           → struct request dari API
   └── dto/responses/*.go          → struct response ke API

5. SERVICES
   └── internal/services/*.go      → business logic

6. HANDLERS
   └── internal/handlers/*.go      → HTTP handler (binding, validasi, panggil service)

7. ROUTES
   └── internal/routes/routes.go   → daftar semua endpoint

8. MIDDLEWARES
   └── internal/middlewares/*.go   → auth, logger, error handler

9. MAIN
   └── cmd/main.go                 → entry point, init semua

10. TESTING
    └── tests/                     → unit, integration, e2e

11. DEPLOY
    └── docker, env, dsb
```