#!/bin/bash

# Fungsi untuk menampilkan bantuan
show_help() {
    echo "Usage: ./task.sh [command]"
    echo ""
    echo "Commands:"
    echo "  start    : Menjalankan server (main.go)"
    echo "  refresh  : Reset database & run seeders"
    echo "  seed     : Menjalankan database seeder saja"
    echo "  dev      : Refresh database lalu jalankan server"
}

case "$1" in
    "start")
        go run cmd/main.go
        ;;
    "refresh")
        go run scripts/refresh.go
        ;;
    "seed")
        go run scripts/seed.go
        ;;
    "dev")
                clear

                            # 1. Infrastructure
                            printf "\033[1;36m>>> STARTING INFRASTRUCTURE\033[0m\n"
                            docker-compose up -d
                            printf "Grafana  : http://localhost:3000 (admin/admin)\n"
                            printf "Loki     : http://loki:3100\n"
                            printf "Metrics  : http://localhost:8080/metrics\n"
                            printf -- "----------------------------------------------------------\n"

                            # 2. Setup Guide (Loki)
                            printf "\033[1;33m[ LOKI SETUP GUIDE ]\033[0m\n"
                            printf "1. Connections > Data Sources > Add Loki\n"
                            printf "2. URL: http://loki:3100 > Save & Test\n"
                            printf "3. Explore > Query: {job=\"parking-api\"}\n"
                            printf -- "----------------------------------------------------------\n"

                            # 3. Database
                            printf "\033[1;36m>>> DATABASE REFRESH\033[0m\n"
                            go run scripts/refresh.go
                            printf -- "----------------------------------------------------------\n"

                            # 4. App
                            printf "\033[1;32m>>> APPLICATION START\033[0m\n"
                            go run cmd/main.go
                ;;
    *)
        show_help
        ;;
esac