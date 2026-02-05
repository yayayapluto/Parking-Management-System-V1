# Stage 1: Build
FROM golang:1.25.1-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
# Build App Utama
RUN go build -o main cmd/main.go
# Build Script Refresh jadi binary mandiri
RUN go build -o refresh-db scripts/refresh.go

# Stage 2: Runtime
FROM alpine:latest
# Kita GA BUTUH 'apk add go' lagi, jadi image lebih enteng & stabil
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/refresh-db .
COPY .env .
# Tetap copy database/ buat seeder (jika seeder baca file sql/csv)
COPY --from=builder /app/database ./database

RUN mkdir -p storage/logs
EXPOSE 8080
CMD ["./main"]