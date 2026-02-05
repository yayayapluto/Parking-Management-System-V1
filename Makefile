# Variables
include .env
export $(shell sed 's/=.*//' .env)

APP_NAME=parking-api
COMPOSE=docker-compose

.PHONY: watch dev docker-dev build down logs refresh seed help docs

help:
	@clear
	@printf "\033[1;33mPARKING SYSTEM AUTOMATION\033[0m\n"
	@printf "----------------------------------------------------------\n"
	@printf "\033[1;36mCommands:\033[0m\n"
	@printf "  make dev          : Mode Hybrid (App Lokal + Infra Docker)\n"
	@printf "  make docker-dev   : Mode Full Docker (Semua di Container)\n"
	@printf "  make build        : Force rebuild docker images\n"
	@printf "  make down         : Stop all containers\n"
	@printf "  make docs         : Generate Swagger only\n"
	@printf "----------------------------------------------------------\n"

watch: docs
	@clear
	@printf "\033[1;36m>>> STARTING INFRASTRUCTURE\033[0m\n"
	@$(COMPOSE) up -d postgres loki promtail grafana
	@printf "\033[1;33m>>> STARTING AIR (HOT RELOAD)...\033[0m\n"
	@air

# --- MODE 1: HYBRID (Development Harian) ---
dev: docs
	@clear
	@# 1. Infrastructure
	@printf "\033[1;36m>>> STARTING INFRASTRUCTURE (HYBRID)\033[0m\n"
	@$(COMPOSE) up -d postgres loki promtail grafana
	@printf "Grafana  : http://localhost:3000 (admin/admin)\n"
	@printf "Loki     : http://localhost:3100\n"
	@printf "Metrics  : http://localhost:$(APP_PORT)/metrics\n"
	@printf "Swagger  : http://localhost:$(APP_PORT)/swagger/index.html\n"
	@printf -- "----------------------------------------------------------\n"

	@# 2. Setup Guide (Loki)
	@printf "\033[1;33m[ LOKI SETUP GUIDE ]\033[0m\n"
	@printf "1. Connections > Data Sources > Add Loki\n"
	@printf "2. URL: http://parking-loki:3100 > Save & Test\n"
	@printf "3. Explore > Query: {job=\"parking-app\"}\n"
	@printf -- "----------------------------------------------------------\n"

	@# 3. Database (Lokal)
	@printf "\033[1;36m>>> DATABASE REFRESH\033[0m\n"
	@go run scripts/refresh.go
	@printf -- "----------------------------------------------------------\n"

	@# 4. Application (Lokal)
	@printf "\033[1;32m>>> SYSTEM IS LIVE (LOCAL)\033[0m\n"
	@go run cmd/main.go

# --- MODE 2: FULL DOCKER (Simulasi Production) ---
dev-docker: docs
	@clear
	@# 1. Infrastructure & App
	@printf "\033[1;35m>>> STARTING FULL DOCKER MODE\033[0m\n"
	@$(COMPOSE) up -d
	@printf "Grafana  : http://localhost:3000 (admin/admin)\n"
	@printf "Loki     : http://localhost:3100\n"
	@printf "Metrics  : http://localhost:$(APP_PORT)/metrics\n"
	@printf "Swagger  : http://localhost:$(APP_PORT)/swagger/index.html\n"
	@printf -- "----------------------------------------------------------\n"

	@# 2. Setup Guide (Loki) - TETAP ADA DI SINI!
	@printf "\033[1;33m[ LOKI SETUP GUIDE ]\033[0m\n"
	@printf "1. Connections > Data Sources > Add Loki\n"
	@printf "2. URL: http://parking-loki:3100 > Save & Test\n"
	@printf "3. Explore > Query: {job=\"parking-app\"}\n"
	@printf -- "----------------------------------------------------------\n"

	@printf "Waiting for container stability (5s)...\n"
	@sleep 5

	@# 3. Database Refresh (Inside Docker)
	@printf "\033[1;36m>>> DATABASE REFRESH (DOCKER)\033[0m\n"
	@docker exec -it $(APP_NAME) ./refresh-db
	@printf -- "----------------------------------------------------------\n"

	@# 4. Application Logs
	@printf "\033[1;32m>>> SYSTEM IS LIVE (DOCKER)\033[0m\n"
	@docker logs -f $(APP_NAME)

# --- UTILITIES ---
docs:
	@printf "\033[1;34m>>> GENERATING SWAGGER DOCS\033[0m\n"
	@swag init -g cmd/main.go

build:
	@printf "\033[1;36m>>> BUILDING DOCKER IMAGES\033[0m\n"
	@$(COMPOSE) up -d --build

down:
	@printf "\033[1;31m>>> STOPPING ALL SERVICES\033[0m\n"
	@$(COMPOSE) down

refresh:
	@go run scripts/refresh.go

seed:
	@go run scripts/seed.go

logs:
	@docker logs -f $(APP_NAME)