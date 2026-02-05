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
        go run scripts/refresh.go && go run cmd/main.go
        ;;
    *)
        show_help
        ;;
esac